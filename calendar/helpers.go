package calendar

import (
	"fmt"
	"time"

	"github.com/adel-habib/calendar/holidays"
	minusculesvg "github.com/adel-habib/calendar/minusculeSVG"
)

func (d dayGroup) FormattedDate() string {
	return d.Date.Format("02-01-2006")
}

// 2-D grid -> column represents a month, rect in column represents a day
// y = f(day) = day*RectHeight
// x = f(month) =  month*RectWidth
// X- And Y-Offset take the margins, headers, margins etc into consideration
func dateToCoordinates(t time.Time, rectWidth float64, rectHeight float64, xOffset float64, yOffset float64) (p position) {
	month := float64(t.Month())
	day := float64(t.Day())
	p.x = ((month - 1) * rectWidth) + xOffset
	p.y = (day * rectHeight) + yOffset
	return
}

func newDayText(d time.Time, p position, props *Props) (label minusculesvg.Text, date minusculesvg.Text) {
	dateOffsetX := props.RectWidth * 0.05
	labelOffsetX := props.FontSize * 2
	date = minusculesvg.NewText(fmt.Sprintf("%02d", d.Day()), p.x+dateOffsetX, p.y+props.FontSize, "dateText")
	label = minusculesvg.NewText(fmt.Sprintf("%s", d.Weekday())[0:2], p.x+labelOffsetX, p.y+props.FontSize, "dayText")
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

func newDayGroup(t time.Time, p position, props *Props, s []holidays.Holiday) (g dayGroup) {
	g.Rect = newDayRect(t, p, props.RectWidth, props.RectHeight)
	g.Date = t
	label, date := newDayText(t, p, props)
	g.Texts = []minusculesvg.Text{label, date}
	// check of the day is a holiday
	idx := holidays.Index(s, func(h holidays.Holiday) bool { return h.Date.Equal(t) })
	if idx != -1 {
		h := s[idx]
		txt := minusculesvg.NewText(h.Name, p.x+props.RectWidth-(0.02*props.RectWidth), p.y+props.RectHeight-(0.15*props.RectHeight), "holidayText")
		g.Texts = append(g.Texts, txt)
		g.Rect.Class = "holidayRect"
		if isWeekend(t) {
			g.Rect.Class = "holidayWeekEndRect"
		}
	}
	return
}

func Newheader(year int, props Props) (header headerGroup) {
	headerWidth := props.Width - 2*props.Margin
	rect := minusculesvg.NewRect(props.Margin, props.Margin, headerWidth, props.HeaderHeight, "headerRect")
	x := props.Margin + headerWidth
	y := props.Margin + (props.HeaderHeight / 2.0)
	text := minusculesvg.NewText(fmt.Sprint(year), x, y, "headerText")
	// center vertically
	text.DominantBaseline = "central"
	// shift the text such that the end of the resulting rendered text is at the initial current text position
	text.TextAnchor = "end"
	header.Rect = rect
	header.Text = text
	return
}

func newFooter(props *Props) (footer headerGroup) {
	y := props.Height - props.Margin - props.FooterHeight
	rect := minusculesvg.NewRect(props.Margin, y, props.Width-2*props.Margin, props.FooterHeight, "headerRect")
	footer.Rect = rect
	return
}

func newMonthsLabels(g Props) (labels []minusculesvg.Text) {
	gapHeight := g.RectHeight
	p := position{x: 0.0, y: g.Margin + g.HeaderHeight + (gapHeight / 2)}
	for month := 1; month <= numMonths; month++ {
		cursor := month
		if month > 12 {
			cursor = month - 12
		}
		p.x = (float64(month-1) * g.RectWidth) + (g.RectWidth / 2) + margin
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

func newProps(width float64, height float64) Props {
	g := Props{Width: width, Height: height}
	g.NumMonths = numMonths
	g.Margin = margin
	g.HeaderHeight = (g.Height - 2*g.Margin) * 0.1
	g.FooterHeight = (g.Height - 2*g.Margin) * 0.08
	// 31 days a month + gap between header and body + gap between body and footer
	g.RectHeight = (g.Height - g.HeaderHeight - g.FooterHeight - 2*margin) / 33
	g.RectWidth = (g.Width - 2*g.Margin) / g.NumMonths
	g.LogoHeight = g.HeaderHeight * 0.8
	g.LogoWidth = g.LogoHeight * 3.5
	g.LogoScalFactor = (g.HeaderHeight / pointsToPixels(itemisLogoHeighthOriginal)) * 0.95
	g.HeaderFonzSize = g.HeaderHeight * 0.95
	g.FontSize = g.RectHeight * 0.7
	g.MonthLabelFonzSize = g.RectHeight * 0.5
	return g
}
func newCalendarObject(year int, s []holidays.Holiday, geometry Props) bodyObject {
	ob := bodyObject{
		Year:         year,
		Header:       Newheader(year, geometry),
		MonthsLabels: newMonthsLabels(geometry),
		MonthGroups:  monthsGroups(year, s, geometry),
		Props:        geometry,
		Footer:       newFooter(&geometry),
	}
	return ob
}

// Create a group element for each day of the year
func monthsGroups(year int, s []holidays.Holiday, g Props) map[string][]dayGroup {
	monthDayGroupsMap := make(map[string][]dayGroup)
	yearCursor := year
	for month := 1; month <= numMonths; month++ {
		var monthDayGroups []dayGroup

		monthCursor := month
		if month > 12 {
			yearCursor++
			monthCursor = month - 12
		}

		// number of days of a month, [go normalises time]
		daysOfMonth := time.Date(year, time.Month(monthCursor+1), 0, 0, 0, 0, 0, time.UTC).Day()

		for day := 1; day <= daysOfMonth; day++ {
			dateCursor := time.Date(yearCursor, time.Month(monthCursor), day, 0, 0, 0, 0, time.UTC)

			xOffset := g.Margin
			yOffset := g.Margin + g.HeaderHeight
			// january of the following year
			if yearCursor > year {
				xOffset += g.RectWidth * 12
			}
			p := dateToCoordinates(dateCursor, g.RectWidth, g.RectHeight, xOffset, yOffset)
			group := newDayGroup(dateCursor, p, &g, s)
			monthDayGroups = append(monthDayGroups, group)
		}
		monthDayGroupsMap[fmt.Sprintf("%d-%02d-month-group", yearCursor, monthCursor)] = monthDayGroups
	}
	return monthDayGroupsMap
}

func pointsToPixels(pt float64) float64 {
	return pt * (96.0 / 72.0)
}

func appendHolidayLabelText(dg *dayGroup, h holidays.Holiday) {
}
