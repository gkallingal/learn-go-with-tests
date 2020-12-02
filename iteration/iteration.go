package iteration

const (
	repeatCount = 5
)

//Repeat function will take any character and repeat it N number of times based on count parameter
func Repeat(letter string, count int) string {
	var repeat string

	for i := 0; i < count; i++ {
		repeat += letter
	}
	return repeat
}
