package tests

import (
	"reflect"
	"testing"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/value"
)

func TestLess(t *testing.T) {
	tests := []struct {
		name  string
		left  value.Value
		right value.Value
		want  value.Value
	}{
		{"Num G", value.NewNumber(7), value.NewNumber(4), value.False},
		{"Num E", value.NewNumber(420), value.NewNumber(420), value.False},
		{"Num L", value.NewNumber(69), value.NewNumber(420), value.True},
		{"Str GS", value.NewString("poggers"), value.NewString("asdf"), value.False},
		{"Str GE", value.NewString("poggers"), value.NewString("aoggers"), value.False},
		{"Str GL", value.NewString("poggers"), value.NewString("asdfasfjaosfjsaf"), value.False},
		{"Str EE", value.NewString("poggers"), value.NewString("poggers"), value.False},
		{"Str LS", value.NewString("asdfg"), value.NewString("xsdf"), value.True},
		{"Str LE", value.NewString("pog"), value.NewString("xog"), value.True},
		{"Str LL", value.NewString("pog"), value.NewString("poggers"), value.True},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operations.Less(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s Less() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
