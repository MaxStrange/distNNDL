#! /usr/bin/env bash
echo "Removing everything in the build directory..."
if [ -d build ]; then
    rm build/*
fi

echo "Running antlr4..."
java -jar /usr/local/lib/antlr-4.7-complete.jar NNDL.g4

echo "Compiling .java files..."
javac -classpath /usr/local/lib/antlr-4.7-complete.jar NNDL*.java

mkdir build

echo "Moving everything to the build directory..."
mv NNDL*.java build/
mv NNDL*.class build/
mv NNDL*.tokens build/
cp NNDL.g4 build/
cp examples/* build/

echo "Done. To grun, cd into build and run: 'grun NNDL prog <EXAMPLE> -gui'"
