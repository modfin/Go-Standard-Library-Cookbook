// +build freebsd linux

package main

import (
	"os"

	"golang.org/x/sys/unix"
)

func diskInfo() (total, free int64, err error) {
	var stat unix.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	// https://pkg.go.dev/golang.org/x/sys@v0.0.0-20211112193437-faf0a1b62c6b/unix#Statfs
	err = unix.Statfs(wd, &stat)
	if err != nil {
		return
	}

	// fmt.Printf("statfs: %#v", stat)
	// https://pkg.go.dev/golang.org/x/sys@v0.0.0-20211112193437-faf0a1b62c6b/unix#Statfs_t
	return int64(stat.Blocks * stat.Bsize), int64(stat.Bfree * stat.Bsize), nil
	// return 0, 0, nil
}
