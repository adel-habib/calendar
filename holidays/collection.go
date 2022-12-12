package holidays

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

func Index(s []Holiday, f func(Holiday) bool) int {
	for i, h := range s {
		if f(h) {
			return i
		}
	}
	return -1
}

func contains(s []Region, e Region) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}