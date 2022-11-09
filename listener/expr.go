package listener

import (
	"github.com/narutopig/neon-lang/listener/operations"
	"github.com/narutopig/neon-lang/parser"
)

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
