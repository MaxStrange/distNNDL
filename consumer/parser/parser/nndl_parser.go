// Generated from NNDL.g4 by ANTLR 4.7.

package parser // NNDL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 137,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 6, 5, 49, 10, 5, 13, 5, 14, 5, 50, 3, 5, 3, 5, 3, 6, 3, 6, 7,
	6, 57, 10, 6, 12, 6, 14, 6, 60, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 7, 8, 80, 10, 8, 12, 8, 14, 8, 83, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 98, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 116, 10, 13, 12, 13, 14,
	13, 119, 11, 13, 3, 14, 3, 14, 3, 14, 5, 14, 124, 10, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 132, 10, 14, 12, 14, 14, 14, 135, 11,
	14, 3, 14, 2, 4, 24, 26, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 2, 6, 3, 2, 13, 17, 3, 2, 18, 19, 3, 2, 20, 22, 3, 2, 23, 24, 2, 133,
	2, 31, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 41, 3, 2, 2, 2, 8, 46, 3, 2, 2,
	2, 10, 54, 3, 2, 2, 2, 12, 63, 3, 2, 2, 2, 14, 73, 3, 2, 2, 2, 16, 88,
	3, 2, 2, 2, 18, 97, 3, 2, 2, 2, 20, 99, 3, 2, 2, 2, 22, 103, 3, 2, 2, 2,
	24, 107, 3, 2, 2, 2, 26, 123, 3, 2, 2, 2, 28, 30, 5, 4, 3, 2, 29, 28, 3,
	2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32,
	34, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 34, 35, 5, 6, 4, 2, 35, 36, 7, 2, 2,
	3, 36, 3, 3, 2, 2, 2, 37, 38, 7, 26, 2, 2, 38, 39, 7, 38, 2, 2, 39, 40,
	7, 25, 2, 2, 40, 5, 3, 2, 2, 2, 41, 42, 7, 27, 2, 2, 42, 43, 5, 8, 5, 2,
	43, 44, 5, 10, 6, 2, 44, 45, 7, 28, 2, 2, 45, 7, 3, 2, 2, 2, 46, 48, 7,
	29, 2, 2, 47, 49, 5, 12, 7, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2,
	50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 7,
	30, 2, 2, 53, 9, 3, 2, 2, 2, 54, 58, 7, 31, 2, 2, 55, 57, 5, 14, 8, 2,
	56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3,
	2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 62, 7, 32, 2, 2, 62,
	11, 3, 2, 2, 2, 63, 64, 7, 33, 2, 2, 64, 65, 7, 3, 2, 2, 65, 66, 7, 38,
	2, 2, 66, 67, 7, 39, 2, 2, 67, 68, 7, 4, 2, 2, 68, 69, 7, 39, 2, 2, 69,
	70, 7, 5, 2, 2, 70, 71, 7, 38, 2, 2, 71, 72, 7, 25, 2, 2, 72, 13, 3, 2,
	2, 2, 73, 74, 5, 16, 9, 2, 74, 75, 7, 3, 2, 2, 75, 76, 7, 35, 2, 2, 76,
	81, 5, 16, 9, 2, 77, 78, 7, 6, 2, 2, 78, 80, 5, 16, 9, 2, 79, 77, 3, 2,
	2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84,
	3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 5, 2, 2, 85, 86, 7, 38, 2, 2,
	86, 87, 7, 25, 2, 2, 87, 15, 3, 2, 2, 2, 88, 89, 7, 7, 2, 2, 89, 90, 5,
	18, 10, 2, 90, 91, 7, 6, 2, 2, 91, 92, 5, 18, 10, 2, 92, 93, 7, 8, 2, 2,
	93, 17, 3, 2, 2, 2, 94, 98, 5, 20, 11, 2, 95, 98, 5, 22, 12, 2, 96, 98,
	5, 26, 14, 2, 97, 94, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 96, 3, 2, 2,
	2, 98, 19, 3, 2, 2, 2, 99, 100, 7, 9, 2, 2, 100, 101, 5, 24, 13, 2, 101,
	102, 7, 10, 2, 2, 102, 21, 3, 2, 2, 2, 103, 104, 7, 11, 2, 2, 104, 105,
	7, 36, 2, 2, 105, 106, 7, 12, 2, 2, 106, 23, 3, 2, 2, 2, 107, 108, 8, 13,
	1, 2, 108, 109, 5, 26, 14, 2, 109, 110, 9, 2, 2, 2, 110, 111, 5, 26, 14,
	2, 111, 117, 3, 2, 2, 2, 112, 113, 12, 3, 2, 2, 113, 114, 9, 3, 2, 2, 114,
	116, 5, 24, 13, 4, 115, 112, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115,
	3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 25, 3, 2, 2, 2, 119, 117, 3, 2,
	2, 2, 120, 121, 8, 14, 1, 2, 121, 124, 7, 37, 2, 2, 122, 124, 7, 39, 2,
	2, 123, 120, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 133, 3, 2, 2, 2, 125,
	126, 12, 6, 2, 2, 126, 127, 9, 4, 2, 2, 127, 132, 5, 26, 14, 7, 128, 129,
	12, 5, 2, 2, 129, 130, 9, 5, 2, 2, 130, 132, 5, 26, 14, 6, 131, 125, 3,
	2, 2, 2, 131, 128, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2,
	2, 133, 134, 3, 2, 2, 2, 134, 27, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 11,
	31, 50, 58, 81, 97, 117, 123, 131, 133,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'x'", "'='", "','", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'<'", "'<='", "'>'", "'>='", "'=='", "'&&'", "'||'", "'*'", "'/'", "'%'",
	"'+'", "'-'", "';'", "'import'", "'network'", "'end network'", "'neurons'",
	"'end neurons'", "'connections'", "'end connections'", "'layer'", "", "'connect'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "STAT_END", "IMPORT", "NETWORK_OPEN", "NETWORK_END",
	"NEURON_OPEN", "NEURON_END", "CONNECT_OPEN", "CONNECT_END", "LAYER", "MAT_DECL",
	"CONNECT", "CLASS", "NEUR_VAR", "ID", "NUM", "WS",
}

