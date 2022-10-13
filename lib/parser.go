package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse returns list of tokens and syntax errors if found
func Parse(content string) ([]Token, error) {
	tokens := make([]Token, 0)

	chars := strings.Split(content, "")

	currBlock := ""
	currType := NONE

	// start := 0 (used for debugging/error messages later)
	line := 1

	for _, char := range chars {
		singleCharToken := singleCharToken(char)

		if char == "\n" {
			tokens = append(tokens, NewToken(NEWLINE, ""))
			if currType == STRINGVALUE {
				// string not ended
				fmt.Println(currType)

				return tokens, fmt.Errorf("Reached end of line while parsing string on line %d", line)
			}
			line++
		} else if singleCharToken.Type != NONE {
			tokens = append(tokens, NewToken(currType, currBlock))
			tokens = append(tokens, singleCharToken)
			currBlock = ""
			currType = NONE
		} else if currType == STRINGVALUE {
			// in string
			if char == "\"" {
				tokens = append(tokens, NewToken(STRINGVALUE, currBlock))
				currBlock = ""
				currType = NONE
			} else {
				currBlock += char
			}
		} else if currType == NUMVALUE {
			if !isNum(char) {
				fmt.Println(currType)

				return tokens, fmt.Errorf("Unexpected character found while parsing number on line %d", line)
			} else {
				currBlock += char
			}
		} else {
			// not in string
			if char == "\"" {
				currType = STRINGVALUE
			} else if char == " " {
				if currType != NONE {
					_, err := strconv.ParseBool(currBlock)
					if err != nil {
						// not bool
						tokens = append(tokens, NewToken(VARIABLE, currBlock))
					} else {
						tokens = append(tokens, NewToken(BOOLEANVALUE, currBlock))
					}
					currBlock = ""
					currType = NONE
				}
			} else if currType == NONE {
				if isNum(char) {
					currType = NUMVALUE
				} else if isAlpha(char) {
					currType = UNKNOWNVALUE
				} else {
					fmt.Println(currType)
					return tokens, fmt.Errorf("Unexpected character %s on line %d", char, line)
				}
				currBlock += char
			} else if currType == UNKNOWNVALUE {
				if isNum(char) || isAlpha(char) || char == "_" {
					currBlock += char
				}
			} else {
				return tokens, fmt.Errorf("Unexpected character %s on line %d", char, line)
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
	return strings.Contains("qwertyuiopasdfghjklzxcvbnm", str)
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
	case ",":
		return NewToken(COMMA, "")
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
