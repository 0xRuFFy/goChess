package piece

type Color int

const (
	NONE Color = iota
	WHITE
	BLACK
)

type Piece interface {
	ToString() string
	PosMoves([]Piece) (int, []int)
	GetColor() Color
}
