package chess

import (
	"errors"
	"goChess/chess/piece"
	. "goChess/util"
	"unicode"
)

func getPieceByRune(r rune, pos int) (piece.Piece, error) {
	color := piece.WHITE
	if unicode.IsLower(r) {
		color = piece.BLACK
	}
	switch unicode.ToUpper(r) {
	case 'P':
		return piece.NewPawn(color, pos), nil
	case 'N':
		return piece.NewKnight(color, pos), nil
	case 'B':
		return piece.NewBishop(color, pos), nil
	case 'R':
		return piece.NewRook(color, pos), nil
	case 'Q':
		return piece.NewQueen(color, pos), nil
	case 'K':
		return piece.NewKing(color, pos), nil
	default:
		return nil, errors.New("invalid piece")
	}
}

func IsValidFen(str string) bool {
	return true
	//panic("not implemented")
}

func FromFEN(fen string) (Board, error) {
	if !IsValidFen(fen) {
		return Board{}, errors.New("invalid FEN")
	}

	board := make([]piece.Piece, 64)
	i := 56
	for _, char := range fen {
		if unicode.IsDigit(char) {
			i += RuneToDigit(char) - 1
		} else if char == '/' {
			// TODO: go to next line
			i = (i/8 - 1) * 8
		} else if p, err := getPieceByRune(char, i); err == nil {
			board[i] = p
			if i%8 != 7 {
				i++
			}
		}
		// TODO: add game infos
	}

	//panic("not implemeted")
	return Board{board: board, selector: -1}, nil
}
