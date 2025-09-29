package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// Ctrl+F5
func main() {
	TypePrint()
	Variable()
}

func TypePrint() {
	fmt.Println(math.Floor(2.91))
	fmt.Println(math.Ceil(2.91))
	fmt.Println(math.Round(2.91))
	fmt.Println(math.Round(2.3))
	fmt.Println(strings.Title("this is head"))

	fmt.Println()

	fmt.Println("Kim\nInha\t\"\\") // C like
	fmt.Println('A', '가')          // Rune

	fmt.Println()

	fmt.Println(reflect.TypeOf(2.31))
	fmt.Println(reflect.TypeOf("this is head"))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(91))
}

func Variable() {
	// 1
	// var name string
	// var id int

	// name = "Kim Inha"
	// id = 1000

	// 2
	// var name string = "Kim Inha"
	// var id int = 1000

	// 3
	// var name = "Kim Inha"
	// var id = 1000

	// 4
	name := "Kim Inha"
	id := 1000

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
}
