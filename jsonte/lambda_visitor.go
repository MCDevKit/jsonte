package jsonte

import (
	"github.com/MCDevKit/jsonte/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type LambdaVisitor struct {
	parser.BaseJsonTemplateVisitor
	usedVariables []string
	arguments     []string
}

func (v *LambdaVisitor) Visit(tree antlr.ParseTree) {
	switch val := tree.(type) {
	case *parser.FieldContext:
		v.VisitField(val)
		break
	case *parser.ArrayContext:
		v.VisitArray(val)
		break
	case *parser.ObjectContext:
		v.VisitObject(val)
		break
	case *parser.Object_fieldContext:
		v.VisitObject_field(val)
		break
	case *parser.ExpressionContext:
		v.VisitExpression(val)
		break
	case *parser.Function_paramContext:
		v.VisitFunction_param(val)
		break
	case *parser.LambdaContext:
		v.VisitLambda(val)
		break
	case *parser.NameContext:
		v.VisitName(val)
		break
	case *parser.IndexContext:
		v.VisitIndex(val)
		break
	}
}

// resolveLambdaTree resolves a string to an AST tree
func (v *LambdaVisitor) resolveLambdaTree(src string) parser.ILambdaContext {
	is := antlr.NewInputStream(src)
	lexer := parser.NewJsonTemplateLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(antlr.NewConsoleErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplateParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	return p.Lambda()
}

func (v *LambdaVisitor) VisitExpression(ctx *parser.ExpressionContext) {
	v.Visit(ctx.Field())
}

func (v *LambdaVisitor) VisitField(context *parser.FieldContext) {
	for _, f := range context.AllField() {
		v.Visit(f)
	}
	for _, f := range context.AllFunction_param() {
		v.Visit(f)
	}
	if context.Name() != nil && len(context.AllField()) == 0 {
		v.usedVariables = append(v.usedVariables, context.Name().GetText())
	}
	if context.Index() != nil {
		v.Visit(context.Index())
	}
	if context.Array() != nil {
		v.Visit(context.Array())
	}
	if context.Object() != nil {
		v.Visit(context.Object())
	}
}

func (v *LambdaVisitor) VisitIndex(context *parser.IndexContext) {
	if context.Field() != nil {
		v.Visit(context.Field())
	}
}

func (v *LambdaVisitor) VisitArray(context *parser.ArrayContext) {
	for _, f := range context.AllField() {
		v.Visit(f)
	}
}

func (v *LambdaVisitor) VisitObject(context *parser.ObjectContext) {
	for _, f := range context.AllObject_field() {
		v.Visit(f)
	}
}

func (v *LambdaVisitor) VisitObject_field(context *parser.Object_fieldContext) {
	v.Visit(context.Field())
}

func (v *LambdaVisitor) VisitLambda(ctx *parser.LambdaContext) {
	for _, context := range ctx.AllName() {
		v.arguments = append(v.arguments, context.GetText())
	}
	v.Visit(ctx.Field())
}

func (v *LambdaVisitor) VisitFunction_param(ctx *parser.Function_paramContext) {
	if ctx.Field() != nil {
		v.Visit(ctx.Field())
	}
	if ctx.Lambda() != nil {
		v.Visit(ctx.Lambda())
	}
}
