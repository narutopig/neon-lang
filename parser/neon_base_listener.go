// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Neon

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseNeonListener is a complete listener for a parse tree produced by NeonParser.
type BaseNeonListener struct{}

var _ NeonListener = &BaseNeonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNeonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNeonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNeonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNeonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseNeonListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseNeonListener) ExitProgram(ctx *ProgramContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseNeonListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseNeonListener) ExitStat(ctx *StatContext) {}

// EnterIf is called when production if is entered.
func (s *BaseNeonListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseNeonListener) ExitIf(ctx *IfContext) {}

// EnterElif is called when production elif is entered.
func (s *BaseNeonListener) EnterElif(ctx *ElifContext) {}

// ExitElif is called when production elif is exited.
func (s *BaseNeonListener) ExitElif(ctx *ElifContext) {}

// EnterElse is called when production else is entered.
func (s *BaseNeonListener) EnterElse(ctx *ElseContext) {}

// ExitElse is called when production else is exited.
func (s *BaseNeonListener) ExitElse(ctx *ElseContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseNeonListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseNeonListener) ExitWhile(ctx *WhileContext) {}

// EnterFunc is called when production func is entered.
func (s *BaseNeonListener) EnterFunc(ctx *FuncContext) {}

// ExitFunc is called when production func is exited.
func (s *BaseNeonListener) ExitFunc(ctx *FuncContext) {}

// EnterFunccall is called when production funccall is entered.
func (s *BaseNeonListener) EnterFunccall(ctx *FunccallContext) {}

// ExitFunccall is called when production funccall is exited.
func (s *BaseNeonListener) ExitFunccall(ctx *FunccallContext) {}

// EnterFuncarg is called when production funcarg is entered.
func (s *BaseNeonListener) EnterFuncarg(ctx *FuncargContext) {}

// ExitFuncarg is called when production funcarg is exited.
func (s *BaseNeonListener) ExitFuncarg(ctx *FuncargContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseNeonListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseNeonListener) ExitDecl(ctx *DeclContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseNeonListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseNeonListener) ExitAssign(ctx *AssignContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseNeonListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseNeonListener) ExitReturn(ctx *ReturnContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (s *BaseNeonListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (s *BaseNeonListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseNeonListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseNeonListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterBool is called when production Bool is entered.
func (s *BaseNeonListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production Bool is exited.
func (s *BaseNeonListener) ExitBool(ctx *BoolContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseNeonListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseNeonListener) ExitAddSub(ctx *AddSubContext) {}

// EnterComparison is called when production Comparison is entered.
func (s *BaseNeonListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production Comparison is exited.
func (s *BaseNeonListener) ExitComparison(ctx *ComparisonContext) {}

// EnterMDM is called when production MDM is entered.
func (s *BaseNeonListener) EnterMDM(ctx *MDMContext) {}

// ExitMDM is called when production MDM is exited.
func (s *BaseNeonListener) ExitMDM(ctx *MDMContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseNeonListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseNeonListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterString is called when production String is entered.
func (s *BaseNeonListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseNeonListener) ExitString(ctx *StringContext) {}

// EnterNoClue is called when production NoClue is entered.
func (s *BaseNeonListener) EnterNoClue(ctx *NoClueContext) {}

// ExitNoClue is called when production NoClue is exited.
func (s *BaseNeonListener) ExitNoClue(ctx *NoClueContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseNeonListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseNeonListener) ExitInt(ctx *IntContext) {}

// EnterType is called when production type is entered.
func (s *BaseNeonListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseNeonListener) ExitType(ctx *TypeContext) {}
