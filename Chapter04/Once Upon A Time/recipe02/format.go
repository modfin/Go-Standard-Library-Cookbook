package main

import (
	"fmt"
	"time"
)

/**
 * Order is according to ANSI C standardized format (which also _happens_ to be quite close to how US people format time):
 * Month	Day of month	Hour	Minute	Second	Year	Timezone	Fraction
 * 1			2							3			4				5				6			7					0 or 9
 */

func main() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 10000000, time.Local)

	// The formatting is done
	// with use of reference value
	// Jan 2 15:04:05 2006 MST
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))

	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))

	//The predefined formats could
	// be used
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))

	// The formatting supports space padding
	//only for days in Go version 1.9.2
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/_2"))

	// The zero padding is done by adding 0
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))

	//The fraction with ~~leading~~trailing zeros use 0s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.000"))

	//The fraction without ~~leading~~trailing zeros use 9s
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))

	// Append format appends the formatted time to given
	// buffer
	fmt.Println(string(tTime.AppendFormat([]byte("The time is up: "), "03:04PM")))
}
