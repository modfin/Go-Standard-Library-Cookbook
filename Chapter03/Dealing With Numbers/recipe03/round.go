package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

var valA float64 = 3.55554444

func main() {

	// Bad assumption on rounding
	// the number by casting it to
	// integer.
	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v\n", intVal)

	fRound := Round(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	a, b, c, d := 1.1, -1.1, 1.6, -1.6
	fmt.Fprintf(w, "\t%f\t%f\t%f\t%f\t\n", a, b, c, d)
	fmt.Fprintf(w, "int\t%d\t%d\t%d\t%d\t\n", int(a), int(b), int(c), int(d))
	fmt.Fprintf(w, "math.Trunc\t%f\t%f\t%f\t%f\t\n", math.Trunc(a), math.Trunc(b), math.Trunc(c), math.Trunc(d))
	fmt.Fprintf(w, "Round\t%f\t%f\t%f\t%f\t\n", Round(a), Round(b), Round(c), Round(d))
	w.Flush()
}

// Round returns the nearest integer.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}
