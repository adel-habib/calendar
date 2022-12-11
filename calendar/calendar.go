package calendar

import (
	"encoding/xml"
	"time"

	minusculesvg "github.com/adel-habib/calendar/minusculeSVG"
)

const (
	NUM_MONTHS         = 13
	WIDTH              = 1920.0
	HEIGHT             = 1080.0
	FRAME              = 15.0
	HEADER_HEIGHT      = 98.0
	FOOTER_HEIGHT      = 62.0
	UPPER_SPACE_HEIGHT = 45.0
	LOWER_SPACE_HEIGHT = 20.0
	RECT_HEIGHT        = (HEIGHT - HEADER_HEIGHT - FOOTER_HEIGHT - UPPER_SPACE_HEIGHT - LOWER_SPACE_HEIGHT - 2*FRAME) / 31
	RECT_WIDTH         = (WIDTH - 2*FRAME) / NUM_MONTHS
	TEST_FZ            = 20
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

type HeaderGroup struct {
	XMLName xml.Name          `xml:"g"`
	Rect    minusculesvg.Rect `xml:"rect"`
	Text    minusculesvg.Text `xml:"text"`
}

type CalendarProps struct {
	Year      int
	Header    HeaderGroup
	DayGroups []DayGroup
}

func CalendarDayGroups(year int) (gs []DayGroup) {
	yearCursor := year
	for month := 1; month <= NUM_MONTHS; month++ {
		monthCursor := month
		if month > 12 {
			yearCursor++
			monthCursor = month - 12
		}
		daysOfMonth := time.Date(year, time.Month(monthCursor+1), 0, 0, 0, 0, 0, time.UTC).Day()
		for day := 1; day <= daysOfMonth; day++ {
			dateCursor := time.Date(yearCursor, time.Month(monthCursor), day, 0, 0, 0, 0, time.UTC)
			xOffset := FRAME
			if yearCursor > year {
				xOffset += RECT_WIDTH * 12
			}
			p := elementCoordinates(dateCursor, RECT_WIDTH, RECT_HEIGHT, xOffset, FRAME+HEADER_HEIGHT)
			group := NewDayGroup(dateCursor, p)
			gs = append(gs, group)
		}
	}
	return
}