package analysis

import (
	"fmt"
	"math/big"
	"numscript/parser"
	"slices"
)

const TypeMonetary = "monetary"
const TypeAccount = "account"
const TypePortion = "portion"
const TypeAsset = "asset"
const TypeNumber = "number"
const TypeString = "string"

const TypeAny = "*"

var AllowedTypes = []string{
	TypeMonetary,
	TypeAccount,
	TypePortion,
	TypeAsset,
	TypeNumber,
	TypeString,
}

const FnSetTxMeta = "set_tx_meta"
const FnSetAccountMeta = "set_account_meta"

var TopLevelFunctionsTypes = map[string][]string{
	FnSetTxMeta:      {TypeString, TypeAny},
	FnSetAccountMeta: {TypeAccount, TypeString, TypeAny},
}

const FnVarOriginMeta = "meta"
const FnVarOriginBalance = "balance"

type OriginFnArity struct {
	Args   []string
	Return string
}

var OriginFunctionArities = map[string]OriginFnArity{
	FnVarOriginMeta: {
		Args:   []string{TypeAccount, TypeString},
		Return: TypeAny,
	},
	FnVarOriginBalance: {
		Args:   []string{TypeAccount, TypeAsset},
		Return: TypeMonetary,
	},
}

type Diagnostic struct {
	Range parser.Range
	Kind  DiagnosticKind
}

type CheckResult struct {
	declaredVars  map[string]parser.VarDeclaration
	unusedVars    map[string]parser.Range
	varResolution map[*parser.VariableLiteral]parser.VarDeclaration
	Diagnostics   []Diagnostic
}

func (r CheckResult) ResolveVar(v *parser.VariableLiteral) *parser.VarDeclaration {
	k, ok := r.varResolution[v]
	if !ok {
		return nil
	}
	return &k
}

func Check(program parser.Program) CheckResult {
	res := CheckResult{
		declaredVars:  make(map[string]parser.VarDeclaration),
		unusedVars:    make(map[string]parser.Range),
		varResolution: make(map[*parser.VariableLiteral]parser.VarDeclaration),
	}
	for _, varDecl := range program.Vars {
		if varDecl.Type != nil {
			res.checkVarType(*varDecl.Type)
		}

		if varDecl.Name != nil {
			res.checkDuplicateVars(*varDecl.Name, varDecl)
		}

		if varDecl.Origin != nil {
			res.checkVarOrigin(*varDecl.Origin, varDecl)
		}
	}
	for _, statement := range program.Statements {
		switch statement := statement.(type) {
		case *parser.SendStatement:
			res.checkLiteral(statement.Monetary, TypeMonetary)
			res.checkSource(statement.Source)
			res.checkDestination(statement.Destination)
		case *parser.FnCallStatement:
			var sig []string
			if sigLookup, ok := TopLevelFunctionsTypes[statement.Caller.Name]; ok {
				sig = sigLookup
			}
			res.checkFnCallArity(statement, sig)
		}
	}

	// after static AST traversal is complete, check for unused vars
	for name, rng := range res.unusedVars {
		res.Diagnostics = append(res.Diagnostics, Diagnostic{
			Range: rng,
			Kind:  &UnusedVar{Name: name},
		})
	}
	return res
}

func (res *CheckResult) checkFnCallArity(statement *parser.FnCallStatement, sig []string) {
	var validArgs []parser.Literal
	for _, lit := range statement.Args {
		if lit != nil {
			validArgs = append(validArgs, lit)
		}
	}

	if sig != nil {
		actualArgs := len(validArgs)
		expectedArgs := len(sig)

		if actualArgs < expectedArgs {
			// Too few args
			res.Diagnostics = append(res.Diagnostics, Diagnostic{
				Range: statement.Range,
				Kind: &BadArity{
					Expected: expectedArgs,
					Actual:   actualArgs,
				},
			})
		} else if actualArgs > expectedArgs {
			// Too many args
			firstIllegalArg := validArgs[expectedArgs]
			lastIllegalArg := validArgs[len(validArgs)-1]

			if lastIllegalArg != nil {
				rng := parser.Range{
					Start: firstIllegalArg.GetRange().Start,
					End:   lastIllegalArg.GetRange().End,
				}

				res.Diagnostics = append(res.Diagnostics, Diagnostic{
					Range: rng,
					Kind: &BadArity{
						Expected: expectedArgs,
						Actual:   actualArgs,
					},
				})
			}

		}

		for index, arg := range validArgs {
			lastElemIndex := len(sig) - 1
			if index > lastElemIndex {
				break
			}

			type_ := sig[index]
			res.checkLiteral(arg, type_)
		}
	} else {
		for _, arg := range validArgs {
			res.checkLiteral(arg, "*")
		}

		res.Diagnostics = append(res.Diagnostics, Diagnostic{
			Range: statement.Caller.Range,
			Kind: &UnknownFunction{
				Name: statement.Caller.Name,
			},
		})
	}
}

