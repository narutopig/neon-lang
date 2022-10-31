package statement

import tokens "github.com/narutopig/neon-lang/token"

// SType is used to determine what the type of the Statement is
type SType int

const (
	DECLARATION SType = 0
)

// Executable adds a execute function to a Statement
type Executable interface {
	Execute()
}

// Statement represents a command
type Statement struct {
	Type SType
	Args []tokens.Token
	Executable
}
