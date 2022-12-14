package main

import (
	"github.com/adel-habib/calendar/calendar"
	"github.com/adel-habib/calendar/holidays"
)

func main() {
	err := calendar.NewCalendar(2022, holidays.SN).SetResolution(1920.0, 1080.0).Export("cal.svg")
	if err != nil {
		return
	}
}