func isTypeAllowed(typeName string) bool {
	return slices.Contains(AllowedTypes, typeName)
}

func (res *CheckResult) checkVarType(typeDecl parser.TypeDecl) {
	if !isTypeAllowed(typeDecl.Name) {
		res.Diagnostics = append(res.Diagnostics, Diagnostic{
			Range: typeDecl.Range,
			Kind:  &InvalidType{Name: typeDecl.Name},
		})
	}
}

func (res *CheckResult) checkDuplicateVars(variableName parser.VariableLiteral, decl parser.VarDeclaration) {
	// check there aren't duplicate variables
	if _, ok := res.declaredVars[variableName.Name]; ok {
		res.Diagnostics = append(res.Diagnostics, Diagnostic{
			Range: variableName.Range,
			Kind:  &DuplicateVariable{Name: variableName.Name},
		})
	} else {
		res.declaredVars[variableName.Name] = decl
		res.unusedVars[variableName.Name] = variableName.Range
	}
}

func (res *CheckResult) checkVarOrigin(fnCall parser.FnCallStatement, decl parser.VarDeclaration) {
	var sig []string
	if sigLookup, ok := OriginFunctionArities[fnCall.Caller.Name]; ok {
		sig = sigLookup.Args
		res.assertHasType(decl.Name, sigLookup.Return, decl.Type.Name)
	}
	res.checkFnCallArity(&fnCall, sig)
}

func (res *CheckResult) checkLiteral(lit parser.Literal, requiredType string) {
	switch lit := lit.(type) {
	case *parser.VariableLiteral:
		if varDeclaration, ok := res.declaredVars[lit.Name]; ok {
			res.varResolution[lit] = varDeclaration
		} else {
			res.Diagnostics = append(res.Diagnostics, Diagnostic{
				Range: lit.Range,
				Kind:  &UnboundVariable{Name: lit.Name},
			})
		}
		delete(res.unusedVars, lit.Name)

		resolved := res.ResolveVar(lit)
		if resolved == nil || resolved.Type == nil || !isTypeAllowed(resolved.Type.Name) {
			return
		}
		res.assertHasType(lit, requiredType, resolved.Type.Name)

	case *parser.MonetaryLiteral:
		res.assertHasType(lit, requiredType, TypeMonetary)
		res.checkLiteral(lit.Asset, TypeAsset)
		res.checkLiteral(lit.Amount, TypeNumber)
	case *parser.AccountLiteral:
		res.assertHasType(lit, requiredType, TypeAccount)
	case *parser.RatioLiteral:
		res.assertHasType(lit, requiredType, TypePortion)
	case *parser.AssetLiteral:
		res.assertHasType(lit, requiredType, TypeAsset)
	case *parser.NumberLiteral:
		res.assertHasType(lit, requiredType, TypeNumber)
	case *parser.StringLiteral:
		res.assertHasType(lit, requiredType, TypeString)
	}
}

func (res *CheckResult) assertHasType(lit parser.Literal, requiredType string, actualType string) {
	if requiredType == "*" || requiredType == actualType {
		return
	}

	res.Diagnostics = append(res.Diagnostics, Diagnostic{
		Range: lit.GetRange(),
		Kind: &TypeMismatch{
			Expected: requiredType,
			Got:      actualType,
		},
	})

}

