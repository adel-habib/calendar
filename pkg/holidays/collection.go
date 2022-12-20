package holidays

import (
	"github.com/adel-habib/calendar/pkg/regions"
	"time"
)

func Filter(s []Holiday, f func(Holiday) bool) (filteredS []Holiday) {
	for _, h := range s {
		if f(h) {
			filteredS = append(filteredS, h)
		}
	}
	return
}

func All(s []Holiday, f func(Holiday) bool) bool {
	for _, h := range s {
		if !f(h) {
			return false
		}
	}
	return true
}

func IsAny(s []Holiday, dates ...time.Time) bool {
	for _, d := range dates {
		idx := Index(s, func(h Holiday) bool { return h.Date.Equal(d) })
		if idx != -1 {
			return true
		}
	}
	return false
}

func Index(s []Holiday, f func(Holiday) bool) int {
	for i, h := range s {
		if f(h) {
			return i
		}
	}
	return -1
}

func contains(s []regions.Region, e regions.Region) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
