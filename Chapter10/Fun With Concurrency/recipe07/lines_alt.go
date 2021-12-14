package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"sync"
)

const data = `line one
line two with more words
error: This is erroneous line
error: What if there is more then one error
Nah, all's good over here`

type goRes struct {
	data string
	err  error
}

func main() {
	log.Printf("Application %s starting.", "Error Detection")
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	var wg sync.WaitGroup
	var errs = make(chan goRes)

	for scanner.Scan() {
		wg.Add(1)
		go func(s string) (err error) {
			defer func() {
				if err != nil {
					errs <- goRes{data: s, err: err}
				}
				wg.Done()
			}()
			if strings.Contains(s, "error:") {
				return fmt.Errorf(s)
			}
			return nil
		}(scanner.Text())
	}
	go func() {
		wg.Wait()
		close(errs)
	}()

	for err := range errs {
		fmt.Printf("Error: %v, while processing row: %s\n", err.err, err.data)
	}
}
