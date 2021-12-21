package custsort

import (
	"reflect"
	"testing"
)

func TestSortByName(t *testing.T) {
	for _, tt := range []struct {
		name      string
		fields    []sortField
		wantOrder []int
	}{
		{
			name:      "by_Name",
			fields:    []sortField{sfName},
			wantOrder: []int{4, 2, 1, 3},
		},
		{
			name:      "by_Name_Then_Place",
			fields:    []sortField{sfName, sfPlace},
			wantOrder: []int{4, 1, 2, 3},
		},
		{
			name:      "by_Age_Then_Name",
			fields:    []sortField{sfAge, sfName},
			wantOrder: []int{3, 4, 2, 1},
		},
		{
			name:      "no_fields",
			fields:    nil,
			wantOrder: []int{4, 3, 2, 1},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			items := SortBy([]Gopher{
				{"Daniel", 25, 4},
				{"Tom", 19, 3},
				{"Murthy", 33, 2},
				{"Murthy", 42, 1},
			}, tt.fields...)
			var po []int
			for _, item := range items {
				po = append(po, item.Place)
			}
			if !reflect.DeepEqual(po, tt.wantOrder) {
				t.Errorf("got order: %v, want: %v", po, tt.wantOrder)
			}
		})
	}
}
