// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // neon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseneonListener is a complete listener for a parse tree produced by neonParser.
type BaseneonListener struct{}

var _ neonListener = &BaseneonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseneonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseneonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseneonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseneonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseneonListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseneonListener) ExitProgram(ctx *ProgramContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseneonListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseneonListener) ExitStat(ctx *StatContext) {}

// EnterIf is called when production if is entered.
func (s *BaseneonListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseneonListener) ExitIf(ctx *IfContext) {}

// EnterElif is called when production elif is entered.
func (s *BaseneonListener) EnterElif(ctx *ElifContext) {}

// ExitElif is called when production elif is exited.
func (s *BaseneonListener) ExitElif(ctx *ElifContext) {}

// EnterElse is called when production else is entered.
func (s *BaseneonListener) EnterElse(ctx *ElseContext) {}

// ExitElse is called when production else is exited.
func (s *BaseneonListener) ExitElse(ctx *ElseContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseneonListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseneonListener) ExitWhile(ctx *WhileContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseneonListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseneonListener) ExitFunc(ctx *FuncContext) {}

// EnterFunccall is called when production funccall is entered.
func (s *BaseneonListener) EnterFunccall(ctx *FunccallContext) {}

// ExitFunccall is called when production funccall is exited.
func (s *BaseneonListener) ExitFunccall(ctx *FunccallContext) {}

// EnterType is called when production type is entered.
func (s *BaseneonListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseneonListener) ExitType(ctx *TypeContext) {}

// EnterArithop is called when production arithop is entered.
func (s *BaseneonListener) EnterArithop(ctx *ArithopContext) {}

// ExitArithop is called when production arithop is exited.
func (s *BaseneonListener) ExitArithop(ctx *ArithopContext) {}

// EnterCompop is called when production compop is entered.
func (s *BaseneonListener) EnterCompop(ctx *CompopContext) {}

// ExitCompop is called when production compop is exited.
func (s *BaseneonListener) ExitCompop(ctx *CompopContext) {}

// EnterOp is called when production op is entered.
func (s *BaseneonListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseneonListener) ExitOp(ctx *OpContext) {}

// EnterFuncarg is called when production funcarg is entered.
func (s *BaseneonListener) EnterFuncarg(ctx *FuncargContext) {}

// ExitFuncarg is called when production funcarg is exited.
func (s *BaseneonListener) ExitFuncarg(ctx *FuncargContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseneonListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseneonListener) ExitDecl(ctx *DeclContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseneonListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseneonListener) ExitAssign(ctx *AssignContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseneonListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseneonListener) ExitReturn(ctx *ReturnContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseneonListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseneonListener) ExitExpr(ctx *ExprContext) {}
