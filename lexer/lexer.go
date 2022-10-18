package lexer

type Lexer struct {
	Text        string
	Position    Position
	currentChar string
}

func NewLexer(text string) Lexer {
	return Lexer{Text: text, Position: NewPosition(-1, 0, -1), currentChar: ""}
}

func (l *Lexer) Advance() {
	l.Position.Advance(l.currentChar)
	if l.Position.Pos < len(l.Text) {
		l.currentChar = string(l.Text[l.Position.Pos])
	} else {
		l.currentChar = ""
	}
}

func (l Lexer) Peek() string {
	return string(l.Text[l.Position.Pos+1])
}

func (l *Lexer) Tokenize() {
	tokens := []Token{}

	for l.currentChar != "" {
		sct := singleCharToken(l.currentChar)

		if sct.Type != NONE {
			if l.Peek() == "=" {
				switch sct.Type {
				case ASSIGN:
					sct.Type = EQUAL
					l.Advance()
				case GREATER:
					sct.Type = GREATER
					l.Advance()
				case LESS:
					sct.Type = LESS
					l.Advance()
				case NOT:
					sct.Type = NOTEQ
					l.Advance()
				}
			}

			tokens = append(tokens, sct)
		}

		l.Advance()
	}
}
