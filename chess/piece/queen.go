package piece

type Queen struct {
	color    Color
	position int
}

func (queen Queen) PosMoves(board []Piece) (int, []int) {
	panic("not implemented")
}

func (queen Queen) ToString() string {
	if queen.color == WHITE {
		return "Q"
	}
	return "q"
}

func (queen Queen) GetColor() Color {
	return queen.color
}

func NewQueen(c Color, pos int) *Queen {
	return &Queen{
		color:    c,
		position: pos,
	}
}
