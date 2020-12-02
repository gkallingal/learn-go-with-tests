package iteration

func Repeat(letter string) string {
	var repeat string

	for i := 0; i < 5; i++ {
		repeat = repeat + letter
	}
	return repeat
}
