package listener

import (
	"fmt"

	"github.com/narutopig/neon-lang/listener/runtime"
	"github.com/narutopig/neon-lang/listener/runtime/function"
	"github.com/narutopig/neon-lang/parser"
	"github.com/narutopig/neon-lang/value"
)

// Interpreter implements BaseNeonListener
type Interpreter struct {
	*parser.BaseNeonListener
	Memory runtime.Memory
	stack  []value.Value
	errors []runtime.Error
	funcs  map[string]function.Function
}

// NewInterpreter returns a new interpreter instance
func NewInterpreter() *Interpreter {
	return &Interpreter{Memory: runtime.M(), funcs: function.Std}
}

func (i *Interpreter) panic(message string, line int) {
	i.errors = append(i.errors, runtime.E(message))
}

func (i *Interpreter) push(val value.Value) {
	i.stack = append(i.stack, val)
}

func (i *Interpreter) pop() value.Value {
	if len(i.stack) < 1 {
		i.panic("stack empty, unable to pop", -1)
	}

	// Get the last value from the stack.
	result := i.stack[len(i.stack)-1]

	// Remove the last element from the stack.
	i.stack = i.stack[:len(i.stack)-1]

	return result
}

// Print should only be used for debug and prints out the current stack
func (i Interpreter) Print() {
	fmt.Println(i.stack)
}

// Errors returns the list of compile/runtime errors
func (i Interpreter) Errors() []runtime.Error {
	return i.errors
}

func (i Interpreter) GetFunction(name string) function.Function {
	return i.funcs[name]
}

func (i Interpreter) HasFunction(name string) bool {
	_, exists := i.funcs[name]

	return exists
}

func (i *Interpreter) AddFunction(name string, function function.Function) {
	i.funcs[name] = function
}
