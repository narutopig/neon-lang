package lexer

// Represents current position in the file

type Position struct {
	Pos  int
	Line int
	Col  int
}

func NewPosition(pos int, line int, col int) Position {
	return Position{Pos: pos, Line: line, Col: col}
}

func (p *Position) Advance(char string) *Position {
	p.Pos += 1
	p.Col += 1

	if char == "\n" {
		p.Line += 1
		p.Col = 0
	}

	return p
}
