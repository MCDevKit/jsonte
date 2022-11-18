// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JsonTemplate

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by JsonTemplateParser.
type JsonTemplateVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonTemplateParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#function_param.
	VisitFunction_param(ctx *Function_paramContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#object_field.
	VisitObject_field(ctx *Object_fieldContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by JsonTemplateParser#index.
	VisitIndex(ctx *IndexContext) interface{}
}
