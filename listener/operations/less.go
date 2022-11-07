package operations

import "github.com/narutopig/neon-lang/value"

func Less(left value.Value, right value.Value) value.Value {
	len1 := len(left.Data)
	len2 := len(right.Data)

	l := len1

	if l > len2 {
		l = len2
	}

	for i := 0; i < l; i++ {
		if left.Data[i] < right.Data[i] {
			return value.True
		} else if left.Data[i] > right.Data[i] {
			return value.False
		}
	}

	// values are equal
	if len1 < len2 {
		return value.True
	}
	return value.False // "abcd" "abc"
}
