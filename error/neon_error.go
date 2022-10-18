package neon_error

import "fmt"

// Represents a basic error
type Error struct {
	Line      int
	ErrorType string
	Details   string
}

const (
	syntax      string = "SyntaxError"
	illegalchar string = "IllegalCharError"
	eol         string = "EndOfLineError"
	eof         string = "reached end of file while parsing string"
	none        string = ""
)

func (e Error) String() string {
	return fmt.Sprintf("%s: %s on line %d", e.ErrorType, e.Details, e.Line)
}

func (e Error) IsNone() bool {
	return e.ErrorType == none
}

func NewError(line int, errorType string, details string) Error {
	return Error{Line: line, ErrorType: errorType, Details: details}
}

func NoError() Error {
	return Error{Line: -1, ErrorType: none, Details: ""}
}

func SyntaxError(line int, details string) Error {
	return Error{Line: line, ErrorType: syntax, Details: details}
}

func IllegalCharError(line int, details string) Error {
	return Error{Line: line, ErrorType: illegalchar, Details: details}
}

func EndOfFileError(line int, details string) Error {
	return Error{Line: line, ErrorType: eof, Details: details}
}
func EndOfLineError(line int, details string) Error {
	return Error{Line: line, ErrorType: eol, Details: details}
}
