import antlr4
from grammar import NNDLLexer
from grammar import NNDLParser
import sys

def compile(nndl_content):
    """
    """
    input = antlr4.InputStream(nndl_content)
    lexer = NNDLLexer.NNDLLexer(input)
    stream = antlr4.CommonTokenStream(lexer)
    parser = NNDLParser.NNDLParser(stream)
    tree = parser.prog()

    # Walk the tree and generate the dot file
    print("Walking tree...")
    dg = dotgen.DotGenerator(file_to_compile, output_file_name + ".dot")
    walker = antlr4.ParseTreeWalker()
    walker.walk(dg, tree)

    # Use the dotgenerator's network that it figured out from the nndl file
    # to generate the cpp file
    print("Generating cpp files...")
    nw = dg._network

if __name__ == "__main__":
    with open("BLAH.txt", 'w') as f:
        f.write("Hello!\n")

    # Read everything from stdin
    nndl_content = [line for line in sys.stdin]
    with open("TEMPFILE.txt", 'w') as f:
        for line in nndl_content:
            f.write(line)
    # print something to stdout
    print("Got NNDL Content:", nndl_content)
