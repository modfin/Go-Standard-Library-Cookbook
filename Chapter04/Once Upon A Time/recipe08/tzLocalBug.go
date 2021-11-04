package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
	// _ "time/tzdata" // Issue not affected if tz data is embedded or not, as far as I could see
)

// From zoneinfo_windows.go in the standard lib:
// BUG(brainman,rsc): On Windows, the operating system does not provide complete
// time zone information.
// The implementation assumes that this year's rules for daylight savings
// time apply to all previous and future years as well.

func main() {
	local, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {
		panic(err)
	}

	// Before 1996, summertime ended in September in Sweden, and in 1995,
	// summertime ended 24th of September, so 1st of October 1995 used CET.
	// https://www.timeanddate.com/time/change/sweden/stockholm
	utcTime := time.Date(1995, 10, 01, 12, 00, 00, 00, time.UTC)

	fmt.Println(utcTime.In(local))
	fmt.Println(utcTime.Local())

	if strings.HasPrefix(runtime.GOOS, "windows") {
		// Replace default `time.Local` instance, with the
		// `time.LoadLocation` instance, and try again...
		time.Local = local
		fmt.Println(utcTime.Local())
	}

	// Output in Windows:
	// 1995-10-01 13:00:00 +0100 CET
	// 1995-10-01 14:00:00 +0200 CEST
	// 1995-10-01 13:00:00 +0100 CET

	// Output in FreeBSD:
	// 1995-10-01 13:00:00 +0100 CET
	// 1995-10-01 13:00:00 +0100 CET
}
