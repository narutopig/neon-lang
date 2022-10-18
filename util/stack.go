package util

import (
	"container/list"
	"fmt"

	tokens "github.com/narutopig/neon-lang/token"
)

// TokenStack is a stack of tokens, implemented with container/list
type TokenStack struct {
	stack *list.List
}

// NewTStack returns a new TokenStack
func NewTStack() TokenStack {
	return TokenStack{stack: list.New()}
}

// Push adds a value to the front of the stack
func (ts *TokenStack) Push(value tokens.Token) {
	ts.stack.PushFront(value)
}

// Pop removes the front of the stack and returns the removed item and an error if the stack length is 0
func (ts *TokenStack) Pop() (tokens.Token, error) {
	if ts.stack.Len() > 0 {
		ele := ts.stack.Front()
		ts.stack.Remove(ele)
		return ele.Value.(tokens.Token), nil
	}

	return tokens.NoneToken, fmt.Errorf("empty stack")
}

// Front returns the item at the front of the list
func (ts *TokenStack) Front() (tokens.Token, error) {
	if ts.stack.Len() > 0 {
		if val, ok := ts.stack.Front().Value.(tokens.Token); ok {
			return val, nil
		}
		return tokens.NoneToken, fmt.Errorf("wrong type of element in stack")
	}
	return tokens.NoneToken, fmt.Errorf("empty stack")
}

// Size returns the size of the stack
func (ts *TokenStack) Size() int {
	return ts.stack.Len()
}

// Empty returns if the size of the stack is 0
func (ts *TokenStack) Empty() bool {
	return ts.stack.Len() == 0
}
