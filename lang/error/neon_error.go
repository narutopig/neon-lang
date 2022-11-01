package neonerror

import "fmt"

// Error represents a basic error
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

// IsNull checks if the error is a "null" error
func (e Error) IsNull() bool {
	return e.ErrorType == none
}

// New constructs a new Error
func New(line int, errorType string, details string) Error {
	return Error{Line: line, ErrorType: errorType, Details: details}
}

// Null is an nonexistent error
var Null Error = New(-1, none, "")

// SyntaxError is thrown when a error in the syntax, e.g. mismatched parentheses
func SyntaxError(line int, details string) Error {
	return New(line, syntax, details)
}

// IllegalCharError is thrown when an illegal character is encountered
func IllegalCharError(line int, details string) Error {
	return New(line, illegalchar, details)
}

// EndOfFileError is thrown when the EOF is reached while parsing a string
func EndOfFileError(line int, details string) Error {
	return New(line, eof, details)
}

// EndOfLineError is thrown when the EOL is reached while parsing a string
func EndOfLineError(line int, details string) Error {
	return New(line, eol, details)
}
