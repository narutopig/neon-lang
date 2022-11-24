package function

import (
	"fmt"

	"github.com/narutopig/neon-lang/value"
)

var Println Function = Function{Void, []value.ValueType{value.String}, println}

func println(args []value.Value) value.Value {
	print(args)
	fmt.Print("\n")
	return value.Value{}
}
