package minusculesvg

import "fmt"

func NewText(text string, x float64, y float64, class string) (textEl Text) {
	textEl = Text{
		X:                     fmt.Sprintf("%.2f", x),
		Y:                     fmt.Sprintf("%.2f", y),
		Text:                  text,
		TextStylingAttributes: TextStylingAttributes{StylingAttributes: StylingAttributes{Class: class}},
	}
	return
}

func NewTextWithStyleAttributes(text string, x float64, y float64, attributes TextStylingAttributes) (textEl Text) {
	textEl = Text{
		X:                     fmt.Sprintf("%.2f", x),
		Y:                     fmt.Sprintf("%.2f", y),
		Text:                  text,
		TextStylingAttributes: attributes}
	return
}

func NewRect(x float64, y float64, width float64, height float64, class string) (rect Rect) {
	rect = Rect{
		X:                 fmt.Sprintf("%.2f", x),
		Y:                 fmt.Sprintf("%.2f", y),
		Width:             fmt.Sprintf("%.2f", width),
		Height:            fmt.Sprintf("%.2f", height),
		StylingAttributes: StylingAttributes{Class: class}}
	return
}
