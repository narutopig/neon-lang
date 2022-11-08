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
	fmt.Println(args[0])
	return value.Value{}
}
