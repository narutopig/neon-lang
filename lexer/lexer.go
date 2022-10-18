package lexer

import (
	"fmt"

	ne "github.com/narutopig/neon-lang/error"
	t "github.com/narutopig/neon-lang/token"
)

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
	if l.Position.Pos+1 >= len(l.Text) {
		return ""
	}
	return string(l.Text[l.Position.Pos+1])
}

func (l *Lexer) Tokenize() ([]t.Token, ne.Error) {
	tokens := []t.Token{}

	l.Advance()

	for l.currentChar != "" {
		sct := singleCharToken(l.currentChar)

		if isDigit(l.currentChar) {
			tokens = append(tokens, l.addNum())
		} else if l.currentChar == "\"" {
			val, err := l.addString()

			l.Advance()
			if err != nil {
				if err.Error() == "eof" {
					return []t.Token{}, ne.EndOfLineError(l.Position.Line, "reached end of file while parsing string")
				} else if err.Error() == "eol" {
					return []t.Token{}, ne.EndOfLineError(l.Position.Line, "reached end of line while parsing string")
				}
			}

			tokens = append(tokens, val)
		} else if l.currentChar == "_" || isAlphaNumeric(l.currentChar) {
			tokens = append(tokens, l.addIdentifier())
		} else if l.currentChar == " " || l.currentChar == "\t" || l.currentChar == "\n" {
			l.Advance()
		} else if sct.Type != t.NONE {
			if l.Peek() == "=" {
				switch sct.Type {
				case t.ASSIGN:
					sct.Type = t.EQUAL
					l.Advance()
				case t.GREATER:
					sct.Type = t.GREATER
					l.Advance()
				case t.LESS:
					sct.Type = t.LESS
					l.Advance()
				case t.NOT:
					sct.Type = t.NOTEQ
					l.Advance()
				}
			}

			tokens = append(tokens, sct)

			l.Advance()
		}
	}

	return tokens, ne.NoError()
}

func (l *Lexer) addNum() t.Token {
	res := ""
	dotc := 0

	for l.currentChar != "" && (isDigit(l.currentChar) || l.currentChar == ".") {
		if l.currentChar == "." {
			if dotc == 1 {
				break
			}
			dotc++
			res += "."
		} else {
			res += l.currentChar
		}
		l.Advance()
	}

	return t.NewToken(t.NUMVALUE, res)
}

func (l *Lexer) addString() (t.Token, error) {
	res := ""

	l.Advance()
	for l.currentChar != "\"" {
		if l.currentChar == "\n" {
			return t.NewToken(t.NONE, ""), fmt.Errorf("eol")
		} else if l.currentChar == "" {
			return t.NewToken(t.NONE, ""), fmt.Errorf("eof")
		}

		res += l.currentChar
		l.Advance()
	}

	return t.NewToken(t.STRINGVALUE, res), nil
}

func (l *Lexer) addIdentifier() t.Token {
	res := l.currentChar

	l.Advance()

	for isAlphaNumeric(l.currentChar) {
		res += l.currentChar

		l.Advance()
	}

	return identifierMagic(t.NewToken(t.IDENTIFIER, res))
}
