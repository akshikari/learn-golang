package iteration

// Repeat takes a character string and a integer n and will generate a new string with the character repeated n times
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
