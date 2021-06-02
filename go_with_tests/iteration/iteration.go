package iteration

func Repeat(s string) string {
	finalstr := ""
	for i := 0; i < 5; i++ {
		finalstr = finalstr + s
	}
	return finalstr
}
