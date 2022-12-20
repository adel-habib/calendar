<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<svg xmlns="http://www.w3.org/2000/svg" version="1.1" viewBox="0 0 {{.Props.Width }}  {{.Props.Height }}">

    <title property="dc:title">Calendar {{.Year}} | {{.Region}}</title>
    <desc property="dc:creator">Itemis Leipzig</desc>

    <!-- graphical elements -->

    <defs>
        <style>
            <!-- Inject css styles -->
            {{ template "styles.css" .Props }}
        </style>

        <!-- gradient to style days that are both a holiday and a weekend day (Sat, Sun) -->
        <linearGradient id="weekEndHolidayGR">
            <stop offset="0%" class="lgs2" />
            <stop offset="100%" class="lgs1" />
        </linearGradient>
      </defs>
      


    <g id="cal-header">
        <desc>Calendar Header, contains logo of the company (Itemis) and the current year ({{.Year}})</desc>
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
        <desc>Months labels. from January {{.Year}} to January {{.Year | Inc}}</desc>
        {{ range $index, $label := .MonthsLabels }}
        <text x="{{$label.X}}" y="{{$label.Y}}"> {{ $label.Text }}</text>
        {{ end }}
    </g>
    <g id="calendar-body">
        <desc>
            Calendar Body, contains 13 columns each of which represents a month. from January {{.Year}} to January {{.Year | Inc}}
            Each column contains as many sqaures as there are days in the corresponding months 
            Each sqaure contains the short day name, date and holiday name if it is a holiday
          </desc>
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
    </g>


    <g id="kw-labels" class="calendarWeekText" dominant-baseline="auto">
        <desc> Each week is labeled with its ISO week date (ISO-8601 Week-Based Calendar)</desc>
        {{ range $ii, $week := .WeekLabels }}
        <text id="{{$week.Id}}" x="{{$week.X}}" y="{{$week.Y}}"> {{ $week.Text }}</text>
        {{ end }}
    </g>

    <g id="footer">
        {{ $hrect := .Footer.Rect}}
        <rect x="{{$hrect.X}}" y="{{$hrect.Y}}" width="{{$hrect.Width}}" height="{{$hrect.Height}}" class="{{$hrect.Class}}"> </rect>
    </g>
</svg>
