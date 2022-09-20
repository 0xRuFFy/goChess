package piece

type Knight struct {
	color    Color
	position int
}

func (knight Knight) PosMoves(board []Piece) (int, []int) {
	panic("not implemented")
}

func (knight Knight) ToString() string {
	if knight.color == WHITE {
		return "K"
	}
	return "k"
}

func (knight Knight) GetColor() Color {
	return knight.color
}

func NewKnight(c Color, pos int) *Knight {
	return &Knight{
		color:    c,
		position: pos,
	}
}
