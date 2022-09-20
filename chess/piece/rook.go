package piece

type Rook struct {
	color    Color
	position int
}

func (rook Rook) PosMoves(board []Piece) (int, []int) {
	var moves []int
	var row, col, i int

	for i = rook.position/8 + 1; i < 8; i++ {
		row = 8 * i
		col = rook.position % 8
		if board[row+col] != nil {
			if board[row+col].GetColor() != rook.color {
				moves = append(moves, row+col)
			}
			break
		}
		moves = append(moves, row+col)
	}

	for i = rook.position/8 - 1; i >= 0; i-- {
		row = 8 * i
		col = rook.position % 8
		if board[row+col] != nil {
			if board[row+col].GetColor() != rook.color {
				moves = append(moves, row+col)
			}
			break
		}
		moves = append(moves, row+col)
	}

	for i = rook.position%8 + 1; i < 8; i++ {
		row = rook.position / 8 * 8
		if board[row+i] != nil {
			if board[row+i].GetColor() != rook.color {
				moves = append(moves, row+i)
			}
			break
		}
		moves = append(moves, row+i)
	}

	for i = rook.position%8 - 1; i >= 0; i-- {
		row = rook.position / 8 * 8
		if board[row+i] != nil {
			if board[row+i].GetColor() != rook.color {
				moves = append(moves, row+i)
			}
			break
		}
		moves = append(moves, row+i)
	}

	//panic("not implemented")
	return rook.position, moves
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
