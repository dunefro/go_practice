package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your first word: ")
	scan1 := bufio.NewScanner(os.Stdin)
	scan1.Scan()
	text1 := scan1.Text()
	fmt.Print("Enter your second word: ")
	scan1.Scan()
	text2 := scan1.Text()
	m1 := getMap(text1)
	m2 := getMap(text2)
	fmt.Println(m1)
	fmt.Println(m2)
	for b, _ := range m1 {
		if m1[b] != m2[b] {
			fmt.Println("Not anagrams")
			return
		}
	}
	fmt.Println("anagrams")
}
func getMap(s string) map[byte]int {
	m := make(map[byte]int)
	b := []byte(s)
	for _, v := range b {
		if v != ' ' {
			m[v]++
		}
	}
	return m
}
