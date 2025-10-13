package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
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
	//InOut()
	fmt.Println("\n------------------\n")
	Shadowing()
	fmt.Println("\n------------------\n")
	ConvStrToFloat()
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
	if err != nil {
		log.Fatal(err) // report err, and exit process
	}
	fmt.Println(i)
}

func Shadowing() {
	// var int int = 99
	// var b int = 8
	// fmt.Println(int, b)

	// var fmt string = "inha"
	// fmt.Println(fmt)
}

func ConvStrToFloat() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter score : ")
	i, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err) // report err, and exit process
	}

	i = strings.TrimSpace(i)
	score, err := strconv.ParseFloat(i, 64)
	if err != nil {
		log.Fatal(err) // report err, and exit process
	}

	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}

	fmt.Println(score, status)
}
