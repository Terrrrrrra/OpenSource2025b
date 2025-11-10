package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	weights, err := GetFloats("meatWeights.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for _, number := range weights {
		sum += number
	}
	weeks := float64(len(weights))
	fmt.Printf("평균 고기 소비량: %0.2f\n", sum/weeks)
}
