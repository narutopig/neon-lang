package tests

import (
	"reflect"
	"testing"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/value"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name  string
		left  value.Value
		right value.Value
		want  value.Value
		err   bool // should there be an error
	}{
		{"Number + string", value.NewNumber(69), value.NewString("hello world"), value.NewString("69.000000hello world"), false},
		{"Number + number", value.NewNumber(14), value.NewNumber(69), value.NewNumber(83), false},
		{"Number + boolean", value.NewNumber(17), value.False, value.Value{}, true},
		{"String + string", value.NewString("AMONGA"), value.NewString("US"), value.NewString("AMONGAUS"), false},
		{"String + number", value.NewString("AMONGA US"), value.NewNumber(3), value.NewString("AMONGA US3.000000"), false},
		{"String + boolean", value.NewString("AMONGA"), value.True, value.Value{}, true},
		{"Boolean + string", value.True, value.NewString("test"), value.Value{}, true},
		{"Boolean + number", value.True, value.NewNumber(3), value.Value{}, true},
		{"Boolean + boolean", value.True, value.True, value.Value{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := operations.Add(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) && ((err != "") == tt.err) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
