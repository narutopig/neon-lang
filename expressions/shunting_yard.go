package expressions

import (
	ne "github.com/narutopig/neon-lang/error"
	g "github.com/narutopig/neon-lang/grammar"
	t "github.com/narutopig/neon-lang/token"
	"github.com/narutopig/neon-lang/util"
)

// Shunt implements the shunting-yard algorithm from https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func Shunt(tokens []t.Token) (util.TokenStack, ne.Error) {
	output := util.NewTStack()
	operator := util.NewTStack()

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
		} else if token.Type == t.LEFTPAREN {
			operator.Push(token)
		} else if token.Type == t.RIGHTPAREN {
			for operator.Front().Type != t.LEFTPAREN {
				if operator.Empty() {
					// assertion
					val := operator.Pop()
					output.Push(val)
				}
			}
			if operator.Front().Type != t.LEFTPAREN {
				// assertion
			}
		}
	}

	return output, ne.Null
}
