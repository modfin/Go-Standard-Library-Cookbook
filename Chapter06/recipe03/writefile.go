package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("sample.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go is awesome!\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Yeah! Go is great.\n"))
	if err != nil {
		panic(err)
	}

	// Add more examples of how to write to the file..
	_, err = fmt.Fprintf(f, "This %s works\n", "also")
	if err != nil {
		panic(err)
	}

	bw := bufio.NewWriter(f)
	_, err = bw.WriteString("bufiowriter was here! 😎\n")
	if err != nil {
		panic(err)
	}
	// Since bw is buffered and may not have been sent all data to the underlying io.Writer,
	// we must call flush when we are done with the all writes.
	err = bw.Flush()
	if err != nil {
		panic(err)
	}

	// There is a lot of other ways we can write to the file, anything that can take an io.Writer
	// can be used to write to the file. We can for example use the charmap package to use a specific
	// encofing when we write the file, by wrapping the writer in a encoder, something like this:
	// newWriter := charmap.Windows1252.NewEncoder().Writer(f)
}
