package main

import (
	"chapter01-recipe03/subpackage"
	"flag"
	"fmt"
	"sync"
)

func main() {
	var f0 string
	var f1 string

	/*
	 * flag package is very stateful
	 * We need to take note that we're binding pointers below.
	 * Any action/mutation on the flag package may now affect these variables
	 */

	flag.StringVar(&f0, "f0", "", "Flag 0")
	flag.StringVar(&f1, "f1", "", "Flag 1")

	_ = flag.Set("f0", "mutate before Parse()")

	fmt.Printf("before parse\t\t\tf0=\"%s\"\tf1=\"%s\"\n", f0, f1)

	flag.Parse()

	_ = flag.Set("f1", "mutate after Parse()")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		subpackage.IPromiseThisIsATotallyPureFunction()
		fmt.Printf("async end (after another Parse)\tf0=\"%s\"\t\t\tf1=\"%s\"\n", f0, f1)
	}()

	fmt.Printf("sync end\t\t\tf0=\"%s\"\t\t\tf1=\"%s\"\n", f0, f1)

	wg.Wait()
}
