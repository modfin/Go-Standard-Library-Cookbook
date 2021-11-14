package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func readBench() (speed bps, err error) {
	var f *os.File
	f, err = os.Open("bigfile.txt")
	if err != nil {
		return 0, fmt.Errorf("readBench: OpenFile error, %w", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			err = fmt.Errorf("readBench: Close file error, %w", err)
		}
	}()

	var buf = make([]byte, 1024*1024)
	var n, read int
	start := time.Now()
	for err == nil {
		n, err = f.Read(buf)
		read += n
	}
	d := time.Since(start)
	if err == io.EOF {
		err = nil
	}
	return bps(float64(read) / d.Seconds()), err
}
