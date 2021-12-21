package custsort

import (
	"sort"
)

type (
	Gopher struct {
		Name  string
		Age   int
		Place int
	}
	sortField int
)

const (
	sfName sortField = iota
	sfAge
	sfPlace
)

func SortBy(items []Gopher, fields ...sortField) []Gopher {
	if len(fields) == 0 {
		return items
	}

	sort.SliceStable(items, func(i, j int) bool {
		i1 := items[i]
		i2 := items[j]
		for _, sf := range fields {
			switch sf {
			case sfName:
				if i1.Name == i2.Name {
					continue
				}
				return i1.Name < i2.Name
			case sfAge:
				if i1.Age == i2.Age {
					continue
				}
				return i1.Age < i2.Age
			case sfPlace:
				if i1.Place == i2.Place {
					continue
				}
				return i1.Place < i2.Place
			}
		}
		return false
	})
	return items
}
