package tests

import (
	"reflect"
	"testing"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/value"
)

func TestSubtract(t *testing.T) {
	tests := []struct {
		name  string
		left  value.Value
		right value.Value
		want  value.Value
		err   bool
	}{
		{"Number - number", value.NewNumber(14), value.NewNumber(69), value.NewNumber(-55), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operations.Subtract(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
