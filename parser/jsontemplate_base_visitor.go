// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JsonTemplate

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseJsonTemplateVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonTemplateVisitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitFunction_param(ctx *Function_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitSpread_field(ctx *Spread_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitObject_field(ctx *Object_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonTemplateVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}
