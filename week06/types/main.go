package main

import (
	"fmt"
	"math"
	"strings"
)

// Ctrl+F5
func main() {
	fmt.Println(math.Floor(2.91))
	fmt.Println(math.Ceil(2.91))
	fmt.Println(math.Round(2.91))
	fmt.Println(math.Round(2.3))
	fmt.Println(strings.Title("this is head"))

	fmt.Println("Kim\nInha\t\"\\") // C like

	fmt.Println('A', '가')
}
