package iteration

func Repeat(char string, no int) string {
	var repeated string
	for i := 0; i < no; i++ {
		repeated += char
	}
	return repeated
}
