package stack

import (
	"container/list"
	"fmt"

	tokens "github.com/narutopig/neon-lang/lang/token"
)

// Stack is a stack of tokens, implemented with container/list
type Stack struct {
	stack *list.List
}

// New returns a new TokenStack
func New() Stack {
	return Stack{stack: list.New()}
}

// Push adds a value to the front of the stack
func (ts *Stack) Push(value tokens.Token) {
	ts.stack.PushFront(value)
}

// Pop removes the front of the stack and returns the removed item and an error if the stack length is 0
func (ts *Stack) Pop() tokens.Token {
	if ts.stack.Len() > 0 {
		ele := ts.stack.Front()
		ts.stack.Remove(ele)
		return ele.Value.(tokens.Token)
	}

	return tokens.NoneToken
}

// Front returns the item at the front of the list
func (ts *Stack) Front() tokens.Token {
	if ts.stack.Len() > 0 {
		if val, ok := ts.stack.Front().Value.(tokens.Token); ok {
			return val
		}
		return tokens.NoneToken
	}
	return tokens.NoneToken
}

// Size returns the size of the stack
func (ts *Stack) Size() int {
	return ts.stack.Len()
}

// Empty returns if the size of the stack is 0
func (ts *Stack) Empty() bool {
	return ts.stack.Len() == 0
}

// Copy returns a copy of the stack
func Copy(ts Stack) Stack {
	res := Stack{stack: list.New()}

	res.stack.PushBackList(ts.stack)
	return res
}

// Reversed returns a reversed copy of the stack
func Reversed(ts Stack) Stack {
	ts2 := Copy(ts)
	res := Stack{stack: list.New()}

	for !ts2.Empty() {
		res.stack.PushFront(ts2.Pop())
	}
	return res
}

// Print outputs the elements of the stack
func (ts Stack) Print() {
	// TODO: Fix the stack clearing after print
	ts2 := Copy(ts)

	for !ts2.Empty() {
		fmt.Println(ts2.Pop())
	}
}
