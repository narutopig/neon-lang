package token

type TokenType byte

//go:generate stringer -type=TokenType
const (
	NumType TokenType = iota
	StrType
	BoolType
	VoidType
	True
	False
	Def
	Return
	If
	Elif
	While
	Equality
	NotEqual
	LessEq
	GreaterEq
	Less
	Greater

	// long
	StrLiteral
	NumLiteral
	Id

	// short
	AddSub
	MulDivMod
	Comma
	Semi
	Lp
	Rp
	Lb
	Rb
	Lc
	Rc
	Equal
	Bang
)
