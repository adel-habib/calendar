package calendar

import (
	"fmt"
	"time"

	minusculesvg "github.com/adel-habib/calendar/minusculeSVG"
)

func (d DayGroup) FormattedDate() string {
	return d.Date.Format("02-01-2006")
}

func elementCoordinates(t time.Time, rectWidth float64, rectHeight float64, xoffset float64, yoffset float64) (p Position) {
	month := float64(t.Month())
	day := float64(t.Day())
	p.x = ((month - 1) * rectWidth) + xoffset
	p.y = (day * rectHeight) + yoffset
	return
}

func newDayText(d time.Time, p Position) (label minusculesvg.Text, date minusculesvg.Text) {
	dateOffsetX := RectWidth * 0.05
	labelOffsetX := TestFz * 1.5
	date = minusculesvg.NewText(fmt.Sprintf("%02d", d.Day()), p.x+dateOffsetX, p.y+RectHeight-(TestFz/2), "dateText")
	label = minusculesvg.NewText(fmt.Sprintf("%s", d.Weekday())[0:2], p.x+labelOffsetX, p.y+RectHeight-(TestFz/2), "dayText")
	return
}

func newDayRect(d time.Time, p Position) (rect minusculesvg.Rect) {
	class := "defaultRect"
	if isWeekend(d) {
		class = "weekendRect"
	}
	rect = minusculesvg.NewRect(p.x, p.y, RectWidth, RectHeight, class)
	return
}

func NewDayGroup(t time.Time, p Position) (g DayGroup) {
	g.Rect = newDayRect(t, p)
	g.Date = t
	label, date := newDayText(t, p)
	g.Texts = []minusculesvg.Text{label, date}
	return
}

func Header(year int) (header HeaderGroup) {
	rw := SVGWidth - 2*FrameOffset
	rect := minusculesvg.NewRect(FrameOffset, FrameOffset, rw, HeaderHeight, "headerRect")
	text := minusculesvg.NewText(fmt.Sprint(year), FrameOffset+rw, HeaderHeight, "headerText")
	text.DominantBaseline = "text-top"
	text.TextAnchor = "end"
	header.Rect = rect
	header.Text = text
	return
}

func NewMonthsLabels() (labels []minusculesvg.Text) {
	gapHeight := RectHeight
	p := Position{x: 0.0, y: FrameOffset + HeaderHeight + (gapHeight / 2)}
	for month := 1; month <= NumMonths; month++ {
		cmonth := month
		if month > 12 {
			cmonth = month - 12
		}
		p.x = (float64(month-1) * RectWidth) + (RectWidth / 2) + FrameOffset
		text := minusculesvg.NewText(fmt.Sprint(time.Month(cmonth)), p.x, p.y, "monthText")
		labels = append(labels, text)
	}
	return
}

func isWeekend(d time.Time) bool {
	weekDay := d.Weekday()
	if weekDay == time.Saturday || weekDay == time.Sunday {
		return true
	}
	return false
}
