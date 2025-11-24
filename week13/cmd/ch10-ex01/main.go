package main

import (
	"fmt"
	"log"
	"main/pkg/calendar"
)

func main() {
	event := calendar.Event{}
	// date.year = 2025
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
	fmt.Println(event.Title())
	fmt.Println(event.Year(), "년", event.Month(), "월", event.Day(), "일")
}
