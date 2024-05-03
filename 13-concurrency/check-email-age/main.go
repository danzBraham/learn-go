package main

import (
	"time"
)

func main() {
	emails := [3]Email{
		{"Hello there!", time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"Hello peace!", time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC)},
		{"Hello dark!", time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC)},
	}

	checkEmailAge(emails)
}

type Email struct {
	Body string
	Date time.Time
}

func sendIsOld(isOldChan chan<- bool, emails [3]Email) {
	go func() {
		for _, em := range emails {
			if em.Date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()
}

func checkEmailAge(emails [3]Email) [3]bool {
	isOldChan := make(chan bool)
	sendIsOld(isOldChan, emails)
	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}
