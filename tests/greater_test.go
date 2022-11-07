package tests

import (
	"reflect"
	"testing"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/value"
)

func TestGreater(t *testing.T) {
	tests := []struct {
		name  string
		left  value.Value
		right value.Value
		want  value.Value
	}{
		{"Num G", value.NewNumber(7), value.NewNumber(4), value.True},
		{"Num E", value.NewNumber(420), value.NewNumber(420), value.False},
		{"Num L", value.NewNumber(69), value.NewNumber(420), value.False},
		{"Str GS", value.NewString("poggers"), value.NewString("asdf"), value.True},
		{"Str GE", value.NewString("poggers"), value.NewString("aoggers"), value.True},
		{"Str GL", value.NewString("poggers"), value.NewString("asdfasfjaosfjsaf"), value.True},
		{"Str EE", value.NewString("poggers"), value.NewString("poggers"), value.False},
		{"Str LS", value.NewString("asdfg"), value.NewString("xsdf"), value.False},
		{"Str LE", value.NewString("pog"), value.NewString("xog"), value.False},
		{"Str LL", value.NewString("pog"), value.NewString("poggers"), value.False},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := operations.Greater(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Greater() = %v, want %v", got, tt.want)
			}
		})
	}
}
