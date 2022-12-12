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
	dateOffsetX := RECT_WIDTH * 0.05
	labelOffsetX := TEST_FZ * 1.5
	date = minusculesvg.NewText(fmt.Sprintf("%02d", d.Day()), p.x+dateOffsetX, p.y+RECT_HEIGHT-(TEST_FZ/2), "dateText")
	label = minusculesvg.NewText(fmt.Sprintf("%s", d.Weekday())[0:2], p.x+labelOffsetX, p.y+RECT_HEIGHT-(TEST_FZ/2), "dayText")
	return
}

func newDayRect(d time.Time, p Position) (rect minusculesvg.Rect) {
	class := "nRect"
	if isWeekend(d) {
		class = "hRect"
	}
	rect = minusculesvg.NewRect(p.x, p.y, RECT_WIDTH, RECT_HEIGHT, class)
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
	rw := WIDTH - 2*FRAME
	rect := minusculesvg.NewRect(FRAME, FRAME, rw, HEADER_HEIGHT, "headerRect")
	text := minusculesvg.NewText(fmt.Sprint(year), FRAME+rw, HEADER_HEIGHT, "headerText")
	text.DominantBaseline = "text-top"
	text.TextAnchor = "end"
	header.Rect = rect
	header.Text = text
	return
}

func NewMonthsLabels() (labels []minusculesvg.Text) {
	gapHeight := RECT_HEIGHT
	p := Position{x: 0.0, y: FRAME + HEADER_HEIGHT + (gapHeight / 2)}
	for month := 1; month <= NUM_MONTHS; month++ {
		cmonth := month
		if month > 12 {
			cmonth = month - 12
		}
		fmt.Println(cmonth)
		p.x = (float64(month-1) * RECT_WIDTH) + (RECT_WIDTH / 2) + FRAME
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