var ruleNames = []string{
	"prog", "import_stat", "network_stat", "neuron_stat", "connection_stat",
	"layer_stat", "con_stat", "neuron_selection", "logical_expr", "log_pred_expr",
	"class_expr", "predicate", "math_expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type NNDLParser struct {
	*antlr.BaseParser
}

func NewNNDLParser(input antlr.TokenStream) *NNDLParser {
	this := new(NNDLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "NNDL.g4"

	return this
}

// NNDLParser tokens.
const (
	NNDLParserEOF          = antlr.TokenEOF
	NNDLParserT__0         = 1
	NNDLParserT__1         = 2
	NNDLParserT__2         = 3
	NNDLParserT__3         = 4
	NNDLParserT__4         = 5
	NNDLParserT__5         = 6
	NNDLParserT__6         = 7
	NNDLParserT__7         = 8
	NNDLParserT__8         = 9
	NNDLParserT__9         = 10
	NNDLParserT__10        = 11
	NNDLParserT__11        = 12
	NNDLParserT__12        = 13
	NNDLParserT__13        = 14
	NNDLParserT__14        = 15
	NNDLParserT__15        = 16
	NNDLParserT__16        = 17
	NNDLParserT__17        = 18
	NNDLParserT__18        = 19
	NNDLParserT__19        = 20
	NNDLParserT__20        = 21
	NNDLParserT__21        = 22
	NNDLParserSTAT_END     = 23
	NNDLParserIMPORT       = 24
	NNDLParserNETWORK_OPEN = 25
	NNDLParserNETWORK_END  = 26
	NNDLParserNEURON_OPEN  = 27
	NNDLParserNEURON_END   = 28
	NNDLParserCONNECT_OPEN = 29
	NNDLParserCONNECT_END  = 30
	NNDLParserLAYER        = 31
	NNDLParserMAT_DECL     = 32
	NNDLParserCONNECT      = 33
	NNDLParserCLASS        = 34
	NNDLParserNEUR_VAR     = 35
	NNDLParserID           = 36
	NNDLParserNUM          = 37
	NNDLParserWS           = 38
)

// NNDLParser rules.
const (
	NNDLParserRULE_prog             = 0
	NNDLParserRULE_import_stat      = 1
	NNDLParserRULE_network_stat     = 2
	NNDLParserRULE_neuron_stat      = 3
	NNDLParserRULE_connection_stat  = 4
	NNDLParserRULE_layer_stat       = 5
	NNDLParserRULE_con_stat         = 6
	NNDLParserRULE_neuron_selection = 7
	NNDLParserRULE_logical_expr     = 8
	NNDLParserRULE_log_pred_expr    = 9
	NNDLParserRULE_class_expr       = 10
	NNDLParserRULE_predicate        = 11
	NNDLParserRULE_math_expr        = 12
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Network_stat() INetwork_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INetwork_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INetwork_statContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(NNDLParserEOF, 0)
}

func (s *ProgContext) AllImport_stat() []IImport_statContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImport_statContext)(nil)).Elem())
	var tst = make([]IImport_statContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_statContext)
		}
	}

	return tst
}

