package main

import (
	"fmt"
	"time"
)

func main() {
	preTotal, preFree, err := diskInfo()
	if err != nil {
		fmt.Printf("Failed to get initial disk info, error: (%T)%[1]v\n", err)
		return
	}

	start := time.Now()
	err = createBigFile(2 * 1024 * 1024 * 1024) // Limit test to a 2GB file to be able to test on FAT
	d := time.Since(start)
	if err != nil {
		fmt.Printf("Failed to create file, error: %v\n", err)
		return
	}
	fmt.Printf("File created in %v\n", d)

	postTotal, postFree, err := diskInfo()
	if err != nil {
		fmt.Printf("Failed to get post disk info, error: (%T)%[1]v\n", err)
		return
	}
	fmt.Println("             Total Bytes       Free Bytes")
	fmt.Printf("Before: %16d %16d\n", preTotal, preFree)
	fmt.Printf(" After: %16d %16d\n", postTotal, postFree)
	fmt.Printf("  Diff: %16d %16d\n", postTotal-preTotal, postFree-preFree)

	readSpeed, err := readBench()
	if err != nil {
		fmt.Printf("Failed to do read benchmark, error: %v\n", err)
		return
	}
	fmt.Printf("Read speed: %s\n", readSpeed)
}

// Windows, NTFS:
// File created in 515.8µs
//              Total Bytes       Free Bytes
// Before:    1000186310656     867116195840
//  After:    1000186310656     864968712192
//   Diff:                0      -2147483648
// Read speed: 3.874 GB/s

// FreeBSD, ZFS:
// File created in 215.08µs
//              Total Bytes       Free Bytes
// Before:     418245324800     416976048128
//  After:     418245324800     416976048128
//   Diff:                0                0
// Read speed: 16.409 GB/s

// FreeBSD, 20GB MD/UFS on ZFS:
// File created in 5.332575ms
//              Total Bytes       Free Bytes
// Before:      20795650048      20762025984
//  After:      20795650048      20761956352
//   Diff:                0           -69632
// Read speed: 42.888 GB/s

// FreeBSD, 20GB MD/FAT32 on ZFS:
// File created in 22.262692108s
//              Total Bytes       Free Bytes
// Before:      21469626368      21469528064
//  After:      21469626368      19322044416
//   Diff:                0      -2147483648
// Read speed: 5.623 GB/s
