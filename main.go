package main

import (
	"stella-implementation-in-go/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func handle_expr_context(ctx parser.IExprContext) {
	_, is_const_true := ctx.(*parser.ConstTrueContext)
	if is_const_true {
		print("I see true\n")
		return
	}
	if_ctx, is_if := ctx.(*parser.IfContext)
	if is_if {
		handle_expr_context(if_ctx.GetCondition())
		handle_expr_context(if_ctx.GetThenExpr())
		handle_expr_context(if_ctx.GetElseExpr())
		return
	}
	var_ctx, is_var := ctx.(*parser.VarContext)
	if is_var {
		print("I see variable ", var_ctx.GetName().GetText(), "\n")
		return
	}
}

func handle_decl_context(ctx parser.IDeclContext) {
	decl_fun_ctx, is_decl_fun := ctx.(*parser.DeclFunContext)
	if is_decl_fun {
		print("Declare function ", decl_fun_ctx.GetName().GetText(), "\n")
		// decl_fun_ctx.GetParamDecls()
		// decl_fun_ctx.GetReturnType()
		handle_expr_context(decl_fun_ctx.GetReturnExpr())
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
