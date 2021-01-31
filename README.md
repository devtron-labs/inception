# inception

Command to build grammar
cd pkg/language/grammar/ && antlr Klang.g4 -o ../parser -package parser -Dlanguage=Go && cd ../../..

from the generated crd remove
subresources:
status: {}