func (res *CheckResult) checkSource(source parser.Source) {
	if source == nil {
		return
	}

	switch source := source.(type) {
	case *parser.AccountLiteral:

	case *parser.VariableLiteral:
		res.checkLiteral(source, TypeAccount)

	case *parser.SourceOverdraft:
		res.checkLiteral(source.Address, TypeAccount)
		if source.Bounded != nil {
			res.checkLiteral(*source.Bounded, TypeMonetary)
		}

	case *parser.SourceInorder:
		for _, source := range source.Sources {
			res.checkSource(source)
		}

	case *parser.SourceCapped:
		res.checkLiteral(source.Cap, TypeMonetary)
		res.checkSource(source.From)

	case *parser.SourceAllotment:
		var remainingAllotment *parser.RemainingAllotment = nil
		var variableLiterals []parser.VariableLiteral

		sum := big.NewRat(0, 1)
		for i, allottedItem := range source.Items {
			isLast := i == len(source.Items)-1

			switch allotment := allottedItem.Allotment.(type) {
			case *parser.VariableLiteral:
				variableLiterals = append(variableLiterals, *allotment)
				res.checkLiteral(allotment, TypePortion)
			case *parser.RatioLiteral:
				sum.Add(sum, allotment.ToRatio())
			case *parser.RemainingAllotment:
				if isLast {
					remainingAllotment = allotment
				} else {
					res.Diagnostics = append(res.Diagnostics, Diagnostic{
						Range: source.Range,
						Kind:  &RemainingIsNotLast{},
					})
				}
			}

			res.checkSource(allottedItem.From)
		}

		res.checkHasBadAllotmentSum(*sum, source.Range, remainingAllotment, variableLiterals)

	default:
		panic(fmt.Sprintf("unhandled clause: %+s", source))
	}
}

func (res *CheckResult) checkDestination(destination parser.Destination) {
	if destination == nil {
		return
	}

	switch destination := destination.(type) {
	case *parser.VariableLiteral:
		res.checkLiteral(destination, TypeAccount)

	case *parser.DestinationInorder:
		for _, clause := range destination.Clauses {
			res.checkLiteral(clause.Cap, "monetary")
			res.checkDestinationInorderTarget(clause.To)
		}
		res.checkDestinationInorderTarget(destination.Remaining)

	case *parser.DestinationAllotment:
		var remainingAllotment *parser.RemainingAllotment
		var variableLiterals []parser.VariableLiteral
		sum := big.NewRat(0, 1)

		for i, allottedItem := range destination.Items {
			isLast := i == len(destination.Items)-1

			switch allotment := allottedItem.Allotment.(type) {
			case *parser.VariableLiteral:
				variableLiterals = append(variableLiterals, *allotment)
				res.checkLiteral(allotment, TypePortion)
			case *parser.RatioLiteral:
				sum.Add(sum, allotment.ToRatio())
			case *parser.RemainingAllotment:
				if isLast {
					remainingAllotment = allotment
				} else {
					res.Diagnostics = append(res.Diagnostics, Diagnostic{
						Range: destination.Range,
						Kind:  &RemainingIsNotLast{},
					})
				}
			}

			res.checkDestination(allottedItem.To)
		}

		res.checkHasBadAllotmentSum(*sum, destination.Range, remainingAllotment, variableLiterals)
	}
}

func (res *CheckResult) checkDestinationInorderTarget(target parser.DestinationInorderTarget) {
	switch target := target.(type) {
	case *parser.DestinationTo:
		res.checkDestination(target.Destination)
	case *parser.DestinationKept:
		// nothing to do
	}
}

func (res *CheckResult) checkHasBadAllotmentSum(
	sum big.Rat,
	rng parser.Range,
	remaining *parser.RemainingAllotment,
	variableLiterals []parser.VariableLiteral,
) {
	cmp := sum.Cmp(big.NewRat(1, 1))
	switch cmp {
	case 1, -1:
		if (cmp == -1 && remaining != nil) || (cmp == -1 && len(variableLiterals) > 1) {
			return
		}

		if cmp == -1 && len(variableLiterals) == 1 {
			var value big.Rat
			res.Diagnostics = append(res.Diagnostics, Diagnostic{
				Range: variableLiterals[0].Range,
				Kind: &FixedPortionVariable{
					Value: *value.Sub(big.NewRat(1, 1), &sum),
				},
			})
		} else {
			res.Diagnostics = append(res.Diagnostics, Diagnostic{
				Range: rng,
				Kind: &BadAllotmentSum{
					Sum: sum,
				},
			})
		}

	// sum == 1
	case 0:
		for _, varLit := range variableLiterals {
			res.Diagnostics = append(res.Diagnostics, Diagnostic{
				Range: varLit.Range,
				Kind: &FixedPortionVariable{
					Value: *big.NewRat(0, 1),
				},
			})
		}
		if remaining != nil {
			res.Diagnostics = append(res.Diagnostics, Diagnostic{
				Range: remaining.Range,
				Kind:  &RedundantRemaining{},
			})
		}
	}
}
