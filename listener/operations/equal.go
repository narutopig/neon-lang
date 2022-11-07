package operations

import "github.com/narutopig/neon-lang/value"

func Equal(left value.Value, right value.Value) value.Value {
	length := len(left.Data)
	temp := len(right.Data)

	if length != temp {
		return value.False
	}

	for i := 0; i < length; i++ {
		if left.Data[i] != right.Data[i] {
			return value.False
		}
	}

	return value.True
}
