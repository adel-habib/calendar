package main

import (
	"github.com/adel-habib/calendar/calendar"
	"github.com/adel-habib/calendar/holidays"
)

func main() {
	err := calendar.NewCalendar(2022, holidays.SN).SetResolution(1280.0, 720.0).Export("cal.svg")
	if err != nil {
		return
	}

}
