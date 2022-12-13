package calendar

import (
	"embed"
	"github.com/adel-habib/calendar/holidays"
	"os"
	"text/template"
)

// embed templates in binary
//go:embed static/*
var efs embed.FS

type Calendar interface {
	Export(name string) error
	SetResolution(width float64, height float64) *calendar
}

var tpl *template.Template

func init() {
	temp, err := template.ParseFS(efs, "static/temp.tpl", "static/styles.css", "static/logo.svg")
	if err != nil {
		panic(err)
	}
	tpl = temp
}

func NewCalendar(year uint, region holidays.Region) *calendar {
	cal := &calendar{year: int(year), region: region, geometry: newGeometry(1920.0, 1080.0)}
	cal.hs = holidays.GetHolidaysList(region, year, year+1)
	return cal
}

func (c *calendar) SetResolution(width float64, height float64) *calendar {
	c.geometry = newGeometry(width, height)
	return c
}

func (c *calendar) Export(name string) error {
	obj := newCalendarObject(int(c.year), c.hs, c.geometry)
	f, err := os.Create(name)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return err
	}
	err = tpl.Execute(f, obj)
	if err != nil {
		return err
	}
	return nil
}
