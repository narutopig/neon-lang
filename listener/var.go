package listener

import (
	"fmt"

	"github.com/narutopig/neon-lang/parser"
)

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
