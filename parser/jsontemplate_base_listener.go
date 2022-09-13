// Code generated from ../grammar/JsonTemplate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JsonTemplate

import "github.com/antlr/antlr4/runtime/Go/antlr"

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