func (s *ProgContext) Import_stat(i int) IImport_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_statContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_statContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *NNDLParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NNDLParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NNDLParserIMPORT {
		{
			p.SetState(26)
			p.Import_stat()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Network_stat()
	}
	{
		p.SetState(33)
		p.Match(NNDLParserEOF)
	}

	return localctx
}

// IImport_statContext is an interface to support dynamic dispatch.
type IImport_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_statContext differentiates from other interfaces.
	IsImport_statContext()
}

type Import_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_statContext() *Import_statContext {
	var p = new(Import_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_import_stat
	return p
}

func (*Import_statContext) IsImport_statContext() {}

func NewImport_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_statContext {
	var p = new(Import_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_import_stat

	return p
}

func (s *Import_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Import_statContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(NNDLParserIMPORT, 0)
}

func (s *Import_statContext) ID() antlr.TerminalNode {
	return s.GetToken(NNDLParserID, 0)
}

func (s *Import_statContext) STAT_END() antlr.TerminalNode {
	return s.GetToken(NNDLParserSTAT_END, 0)
}

func (s *Import_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterImport_stat(s)
	}
}

func (s *Import_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitImport_stat(s)
	}
}

func (p *NNDLParser) Import_stat() (localctx IImport_statContext) {
	localctx = NewImport_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NNDLParserRULE_import_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(NNDLParserIMPORT)
	}
	{
		p.SetState(36)
		p.Match(NNDLParserID)
	}
	{
		p.SetState(37)
		p.Match(NNDLParserSTAT_END)
	}

	return localctx
}

// INetwork_statContext is an interface to support dynamic dispatch.
type INetwork_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNetwork_statContext differentiates from other interfaces.
	IsNetwork_statContext()
}

type Network_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNetwork_statContext() *Network_statContext {
	var p = new(Network_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_network_stat
	return p
}

func (*Network_statContext) IsNetwork_statContext() {}

func NewNetwork_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Network_statContext {
	var p = new(Network_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_network_stat

	return p
}

func (s *Network_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Network_statContext) NETWORK_OPEN() antlr.TerminalNode {
	return s.GetToken(NNDLParserNETWORK_OPEN, 0)
}

func (s *Network_statContext) Neuron_stat() INeuron_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INeuron_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INeuron_statContext)
}

func (s *Network_statContext) Connection_stat() IConnection_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnection_statContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnection_statContext)
}

func (s *Network_statContext) NETWORK_END() antlr.TerminalNode {
	return s.GetToken(NNDLParserNETWORK_END, 0)
}

func (s *Network_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Network_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Network_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterNetwork_stat(s)
	}
}

func (s *Network_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitNetwork_stat(s)
	}
}

func (p *NNDLParser) Network_stat() (localctx INetwork_statContext) {
	localctx = NewNetwork_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NNDLParserRULE_network_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(NNDLParserNETWORK_OPEN)
	}
	{
		p.SetState(40)
		p.Neuron_stat()
	}
	{
		p.SetState(41)
		p.Connection_stat()
	}
	{
		p.SetState(42)
		p.Match(NNDLParserNETWORK_END)
	}

	return localctx
}

