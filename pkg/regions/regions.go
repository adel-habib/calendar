package regions

import "strings"

type Region string

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

var SupportedRegions = [...]Region{
	DE,
	BW,
	BY,
	ST,
	BB,
	HE,
	NW,
	RP,
	SL,
	BE,
	HB,
	HH,
	SN,
	TH,
	SH,
	MV,
	NV,
	NI,
}

func RegionByName(name string) (int, *Region) {
	for index, region := range SupportedRegions {
		r := string(region)
		if strings.EqualFold(r, name) {
			return index, &region
		}
	}
	return -1, nil
}
