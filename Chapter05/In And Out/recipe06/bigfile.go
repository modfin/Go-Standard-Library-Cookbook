package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a 10GB file by writing one byte
	f, e := os.OpenFile("bigfile.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	size := 10 * 1024 * 1024 * 1024
	n, err := f.WriteAt([]byte{0}, int64(size-1))
	if err != nil {
		panic(e)
	}
	fmt.Printf("Written %d bytes", n)
}
