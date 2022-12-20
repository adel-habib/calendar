package main

import (
	"errors"
	"fmt"
	"github.com/adel-habib/calendar/pkg/regions"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
	"time"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) parseCalendarParams(r *url.URL) (error, *CalendarQueryParams) {
	params := CalendarQueryParams{}
	yearString := r.Query().Get("year")
	regionString := r.Query().Get("region")
	if len(yearString) == 0 {
		currentYear := time.Now().Year()
		if time.Now().Month() == time.December {
			params.Year = currentYear + 1
		} else {
			params.Year = currentYear
		}
	} else {
		y, err := strconv.Atoi(yearString)
		if err != nil || y < 1 || y > 3000 {
			return errors.New("invalid query param, year must be positive integer between 1 and 3000"), nil
		}
		params.Year = y
	}
	if len(regionString) == 0 {
		params.Region = regions.DE
	} else {
		ind, reg := regions.RegionByName(regionString)
		if ind == -1 {
			return errors.New("invalid query param, unsupported region"), nil
		}
		params.Region = *reg
	}
	return nil, &params
}
