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

// EnterFunc is called when production func is entered.
func (s *BaseneonListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseneonListener) ExitFunc(ctx *FuncContext) {}

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

// EnterExpr is called when production expr is entered.
func (s *BaseneonListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseneonListener) ExitExpr(ctx *ExprContext) {}
