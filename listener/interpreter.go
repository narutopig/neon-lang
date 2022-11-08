package listener

import (
	"fmt"
	"strconv"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/listener/runtime"
	"github.com/narutopig/neon-lang/parser"
	"github.com/narutopig/neon-lang/value"
)

// Interpreter implements BaseNeonListener
type Interpreter struct {
	*parser.BaseNeonListener
	Memory runtime.Memory
	stack  []value.Value
	errors []runtime.Error
}

// NewInterpreter returns a new interpreter instance
func NewInterpreter() *Interpreter {
	return &Interpreter{Memory: runtime.M()}
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

// EnterNumVar is called when production NumVar is entered.
func (i *Interpreter) EnterNumVar(ctx *parser.NumVarContext) {
	varname := ctx.GetVar_().GetText()
	_, exists := i.Memory.Get(varname)

	if exists {
		i.panic(fmt.Sprintf("Variable %s already exists", varname), ctx.GetStart().GetLine())
	}
}

// ExitNumVar is called when production NumVar is exited.
func (i *Interpreter) ExitNumVar(ctx *parser.NumVarContext) {
	varname := ctx.GetVar_().GetText()

	i.Memory.Put(varname, i.pop())
}

// EnterStrVar is called when production StrVar is entered.
func (i *Interpreter) EnterStrVar(ctx *parser.StrVarContext) {
	varname := ctx.GetVar_().GetText()
	_, exists := i.Memory.Get(varname)

	if exists {
		i.panic(fmt.Sprintf("Variable %s already exists", varname), ctx.GetStart().GetLine())
	}
}

// ExitStrVar is called when production StrVar is exited.
func (i *Interpreter) ExitStrVar(ctx *parser.StrVarContext) {
	varname := ctx.GetVar_().GetText()

	i.Memory.Put(varname, i.pop())
}

// EnterBoolVar is called when production BoolVar is entered.
func (i *Interpreter) EnterBoolVar(ctx *parser.BoolVarContext) {
	varname := ctx.GetVar_().GetText()
	_, exists := i.Memory.Get(varname)

	if exists {
		i.panic(fmt.Sprintf("Variable %s already exists", varname), ctx.GetStart().GetLine())
	}
}

// ExitBoolVar is called when production BoolVar is exited.
func (i *Interpreter) ExitBoolVar(ctx *parser.BoolVarContext) {
	varname := ctx.GetVar_().GetText()

	i.Memory.Put(varname, i.pop())
}

// ExitAddSub is called when production AddSub is exited.
func (i *Interpreter) ExitAddSub(ctx *parser.AddSubContext) {
	right, left := i.pop(), i.pop()

	switch ctx.GetOp().GetText() {
	case "+":
		val, err := operations.Add(left, right)
		if err != "" {
			i.panic(err, ctx.GetStart().GetLine())
		} else {
			i.push(val)
		}
	case "-":
		val, err := operations.Subtract(left, right)
		if err != "" {
			i.panic(err, ctx.GetStart().GetLine())
		} else {
			i.push(val)
		}
	}
}

// ExitComparison is called when production Comparison is exited.
func (i *Interpreter) ExitComparison(ctx *parser.ComparisonContext) {
	right, left := i.pop(), i.pop()

	op := ctx.GetOp().GetText()

	if op == "==" {
		i.push(operations.Equal(left, right))
	} else if op == "<=" {
		i.push(operations.LessEq(left, right))
	} else if op == ">=" {
		i.push(operations.GreaterEq(left, right))
	} else if op == "<" {
		i.push(operations.Less(left, right))
	} else if op == ">" {
		i.push(operations.Greater(left, right))
	}
}

// ExitMDM is called when production MDM is exited.
func (i *Interpreter) ExitMDM(ctx *parser.MDMContext) {
	right, left := i.pop(), i.pop()

	op := ctx.GetOp().GetText()

	if op == "*" {
		val, err := operations.Multiply(left, right)
		if err != "" {
			i.panic(err, ctx.GetStart().GetLine())
		} else {
			i.push(val)
		}
	} else if op == "/" {
		val, err := operations.Divide(left, right)
		if err != "" {
			i.panic(err, ctx.GetStart().GetLine())
		} else {
			i.push(val)
		}
	} else if op == "%" {
		val, err := operations.Modulus(left, right)
		if err != "" {
			i.panic(err, ctx.GetStart().GetLine())
		} else {
			i.push(val)
		}
	}
}

// ExitNotExpr is called when production NotExpr is exited.
func (i *Interpreter) ExitNotExpr(ctx *parser.NotExprContext) {
}

// EnterString is called when production String is entered.
func (i *Interpreter) EnterString(ctx *parser.StringContext) {
	i.stack = append(i.stack, value.NewString(ctx.GetText()))
}

// EnterIdentifier is called when production Identifier is entered.
func (i *Interpreter) EnterIdentifier(ctx *parser.IdentifierContext) {
	i.stack = append(i.stack, value.NewIdentifier(ctx.GetText()))
}

// EnterInt is called when production Int is entered.
func (i *Interpreter) EnterInt(ctx *parser.IntContext) {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		i.panic("Error while converting number to string", ctx.GetStart().GetLine())
	}
	i.stack = append(i.stack, value.NewNumber(val))
}

// EnterBool is called when production Bool is entered.
func (i *Interpreter) EnterBool(ctx *parser.BoolContext) {
	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		i.panic("Error while converting boolean to string", ctx.GetStart().GetLine())
	}
	i.stack = append(i.stack, value.NewBoolean(val))
}

// Print should only be used for debug and prints out the current stack
func (i Interpreter) Print() {
	fmt.Println(i.stack)
}

// Errors returns the list of compile/runtime errors
func (i Interpreter) Errors() []runtime.Error {
	return i.errors
}
