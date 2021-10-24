package main

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// Epsilon float compare const
const TOLERANCE = 1e-8

func TestRound(t *testing.T) {
	// Try out propertybased testing in GO using testing/quick package: https://pkg.go.dev/testing/quick

	// Test function that validate that differance between "our" Round and math.Round is < 1e-8
	f := func(x float64) bool {
		t.Logf("Round(%f)\t = %f", x, Round(x))
		return math.Abs(Round(x)-math.Round(x)) < TOLERANCE
	}

	// Function that generate small testvalues between -1 and +1.
	tv := func(args []reflect.Value, r *rand.Rand) {
		// Guessing there can be a better way to fill args array..
		floatType := reflect.TypeOf(0.0)
		for i := 0; i < len(args); i++ {
			v := reflect.New(floatType).Elem()
			v.SetFloat(r.Float64()*2 - 1)
			args[i] = v
		}
	}

	// By default, quick.Check will create very large numbers with no room for decimals
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}

	// Since we are testing two different round methods, use our test values generator that generate small values
	if err := quick.Check(f, &quick.Config{Values: tv}); err != nil {
		t.Error(err)
	}

	// And just for kicks, also test to use quick.CheckEqual
	if err := quick.CheckEqual(Round, math.Round, &quick.Config{Values: tv}); err != nil {
		t.Error(err)
	}
}
