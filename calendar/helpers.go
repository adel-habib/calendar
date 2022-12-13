package calendar

import (
	"fmt"
	"github.com/adel-habib/calendar/holidays"
	"github.com/adel-habib/calendar/minusculeSVG"
	"time"
)

func (d dayGroup) FormattedDate() string {
	return d.Date.Format("02-01-2006")
}

func elementCoordinates(t time.Time, rectWidth float64, rectHeight float64, xOffset float64, yOffset float64) (p position) {
	month := float64(t.Month())
	day := float64(t.Day())
	p.x = ((month - 1) * rectWidth) + xOffset
	p.y = (day * rectHeight) + yOffset
	return
}

func newDayText(d time.Time, p position, rectWidth float64, rectHeight float64) (label minusculesvg.Text, date minusculesvg.Text) {
	dateOffsetX := rectWidth * 0.05
	labelOffsetX := testFz * 1.5
	date = minusculesvg.NewText(fmt.Sprintf("%02d", d.Day()), p.x+dateOffsetX, p.y+rectHeight-(testFz/2), "dateText")
	label = minusculesvg.NewText(fmt.Sprintf("%s", d.Weekday())[0:2], p.x+labelOffsetX, p.y+rectHeight-(testFz/2), "dayText")
	return
}

func newDayRect(d time.Time, p position, rectWidth float64, rectHeight float64) (rect minusculesvg.Rect) {
	class := "defaultRect"
	if isWeekend(d) {
		class = "weekendRect"
	}
	rect = minusculesvg.NewRect(p.x, p.y, rectWidth, rectHeight, class)
	return
}

func newDayGroup(t time.Time, p position, rectWidth float64, rectHeight float64) (g dayGroup) {
	g.Rect = newDayRect(t, p, rectWidth, rectHeight)
	g.Date = t
	label, date := newDayText(t, p, rectWidth, rectHeight)
	g.Texts = []minusculesvg.Text{label, date}
	return
}

func header(year int) (header headerGroup) {
	rw := svgWidth - 2*margin
	rect := minusculesvg.NewRect(margin, margin, rw, headerHeight, "headerRect")
	text := minusculesvg.NewText(fmt.Sprint(year), margin+rw, headerHeight, "headerText")
	text.DominantBaseline = "text-top"
	text.TextAnchor = "end"
	header.Rect = rect
	header.Text = text
	return
}

func newMonthsLabels(rectWidth float64, rectHeight float64) (labels []minusculesvg.Text) {
	gapHeight := rectHeight
	p := position{x: 0.0, y: margin + headerHeight + (gapHeight / 2)}
	for month := 1; month <= numMonths; month++ {
		cursor := month
		if month > 12 {
			cursor = month - 12
		}
		p.x = (float64(month-1) * rectWidth) + (rectWidth / 2) + margin
		text := minusculesvg.NewText(fmt.Sprint(time.Month(cursor)), p.x, p.y, "monthText")
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

func newGeometry(width float64, height float64) Geometry {
	g := Geometry{Width: width, Height: height}
	g.NumMonths = numMonths
	g.Margin = margin
	g.HeaderHeight = (g.Height - g.Margin) * 0.1
	g.FooterHeight = (g.Height - g.Margin) * 0.05
	g.RectHeight = (g.Height - g.HeaderHeight - g.FooterHeight - 2*margin) / 31
	g.RectWidth = (g.Width - 2*g.Margin) / g.NumMonths
	return g
}
func newCalendarObject(year int, s []holidays.Holiday, geometry Geometry) bodyObject {
	ob := bodyObject{
		Year:         year,
		Header:       header(year),
		MonthsLabels: newMonthsLabels(geometry.RectWidth, geometry.RectHeight),
		MonthGroups:  monthsGroups(year, s, geometry.RectWidth, geometry.RectHeight),
	}
	return ob
}
func monthsGroups(year int, s []holidays.Holiday, rectWidth float64, rectHeight float64) (ms map[string][]dayGroup) {
	ms = make(map[string][]dayGroup)
	yearCursor := year
	for month := 1; month <= numMonths; month++ {
		var monthDayGroups []dayGroup

		monthCursor := month
		if month > 12 {
			yearCursor++
			monthCursor = month - 12
		}

		daysOfMonth := time.Date(year, time.Month(monthCursor+1), 0, 0, 0, 0, 0, time.UTC).Day()

		for day := 1; day <= daysOfMonth; day++ {
			dateCursor := time.Date(yearCursor, time.Month(monthCursor), day, 0, 0, 0, 0, time.UTC)

			xOffset := margin
			if yearCursor > year {
				xOffset += rectWidth * 12
			}
			p := elementCoordinates(dateCursor, rectWidth, rectHeight, xOffset, margin+headerHeight)
			group := newDayGroup(dateCursor, p, rectWidth, rectHeight)

			idx := holidays.Index(s, func(h holidays.Holiday) bool { return h.Date.Equal(dateCursor) })
			if idx != -1 {
				h := s[idx]
				txt := minusculesvg.NewText(h.Name, p.x+rectWidth-(0.02*rectWidth), p.y+rectHeight-(0.15*rectHeight), "holidayText")
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
