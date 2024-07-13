package analysis

import "fmt"

type Severity = byte

// !important! keep in sync with LSP specs
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnosticSeverity
const (
	_ Severity = iota
	ErrorSeverity
	WarningSeverity
	Information
	Hint
)

type DiagnosticKind interface {
	Message() string
	Severity() Severity
}

// ###### Diagnostics

type InvalidType struct {
	Name string
}

// TODO evaluate suggestion using Levenshtein distance
func (e *InvalidType) Message() string {
	allowedTypeList := ""
	for index, t := range AllowedTypes {
		if index != 0 {
			allowedTypeList += ", "
		}
		allowedTypeList += t
	}

	return fmt.Sprintf("'%s' is not a valid type. Allowed types are: %s", e.Name, allowedTypeList)
}

func (*InvalidType) Severity() Severity {
	return ErrorSeverity
}

type DuplicateVariable struct {
	Name string
}

func (e *DuplicateVariable) Message() string {
	return fmt.Sprintf("A variable with the name '$%s' was already declared", e.Name)
}

func (*DuplicateVariable) Severity() Severity {
	return ErrorSeverity
}

type UnboundVariable struct {
	Name string
}

// TODO evaluate suggestion using Levenshtein distance
func (e *UnboundVariable) Message() string {
	return fmt.Sprintf("The variable '$%s' was not declared", e.Name)
}

func (*UnboundVariable) Severity() Severity {
	return ErrorSeverity
}

type UnusedVar struct {
	Name string
}

func (e *UnusedVar) Message() string {
	return fmt.Sprintf("The variable '$%s' is never used", e.Name)
}

func (*UnusedVar) Severity() Severity {
	return WarningSeverity
}

type TypeMismatch struct {
	Expected string
	Got      string
}

func (e *TypeMismatch) Message() string {
	return fmt.Sprintf("Type mismatch (expected '%s', got '%s' instead)", e.Expected, e.Got)
}

func (*TypeMismatch) Severity() Severity {
	return ErrorSeverity
}
