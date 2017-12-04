package main

import (
	"github.com/MaxStrange/distNNDL/consumer/parser/parser"
)

type treeShapeListener struct {
	*parser.BaseNNDLListener
	net *network
}

func newTreeShapeListener(net *network) *treeShapeListener {
	return &treeShapeListener{
		net: net,
	}
}

////////////////////////////////////////////////////////////////////////////
/*
func (tsl *treeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func (tsl *treeShapeListener) ExitLayer_stat(ctx *parser.Layer_statContext) {
	layerName := ctx.ID(0).GetText()
	numRows, _ := strconv.ParseInt(ctx.NUM(0).GetText(), 10, 64)
	numCols, _ := strconv.ParseInt(ctx.NUM(1).GetText(), 10, 64)
	neuronType := ctx.ID(1).GetText()
	tsl.net.addLayer(numRows, numCols, neuronType, layerName)
}

func (tsl *treeShapeListener) EnterConnection_stat(ctx *parser.Connection_statContext) {
	//self._conx = []
}

func (tsl *treeShapeListener) ExitConnection_stat(ctx *parser.Connection_statContext) {
}
*/
////////////////////////////////////////////////////////////////////////////

func parseNNDLIntoNetwork(nndlContent string, net *network) {
	// input := antlr.NewInputStream(nndlContent)
	// lexer := parser.NewNNDLLexer(input)
	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// p := parser.NewNNDLParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true
	// tree := p.Prog()
	// antlr.ParseTreeWalkerDefault.Walk(newTreeShapeListener(net), tree)

	// TODO: print nndlContent to a file, then hand it over to Python to parse.
}
