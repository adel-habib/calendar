package holidays

import (
	"fmt"
	"github.com/adel-habib/calendar/pkg/regions"
	"time"
)

type Holiday struct {
	Date    time.Time
	Name    string
	Regions []regions.Region
	Federal bool
}

func (h Holiday) String() string {
	if h.Federal {
		return fmt.Sprintf("{%v, %s, federal}", h.Name, h.Date.Format("2006-01-02"))
	} else {
		return fmt.Sprintf("{%v, %s,regions: %v}", h.Name, h.Date.Format("2006-01-02"), h.Regions)
	}
}
