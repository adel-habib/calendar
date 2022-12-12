package main

import (
	"log"
	"os"
	"text/template"

	"github.com/adel-habib/calendar/calendar"
	"github.com/adel-habib/calendar/holidays"
)

func main() {

	tpl, err := template.ParseFiles("./static/templates/temp.tpl", "./static/css/styles.css", "./static/logo.svg")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create("out.svg")
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	hs := holidays.GermanHolidays(2022)
	monthGroups := calendar.CalendarDayGroups(2022, hs)
	header := calendar.Header(2022)
	cal := calendar.CalendarProps{Year: 2022, Header: header, MonthGroups: monthGroups, MonthsLabels: calendar.NewMonthsLabels()}
	tpl.Execute(f, cal)

}
