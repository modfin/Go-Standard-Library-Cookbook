package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Default date is: %v\n", t)

	// Add 3 days
	r1 := t.Add(72 * time.Hour) // Time Add Duration -> Time
	fmt.Printf("Default date +3D is: %v\n", r1)

	// Subtract 3 days
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Default date -3D is: %v\n", r1)

	// A Duration is just a typed int64 with some added formatting
	h := time.Hour
	m := time.Minute
	fmt.Printf("time.Duration + time.Duration = %v\n", reflect.TypeOf(h+m))
	fmt.Printf("time.Duration - time.Duration = %v\n", reflect.TypeOf(h-m))
	fmt.Printf("A quarter = %v\n", h/4)
	fmt.Printf("Another quarter = %v\n", m*15)

	// More comfortable api
	// to add days/months/years
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1YR +3MTH +2D is: %v\n", r1)
}
