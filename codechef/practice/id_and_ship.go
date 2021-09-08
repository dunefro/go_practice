package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	for i := int64(0); i < T; i++ {
		scanner.Scan()
		str := scanner.Text()
		switch str {
		case "B", "b":
			fmt.Println("BattleShip")
		case "C", "c":
			fmt.Println("Cruiser")
		case "D", "d":
			fmt.Println("Destroyer")
		case "F", "f":
			fmt.Println("Frigate")
		}
	}
}
