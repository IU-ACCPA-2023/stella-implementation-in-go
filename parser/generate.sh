#!/bin/sh

set -e
set -x

JAVA=java
JAVA_FLAGS="-Xmx500M -cp ${HOME}/.java/jar/antlr-4.11.1-complete.jar:${CLASSPATH}"
PARSER="${JAVA} ${JAVA_FLAGS} org.antlr.v4.Tool"
PARSER_FLAGS="-Dlanguage=Go -no-visitor -package parser"

${PARSER} ${PARSER_FLAGS} *.g4
