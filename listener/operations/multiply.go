package operations

import (
	"github.com/narutopig/neon-lang/value"
)

func Multiply(left value.Value, right value.Value) (value.Value, string) {
	if left.Type == value.Number {
		lval := floatFromBytes(left.Data)

		if right.Type == value.Number {
			rval := floatFromBytes(right.Data)

			return value.NewNumber(lval * rval), ""
		}
	} else if left.Type == value.String {
		lval := string(left.Data)

		if right.Type == value.Number {
			rval := floatFromBytes(right.Data)

			temp := ""
			for i := 0; i < int(rval); i++ {
				temp += lval
			}

			return value.NewString(temp), ""
		}
	}

	return value.Value{}, invalidOp(left.Type, right.Type, "*")
}
