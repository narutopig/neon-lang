package expressions

import (
	neonerror "github.com/narutopig/neon-lang/error"
	"github.com/narutopig/neon-lang/grammar"
	"github.com/narutopig/neon-lang/stack"
	tokens "github.com/narutopig/neon-lang/token"
)

// Eval evaluates the TokenStack returned from Shunt()
func Eval(ts stack.Stack) (float64, neonerror.Error) {
	rev := stack.Reversed(ts)

	for !rev.Empty() {
		token := rev.Front()

		if !tokens.ContainsTT(token.Type, grammar.ARITHOPS) {
			
		}
	}
	return 0, neonerror.Null
}
