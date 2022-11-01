package lexer

import (
	"fmt"

	ne "github.com/narutopig/neon-lang/lang/error"
	t "github.com/narutopig/neon-lang/lang/token"
)

// Lexer encapsulates the lexing process into a class
type Lexer struct {
	Text        string
	Position    Position
	currentChar string
}

// New returns a new Lexer from the provided content
func New(text string) Lexer {
	return Lexer{Text: text, Position: NewPosition(-1, 0, -1), currentChar: ""}
}

func (l *Lexer) advance() {
	l.Position.Advance(l.currentChar)
	if l.Position.Pos < len(l.Text) {
		l.currentChar = string(l.Text[l.Position.Pos])
	} else {
		l.currentChar = ""
	}
}

func (l Lexer) peek() string {
	if l.Position.Pos+1 >= len(l.Text) {
		return ""
	}
	return string(l.Text[l.Position.Pos+1])
}

// Tokenize returns the list of tokens from parsing the content
func (l *Lexer) Tokenize() ([]t.Token, ne.Error) {
	tokens := []t.Token{}

	l.advance()

	for l.currentChar != "" {
		sct := singleCharToken(l.currentChar)

		if isDigit(l.currentChar) {
			tokens = append(tokens, l.addNum())
		} else if l.currentChar == "\"" {
			val, err := l.addString()

			l.advance()
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
			l.advance()
		} else if sct.Type != t.NONE {
			if l.peek() == "=" {
				switch sct.Type {
				case t.ASSIGN:
					sct.Type = t.EQUAL
					l.advance()
				case t.GREATER:
					sct.Type = t.GREATER
					l.advance()
				case t.LESS:
					sct.Type = t.LESS
					l.advance()
				case t.NOT:
					sct.Type = t.NOTEQ
					l.advance()
				}
			}

			tokens = append(tokens, sct)

			l.advance()
		}
	}

	return tokens, ne.Null
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
		l.advance()
	}

	return t.New(t.NUMVALUE, res)
}

func (l *Lexer) addString() (t.Token, error) {
	res := ""

	l.advance()
	for l.currentChar != "\"" {
		if l.currentChar == "\n" {
			return t.NoneToken, fmt.Errorf("eol")
		} else if l.currentChar == "" {
			return t.NoneToken, fmt.Errorf("eof")
		}

		res += l.currentChar
		l.advance()
	}

	return t.New(t.STRINGVALUE, res), nil
}

func (l *Lexer) addIdentifier() t.Token {
	res := l.currentChar

	l.advance()

	for isAlphaNumeric(l.currentChar) {
		res += l.currentChar

		l.advance()
	}

	return identifierMagic(t.New(t.IDENTIFIER, res))
}
