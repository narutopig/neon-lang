package listener

import (
	"fmt"
	"strconv"

	"github.com/narutopig/neon-lang/parser"
	"github.com/narutopig/neon-lang/value"
)

// EnterString is called when production String is entered.
func (i *Interpreter) EnterString(ctx *parser.StringContext) {
	text := ctx.GetText()
	i.stack = append(i.stack, value.NewString(text[1:len(text)-1]))
}

// EnterIdentifier is called when production Identifier is entered.
func (i *Interpreter) EnterIdentifier(ctx *parser.IdentifierContext) {
	val, exists := i.Memory.Get(ctx.GetText())
	if !exists {
		i.panic(fmt.Sprintf("Variable %s not found", ctx.GetText()), ctx.GetStart().GetLine())
	}
	i.stack = append(i.stack, val)
}

// EnterInt is called when production Int is entered.
func (i *Interpreter) EnterInt(ctx *parser.IntContext) {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		i.panic("Error while converting number to string", ctx.GetStart().GetLine())
	}
	i.stack = append(i.stack, value.NewNumber(val))
}

// EnterBool is called when production Bool is entered.
func (i *Interpreter) EnterBool(ctx *parser.BoolContext) {
	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		i.panic("Error while converting boolean to string", ctx.GetStart().GetLine())
	}
	i.stack = append(i.stack, value.NewBoolean(val))
}
