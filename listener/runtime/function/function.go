package function

import (
	"github.com/narutopig/neon-lang/value"
)

type Function struct {
	Type ReturnType
	Args []value.ValueType
	Exec ExecFunc
}

type (
	ExecFunc   func(args []value.Value) value.Value
	ReturnType byte
)

const (
	String ReturnType = iota
	Number
	Bool
	Void
)

var Std map[string]Function = map[string]Function{
	"print":   Print,
	"println": Println,
}
