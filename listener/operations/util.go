package operations

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/narutopig/neon-lang/value"
)

func invalidOp(left, right value.ValueType, op string) string {
	return fmt.Sprintf("Invalid operation %s between variable of type %s and %s", op, left, right)
}

func floatFromBytes(bytes []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(bytes))
}
