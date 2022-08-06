package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(convert([]int{1, 2, 3}))
}

func convert(x []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range x {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
