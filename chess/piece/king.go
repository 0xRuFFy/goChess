package piece

type King struct {
	color    Color
	position int
}

func (king King) PosMoves(board []Piece) (int, []int) {
	panic("not implemented")
}

func (king King) ToString() string {
	if king.color == WHITE {
		return "K"
	}
	return "k"
}

func (king King) GetColor() Color {
	return king.color
}

func NewKing(c Color, pos int) *King {
	return &King{
		color:    c,
		position: pos,
	}
}
