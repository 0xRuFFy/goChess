package chess

import (
	"errors"
	"fmt"
	"goChess/chess/piece"
	"goChess/util"
	"unicode"
)

type Board struct {
	board     []piece.Piece
	selector  int
	selection map[int]bool
}

func NewBoard() Board {
	return Board{
		board: newBoard(),
	}
}

func (b *Board) ToString() string {
	// TODO: remove after testing
	b.SetSelection(b.board[1].PosMoves(b.board))

	boardStr := ""
	for i := 0; i < len(b.board); i++ {
		if i%8 == 0 {
			boardStr += fmt.Sprintf(" %d ", i/8+1)
		}

		symbol := "*"
		if b.board[i] != nil {
			symbol = b.board[i].ToString()
		}
		color := "\x1b[36m"
		if i == b.selector {
			color = "\x1b[33m"
		} else if _, ok := b.selection[i]; !ok {
			color = "\x1b[0m"
		}
		boardStr += fmt.Sprintf(" %s%2s%s ", color, symbol, "\x1b[0m")

		if i%8 == 7 {
			boardStr += "\n"
		}
	}
	boardStr += "     A   B   C   D   E   F   G   H"
	return boardStr
}

func (b *Board) SetSelection(selector int, selection []int) {
	m := make(map[int]bool)
	for _, s := range selection {
		m[s] = true
	}
	b.selector = selector
	b.selection = m
}

func IsValidFieldIndex(index int) bool {
	return index >= 0 && index < 64
}

func FromFEN(fen string) error {
	if !IsValidFen(fen) {
		return errors.New("invalid FEN")
	}

	i := 0
	for j, char := range fen {
		if unicode.IsDigit(char) {
			i += util.RuneToDigit(char)
		} else if char == '/' {
			// TODO: go to next line
		}
		// TODO: add pieces
		// TODO: add game infos
	}

	panic("not implemeted")
	return nil
}

func newBoard() []piece.Piece {
	b := make([]piece.Piece, 64)

	b[0] = piece.NewRook(piece.WHITE, 0)
	b[1] = piece.NewKnight(piece.WHITE, 1)
	b[2] = piece.NewBishop(piece.WHITE, 2)
	b[3] = piece.NewQueen(piece.WHITE, 3)
	b[4] = piece.NewKing(piece.WHITE, 4)
	b[5] = piece.NewBishop(piece.WHITE, 5)
	b[6] = piece.NewKnight(piece.WHITE, 6)
	b[7] = piece.NewRook(piece.WHITE, 7)
	for i := 0; i < 8; i++ {
		b[8+i] = piece.NewPawn(piece.WHITE, 8+i)
	}

	for i := 0; i < 8; i++ {
		b[48+i] = piece.NewPawn(piece.BLACK, 48+i)
	}
	b[56] = piece.NewRook(piece.BLACK, 0)
	b[57] = piece.NewKnight(piece.BLACK, 1)
	b[58] = piece.NewBishop(piece.BLACK, 2)
	b[59] = piece.NewQueen(piece.BLACK, 3)
	b[60] = piece.NewKing(piece.BLACK, 4)
	b[61] = piece.NewBishop(piece.BLACK, 5)
	b[62] = piece.NewKnight(piece.BLACK, 6)
	b[63] = piece.NewRook(piece.BLACK, 7)
	return b
}
