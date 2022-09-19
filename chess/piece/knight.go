package piece

type Knight struct {
	color    Color
	position int
}

func (knight Knight) PosMoves(board []Piece) (int, []int) {
	var moves []int
	offsets := []int{6, 15, 17, 10, -6, -15, -17, -10}
	
	for _, i := range offsets {
		pos := knight.position + i

		if pos > 0 && pos < len(board) && (board[pos] == nil || board[pos].GetColor() != knight.color) {
			moves = append(moves, pos)
		}
	}

	return knight.position, moves
}

func (knight Knight) ToString() string {
	if knight.color == WHITE {
		return "S"
	}
	return "s"
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
