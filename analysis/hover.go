package analysis

import (
	"numscript/parser"
)

type Hover interface{ hover() }

type VariableHover struct {
	Range parser.Range
	Node  *parser.VariableLiteral
}

func (*VariableHover) hover() {}

func HoverOn(program parser.Program, position parser.Position) Hover {
	for _, statement := range program.Statements {
		// TODO binary search into statements
		if statement == nil || !statement.GetRange().Contains(position) {
			continue
		}

		switch statement := statement.(type) {
		case *parser.SendStatement:
			hover := hoverOnSendStatement(*statement, position)
			if hover != nil {
				return hover
			}

		case *parser.FnCallStatement:
			hover := hoverOnFnCall(*statement, position)
			if hover != nil {
				return hover
			}
		}

	}
	return nil
}

func hoverOnSendStatement(sendStatement parser.SendStatement, position parser.Position) Hover {
	if !sendStatement.Range.Contains(position) {
		return nil
	}

	hover := hoverOnLiteral(sendStatement.Monetary, position)
	if hover != nil {
		return hover
	}

	hover = hoverOnSource(sendStatement.Source, position)
	if hover != nil {
		return hover
	}

	hover = hoverOnDestination(sendStatement.Destination, position)
	if hover != nil {
		return hover
	}

	return nil
}

func hoverOnLiteral(lit parser.Literal, position parser.Position) Hover {
	if !lit.GetRange().Contains(position) {
		return nil
	}

	switch lit := lit.(type) {
	case *parser.VariableLiteral:
		return &VariableHover{
			Range: lit.Range,
			Node:  lit,
		}
	case *parser.MonetaryLiteral:
		hover := hoverOnLiteral(lit.Amount, position)
		if hover != nil {
			return hover
		}

		hover = hoverOnLiteral(lit.Asset, position)
		if hover != nil {
			return hover
		}

	}

	return nil
}

func hoverOnSource(source parser.Source, position parser.Position) Hover {
	if !source.GetRange().Contains(position) {
		return nil
	}

	switch source := source.(type) {
	case *parser.SourceCapped:
		hover := hoverOnLiteral(source.Cap, position)
		if hover != nil {
			return hover
		}
		hover = hoverOnSource(source.From, position)
		if hover != nil {
			return hover
		}
		return nil

	case *parser.SourceOverdraft:
		hover := hoverOnLiteral(source.Address, position)
		if hover != nil {
			return hover
		}

		if source.Bounded != nil {
			hover := hoverOnLiteral(*source.Bounded, position)
			if hover != nil {
				return hover
			}
		}

		return nil

	case *parser.SourceSeq:
		for _, source := range source.Sources {
			// TODO binary search
			if !source.GetRange().Contains(position) {
				continue
			}

			hover := hoverOnSource(source, position)
			if hover != nil {
				return hover
			}
		}

	case *parser.SourceAllotment:
		for _, item := range source.Items {
			// TODO binary search
			if !item.Range.Contains(position) {
				continue
			}

			switch allot := item.Allotment.(type) {
			case parser.Literal:
				hover := hoverOnLiteral(allot, position)
				if hover != nil {
					return hover
				}
			}

			hover := hoverOnSource(item.From, position)
			if hover != nil {
				return hover
			}
		}

	case *parser.VariableLiteral:
		return hoverOnLiteral(source, position)

	}

	return nil

}

func hoverOnDestination(destination parser.Destination, position parser.Position) Hover {
	if !destination.GetRange().Contains(position) {
		return nil
	}

	switch source := destination.(type) {
	case *parser.DestinationSeq:
		for _, destination := range source.Destinations {
			// TODO binary search
			if !destination.GetRange().Contains(position) {
				continue
			}

			hover := hoverOnDestination(destination, position)
			if hover != nil {
				return hover
			}
		}

	case *parser.DestinationAllotment:
		for _, item := range source.Items {
			// TODO binary search
			if !item.Range.Contains(position) {
				continue
			}

			switch allot := item.Allotment.(type) {
			case parser.Literal:
				hover := hoverOnLiteral(allot, position)
				if hover != nil {
					return hover
				}
			}

			hover := hoverOnDestination(item.To, position)
			if hover != nil {
				return hover
			}
		}

	case *parser.VariableLiteral:
		return hoverOnLiteral(source, position)
	}

	return nil

}

func hoverOnFnCall(callStatement parser.FnCallStatement, position parser.Position) Hover {
	if !callStatement.Range.Contains(position) {
		return nil
	}

	for _, arg := range callStatement.Args {
		hover := hoverOnLiteral(arg, position)
		if hover != nil {
			return hover
		}
	}

	return nil

}
