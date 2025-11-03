package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arrayBool [3]bool = [3]bool{true, false, true}
	// arrayBool := [3]bool{true, false, true} 이것도 됨
	var arrayInt [3]int
	fmt.Println(arrayBool[1])

	fmt.Println(arrayBool[1]) // zero value (flase)
	arrayInt[1]++
	arrayInt[1] = arrayInt[1] + 1
	fmt.Println(arrayInt[1]) // zero value + 2

	arrayInt2 := [3]int{-9, 11, 7}
	for i := 0; i < 3; i++ {
		fmt.Println(i, arrayInt2[i])
		fmt.Println(i, arrayBool[i])
	}
	fmt.Printf("%#v\n", arrayBool)
	fmt.Printf("%#v\n", arrayInt2)
	fmt.Println(reflect.TypeOf(arrayInt2))

	for i, number := range arrayInt2 {
		fmt.Println(i, number)
	}
}
