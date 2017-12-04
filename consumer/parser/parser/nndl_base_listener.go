// Generated from NNDL.g4 by ANTLR 4.7.

package parser // NNDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNNDLListener is a complete listener for a parse tree produced by NNDLParser.
type BaseNNDLListener struct{}

var _ NNDLListener = &BaseNNDLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNNDLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNNDLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNNDLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNNDLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseNNDLListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseNNDLListener) ExitProg(ctx *ProgContext) {}

// EnterImport_stat is called when production import_stat is entered.
func (s *BaseNNDLListener) EnterImport_stat(ctx *Import_statContext) {}

// ExitImport_stat is called when production import_stat is exited.
func (s *BaseNNDLListener) ExitImport_stat(ctx *Import_statContext) {}

// EnterNetwork_stat is called when production network_stat is entered.
func (s *BaseNNDLListener) EnterNetwork_stat(ctx *Network_statContext) {}

// ExitNetwork_stat is called when production network_stat is exited.
func (s *BaseNNDLListener) ExitNetwork_stat(ctx *Network_statContext) {}

// EnterNeuron_stat is called when production neuron_stat is entered.
func (s *BaseNNDLListener) EnterNeuron_stat(ctx *Neuron_statContext) {}

// ExitNeuron_stat is called when production neuron_stat is exited.
func (s *BaseNNDLListener) ExitNeuron_stat(ctx *Neuron_statContext) {}

// EnterConnection_stat is called when production connection_stat is entered.
func (s *BaseNNDLListener) EnterConnection_stat(ctx *Connection_statContext) {}

// ExitConnection_stat is called when production connection_stat is exited.
func (s *BaseNNDLListener) ExitConnection_stat(ctx *Connection_statContext) {}

// EnterLayer_stat is called when production layer_stat is entered.
func (s *BaseNNDLListener) EnterLayer_stat(ctx *Layer_statContext) {}

// ExitLayer_stat is called when production layer_stat is exited.
func (s *BaseNNDLListener) ExitLayer_stat(ctx *Layer_statContext) {}

// EnterCon_stat is called when production con_stat is entered.
func (s *BaseNNDLListener) EnterCon_stat(ctx *Con_statContext) {}

// ExitCon_stat is called when production con_stat is exited.
func (s *BaseNNDLListener) ExitCon_stat(ctx *Con_statContext) {}

// EnterNeuron_selection is called when production neuron_selection is entered.
func (s *BaseNNDLListener) EnterNeuron_selection(ctx *Neuron_selectionContext) {}

// ExitNeuron_selection is called when production neuron_selection is exited.
func (s *BaseNNDLListener) ExitNeuron_selection(ctx *Neuron_selectionContext) {}

// EnterLogical_expr is called when production logical_expr is entered.
func (s *BaseNNDLListener) EnterLogical_expr(ctx *Logical_exprContext) {}

// ExitLogical_expr is called when production logical_expr is exited.
func (s *BaseNNDLListener) ExitLogical_expr(ctx *Logical_exprContext) {}

// EnterLog_pred_expr is called when production log_pred_expr is entered.
func (s *BaseNNDLListener) EnterLog_pred_expr(ctx *Log_pred_exprContext) {}

// ExitLog_pred_expr is called when production log_pred_expr is exited.
func (s *BaseNNDLListener) ExitLog_pred_expr(ctx *Log_pred_exprContext) {}

// EnterClass_expr is called when production class_expr is entered.
func (s *BaseNNDLListener) EnterClass_expr(ctx *Class_exprContext) {}

// ExitClass_expr is called when production class_expr is exited.
func (s *BaseNNDLListener) ExitClass_expr(ctx *Class_exprContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseNNDLListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseNNDLListener) ExitPredicate(ctx *PredicateContext) {}

// EnterMath_expr is called when production math_expr is entered.
func (s *BaseNNDLListener) EnterMath_expr(ctx *Math_exprContext) {}

// ExitMath_expr is called when production math_expr is exited.
func (s *BaseNNDLListener) ExitMath_expr(ctx *Math_exprContext) {}
