// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Neon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// NeonListener is a complete listener for a parse tree produced by NeonParser.
type NeonListener interface {
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

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterBool is called when entering the Bool production.
	EnterBool(c *BoolContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterComparison is called when entering the Comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterMDM is called when entering the MDM production.
	EnterMDM(c *MDMContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterNoClue is called when entering the NoClue production.
	EnterNoClue(c *NoClueContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

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

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitBool is called when exiting the Bool production.
	ExitBool(c *BoolContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitComparison is called when exiting the Comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitMDM is called when exiting the MDM production.
	ExitMDM(c *MDMContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitNoClue is called when exiting the NoClue production.
	ExitNoClue(c *NoClueContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)
}
