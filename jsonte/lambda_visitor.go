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
	case *parser.StatementsContext:
		v.VisitStatements(val)
		break
	case *parser.StatementContext:
		v.VisitStatement(val)
		break
	case *parser.Template_stringContext:
		v.VisitTemplate_string(val)
		break
	case *parser.Template_string_partContext:
		v.VisitTemplate_string_part(val)
		break
	}
}

// resolveLambdaTree resolves a string to an AST tree
func (v *LambdaVisitor) resolveLambdaTree(src string) parser.ILambdaContext {
	is := antlr.NewInputStream(src)
	lexer := NewTemplateAwareLexer(is)
	lexer.JsonTemplateLexer.RemoveErrorListeners()
	lexer.JsonTemplateLexer.AddErrorListener(antlr.NewConsoleErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewJsonTemplate(stream)
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
	if context.Lambda() != nil {
		v.Visit(context.Lambda())
	}
	if context.Template_string() != nil {
		v.Visit(context.Template_string())
	}
}

func (v *LambdaVisitor) VisitTemplate_string(ctx *parser.Template_stringContext) {
	for _, part := range ctx.AllTemplate_string_part() {
		v.Visit(part)
	}
}

func (v *LambdaVisitor) VisitTemplate_string_part(ctx *parser.Template_string_partContext) {
	if ctx.Field() != nil {
		v.Visit(ctx.Field())
	}
}

func (v *LambdaVisitor) VisitIndex(context *parser.IndexContext) {
	if context.Field() != nil {
		v.Visit(context.Field())
	}
}

func (v *LambdaVisitor) VisitArray(context *parser.ArrayContext) {
	for _, f := range context.AllSpread_field() {
		v.Visit(f.(*parser.Spread_fieldContext).Field())
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
	if ctx.Field() != nil {
		v.Visit(ctx.Field())
	}
	if ctx.Statements() != nil {
		v.Visit(ctx.Statements())
	}
}

func (v *LambdaVisitor) VisitStatements(ctx *parser.StatementsContext) {
	for _, s := range ctx.AllStatement() {
		v.Visit(s)
	}
}

func (v *LambdaVisitor) VisitStatement(ctx *parser.StatementContext) {
	for _, f := range ctx.AllField() {
		v.Visit(f)
	}
	for _, s := range ctx.AllStatements() {
		v.Visit(s)
	}
}

func (v *LambdaVisitor) VisitName(ctx *parser.NameContext) {
	v.usedVariables = append(v.usedVariables, ctx.GetText())
}

func (v *LambdaVisitor) VisitFunction_param(ctx *parser.Function_paramContext) {
	if ctx.Field() != nil {
		v.Visit(ctx.Field())
	}
	if ctx.Lambda() != nil {
		v.Visit(ctx.Lambda())
	}
}


