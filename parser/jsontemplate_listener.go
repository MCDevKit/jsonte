// Code generated from ../grammar/JsonTemplate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JsonTemplate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JsonTemplateListener is a complete listener for a parse tree produced by JsonTemplateParser.
type JsonTemplateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLambda is called when entering the lambda production.
	EnterLambda(c *LambdaContext)

	// EnterFunction_param is called when entering the function_param production.
	EnterFunction_param(c *Function_paramContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterObject_field is called when entering the object_field production.
	EnterObject_field(c *Object_fieldContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLambda is called when exiting the lambda production.
	ExitLambda(c *LambdaContext)

	// ExitFunction_param is called when exiting the function_param production.
	ExitFunction_param(c *Function_paramContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitObject_field is called when exiting the object_field production.
	ExitObject_field(c *Object_fieldContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)
}
