package main

import (
	"fmt"
	"goChess/chess"
	. "goChess/util"
)

func main() {
	//board := chess.NewBoard()
	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
	board, err := chess.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	ErrorCheck(err)
	fmt.Println(board.ToString())
	board.Turn()
}
