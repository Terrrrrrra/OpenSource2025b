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
	fmt.Println("\n---------------\n")
	Variable()
	fmt.Println("\n---------------\n")
	Variable2()
	fmt.Println("\n---------------\n")
	ZeroValues()
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

func Variable2() {
	//var name string
	name := "Inha"
	var id int
	id = 1000

	var gpa float32 = 3.99

	fmt.Println(id, reflect.TypeOf(id))
	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(gpa, reflect.TypeOf(gpa))
}

func ZeroValues() {
	var f64 float64
	var i16 int16
	var t bool
	var s string
	var i int

	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(i16, reflect.TypeOf(i16))
	fmt.Println(t, reflect.TypeOf(t))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))
}
