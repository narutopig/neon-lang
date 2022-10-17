package lexer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/narutopig/go-utils"
)

// Lex returns list of tokens and syntax errors if found
func Lex(content string) ([]Token, error) {
	tokens := make([]Token, 0)

	chars := strings.Split(content, "")

	currBlock := ""
	currType := NONE

	// start := 0 (used for debugging/error messages later)
	line := 1

	for i := 0; i < len(chars); i++ {
		char := chars[i]
		sct := singleCharToken(char)

		if currType == STRINGVALUE {
			// in string
			if char == "\"" {
				// end string
				tokens = append(tokens, NewToken(STRINGVALUE, currBlock))
				currBlock = ""
				currType = NONE
			} else if char == "\n" {
				if i+1 >= len(chars) {
					return tokens, fmt.Errorf("reached end of file while parsing string on line %d", line)
				}
				return tokens, fmt.Errorf("reached end of line while parsing string on line %d", line)

			} else {
				// part of string
				currBlock += char
			}

			continue // avoid else if nesting even though theres still going to be a lot of them
		}
		if char == "\n" {
			if currType != NONE {
				tokens = append(tokens, NewToken(currType, currBlock))
				currType = NONE
				currBlock = ""
			}

			line++
			continue // none of the other cases match
		}
		if sct.Type != NONE {
			// the char is a single char token
			if sct.Type == PERIOD {
				if currType == NUMVALUE {
					if strings.Contains(currBlock, ".") {
						return tokens, fmt.Errorf("unexpected token . on line %d", line)
					}
					currBlock += "."
					continue
				}
			} else {
				if currType == UNKNOWNVALUE {
					tokens = append(tokens, unknown(currBlock))
				} else if currType != NONE {
					tokens = append(tokens, NewToken(currType, currBlock))
				}
				if i+1 < len(chars) && chars[i+1] == "=" {
					switch sct.Type {
					case ASSIGN:
						sct.Type = EQUAL
						i++
					case LESS:
						sct.Type = LESSEQ
						i++
					case GREATER:
						sct.Type = GREATEREQ
						i++
					}
				}
			}
			tokens = append(tokens, sct)
			currBlock = ""
			currType = NONE
		} else if currType == NUMVALUE {
			if !isNum(char) {
				// not digit in num
				return tokens, fmt.Errorf("unexpected character found while parsing number on line %d", line)
			}
			currBlock += char
		} else {
			// not in string
			if char == "\"" {
				// beginning of string
				currType = STRINGVALUE
			} else if char == " " {
				// space
				if currType != NONE {
					// could be bool or identifier
					tokens = append(tokens, unknown(currBlock))
					currBlock = ""
					currType = NONE
				}
			} else if currType == NONE {
				if isNum(char) {
					currType = NUMVALUE
				} else if isAlpha(char) {
					currType = UNKNOWNVALUE
				} else {
					return tokens, fmt.Errorf("unexpected character %s on line %d", char, line)
				}
				currBlock += char
			} else if currType == UNKNOWNVALUE {
				if isNum(char) || isAlpha(char) || char == "_" {
					currBlock += char
				}
			} else {
				return tokens, fmt.Errorf("unexpected character %s on line %d", char, line)
			}
		}
	}

	return tokens, nil
}

func isNum(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func isAlpha(str string) bool {
	// could be optimized with binary search
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	return utils.BinarySearchS(strings.ToLower(str), letters) != -1
}

func unknown(currBlock string) Token {
	_, err := strconv.ParseBool(currBlock)
	typeTok := typeToken(currBlock)
	if err != nil {
		// not bool
		if typeTok.Type != NONE {
			return typeTok
		}
		return NewToken(IDENTIFIER, currBlock)
	}
	return NewToken(BOOLEANVALUE, currBlock)
}

func typeToken(str string) Token {
	switch str {
	case "int":
		return NewToken(INTTYPE, "")
	case "string":
		return NewToken(STRINGTYPE, "")
	case "float":
		return NewToken(FLOATTYPE, "")
	case "bool":
		return NewToken(BOOLEANTYPE, "")
	}
	return NewToken(NONE, "")
}

func singleCharToken(str string) Token {
	switch str {
	case "(":
		return NewToken(LEFTPAREN, "")
	case ")":
		return NewToken(RIGHTPAREN, "")
	case "[":
		return NewToken(LEFTBRACKET, "")
	case "]":
		return NewToken(RIGHTBRACKET, "")
	case "{":
		return NewToken(LEFTCURLY, "")
	case "}":
		return NewToken(RIGHTCURLY, "")
	case ",":
		return NewToken(COMMA, "")
	case ".":
		return NewToken(PERIOD, "")
	case ";":
		return NewToken(SEMI, "")
	case "+":
		return NewToken(ADD, "")
	case "-":
		return NewToken(SUBTRACT, "")
	case "*":
		return NewToken(MULTIPLY, "")
	case "/":
		return NewToken(DIVIDE, "")
	case "%":
		return NewToken(MODULUS, "")
	case "=":
		return NewToken(ASSIGN, "")
	case "<":
		return NewToken(LESS, "")
	case ">":
		return NewToken(GREATER, "")
	case "!":
		return NewToken(NOT, "")
	case "&":
		return NewToken(AND, "")
	case "|":
		return NewToken(OR, "")
	}
	return NewToken(NONE, "")
}
