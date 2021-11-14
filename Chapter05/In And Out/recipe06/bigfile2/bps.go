package main

import (
	"fmt"
)

type (
	bps int
)

func (b bps) String() string {
	var units = []string{"B/s", "kB/s", "MB/s", "GB/s", "TB/s", "PB/s"}
	var i, u = float64(b), 0
	for i > 1000 && u < 5 {
		i = i / 1024
		u++
	}
	return fmt.Sprintf("%.3f %s", i, units[u])
}
