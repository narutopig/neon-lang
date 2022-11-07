package tests

import (
	"reflect"
	"testing"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/value"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name  string
		left  value.Value
		right value.Value
		want  value.Value
	}{
		// G = greater, E = equal, L = less, S = shorter, L (in second) = longer, NE = not equal
		{"Num G", value.NewNumber(7), value.NewNumber(4), value.False},
		{"Num E", value.NewNumber(420), value.NewNumber(420), value.True},
		{"Num L", value.NewNumber(69), value.NewNumber(420), value.False},
		{"Str GS", value.NewString("poggers"), value.NewString("asdf"), value.False},
		{"Str GE", value.NewString("poggers"), value.NewString("aoggers"), value.False},
		{"Str GL", value.NewString("poggers"), value.NewString("asdfasfjaosfjsaf"), value.False},
		{"Str EE", value.NewString("poggers"), value.NewString("poggers"), value.True},
		{"Str LS", value.NewString("asdfg"), value.NewString("xsdf"), value.False},
		{"Str LE", value.NewString("pog"), value.NewString("xog"), value.False},
		{"Str LL", value.NewString("pog"), value.NewString("poggers"), value.False},
		{"Bool E 1", value.True, value.True, value.True},
		{"Bool E 2", value.False, value.False, value.True},
		{"Bool NE 1", value.True, value.False, value.False},
		{"Bool NE 2", value.False, value.True, value.False},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operations.Equal(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
