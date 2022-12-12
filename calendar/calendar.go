package calendar

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/adel-habib/calendar/holidays"
	minusculesvg "github.com/adel-habib/calendar/minusculeSVG"
)

const (
	NumMonths    = 13
	SVGWidth     = 1920.0
	SVGHeight    = 1080.0
	FrameOffset  = 15.0
	HeaderHeight = 98.0
	FooterHeight = 62.0
	RectHeight   = (SVGHeight - HeaderHeight - FooterHeight - 2*FrameOffset) / 31
	RectWidth    = (SVGWidth - 2*FrameOffset) / NumMonths
	TestFz       = 20
)

type Position struct {
	x float64
	y float64
}

type DayGroup struct {
	XMLName xml.Name            `xml:"g"`
	Rect    minusculesvg.Rect   `xml:"rect"`
	Texts   []minusculesvg.Text `xml:"text"`
	Date    time.Time
}

type MonthGroup struct {
	DGs []DayGroup `xml:"g"`
}

type HeaderGroup struct {
	XMLName xml.Name          `xml:"g"`
	Rect    minusculesvg.Rect `xml:"rect"`
	Text    minusculesvg.Text `xml:"text"`
}

type CalendarProps struct {
	Year         int
	Header       HeaderGroup
	MonthsLabels []minusculesvg.Text
	MonthGroups  map[string][]DayGroup
}

func CalendarDayGroups(year int, s []holidays.Holiday) (ms map[string][]DayGroup) {
	ms = make(map[string][]DayGroup)
	yearCursor := year
	for month := 1; month <= NumMonths; month++ {
		var monthDayGroups []DayGroup

		monthCursor := month
		if month > 12 {
			yearCursor++
			monthCursor = month - 12
		}

		daysOfMonth := time.Date(year, time.Month(monthCursor+1), 0, 0, 0, 0, 0, time.UTC).Day()

		for day := 1; day <= daysOfMonth; day++ {
			dateCursor := time.Date(yearCursor, time.Month(monthCursor), day, 0, 0, 0, 0, time.UTC)

			xOffset := FrameOffset
			if yearCursor > year {
				xOffset += RectWidth * 12
			}
			p := elementCoordinates(dateCursor, RectWidth, RectHeight, xOffset, FrameOffset+HeaderHeight)
			group := NewDayGroup(dateCursor, p)

			idx := holidays.Index(s, func(h holidays.Holiday) bool { return h.Date.Equal(dateCursor) })
			if idx != -1 {
				h := s[idx]
				txt := minusculesvg.NewText(h.Name, p.x+RectWidth-(0.02*RectWidth), p.y+RectHeight-(0.1*RectHeight), "holidayText")
				group.Texts = append(group.Texts, txt)
				group.Rect.Class = "holidayRect"
				if isWeekend(dateCursor) {
					group.Rect.Class = "holidayWeekEndRect"
				}
			}
			monthDayGroups = append(monthDayGroups, group)
		}
		ms[fmt.Sprintf("%d-%02d-month-group", yearCursor, monthCursor)] = monthDayGroups
	}
	return
}
