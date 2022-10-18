package lexer

import (
	"fmt"

	ne "github.com/narutopig/neon-lang/error"
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

func (l *Lexer) Tokenize() ([]Token, ne.Error) {
	tokens := []Token{}

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
					return []Token{}, ne.EndOfLineError(l.Position.Line, "reached end of file while parsing string")

				} else if err.Error() == "eol" {
					return []Token{}, ne.EndOfLineError(l.Position.Line, "reached end of line while parsing string")
				}
			}

			tokens = append(tokens, val)
		} else if l.currentChar == "_" || isAlphaNumeric(l.currentChar) {
			tokens = append(tokens, l.addIdentifier())
		} else if l.currentChar == " " || l.currentChar == "\t" || l.currentChar == "\n" {
			l.Advance()
		} else if sct.Type != NONE {
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

			l.Advance()
		}
	}

	return tokens, ne.NoError()
}

func (l *Lexer) addNum() Token {
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

	return NewToken(NUMVALUE, res)
}

func (l *Lexer) addString() (Token, error) {
	res := ""

	l.Advance()
	for l.currentChar != "\"" {
		if l.currentChar == "\n" {
			return NewToken(NONE, ""), fmt.Errorf("eol")
		} else if l.currentChar == "" {
			return NewToken(NONE, ""), fmt.Errorf("eof")
		}

		res += l.currentChar
		l.Advance()
	}

	return NewToken(STRINGVALUE, res), nil
}

func (l *Lexer) addIdentifier() Token {
	res := l.currentChar

	l.Advance()

	for isAlphaNumeric(l.currentChar) {
		res += l.currentChar

		l.Advance()
	}

	return identifierMagic(NewToken(IDENTIFIER, res))
}
