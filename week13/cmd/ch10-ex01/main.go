package main

import (
	"fmt"
	"log"
	"main/pkg/calendar"
)

func main() {
	today := calendar.Date{}
	// date.year = 2025

	err := today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today)

	fmt.Println(today.Year(), "년", today.Month(), "월", today.Day(), "일")
}
