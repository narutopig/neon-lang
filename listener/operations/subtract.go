package operations

import (
	"github.com/narutopig/neon-lang/value"
)

func Subtract(left value.Value, right value.Value) (value.Value, string) {
	if left.Type == value.Number {
		lval := floatFromBytes(left.Data)

		if right.Type == value.Number {
			rval := floatFromBytes(right.Data)

			return value.NewNumber(lval - rval), ""
		}
	}
	return value.Value{}, invalidOp(left.Type, right.Type, "+")
}
