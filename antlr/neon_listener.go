// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // neon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// neonListener is a complete listener for a parse tree produced by neonParser.
type neonListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterElif is called when entering the elif production.
	EnterElif(c *ElifContext)

	// EnterElse is called when entering the else production.
	EnterElse(c *ElseContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// EnterFunc is called when entering the func production.
	EnterFunc(c *FuncContext)

	// EnterFunccall is called when entering the funccall production.
	EnterFunccall(c *FunccallContext)

	// EnterFuncarg is called when entering the funcarg production.
	EnterFuncarg(c *FuncargContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterReturn is called when entering the return production.
	EnterReturn(c *ReturnContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterArithop is called when entering the arithop production.
	EnterArithop(c *ArithopContext)

	// EnterCompop is called when entering the compop production.
	EnterCompop(c *CompopContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitElif is called when exiting the elif production.
	ExitElif(c *ElifContext)

	// ExitElse is called when exiting the else production.
	ExitElse(c *ElseContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)

	// ExitFunc is called when exiting the func production.
	ExitFunc(c *FuncContext)

	// ExitFunccall is called when exiting the funccall production.
	ExitFunccall(c *FunccallContext)

	// ExitFuncarg is called when exiting the funcarg production.
	ExitFuncarg(c *FuncargContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitReturn is called when exiting the return production.
	ExitReturn(c *ReturnContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitArithop is called when exiting the arithop production.
	ExitArithop(c *ArithopContext)

	// ExitCompop is called when exiting the compop production.
	ExitCompop(c *CompopContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)
}
