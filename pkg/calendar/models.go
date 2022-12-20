package calendar

import (
	"encoding/xml"
	"github.com/adel-habib/calendar/pkg/holidays"
	"github.com/adel-habib/calendar/pkg/minusculeSVG"
	"github.com/adel-habib/calendar/pkg/regions"
	"time"
)

const (
	numMonths                 = 13
	svgWidth                  = 1920.0
	svgHeight                 = 1080.0
	margin                    = 15.0
	headerHeight              = 98.0
	footerHeight              = 62.0
	testFz                    = 20
	itemisLogoWidthOriginal   = 283.46
	itemisLogoHeighthOriginal = 81.86
)

type calendar struct {
	year   int
	region regions.Region
	hs     []holidays.Holiday
	Props  Props
}

type Props struct {
	Width              float64
	Height             float64
	Margin             float64
	HeaderHeight       float64
	FooterHeight       float64
	RectHeight         float64
	RectWidth          float64
	NumMonths          float64
	LogoWidth          float64
	LogoHeight         float64
	LogoScalFactor     float64
	FontSize           float64
	MonthLabelFonzSize float64
	HeaderFonzSize     float64
}

type bodyObject struct {
	Year         int
	Region       regions.Region
	Props        Props
	Header       headerGroup
	MonthsLabels []minusculesvg.Text
	MonthGroups  map[string][]dayGroup
	WeekLabels   []minusculesvg.Text
	Footer       headerGroup
}

type position struct {
	x float64
	y float64
}

type dayGroup struct {
	XMLName xml.Name            `xml:"g"`
	Rect    minusculesvg.Rect   `xml:"rect"`
	Texts   []minusculesvg.Text `xml:"text"`
	Date    time.Time
}

type headerGroup struct {
	XMLName xml.Name          `xml:"g"`
	Rect    minusculesvg.Rect `xml:"rect"`
	Text    minusculesvg.Text `xml:"text"`
}
