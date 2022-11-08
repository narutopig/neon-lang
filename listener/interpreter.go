package listener

import (
	"fmt"
	"strconv"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/listener/runtime"
	"github.com/narutopig/neon-lang/parser"
	"github.com/narutopig/neon-lang/util"
	"github.com/narutopig/neon-lang/value"
)

// Interpreter implements BaseNeonListener
type Interpreter struct {
	*parser.BaseNeonListener
	Memory runtime.Memory
	stack  []value.Value
}

// NewInterpreter returns a new interpreter instance
func NewInterpreter() *Interpreter {
	return &Interpreter{Memory: *runtime.NewMemory()}
}

func (i *Interpreter) push(val value.Value) {
	i.stack = append(i.stack, val)
}

func (i *Interpreter) pop() value.Value {
	if len(i.stack) < 1 {
		panic("stack is empty unable to pop")
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
		panic(fmt.Sprintf("Variable %s already exists", varname))
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
		panic(fmt.Sprintf("Variable %s already exists", varname))
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
		panic(fmt.Sprintf("Variable %s already exists", varname))
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
		i.push(operations.Add(left, right))
	case "-":
		i.push(operations.Subtract(left, right))
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
		i.push(operations.Multiply(left, right))
	} else if op == "/" {
		i.push(operations.Divide(left, right))
	} else if op == "%" {
		i.push(operations.Modulus(left, right))
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
	util.Assert(err == nil, "Error while converting number to string")
	i.stack = append(i.stack, value.NewNumber(val))
}

// EnterBool is called when production Bool is entered.
func (i *Interpreter) EnterBool(ctx *parser.BoolContext) {
	val, err := strconv.ParseBool(ctx.GetText())
	util.Assert(err == nil, "Error while converting boolean to string")
	i.stack = append(i.stack, value.NewBoolean(val))
}

func (i Interpreter) Print() {
	fmt.Println(i.stack)
}