// INeuron_statContext is an interface to support dynamic dispatch.
type INeuron_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeuron_statContext differentiates from other interfaces.
	IsNeuron_statContext()
}

type Neuron_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeuron_statContext() *Neuron_statContext {
	var p = new(Neuron_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_neuron_stat
	return p
}

func (*Neuron_statContext) IsNeuron_statContext() {}

func NewNeuron_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Neuron_statContext {
	var p = new(Neuron_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_neuron_stat

	return p
}

func (s *Neuron_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Neuron_statContext) NEURON_OPEN() antlr.TerminalNode {
	return s.GetToken(NNDLParserNEURON_OPEN, 0)
}

func (s *Neuron_statContext) NEURON_END() antlr.TerminalNode {
	return s.GetToken(NNDLParserNEURON_END, 0)
}

func (s *Neuron_statContext) AllLayer_stat() []ILayer_statContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILayer_statContext)(nil)).Elem())
	var tst = make([]ILayer_statContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILayer_statContext)
		}
	}

	return tst
}

func (s *Neuron_statContext) Layer_stat(i int) ILayer_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILayer_statContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILayer_statContext)
}

func (s *Neuron_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Neuron_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Neuron_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterNeuron_stat(s)
	}
}

func (s *Neuron_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitNeuron_stat(s)
	}
}

func (p *NNDLParser) Neuron_stat() (localctx INeuron_statContext) {
	localctx = NewNeuron_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NNDLParserRULE_neuron_stat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(NNDLParserNEURON_OPEN)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == NNDLParserLAYER {
		{
			p.SetState(45)
			p.Layer_stat()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(NNDLParserNEURON_END)
	}

	return localctx
}

// IConnection_statContext is an interface to support dynamic dispatch.
type IConnection_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnection_statContext differentiates from other interfaces.
	IsConnection_statContext()
}

type Connection_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnection_statContext() *Connection_statContext {
	var p = new(Connection_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_connection_stat
	return p
}

func (*Connection_statContext) IsConnection_statContext() {}

func NewConnection_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Connection_statContext {
	var p = new(Connection_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_connection_stat

	return p
}

func (s *Connection_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Connection_statContext) CONNECT_OPEN() antlr.TerminalNode {
	return s.GetToken(NNDLParserCONNECT_OPEN, 0)
}

func (s *Connection_statContext) CONNECT_END() antlr.TerminalNode {
	return s.GetToken(NNDLParserCONNECT_END, 0)
}

func (s *Connection_statContext) AllCon_stat() []ICon_statContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICon_statContext)(nil)).Elem())
	var tst = make([]ICon_statContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICon_statContext)
		}
	}

	return tst
}

func (s *Connection_statContext) Con_stat(i int) ICon_statContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICon_statContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICon_statContext)
}

func (s *Connection_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Connection_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Connection_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterConnection_stat(s)
	}
}

func (s *Connection_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitConnection_stat(s)
	}
}

func (p *NNDLParser) Connection_stat() (localctx IConnection_statContext) {
	localctx = NewConnection_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NNDLParserRULE_connection_stat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(NNDLParserCONNECT_OPEN)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NNDLParserT__4 {
		{
			p.SetState(53)
			p.Con_stat()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(NNDLParserCONNECT_END)
	}

	return localctx
}

// ILayer_statContext is an interface to support dynamic dispatch.
type ILayer_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLayer_statContext differentiates from other interfaces.
	IsLayer_statContext()
}

type Layer_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLayer_statContext() *Layer_statContext {
	var p = new(Layer_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_layer_stat
	return p
}

func (*Layer_statContext) IsLayer_statContext() {}

func NewLayer_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Layer_statContext {
	var p = new(Layer_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_layer_stat

	return p
}

func (s *Layer_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Layer_statContext) LAYER() antlr.TerminalNode {
	return s.GetToken(NNDLParserLAYER, 0)
}

func (s *Layer_statContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NNDLParserID)
}

func (s *Layer_statContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NNDLParserID, i)
}

