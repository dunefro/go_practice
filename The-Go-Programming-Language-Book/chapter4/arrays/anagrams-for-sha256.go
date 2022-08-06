package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	s := make([]rune, 0, len(c1))
	val1 := fmt.Sprintf("%x\n", c1)
	val2 := fmt.Sprintf("%x\n", c2)
	m1 := getMap(val1)
	m2 := getMap(val2)
	fmt.Println(m1)
	fmt.Println(m2)
	count := 0
	for k, v := range m1 {
		if v == m2[k] {
			count++
			s = append(s, k)
		}
	}
	fmt.Printf("Total matching count: %d\n", count)
	fmt.Println("matching characters ", s)
}

func getMap(s string) map[rune]int {
	m := map[rune]int{}
	for _, v := range s {
		m[v]++
	}
	return m
}
