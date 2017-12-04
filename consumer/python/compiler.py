import antlr4
from grammar import NNDLLexer
from grammar import NNDLParser
from dotgen import DotGenerator
import sys

def compile(nndl_content):
    """
    """
    input_stream = antlr4.InputStream(nndl_content)
    lexer = NNDLLexer.NNDLLexer(input_stream)
    stream = antlr4.CommonTokenStream(lexer)
    parser = NNDLParser.NNDLParser(stream)
    tree = parser.prog()

    # Walk the tree and generate the dot file
    dg = DotGenerator()
    walker = antlr4.ParseTreeWalker()
    walker.walk(dg, tree)

    # Use the dotgenerator's network that it figured out from the nndl file
    # to generate the cpp file
    nw = dg._network
    return nw

if __name__ == "__main__":
    nndl_content = [line for line in sys.stdin]
    nw  = compile("".join(nndl_content))
    print(nw.to_json())
