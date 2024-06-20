package main

import (
	"fmt"
	"time"
)

func main() {
	var login, logout string
	fmt.Print("Enter login time (HH:MM) and logout time (HH:MM): ")
	fmt.Scanf("%s %s", &login, &logout)

	// Define the layout for parsing time
	layout := "15:04"

	// Parse the login and logout times
	loginTime, err := time.Parse(layout, login)
	if err != nil {
		fmt.Println("Error parsing login time:", err)
		return
	}
	logoutTime, err := time.Parse(layout, logout)
	if err != nil {
		fmt.Println("Error parsing logout time:", err)
		return
	}

	// Handle case where logout is before login (spans midnight)
	if logoutTime.Before(loginTime) {
		logoutTime = logoutTime.Add(24 * time.Hour)
	}

	// Calculate the number of full 15-minute rounds
	fullRounds := countFullRounds(loginTime, logoutTime)
	fmt.Println("Full rounds played:", fullRounds)
}

func countFullRounds(start, end time.Time) int {
	// Start the first round at the next 15-minute mark after login
	roundedStart := start.Add(time.Minute * time.Duration(15-start.Minute()%15)).Truncate(time.Minute * 15)
	if roundedStart.Before(start) {
		roundedStart = roundedStart.Add(time.Minute * 15)
	}

	// Initialize the count of full rounds
	fullRounds := 0

	// Increment the start time by 15 minutes and count full rounds until it surpasses the end time
	for {
		roundedStart = roundedStart.Add(time.Minute * 15)
		if roundedStart.After(end) || roundedStart.Equal(end) {
			break
		}
		fullRounds++
	}

	return fullRounds
}
