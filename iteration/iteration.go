package iteration

const (
	repeatCount = 5
)

//Repeat function will take any character and repeat it 5 times
func Repeat(letter string) string {
	var repeat string

	for i := 0; i < repeatCount; i++ {
		repeat += letter
	}
	return repeat
}
