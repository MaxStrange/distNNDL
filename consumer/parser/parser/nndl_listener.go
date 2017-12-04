// Generated from NNDL.g4 by ANTLR 4.7.

package parser // NNDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NNDLListener is a complete listener for a parse tree produced by NNDLParser.
type NNDLListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterImport_stat is called when entering the import_stat production.
	EnterImport_stat(c *Import_statContext)

	// EnterNetwork_stat is called when entering the network_stat production.
	EnterNetwork_stat(c *Network_statContext)

	// EnterNeuron_stat is called when entering the neuron_stat production.
	EnterNeuron_stat(c *Neuron_statContext)

	// EnterConnection_stat is called when entering the connection_stat production.
	EnterConnection_stat(c *Connection_statContext)

	// EnterLayer_stat is called when entering the layer_stat production.
	EnterLayer_stat(c *Layer_statContext)

	// EnterCon_stat is called when entering the con_stat production.
	EnterCon_stat(c *Con_statContext)

	// EnterNeuron_selection is called when entering the neuron_selection production.
	EnterNeuron_selection(c *Neuron_selectionContext)

	// EnterLogical_expr is called when entering the logical_expr production.
	EnterLogical_expr(c *Logical_exprContext)

	// EnterLog_pred_expr is called when entering the log_pred_expr production.
	EnterLog_pred_expr(c *Log_pred_exprContext)

	// EnterClass_expr is called when entering the class_expr production.
	EnterClass_expr(c *Class_exprContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterMath_expr is called when entering the math_expr production.
	EnterMath_expr(c *Math_exprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitImport_stat is called when exiting the import_stat production.
	ExitImport_stat(c *Import_statContext)

	// ExitNetwork_stat is called when exiting the network_stat production.
	ExitNetwork_stat(c *Network_statContext)

	// ExitNeuron_stat is called when exiting the neuron_stat production.
	ExitNeuron_stat(c *Neuron_statContext)

	// ExitConnection_stat is called when exiting the connection_stat production.
	ExitConnection_stat(c *Connection_statContext)

	// ExitLayer_stat is called when exiting the layer_stat production.
	ExitLayer_stat(c *Layer_statContext)

	// ExitCon_stat is called when exiting the con_stat production.
	ExitCon_stat(c *Con_statContext)

	// ExitNeuron_selection is called when exiting the neuron_selection production.
	ExitNeuron_selection(c *Neuron_selectionContext)

	// ExitLogical_expr is called when exiting the logical_expr production.
	ExitLogical_expr(c *Logical_exprContext)

	// ExitLog_pred_expr is called when exiting the log_pred_expr production.
	ExitLog_pred_expr(c *Log_pred_exprContext)

	// ExitClass_expr is called when exiting the class_expr production.
	ExitClass_expr(c *Class_exprContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitMath_expr is called when exiting the math_expr production.
	ExitMath_expr(c *Math_exprContext)
}
