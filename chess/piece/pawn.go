package piece

type Pawn struct {
	color    Color
	position int
	moved    bool
}

func (pawn Pawn) PosMoves(board []Piece) (int, []int) {
	var moves []int
	dir := 1
	if pawn.color == BLACK {
		dir = -1
	}

	spaceToEdge := pawn.spaceToEdge()
	if spaceToEdge > 0 && board[pawn.position+8*dir] == nil {
		moves = append(moves, pawn.position+8*dir)
		if spaceToEdge > 1 && !pawn.moved && board[pawn.position+dir*16] == nil {
			moves = append(moves, pawn.position+dir*16)
		}
	}

	if pawn.position%8 != 0 &&
		board[pawn.position+8*dir-1] != nil &&
		board[pawn.position+8*dir-1].GetColor() != pawn.color {
		moves = append(moves, pawn.position+8*dir-1)
	}

	if pawn.position%8 != 7 &&
		board[pawn.position+8*dir+1] != nil &&
		board[pawn.position+8*dir+1].GetColor() != pawn.color {
		moves = append(moves, pawn.position+8*dir+1)
	}

	return pawn.position, moves
}

func (pawn Pawn) spaceToEdge() int {
	if pawn.color == WHITE {
		return 7 - pawn.position/8
	}
	return pawn.position / 8
}

func (pawn Pawn) ToString() string {
	if pawn.color == WHITE {
		return "P"
	}
	return "p"
}

func (pawn Pawn) GetColor() Color {
	return pawn.color
}

func NewPawn(c Color, pos int) *Pawn {
	return &Pawn{
		color:    c,
		position: pos,
	}
}
