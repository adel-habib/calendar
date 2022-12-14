package holidays

import (
	"time"

	"github.com/adel-habib/calendar/regions"
)

func calculateEasterDate(year int) (easterDate time.Time) {
	a := year % 19
	b := year % 4
	c := year % 7
	k := year / 100
	p := k / 3
	q := k / 4
	m := (15 + k - p - q) % 30
	d := (19*a + m) % 30
	n := (4 + k - q) % 7
	e := (2*b + 4*c + 6*d + n) % 7
	day := 22 + d + e
	month := time.March
	easterDate = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return
}

func newDate(year int, month time.Month, day int) (d time.Time) {
	d = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return
}
func GermanHolidays(year int) (holidays []Holiday) {
	// Fixed Holidays
	Neujahr := newDate(year, time.January, 1)
	HeiligeDerDreiKoenige := newDate(year, time.January, 6)
	FrauenTag := newDate(year, time.March, 8)
	TagDerArbeit := newDate(year, time.May, 1)
	MariaHimmelfahrt := newDate(year, time.August, 15)
	WeltKinderTag := newDate(year, time.September, 20)
	TagDerDeutschenEinheit := newDate(year, time.October, 3)
	Reformationstag := newDate(year, time.October, 31)
	Allerheiligen := newDate(year, time.November, 1)
	ErsterWeihnachtstag := newDate(year, time.December, 25)
	ZweiterWeihnachtstag := newDate(year, time.December, 26)

	// Easter-based Holidays
	OsterSonntag := calculateEasterDate(year)
	y, m, d := OsterSonntag.Date()
	OsterMontag := newDate(y, m, d+1)
	KarFreitag := newDate(y, m, d-2)
	ChristiHimmelFahrt := newDate(y, m, d+39)
	PfingstSonntag := newDate(y, m, d+49)
	PfingstMontag := newDate(y, m, d+50)
	Fronleichnam := newDate(y, m, d+60)
	Aschermittwoch := newDate(y, m, d-46)

	NOV23 := newDate(year, time.November, 23)
	BussUndBettag := PreviousDayOfWeek(NOV23, time.Wednesday)

	WEIBERFASTNACHT := PreviousDayOfWeek(Aschermittwoch, time.Thursday)
	ROSENMONTAG := PreviousDayOfWeek(Aschermittwoch, time.Monday)
	y, m, d = ROSENMONTAG.Date()
	FASTNACHT := newDate(y, m, d+1)

	nationalHolidays := map[string]time.Time{
		"Neujahr":             Neujahr,
		"Karfreitag":          KarFreitag,
		"Ostermontag":         OsterMontag,
		"Christi Himmelfahrt": ChristiHimmelFahrt,
		"Pfingstmontag":       PfingstMontag,
		"Tag der Arbeit":      TagDerArbeit,
		"Deutsche Einheit":    TagDerDeutschenEinheit,
		"1. Weihnachtstag":    ErsterWeihnachtstag,
		"2. Weihnachtstag":    ZweiterWeihnachtstag,
	}
	// Date of local holidays.
	regionalHolidays := map[string]time.Time{
		"Heilige Drei K??nige": HeiligeDerDreiKoenige,
		"Frauen Tag":          FrauenTag,
		"Bu??- und Bettag":     BussUndBettag,
		"Weltkindertag":       WeltKinderTag,
		"Ostersonntag":        OsterSonntag,
		"Pfingstsonntag":      PfingstSonntag,
		"Fronleichnam":        Fronleichnam,
		"Mari?? Himmelfahrt":   MariaHimmelfahrt,
		"Reformationstag":     Reformationstag,
		"Allerheiligen":       Allerheiligen,
		"Rosenmontag":         ROSENMONTAG,
		"Fastnacht":           FASTNACHT,
		"Weiberfastnacht":     WEIBERFASTNACHT,
	}

	regionalHolidaysRegions := map[string][]regions.Region{
		"Heilige Drei K??nige": {regions.BW, regions.BY, regions.ST},
		"Frauen Tag":          {regions.BE},
		"Bu??- und Bettag":     {regions.SN},
		"Weltkindertag":       {regions.TH},
		"Ostersonntag":        {regions.BB},
		"Pfingstsonntag":      {regions.BB},
		"Fronleichnam":        {regions.BW, regions.BY, regions.HE, regions.ST, regions.NW, regions.RP, regions.SL},
		"Mari?? Himmelfahrt":   {regions.SL},
		"Reformationstag":     {regions.BB, regions.HE, regions.HB, regions.HH, regions.MV, regions.NI, regions.SN, regions.ST, regions.SH, regions.TH},
		"Allerheiligen":       {regions.BW, regions.BY, regions.NV, regions.RP, regions.SL},
		"Rosenmontag":         {regions.BW},
		"Fastnacht":           {regions.BW},
		"Weiberfastnacht":     {regions.BW},
		// Rosenmontag, Fastnacht, Weiberfastnacht
		// not an actual regional holiday but treated as such in regions where carnival is a big thing, e.g. K??lln
	}

	for holiday, date := range nationalHolidays {
		h := Holiday{Name: holiday, Date: date, Federal: true, Regions: []regions.Region{regions.DE}}
		holidays = append(holidays, h)
	}

	for holiday, date := range regionalHolidays {
		regions := regionalHolidaysRegions[holiday]
		h := Holiday{Name: holiday, Date: date, Federal: false, Regions: regions}
		holidays = append(holidays, h)
	}

	return
}

func PreviousDayOfWeek(date time.Time, day time.Weekday) (prevDate time.Time) {
	prevDate = date
	y, m, d := date.Date()
	for i := 1; i < 8; i++ {
		cursor := newDate(y, m, d-i)
		if cursor.Weekday() == day {
			prevDate = cursor
			break
		}

	}
	return
}

func GermanHolidaysByRegion(year int, region regions.Region) (holidays []Holiday) {
	hs := GermanHolidays(year)
	for _, h := range hs {
		if contains(h.Regions, region) || h.Federal {
			holidays = append(holidays, h)
		}
	}
	return
}

func GetHolidaysList(r regions.Region, years ...uint) (s []Holiday) {
	for _, v := range years {
		s = append(s, GermanHolidaysByRegion(int(v), r)...)
	}
	return
}

func GermanHolidaysNotInRegion(year int, region regions.Region) (holidays []Holiday) {
	hs := GermanHolidays(year)
	for _, h := range hs {
		if !contains(h.Regions, region) && !h.Federal {
			holidays = append(holidays, h)
		}
	}
	return
}
