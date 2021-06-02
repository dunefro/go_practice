package iteration

func Repeat(s string, n int) string {
	finalstr := ""
	if n == 0 {
		n = 5
	}
	for i := 0; i < n; i++ {
		finalstr = finalstr + s
	}
	return finalstr
}