func (s *Layer_statContext) AllNUM() []antlr.TerminalNode {
	return s.GetTokens(NNDLParserNUM)
}

func (s *Layer_statContext) NUM(i int) antlr.TerminalNode {
	return s.GetToken(NNDLParserNUM, i)
}

func (s *Layer_statContext) STAT_END() antlr.TerminalNode {
	return s.GetToken(NNDLParserSTAT_END, 0)
}

func (s *Layer_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Layer_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Layer_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterLayer_stat(s)
	}
}

func (s *Layer_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitLayer_stat(s)
	}
}

func (p *NNDLParser) Layer_stat() (localctx ILayer_statContext) {
	localctx = NewLayer_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NNDLParserRULE_layer_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(NNDLParserLAYER)
	}
	{
		p.SetState(62)
		p.Match(NNDLParserT__0)
	}
	{
		p.SetState(63)
		p.Match(NNDLParserID)
	}
	{
		p.SetState(64)
		p.Match(NNDLParserNUM)
	}
	{
		p.SetState(65)
		p.Match(NNDLParserT__1)
	}
	{
		p.SetState(66)
		p.Match(NNDLParserNUM)
	}
	{
		p.SetState(67)
		p.Match(NNDLParserT__2)
	}
	{
		p.SetState(68)
		p.Match(NNDLParserID)
	}
	{
		p.SetState(69)
		p.Match(NNDLParserSTAT_END)
	}

	return localctx
}

// ICon_statContext is an interface to support dynamic dispatch.
type ICon_statContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCon_statContext differentiates from other interfaces.
	IsCon_statContext()
}

type Con_statContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCon_statContext() *Con_statContext {
	var p = new(Con_statContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_con_stat
	return p
}

func (*Con_statContext) IsCon_statContext() {}

func NewCon_statContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Con_statContext {
	var p = new(Con_statContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_con_stat

	return p
}

func (s *Con_statContext) GetParser() antlr.Parser { return s.parser }

func (s *Con_statContext) AllNeuron_selection() []INeuron_selectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INeuron_selectionContext)(nil)).Elem())
	var tst = make([]INeuron_selectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INeuron_selectionContext)
		}
	}

	return tst
}

func (s *Con_statContext) Neuron_selection(i int) INeuron_selectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INeuron_selectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INeuron_selectionContext)
}

func (s *Con_statContext) CONNECT() antlr.TerminalNode {
	return s.GetToken(NNDLParserCONNECT, 0)
}

func (s *Con_statContext) ID() antlr.TerminalNode {
	return s.GetToken(NNDLParserID, 0)
}

func (s *Con_statContext) STAT_END() antlr.TerminalNode {
	return s.GetToken(NNDLParserSTAT_END, 0)
}

func (s *Con_statContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Con_statContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Con_statContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterCon_stat(s)
	}
}

func (s *Con_statContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitCon_stat(s)
	}
}

func (p *NNDLParser) Con_stat() (localctx ICon_statContext) {
	localctx = NewCon_statContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NNDLParserRULE_con_stat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Neuron_selection()
	}
	{
		p.SetState(72)
		p.Match(NNDLParserT__0)
	}
	{
		p.SetState(73)
		p.Match(NNDLParserCONNECT)
	}
	{
		p.SetState(74)
		p.Neuron_selection()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NNDLParserT__3 {
		{
			p.SetState(75)
			p.Match(NNDLParserT__3)
		}
		{
			p.SetState(76)
			p.Neuron_selection()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(NNDLParserT__2)
	}
	{
		p.SetState(83)
		p.Match(NNDLParserID)
	}
	{
		p.SetState(84)
		p.Match(NNDLParserSTAT_END)
	}

	return localctx
}

// INeuron_selectionContext is an interface to support dynamic dispatch.
type INeuron_selectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeuron_selectionContext differentiates from other interfaces.
	IsNeuron_selectionContext()
}

