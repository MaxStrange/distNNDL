## antlr

This directory is where I mess around with the grammar and where
the lexers and parsers are generated.

### Set up

1. Download ANTLR4's jar file: `curl -O http://www.antlr.org/download/antlr-4.7-complete.jar`
1. `mv antlr-4.7-complete.jar /usr/local/lib/`
1. Add `alias grun='java -cp .:/usr/local/lib/antlr-4.7-complete.jar org.antlr.v4.gui.TestRig` to your .bashrc or .bash_profile

