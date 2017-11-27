#! /usr/bin/env bash
echo "Removing everything in the build directory..."
if [ -d build ]; then
    rm build/*
fi

echo "Running antlr4..."
java -jar /usr/local/lib/antlr-4.7-complete.jar -Dlanguage=Go NNDL.g4

echo "Moving everything to the build directory..."
mv nndl*.go build/
mv NNDL*.tokens build/
cp NNDL.g4 build/
cp examples/* build/

echo "Copying everything to the compiler directory.."
cp build/* ../compiler/parser/
rm ../compiler/parser/*.ndl
cp examples/* ../compiler/examples/
