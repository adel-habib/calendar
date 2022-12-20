package main

import (
	"bytes"
	"github.com/adel-habib/calendar/pkg/calendar"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	err, params := app.parseCalendarParams(r.URL)
	if err != nil {
		app.clientError(w, 400)
		return
	}

	buf := &bytes.Buffer{}

	cal := calendar.NewCalendar(uint(params.Year), params.Region)
	err = cal.Write(buf)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "image/svg+xml")
	_, err = w.Write(buf.Bytes())
	if err != nil {
		app.serverError(w, err)
		return
	}

}
