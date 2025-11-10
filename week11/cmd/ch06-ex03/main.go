package main

import (
	"fmt"
	"os"
	"strconv"
)

func GetFloats() []float64 {
	var numbers []float64
	file := os.Args[1:]

	for _, num := range file {
		number, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return nil
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	weights := GetFloats()

	sum := 0.0
	for _, number := range weights {
		sum += number
	}
	weeks := float64(len(weights))
	fmt.Printf("평균 고기 소비량: %0.2f\n", sum/weeks)
}
