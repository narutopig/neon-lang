package expressions

import (
	ne "github.com/narutopig/neon-lang/error"
	g "github.com/narutopig/neon-lang/grammar"
	"github.com/narutopig/neon-lang/stack"
	t "github.com/narutopig/neon-lang/token"
)

// Shunt implements the shunting-yard algorithm from https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func Shunt(tokens []t.Token, line int) (stack.TokenStack, ne.Error) {
	output := stack.NewTStack()
	operator := stack.NewTStack()

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type == t.NUMVALUE {
			output.Push(token)
		} else if token.Type == t.IDENTIFIER {
			if i+1 < len(tokens) && tokens[i+1].Type == t.LEFTPAREN {
				operator.Push(token)
			} else {
				output.Push(token)
			}
		} else if t.ContainsTT(token.Type, g.ARITHOPS) {
			o1 := token
			o2 := operator.Front()
			// assumes all operators are left-associative
			for o2.Type != t.LEFTPAREN && t.Precedence(o2.Type) >= t.Precedence(o1.Type) {
				top := operator.Pop()
				if top.Type == t.NONE {
					return output, ne.SyntaxError(line, "missing expression on right hand side")
				}

				output.Push(top)

				o2 = operator.Front()
			}

		} else if token.Type == t.LEFTPAREN {
			operator.Push(token)
		} else if token.Type == t.RIGHTPAREN {
			for operator.Front().Type != t.LEFTPAREN {
				if operator.Empty() {
					return output, ne.SyntaxError(line, "mismatched parentheses")
				}
				val := operator.Pop()
				output.Push(val)
			}
			if operator.Front().Type != t.LEFTPAREN {
				return output, ne.SyntaxError(line, "mismatched parentheses")
			}
			operator.Pop()
			front := operator.Front()
			if front.Type == t.IDENTIFIER {
				operator.Pop()
				output.Push(front)
			}
		}
	}

	for !operator.Empty() {
		front := operator.Front()
		if front.Type == t.LEFTPAREN || front.Type == t.RIGHTPAREN {
			return output, ne.SyntaxError(line, "mismatched parentheses")
		}

		operator.Pop()
		output.Push(front)
	}

	return output, ne.Null
}
