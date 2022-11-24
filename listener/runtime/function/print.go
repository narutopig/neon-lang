package function

import (
	"fmt"

	"github.com/narutopig/neon-lang/value"
)

var Print Function = Function{Void, []value.ValueType{value.String}, print}

func print(args []value.Value) value.Value {
	if len(args) == 0 {
		fmt.Print()
	}
	switch args[0].Type {
	case value.String:
		fmt.Print(string(args[0].Data))
	case value.Number:
		fmt.Print(value.FloatFromBytes(args[0].Data))
	case value.Boolean:
		{
			if args[0].Data[0] == 1 {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}
		}
	}
	return value.Value{}
}
