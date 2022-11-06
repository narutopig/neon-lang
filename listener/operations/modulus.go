package operations

import (
	"math"

	"github.com/narutopig/neon-lang/value"
)

func Modulus(left value.Value, right value.Value) value.Value {
	if left.Type == value.Number {
		lval := floatFromBytes(left.Data)

		if right.Type == value.Number {
			rval := floatFromBytes(right.Data)

			res := math.Mod(lval, rval)
			return value.NewNumber(res)
		}
	}

	invalidOp(left.Type, right.Type, "%")

	return value.Value{}
}
