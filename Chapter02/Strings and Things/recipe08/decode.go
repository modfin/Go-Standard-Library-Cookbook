package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

type transformReader struct {
	io.Reader
	f *os.File
}

func (t transformReader) Close() error { return t.f.Close() }

func open(fileName string, charmap *charmap.Charmap) (io.ReadCloser, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	if charmap == nil {
		return f, nil
	}
	return &transformReader{Reader: charmap.NewDecoder().Reader(f), f: f}, nil
}

func main() {

	// Open windows-1250 file.
	f, err := open("win1250.txt", nil)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read all in raw form.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("Without decode: " + string(b))

	// Decode to unicode
	f2, err := open("win1250.txt", charmap.Windows1250)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	b, err = ioutil.ReadAll(f2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded: " + string(b))

}
