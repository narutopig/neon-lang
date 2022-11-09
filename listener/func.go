package listener

import (
	"fmt"

	"github.com/narutopig/neon-lang/parser"
)

// EnterFunccall is called when production funccall is entered.
func (i *Interpreter) EnterFunccall(ctx *parser.FunccallContext) {
	funcName := ctx.GetFunction().GetText()

	if !i.HasFunction(funcName) {
		i.panic(fmt.Sprintf("Function %s not found", funcName), ctx.GetStart().GetLine())
	}
}

// ExitFunccall is called when production funccall is exited.
func (i *Interpreter) ExitFunccall(ctx *parser.FunccallContext) {
	funcName := ctx.GetFunction().GetText()

	theFunc := i.GetFunction(funcName)

	if len(i.stack) != len(theFunc.Args) {
		i.panic(fmt.Sprintf("Missing arguments for function %s", funcName), ctx.GetStart().GetLine())
	} else {
		theFunc.Exec(i.stack)
	}
}
