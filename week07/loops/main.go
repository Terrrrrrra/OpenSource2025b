package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

func main() {
	Casting()
	fmt.Println("\n------------------\n")
	TimeEX()
	fmt.Println("\n------------------\n")
	Swap()
	fmt.Println("\n------------------\n")
	InOut()
}

func Casting() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("면적은", int(length)*width)
	fmt.Println("length > width", int(length) > width)
	fmt.Println("원본", reflect.TypeOf(length))
}

func TimeEX() {
	var now time.Time = time.Now()
	var month time.Month = now.Month() //month := now.Month()
	var day int = now.Day()
	fmt.Println(month, day)
}

func Swap() {
	broken := "G# r#ck!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	univ := "Go$ inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}

func InOut() {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	log.Fatal(err) // report err, and exit process
	fmt.Println(i)
}
