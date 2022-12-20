package main

import (
	"bytes"
	"io"
	"os"

	"github.com/adel-habib/calendar/calendar"
	"github.com/adel-habib/calendar/regions"
)

func main() {
	b := new(bytes.Buffer)
	err := calendar.NewCalendar(2022, regions.SN).SetResolution(1920.0, 1080.0).Write(b)
	if err != nil {
		return
	}
	f, err := os.Create("calendar.svg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, b)
}
