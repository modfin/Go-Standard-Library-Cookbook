package main

import (
	"fmt"
	"os"
)

func createBigFile(size int64) (err error) {
	// Create a 10GB file by writing one byte
	var f *os.File
	f, err = os.OpenFile("bigfile.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("createBigFile: OpenFile error, %w", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			err = fmt.Errorf("createBigFile: Close file error, %w", err)
		}
	}()

	_, err = f.WriteAt([]byte{0}, size-1)
	if err != nil {
		err = fmt.Errorf("createBigFile: WriteAt error, %w", err)
	}
	return nil
}
