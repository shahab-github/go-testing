package iteration

const repeatCount = 5

func Repeat(s string) string {
	var result string
	for i := 0; i < repeatCount; i++ {
		result += s
	}
	return result
}
