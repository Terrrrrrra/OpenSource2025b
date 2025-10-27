package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp
	fmt.Println(*first, *second)
}

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	amount, err := paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	// amount, err = paintNeeded(4.2, -3.0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%0.2f liters needed\n", amount)

	a, b := 10, 20
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

	fmt.Println("\n\n\n\n\n")

	fmt.Print("점수 입력 : ")
	score, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	status := ""
	if score >= 90 {
		status = "합격!"
	} else {
		status = "불합격"
	}
	fmt.Printf("%.2f점은 %v\n", score, status)
}
