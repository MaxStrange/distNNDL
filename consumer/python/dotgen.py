"""
"""
import logic
import network
from grammar import NNDLListener


class DotGenerator(NNDLListener.NNDLListener):
    """
    Listener that generates dot files when walked over the nndl tree.
    """
    def __init__(self):
        self._network = network.Network()

    def enterProg(self, ctx):
        pass

    def exitProg(self, ctx):
        pass

    def exitLayer_stat(self, ctx):
        layer_name = ctx.ID()[0].getText()
        num_rows = int(ctx.NUM()[0].getText())
        num_cols = int(ctx.NUM()[1].getText())
        neuron_type = ctx.ID()[1].getText()

        self._network.add_layer(nrows=num_rows, ncols=num_cols,
                neurtype=neuron_type, name=layer_name)

    def enterConnection_stat(self, ctx):
        self._conx = []

    def exitConnection_stat(self, ctx):
        cons = ctx.con_stat()
        # A list of from rules and their corresponding to rules (which
        # is a list)
        rules = []
        for con_stat in cons:
            connect_from = con_stat.neuron_selection()[0]
            connect_to = con_stat.neuron_selection()[1:]

            from_logic_layer = connect_from.logical_expr()[0].getText()
            log_type = self._lookup_log_type(connect_from.logical_expr()[0])
            from_logic_layer = logic.py(from_logic_layer, log_type, "i")

            from_logic_neurs = connect_from.logical_expr()[1].getText()
            log_type = self._lookup_log_type(connect_from.logical_expr()[1])
            from_logic_neurs = logic.py(from_logic_neurs, log_type, "j")

            from_rule = logic.Rule(from_logic_layer, from_logic_neurs)

            to_rules = []
            for to in connect_to:
                to_logic_layer = to.logical_expr()[0].getText()
                log_type = self._lookup_log_type(to.logical_expr()[0])
                to_logic_layer = logic.py(to_logic_layer, log_type, "k")

                to_logic_neurs = to.logical_expr()[1].getText()
                log_type = self._lookup_log_type(to.logical_expr()[1])
                to_logic_neurs = logic.py(to_logic_neurs, log_type, "l")

                to_rule = logic.Rule(to_logic_layer, to_logic_neurs)
                to_rules.append(to_rule)
            rules.append((from_rule, to_rules))

        self._network = logic.pygen(rules, self._network)

    def _lookup_log_type(self, logical_expr):
        """
        Looks up whether the logical_expr is log_pred_expr
        or class_expr, or math_expr.
        """
        if logical_expr.log_pred_expr():
            return "pred"
        elif logical_expr.class_expr():
            return "class"
        else:
            return "math"
