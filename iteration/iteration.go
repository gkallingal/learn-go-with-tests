package iteration

import "strings"

const (
	repeatCount = 5
)

//Repeat function will take any character and repeat it N number of times based on count parameter
func Repeat(letter string, count int) string {
	repeat := strings.Repeat(letter, count)
	return repeat
}
