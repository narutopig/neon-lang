package listener

import (
	"fmt"
	"strconv"

	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/parser"
	"github.com/narutopig/neon-lang/util"
	"github.com/narutopig/neon-lang/value"
)

// Interpreter implements BaseNeonListener
type Interpreter struct {
	*parser.BaseNeonListener
	stack []value.Value
}

// NewInterpreter returns a new interpreter instance
func NewInterpreter() Interpreter {
	return Interpreter{}
}

func (l *Interpreter) push(val value.Value) {
	l.stack = append(l.stack, val)
}

func (l *Interpreter) pop() value.Value {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

// EnterIf is called when production if is entered.
func (l *Interpreter) EnterIf(ctx *parser.IfContext) {}

// ExitIf is called when production if is exited.
func (l *Interpreter) ExitIf(ctx *parser.IfContext) {}

// EnterElif is called when production elif is entered.
func (l *Interpreter) EnterElif(ctx *parser.ElifContext) {}

// ExitElif is called when production elif is exited.
func (l *Interpreter) ExitElif(ctx *parser.ElifContext) {}

// EnterElse is called when production else is entered.
func (l *Interpreter) EnterElse(ctx *parser.ElseContext) {}

// ExitElse is called when production else is exited.
func (l *Interpreter) ExitElse(ctx *parser.ElseContext) {}

// EnterWhile is called when production while is entered.
func (l *Interpreter) EnterWhile(ctx *parser.WhileContext) {}

// ExitWhile is called when production while is exited.
func (l *Interpreter) ExitWhile(ctx *parser.WhileContext) {}

// EnterFunc is called when production func is entered.
func (l *Interpreter) EnterFunc(ctx *parser.FuncContext) {}

// ExitFunc is called when production func is exited.
func (l *Interpreter) ExitFunc(ctx *parser.FuncContext) {}

// EnterFunccall is called when production funccall is entered.
func (l *Interpreter) EnterFunccall(ctx *parser.FunccallContext) {}

// ExitFunccall is called when production funccall is exited.
func (l *Interpreter) ExitFunccall(ctx *parser.FunccallContext) {}

// EnterFuncarg is called when production funcarg is entered.
func (l *Interpreter) EnterFuncarg(ctx *parser.FuncargContext) {}

// ExitFuncarg is called when production funcarg is exited.
func (l *Interpreter) ExitFuncarg(ctx *parser.FuncargContext) {}

// EnterDecl is called when production decl is entered.
func (l *Interpreter) EnterDecl(ctx *parser.DeclContext) {}

// ExitDecl is called when production decl is exited.
func (l *Interpreter) ExitDecl(ctx *parser.DeclContext) {}

// EnterAssign is called when production assign is entered.
func (l *Interpreter) EnterAssign(ctx *parser.AssignContext) {}

// ExitAssign is called when production assign is exited.
func (l *Interpreter) ExitAssign(ctx *parser.AssignContext) {}

// EnterReturn is called when production return is entered.
func (l *Interpreter) EnterReturn(ctx *parser.ReturnContext) {}

// ExitReturn is called when production return is exited.
func (l *Interpreter) ExitReturn(ctx *parser.ReturnContext) {}

// EnterExpr is called when production expr is entered.
func (l *Interpreter) EnterExpr(ctx *parser.ExprContext) {}

// ExitExpr is called when production expr is exited.
func (l *Interpreter) ExitExpr(ctx *parser.ExprContext) {}

// EnterType is called when production type is entered.
func (l *Interpreter) EnterType(ctx *parser.TypeContext) {}

// ExitType is called when production type is exited.
func (l *Interpreter) ExitType(ctx *parser.TypeContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (l *Interpreter) EnterFuncCall(ctx *parser.FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (l *Interpreter) ExitFuncCall(ctx *parser.FuncCallContext) {}

// ExitAddSub is called when production AddSub is exited.
func (l *Interpreter) ExitAddSub(ctx *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch ctx.GetOp().GetText() {
	case "+":
		l.push(operations.Add(left, right))
	case "-":
		l.push(operations.Subtract(left, right))
	}
}

// ExitComparison is called when production Comparison is exited.
func (l *Interpreter) ExitComparison(ctx *parser.ComparisonContext) {
	right, left := l.pop(), l.pop()

	op := ctx.GetOp().GetText()

	if op == "==" {
		l.push(operations.Equal(left, right))
	} else if op == "<=" {
	} else if op == ">=" {
	} else if op == "<" {
	} else if op == ">" {
	}
}

// ExitMDM is called when production MDM is exited.
func (l *Interpreter) ExitMDM(ctx *parser.MDMContext) {
	right, left := l.pop(), l.pop()

	op := ctx.GetOp().GetText()

	if op == "*" {
		l.push(operations.Multiply(left, right))
	} else if op == "/" {
		l.push(operations.Divide(left, right))
	} else if op == "%" {
		l.push(operations.Modulus(left, right))
	} else {
		panic(fmt.Sprintf("Unexpected operator %s", op))
	}
}

// ExitNotExpr is called when production NotExpr is exited.
func (l *Interpreter) ExitNotExpr(ctx *parser.NotExprContext) {
}

// EnterString is called when production String is entered.
func (l *Interpreter) EnterString(ctx *parser.StringContext) {
	l.stack = append(l.stack, value.NewString(ctx.GetText()))
}

// EnterIdentifier is called when production Identifier is entered.
func (l *Interpreter) EnterIdentifier(ctx *parser.IdentifierContext) {
	l.stack = append(l.stack, value.NewIdentifier(ctx.GetText()))
}

// EnterInt is called when production Int is entered.
func (l *Interpreter) EnterInt(ctx *parser.IntContext) {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	util.Assert(err != nil, "Error while converting number to string")
	l.stack = append(l.stack, value.NewNumber(val))
}

// EnterBool is called when production Bool is entered.
func (l *Interpreter) EnterBool(ctx *parser.BoolContext) {
	val, err := strconv.ParseBool(ctx.GetText())
	util.Assert(err != nil, "Error while converting boolean to string")
	l.stack = append(l.stack, value.NewBoolean(val))
}

/*
// EnterNoClue is called when production NoClue is entered.
func (l *Listener) EnterNoClue(ctx *parser.NoClueContext) {}

// ExitNoClue is called when production NoClue is exited.
func (l *Listener) ExitNoClue(ctx *parser.NoClueContext) {}

above is for ( expr )
*/
