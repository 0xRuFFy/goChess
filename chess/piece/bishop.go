package piece

type Bishop struct {
	color    Color
	position int
}

func (bishop Bishop) PosMoves(board []Piece) (int, []int) {
	panic("not implemented")
}

func (bishop Bishop) ToString() string {
	if bishop.color == WHITE {
		return "B"
	}
	return "b"
}

func (bishop Bishop) GetColor() Color {
	return bishop.color
}

func NewBishop(c Color, pos int) *Bishop {
	return &Bishop{
		color:    c,
		position: pos,
	}
}
