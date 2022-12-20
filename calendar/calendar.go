package calendar

import (
	"embed"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/adel-habib/calendar/holidays"
	"github.com/adel-habib/calendar/regions"
)

// embed templates in binary
//go:embed static/*
var efs embed.FS

type Calendar interface {
	Write(w io.Writer) error
	Export(name string) error
	SetResolution(width float64, height float64) *calendar
}

var tpl *template.Template

// template functions
var funcMap = template.FuncMap{
	"RoundFloat": func(value float64) string {
		return fmt.Sprintf("%.2f", value)
	},
	"ToInt": func(value float64) string {
		return fmt.Sprintf("%d", int(value))
	},
	"halve": func(value float64) string {
		return fmt.Sprintf("%d", int(value/2))
	},
	"Inc": func(value int) int {
		return value + 1
	},
	"Double": func(value float64) string {
		return fmt.Sprintf("%.2f", value*2)
	},
}

func init() {
	// parse templates on startup
	temp, err := template.New("temp.tpl").Funcs(funcMap).ParseFS(efs, "static/temp.tpl", "static/styles.css", "static/logo.svg")
	if err != nil {
		panic(err)
	}
	tpl = temp
}

func NewCalendar(year uint, region regions.Region) *calendar {
	cal := &calendar{year: int(year), region: region, Props: newProps(1920.0, 1080.0)}
	cal.hs = holidays.GetHolidaysList(region, year, year+1)
	return cal
}

func (c *calendar) SetResolution(width float64, height float64) *calendar {
	c.Props = newProps(width, height)
	return c
}

func (c *calendar) Export(name string) error {
	obj := newCalendarObject(int(c.year), c.region, c.hs, c.Props)
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
func (c *calendar) Write(w io.Writer) error {
	obj := newCalendarObject(int(c.year), c.region, c.hs, c.Props)
	err := tpl.Execute(w, obj)
	if err != nil {
		return err
	}
	return nil
}
