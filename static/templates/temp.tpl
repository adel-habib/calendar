<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" viewBox="0 0 1920 1080">
    <style>
        {{ template "styles.css" }}
    </style>
    <g id="cal-header">
        {{ $hrect := .Header.Rect}}
        <rect x="{{$hrect.X}}" y="{{$hrect.Y}}" width="{{$hrect.Width}}" height="{{$hrect.Height}}" class="{{$hrect.Class}}"> </rect>
        {{ $htext := .Header.Text}}
        <text x="{{$htext.X}}" y="{{$htext.Y}}" class="{{$htext.Class}}" dominant-baseline="{{ $htext.DominantBaseline }}" text-anchor="{{ $htext.TextAnchor }}"> {{ $htext.Text }}</text>
        {{ template "logo.svg" }}
    </g>
    {{ range $index, $dayGroup := .DayGroups }} 
        <g id="{{$dayGroup.FormattedDate}}">
            <rect x="{{$dayGroup.Rect.X}}" y="{{$dayGroup.Rect.Y}}" width="{{$dayGroup.Rect.Width}}" height="{{$dayGroup.Rect.Height}}" class="{{$dayGroup.Rect.Class}}"> </rect>
            {{ range $index, $text := $dayGroup.Texts }}
            <text x="{{$text.X}}" y="{{$text.Y}}" class="{{$text.Class}}"> {{ $text.Text }}</text>
            {{ end }}
        </g>
    {{ end }}
</svg>
