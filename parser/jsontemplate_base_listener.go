// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JsonTemplate

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseJsonTemplateListener is a complete listener for a parse tree produced by JsonTemplateParser.
type BaseJsonTemplateListener struct{}

var _ JsonTemplateListener = &BaseJsonTemplateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonTemplateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonTemplateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonTemplateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonTemplateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseJsonTemplateListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseJsonTemplateListener) ExitScript(ctx *ScriptContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJsonTemplateListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJsonTemplateListener) ExitStatement(ctx *StatementContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseJsonTemplateListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseJsonTemplateListener) ExitStatements(ctx *StatementsContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJsonTemplateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJsonTemplateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BaseJsonTemplateListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BaseJsonTemplateListener) ExitLambda(ctx *LambdaContext) {}

// EnterFunction_param is called when production function_param is entered.
func (s *BaseJsonTemplateListener) EnterFunction_param(ctx *Function_paramContext) {}

// ExitFunction_param is called when production function_param is exited.
func (s *BaseJsonTemplateListener) ExitFunction_param(ctx *Function_paramContext) {}

// EnterField is called when production field is entered.
func (s *BaseJsonTemplateListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseJsonTemplateListener) ExitField(ctx *FieldContext) {}

// EnterArray is called when production array is entered.
func (s *BaseJsonTemplateListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseJsonTemplateListener) ExitArray(ctx *ArrayContext) {}

// EnterSpread_field is called when production spread_field is entered.
func (s *BaseJsonTemplateListener) EnterSpread_field(ctx *Spread_fieldContext) {}

// ExitSpread_field is called when production spread_field is exited.
func (s *BaseJsonTemplateListener) ExitSpread_field(ctx *Spread_fieldContext) {}

// EnterObject is called when production object is entered.
func (s *BaseJsonTemplateListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseJsonTemplateListener) ExitObject(ctx *ObjectContext) {}

// EnterObject_field is called when production object_field is entered.
func (s *BaseJsonTemplateListener) EnterObject_field(ctx *Object_fieldContext) {}

// ExitObject_field is called when production object_field is exited.
func (s *BaseJsonTemplateListener) ExitObject_field(ctx *Object_fieldContext) {}

// EnterName is called when production name is entered.
func (s *BaseJsonTemplateListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseJsonTemplateListener) ExitName(ctx *NameContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseJsonTemplateListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseJsonTemplateListener) ExitIndex(ctx *IndexContext) {}
