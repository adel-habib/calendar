<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" viewBox="0 0 {{.Props.Width }}  {{.Props.Height }}">

    <defs>
        <linearGradient id="weekEndHolidayGR">
            <stop offset="0%" class="lgs2" />
            <stop offset="100%" class="lgs1" />
        </linearGradient>
      </defs>
      
    <style>
        {{ template "styles.css" .Props }}
    </style>

    <g id="cal-header">
        {{ $hrect := .Header.Rect}}
        <rect id="header" x="{{$hrect.X}}" y="{{$hrect.Y}}" width="{{$hrect.Width}}" height="{{$hrect.Height}}" class="{{$hrect.Class}}"> </rect>
        {{ $htext := .Header.Text}}
        <text x="{{$htext.X}}" y="{{$htext.Y}}" class="{{$htext.Class}}" dominant-baseline="{{ $htext.DominantBaseline }}" text-anchor="{{ $htext.TextAnchor }}"> {{ $htext.Text }}</text>
        {{$fa := .Props.LogoScalFactor | RoundFloat}}
        <g id="itemis-logo" transform="translate(25,15) scale({{$fa}})">
            {{ template "logo.svg" }}
        </g>
    </g>

    <g id="months-labels" class="monthText" dominant-baseline="middle" text-anchor="middle">
        {{ range $index, $label := .MonthsLabels }}
        <text x="{{$label.X}}" y="{{$label.Y}}"> {{ $label.Text }}</text>
        {{ end }}
    </g>

    {{ range $key, $dayGroups := .MonthGroups }}
        <g id="{{$key}}">
            {{ range $index, $dayGroup := $dayGroups }}
            <g id="{{$dayGroup.FormattedDate}}">
                <rect x="{{$dayGroup.Rect.X}}" y="{{$dayGroup.Rect.Y}}" width="{{$dayGroup.Rect.Width}}" height="{{$dayGroup.Rect.Height}}" class="{{$dayGroup.Rect.Class}}"> </rect>
                {{ range $index, $text := $dayGroup.Texts }}
                <text x="{{$text.X}}" y="{{$text.Y}}" class="{{$text.Class}}">{{ $text.Text }}</text>
                {{ end }}
            </g>
            {{ end }}
        </g> 

    {{end }}

    <g id="kw-labels" class="calendarWeekText" dominant-baseline="auto">
        {{ range $ii, $week := .WeekLabels }}
        <text x="{{$week.X}}" y="{{$week.Y}}"> {{ $week.Text }}</text>
        {{ end }}
    </g>

    <g id="footer">
        {{ $hrect := .Footer.Rect}}
        <rect x="{{$hrect.X}}" y="{{$hrect.Y}}" width="{{$hrect.Width}}" height="{{$hrect.Height}}" class="{{$hrect.Class}}"> </rect>
    </g>
</svg>
