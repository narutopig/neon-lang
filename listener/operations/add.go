package operations

import (
	"fmt"

	"github.com/narutopig/neon-lang/value"
)

func Add(left value.Value, right value.Value) value.Value {
	if left.Type == value.Number {
		lval := floatFromBytes(left.Data)

		if right.Type == value.Number {
			rval := floatFromBytes(right.Data)

			return value.NewNumber(lval + rval)
		} else if right.Type == value.String {
			return value.NewString(fmt.Sprintf("%f%s", lval, string(right.Data)))
		}
	} else if left.Type == value.String {
		lval := string(left.Data)

		if right.Type == value.Number {
			return value.NewString(fmt.Sprintf("%s%f", lval, floatFromBytes(right.Data)))
		} else if right.Type == value.String {
			return value.NewString(fmt.Sprintf("%s%s", lval, string(right.Data)))
		}
	}

	invalidOp(left.Type, right.Type, "+")

	return value.Value{}
}
