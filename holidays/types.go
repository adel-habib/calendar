package holidays

import (
	"fmt"
	"time"
)

type Region string

type Holiday struct {
	Date    time.Time
	Name    string
	Regions []Region
	Federal bool
}

func (h Holiday) String() string {
	if h.Federal {
		return fmt.Sprintf("{%v, %s, federal}", h.Name, h.Date.Format("2006-01-02"))
	} else {
		return fmt.Sprintf("{%v, %s,regions: %v}", h.Name, h.Date.Format("2006-01-02"), h.Regions)
	}
}

const (
	DE Region = "Deutschland"
	BW Region = "Baden-Württemberg"
	BY Region = "Bayern"
	ST Region = "Sachsen-Anhalt"
	BB Region = "Brandenburg"
	HE Region = "Hessen"
	NW Region = "Nordrhein-Westfalen"
	RP Region = "Rheinland-Pfalz"
	SL Region = "Saarland"
	BE Region = "Berlin"
	HB Region = "Bremen"
	HH Region = "Hamburg"
	SN Region = "Sachsen"
	TH Region = "Thüringen"
	SH Region = "Schleswig-Holstein"
	MV Region = "Mecklenburg-Vorpommern"
	NV Region = "Nordrhein-Westfalen"
	NI Region = "Niedersachsen"
)
