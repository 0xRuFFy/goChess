package main

import (
	"fmt"
	"goChess/chess"
)

func main() {
	board := chess.NewBoard()
	fmt.Println(board.ToString())
}
