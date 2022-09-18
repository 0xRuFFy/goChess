package piece

type Color int

const (
	WHITE Color = iota
	BLACK
	NONE
)

type Piece interface {
	ToString() string
	PosMoves([]Piece) (int, []int)
	GetColor() Color
}
