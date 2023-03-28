package main

import (
	"stella-implementation-in-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func handle_expr_context(ctx parser.IExprContext) {
	switch ctx := ctx.(type) {
	default:
		print("unsupported syntax\n")
	case *parser.ConstTrueContext:
		print("I see true\n")
	case *parser.IfContext:
		handle_expr_context(ctx.GetCondition())
		handle_expr_context(ctx.GetThenExpr())
		handle_expr_context(ctx.GetElseExpr())
	case *parser.VarContext:
		print("I see variable ", ctx.GetName().GetText(), "\n")
	}
}

func handle_decl_context(ctx parser.IDeclContext) {
	switch ctx := ctx.(type) {
	default:
		print("unsupported syntax\n")
	case *parser.DeclFunContext:
		print("Declare function ", ctx.GetName().GetText(), "\n")
		// decl_fun_ctx.GetParamDecls()
		// decl_fun_ctx.GetReturnType()
		handle_expr_context(ctx.GetReturnExpr())
	}
}

func handle_program_context(ctx parser.IProgramContext) {
	for _, decl := range ctx.GetDecls() {
		handle_decl_context(decl)
	}
}

func main() {
	input := antlr.NewInputStream("language core; fn main(x : Nat) -> Nat { return x; }")
	lexer := parser.NewStellaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewStellaParser(stream)
	tree := p.Program()
	handle_program_context(tree)
}