type Neuron_selectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeuron_selectionContext() *Neuron_selectionContext {
	var p = new(Neuron_selectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_neuron_selection
	return p
}

func (*Neuron_selectionContext) IsNeuron_selectionContext() {}

func NewNeuron_selectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Neuron_selectionContext {
	var p = new(Neuron_selectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_neuron_selection

	return p
}

func (s *Neuron_selectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Neuron_selectionContext) AllLogical_expr() []ILogical_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem())
	var tst = make([]ILogical_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILogical_exprContext)
		}
	}

	return tst
}

func (s *Neuron_selectionContext) Logical_expr(i int) ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Neuron_selectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Neuron_selectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Neuron_selectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterNeuron_selection(s)
	}
}

func (s *Neuron_selectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitNeuron_selection(s)
	}
}

func (p *NNDLParser) Neuron_selection() (localctx INeuron_selectionContext) {
	localctx = NewNeuron_selectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NNDLParserRULE_neuron_selection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(NNDLParserT__4)
	}
	{
		p.SetState(87)
		p.Logical_expr()
	}
	{
		p.SetState(88)
		p.Match(NNDLParserT__3)
	}
	{
		p.SetState(89)
		p.Logical_expr()
	}
	{
		p.SetState(90)
		p.Match(NNDLParserT__5)
	}

	return localctx
}

// ILogical_exprContext is an interface to support dynamic dispatch.
type ILogical_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogical_exprContext differentiates from other interfaces.
	IsLogical_exprContext()
}

type Logical_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_exprContext() *Logical_exprContext {
	var p = new(Logical_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_logical_expr
	return p
}

func (*Logical_exprContext) IsLogical_exprContext() {}

func NewLogical_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_exprContext {
	var p = new(Logical_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_logical_expr

	return p
}

func (s *Logical_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_exprContext) Log_pred_expr() ILog_pred_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILog_pred_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILog_pred_exprContext)
}

func (s *Logical_exprContext) Class_expr() IClass_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClass_exprContext)
}

func (s *Logical_exprContext) Math_expr() IMath_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMath_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMath_exprContext)
}

func (s *Logical_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterLogical_expr(s)
	}
}

func (s *Logical_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitLogical_expr(s)
	}
}

func (p *NNDLParser) Logical_expr() (localctx ILogical_exprContext) {
	localctx = NewLogical_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NNDLParserRULE_logical_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NNDLParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Log_pred_expr()
		}

	case NNDLParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Class_expr()
		}

	case NNDLParserNEUR_VAR, NNDLParserNUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.math_expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILog_pred_exprContext is an interface to support dynamic dispatch.
type ILog_pred_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLog_pred_exprContext differentiates from other interfaces.
	IsLog_pred_exprContext()
}

type Log_pred_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLog_pred_exprContext() *Log_pred_exprContext {
	var p = new(Log_pred_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_log_pred_expr
	return p
}

func (*Log_pred_exprContext) IsLog_pred_exprContext() {}

func NewLog_pred_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Log_pred_exprContext {
	var p = new(Log_pred_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_log_pred_expr

	return p
}

func (s *Log_pred_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Log_pred_exprContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *Log_pred_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log_pred_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Log_pred_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterLog_pred_expr(s)
	}
}

func (s *Log_pred_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitLog_pred_expr(s)
	}
}

func (p *NNDLParser) Log_pred_expr() (localctx ILog_pred_exprContext) {
	localctx = NewLog_pred_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NNDLParserRULE_log_pred_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(NNDLParserT__6)
	}
	{
		p.SetState(98)
		p.predicate(0)
	}
	{
		p.SetState(99)
		p.Match(NNDLParserT__7)
	}

	return localctx
}

// IClass_exprContext is an interface to support dynamic dispatch.
type IClass_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClass_exprContext differentiates from other interfaces.
	IsClass_exprContext()
}

type Class_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_exprContext() *Class_exprContext {
	var p = new(Class_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_class_expr
	return p
}

func (*Class_exprContext) IsClass_exprContext() {}

func NewClass_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_exprContext {
	var p = new(Class_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_class_expr

	return p
}

func (s *Class_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Class_exprContext) CLASS() antlr.TerminalNode {
	return s.GetToken(NNDLParserCLASS, 0)
}

func (s *Class_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterClass_expr(s)
	}
}

func (s *Class_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitClass_expr(s)
	}
}

