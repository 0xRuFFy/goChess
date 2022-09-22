package util

func RuneToDigit(c rune) int {
	return int(c) - '0'
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
