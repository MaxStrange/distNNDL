package main

import (
    "github.com/antlr/antlr4/runtime/Go/antlr"
    "github.com/MaxStrange/distNNDL/compiler/parser"
    "os"
    "fmt"
)

type TreeShapeListener struct {
    *parser.BaseNNDLListener
}

func NewTreeShapeListener() *TreeShapeListener {
    return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
    fmt.Println(ctx.GetText())
}

func main() {
    input := antlr.NewFileStream(os.Args[1])
    lexer := parser.NewNNDLLexer(input)
    stream := antlr.NewCommonTokenStream(lexer, 0)
    p := parser.NewNNDLParser(stream)
    p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
    p.BuildParseTrees = true
    tree := p.Prog()
    antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
