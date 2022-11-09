package function

import (
	"fmt"

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

var Print Function = Function{Void, []value.ValueType{value.String}, print}

func print(args []value.Value) value.Value {
	if len(args) == 0 {
		fmt.Println()
	}
	switch args[0].Type {
	case value.String:
		fmt.Println(string(args[0].Data))
	}
	return value.Value{}
}

var Std map[string]Function = map[string]Function{
	"print": Print,
}