func (p *NNDLParser) Class_expr() (localctx IClass_exprContext) {
	localctx = NewClass_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NNDLParserRULE_class_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(NNDLParserT__8)
	}
	{
		p.SetState(102)
		p.Match(NNDLParserCLASS)
	}
	{
		p.SetState(103)
		p.Match(NNDLParserT__9)
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) AllMath_expr() []IMath_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMath_exprContext)(nil)).Elem())
	var tst = make([]IMath_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMath_exprContext)
		}
	}

	return tst
}

func (s *PredicateContext) Math_expr(i int) IMath_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMath_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMath_exprContext)
}

func (s *PredicateContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *PredicateContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *NNDLParser) Predicate() (localctx IPredicateContext) {
	return p.predicate(0)
}

func (p *NNDLParser) predicate(_p int) (localctx IPredicateContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPredicateContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, NNDLParserRULE_predicate, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.math_expr(0)
	}
	p.SetState(107)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NNDLParserT__10)|(1<<NNDLParserT__11)|(1<<NNDLParserT__12)|(1<<NNDLParserT__13)|(1<<NNDLParserT__14))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(108)
		p.math_expr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPredicateContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, NNDLParserRULE_predicate)
			p.SetState(110)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(111)
			_la = p.GetTokenStream().LA(1)

			if !(_la == NNDLParserT__15 || _la == NNDLParserT__16) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
			{
				p.SetState(112)
				p.predicate(2)
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IMath_exprContext is an interface to support dynamic dispatch.
type IMath_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMath_exprContext differentiates from other interfaces.
	IsMath_exprContext()
}

type Math_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_exprContext() *Math_exprContext {
	var p = new(Math_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NNDLParserRULE_math_expr
	return p
}

func (*Math_exprContext) IsMath_exprContext() {}

func NewMath_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_exprContext {
	var p = new(Math_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NNDLParserRULE_math_expr

	return p
}

func (s *Math_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_exprContext) NEUR_VAR() antlr.TerminalNode {
	return s.GetToken(NNDLParserNEUR_VAR, 0)
}

func (s *Math_exprContext) NUM() antlr.TerminalNode {
	return s.GetToken(NNDLParserNUM, 0)
}

func (s *Math_exprContext) AllMath_expr() []IMath_exprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMath_exprContext)(nil)).Elem())
	var tst = make([]IMath_exprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMath_exprContext)
		}
	}

	return tst
}

func (s *Math_exprContext) Math_expr(i int) IMath_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMath_exprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMath_exprContext)
}

func (s *Math_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.EnterMath_expr(s)
	}
}

func (s *Math_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NNDLListener); ok {
		listenerT.ExitMath_expr(s)
	}
}

func (p *NNDLParser) Math_expr() (localctx IMath_exprContext) {
	return p.math_expr(0)
}

func (p *NNDLParser) math_expr(_p int) (localctx IMath_exprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMath_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMath_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, NNDLParserRULE_math_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NNDLParserNEUR_VAR:
		{
			p.SetState(119)
			p.Match(NNDLParserNEUR_VAR)
		}

	case NNDLParserNUM:
		{
			p.SetState(120)
			p.Match(NNDLParserNUM)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMath_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NNDLParserRULE_math_expr)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(124)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NNDLParserT__17)|(1<<NNDLParserT__18)|(1<<NNDLParserT__19))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(125)
					p.math_expr(5)
				}

			case 2:
				localctx = NewMath_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, NNDLParserRULE_math_expr)
				p.SetState(126)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(127)
				_la = p.GetTokenStream().LA(1)

				if !(_la == NNDLParserT__20 || _la == NNDLParserT__21) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(128)
					p.math_expr(4)
				}

			}

		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

func (p *NNDLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *PredicateContext = nil
		if localctx != nil {
			t = localctx.(*PredicateContext)
		}
		return p.Predicate_Sempred(t, predIndex)

	case 12:
		var t *Math_exprContext = nil
		if localctx != nil {
			t = localctx.(*Math_exprContext)
		}
		return p.Math_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NNDLParser) Predicate_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *NNDLParser) Math_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
