package holidays

import "time"

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
	day := (22 + d + e)
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
	NEUJAHR := newDate(year, time.January, 1)
	HEILIGE_DER_DREI_KOENIGE := newDate(year, time.January, 6)
	FRAUEN_TAG := newDate(year, time.March, 8)
	TAG_DER_ARBEIT := newDate(year, time.May, 1)
	MARIA_HIMMELFAHRT := newDate(year, time.August, 15)
	WELT_KINDER_TAG := newDate(year, time.September, 20)
	TAG_DER_DEUTSCHEN_EINHEIT := newDate(year, time.October, 3)
	REFORMATIONSTAG := newDate(year, time.October, 31)
	ALLERHEILIGEN := newDate(year, time.November, 1)
	ERSTER_WEIHNACHTSTAG := newDate(year, time.December, 25)
	ZWEITER_WEIHNACHTSTAG := newDate(year, time.December, 26)

	// Easter-based Holidays
	OSTER_SONNTAG := calculateEasterDate(year)
	y, m, d := OSTER_SONNTAG.Date()
	OSTER_MONTAG := newDate(y, m, d+1)
	KAR_FREITAG := newDate(y, m, d-2)
	CHRISTI_HIMMEL_FAHRT := newDate(y, m, d+39)
	PFINGST_SONNTAG := newDate(y, m, d+49)
	PFINGST_MONTAG := newDate(y, m, d+50)
	FRONLEICHNAM := newDate(y, m, d+60)
	ASCHERMITTWOCH := newDate(y, m, d-46)

	NOV23 := newDate(year, time.November, 23)
	BUSS_UND_BETTAG := PreviousDayOfWeek(NOV23, time.Wednesday)

	WEIBERFASTNACHT := PreviousDayOfWeek(ASCHERMITTWOCH, time.Thursday)
	ROSENMONTAG := PreviousDayOfWeek(ASCHERMITTWOCH, time.Monday)
	y, m, d = ROSENMONTAG.Date()
	FASTNACHT := newDate(y, m, d+1)

	nationalHolidays := map[string]time.Time{
		"Neujahr":             NEUJAHR,
		"Karfreitag":          KAR_FREITAG,
		"Ostermontag":         OSTER_MONTAG,
		"Christi Himmelfahrt": CHRISTI_HIMMEL_FAHRT,
		"Pfingstmontag":       PFINGST_MONTAG,
		"Tag der Arbeit":      TAG_DER_ARBEIT,
		"Deutsche Einheit":    TAG_DER_DEUTSCHEN_EINHEIT,
		"1. Weihnachtstag":    ERSTER_WEIHNACHTSTAG,
		"2. Weihnachtstag":    ZWEITER_WEIHNACHTSTAG,
	}
	// Date of local holidays.
	regionalHolidays := map[string]time.Time{
		"Heilige Drei Könige": HEILIGE_DER_DREI_KOENIGE,
		"Frauen Tag":          FRAUEN_TAG,
		"Buß- und Bettag":     BUSS_UND_BETTAG,
		"Weltkindertag":       WELT_KINDER_TAG,
		"Ostersonntag":        OSTER_SONNTAG,
		"Pfingstsonntag":      PFINGST_SONNTAG,
		"Fronleichnam":        FRONLEICHNAM,
		"Mariä Himmelfahrt":   MARIA_HIMMELFAHRT,
		"Reformationstag":     REFORMATIONSTAG,
		"Allerheiligen":       ALLERHEILIGEN,
		"Rosenmontag":         ROSENMONTAG,
		"Fastnacht":           FASTNACHT,
		"Weiberfastnacht":     WEIBERFASTNACHT,
	}

	regionalHolidaysRegions := map[string][]Region{
		"Heilige Drei Könige": {BW, BY, ST},
		"Frauen Tag":          {BE},
		"Buß- und Bettag":     {SN},
		"Weltkindertag":       {TH},
		"Ostersonntag":        {BB},
		"Pfingstsonntag":      {BB},
		"Fronleichnam":        {BW, BY, HE, ST, NW, RP, SL},
		"Mariä Himmelfahrt":   {SL},
		"Reformationstag":     {BB, HE, HB, HH, MV, NI, SN, ST, SH, TH},
		"Allerheiligen":       {BW, BY, NV, RP, SL},
		"Rosenmontag":         {BW},
		"Fastnacht":           {BW},
		"Weiberfastnacht":     {BW},
		// Rosenmontag, Fastnacht, Weiberfastnacht
		// not an actual regional holiday but treated as such in regions where carnival is a big thing, e.g. Kölln
	}

	for holiday, date := range nationalHolidays {
		h := Holiday{Name: holiday, Date: date, Federal: true, Regions: []Region{DE}}
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

func GermanHolidaysByRegion(year int, region Region) (holidays []Holiday) {
	hs := GermanHolidays(year)
	for _, h := range hs {
		if contains(h.Regions, region) || h.Federal {
			holidays = append(holidays, h)
		}
	}
	return
}

func GermanHolidaysNotInRegion(year int, region Region) (holidays []Holiday) {
	hs := GermanHolidays(year)
	for _, h := range hs {
		if !contains(h.Regions, region) && !h.Federal {
			holidays = append(holidays, h)
		}
	}
	return
}
