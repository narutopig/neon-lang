package lib

// Node describes a node in an abstract syntax tree (at least a bad implementation of it)
type Node struct {
	Token    Token
	Children []Token
}
