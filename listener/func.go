package listener

import "github.com/narutopig/neon-lang/parser"

// EnterFunccall is called when production funccall is entered.
func (i *Interpreter) EnterFunccall(ctx *parser.FunccallContext) {}

// ExitFunccall is called when production funccall is exited.
func (i *Interpreter) ExitFunccall(ctx *parser.FunccallContext) {}
