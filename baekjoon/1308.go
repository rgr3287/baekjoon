package main

import (
	"fmt"
	"time"
)

func main() {
	var todayYear, todayMonth, todayDay, dDayYear, dDayMonth, dDayDay int
	fmt.Scanf("%d %d %d", &todayYear, &todayMonth, &todayDay)
	fmt.Scanf("%d %d %d", &dDayYear, &dDayMonth, &dDayDay)

	today := time.Date(todayYear, time.Month(todayMonth), todayDay, 0, 0, 0, 0, time.UTC)
	dDay := time.Date(dDayYear, time.Month(dDayMonth), dDayDay, 0, 0, 0, 0, time.UTC)

	if today.Month() == 2 && today.Day() == 29 {
		// Adjust today to March 1st for calculation purposes
		today = today.AddDate(0, 0, 1)
	}

	days := int(dDay.Sub(today).Hours() / 24)

	// Check if the camp lasts more than 1000 years
	if dDayYear-todayYear >= 1000 {
		fmt.Println("gg")
	} else {
		fmt.Printf("D-%d", days)
	}
}
