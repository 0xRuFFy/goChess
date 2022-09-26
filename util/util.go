package util

import "unicode"

func RuneToDigit(c rune) int {
	return int(c) - '0'
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func IsValidSquare(s string) bool {
	split := []rune(s)
	if len(split) != 2 {
		return false
	}
	upper := unicode.ToUpper(split[0])
	if upper < 'A' || upper > 'H' {
		return false
	}
	if !unicode.IsDigit(split[1]) {
		return false
	}

	if split[1] < '1' || split[1] > '8' {
		return false
	}

	return true
}

func ConvertChessSquareToIndex(square string) int {
	return 1
	panic("not implemented")
}
