// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JsonTemplate

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by JsonTemplate.
type JsonTemplateVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonTemplate#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by JsonTemplate#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by JsonTemplate#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by JsonTemplate#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by JsonTemplate#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by JsonTemplate#function_param.
	VisitFunction_param(ctx *Function_paramContext) interface{}

	// Visit a parse tree produced by JsonTemplate#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by JsonTemplate#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by JsonTemplate#spread_field.
	VisitSpread_field(ctx *Spread_fieldContext) interface{}

	// Visit a parse tree produced by JsonTemplate#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by JsonTemplate#object_field.
	VisitObject_field(ctx *Object_fieldContext) interface{}

	// Visit a parse tree produced by JsonTemplate#template_string.
	VisitTemplate_string(ctx *Template_stringContext) interface{}

	// Visit a parse tree produced by JsonTemplate#template_string_part.
	VisitTemplate_string_part(ctx *Template_string_partContext) interface{}

	// Visit a parse tree produced by JsonTemplate#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by JsonTemplate#index.
	VisitIndex(ctx *IndexContext) interface{}
}
