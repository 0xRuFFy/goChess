package piece

type Rook struct {
	color    Color
	position int
}

func (rook Rook) PosMoves(board []Piece) (int, []int) {
	panic("not implemented")
}

func (rook Rook) ToString() string {
	if rook.color == WHITE {
		return "R"
	}
	return "r"
}

func (rook Rook) GetColor() Color {
	return rook.color
}

func NewRook(c Color, pos int) *Rook {
	return &Rook{
		color:    c,
		position: pos,
	}
}
