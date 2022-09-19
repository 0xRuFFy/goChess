package piece


type King struct {
	color    Color
	position int
}

func (king King) PosMoves(board []Piece) (int, []int) {
	var moves []int

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			pos := king.position+8*i+j
			if pos >= 0 && pos < len(board) {
				p := board[pos]
				if p == nil || p.GetColor() != king.color {
					moves = append(moves, pos)
				}
			}
		}
	}
	
	return king.position, moves
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
