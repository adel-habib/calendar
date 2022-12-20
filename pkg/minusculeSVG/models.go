package minusculesvg

type CoreAttributes struct {
	Id       string `xml:"id,attr,omitempty"`
	TabIndex string `xml:"tabindex,attr,omitempty"`
	Lang     string `xml:"lang,attr,omitempty"`
}

type StylingAttributes struct {
	Class string `xml:"class,attr,omitempty"`
	Style string `xml:"style,attr,omitempty"`
}

type TextStylingAttributes struct {
	StylingAttributes
	FontFamily     string `xml:"font-family,attr,omitempty"`
	FontSize       string `xml:"font-size,attr,omitempty"`
	FontSizeAdjust string `xml:"font-size-adjust,attr,omitempty"`
	FontStretch    string `xml:"font-stretch,attr,omitempty"`
	FontStyle      string `xml:"font-style,attr,omitempty"`
	FontVariant    string `xml:"font-variant,attr,omitempty"`
	Fontweight     string `xml:"font-font-weight,attr,omitempty"`
}

type Circle struct {
	CoreAttributes
	StylingAttributes
	Cx string `xml:"cx,attr,omitempty"`
	Cy string `xml:"cy,attr,omitempty"`
	R  string `xml:"r,attr,omitempty"`
}

type Text struct {
	CoreAttributes
	TextStylingAttributes
	Text             string `xml:",chardata"`
	X                string `xml:"x,attr,omitempty"`
	Y                string `xml:"y,attr,omitempty"`
	DX               string `xml:"dx,attr,omitempty"`
	DY               string `xml:"dy,attr,omitempty"`
	Rotate           string `xml:"rotate,attr,omitempty"`
	DominantBaseline string `xml:"dominant-baseline,attr,omitempty"`
	TextAnchor       string `xml:"text-anchor,attr,omitempty"`
}

type Path struct {
	CoreAttributes
	StylingAttributes
	D string `xml:"d,attr"`
}

type Rect struct {
	CoreAttributes
	StylingAttributes
	X      string `xml:"x,attr,omitempty"`
	Y      string `xml:"y,attr,omitempty"`
	RX     string `xml:"rx,attr,omitempty"`
	RY     string `xml:"ry,attr,omitempty"`
	Width  string `xml:"width,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
}
