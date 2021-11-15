package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

type walkerFunc func(string, error) (int, int)

var (
	walkers = map[string]walkerFunc{
		"walker": readDirWalker,
		"texas":  filepathWalker,
		"ranger": chuckNorris,
	}
)

// Score:
// readDirWalker:  135841 directories and 1867460 files scanned in 39.845354462s
// filepathWalker: 135841 directories and 1867460 files scanned in 40.622839502s
// ChuckNorris:    135841 directories and 1867460 files scanned in 7.665811978s

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must specify walker type to use for test")
		return
	}
	dirWalker, ok := walkers[os.Args[1]]
	if !ok {
		fmt.Printf("%s not known\n", os.Args[1])
		return
	}
	start := time.Now()
	dirs, files := dirWalker(os.Getwd())
	stop := time.Since(start)
	fmt.Printf("%d directories and %d files scanned in %v\n", dirs, files, stop)
}

func readDirWalker(root string, err error) (int, int) {
	var dc, fc int
	var dw func(string)
	dw = func(dir string) {
		lst, err := ioutil.ReadDir(dir)
		if err != nil {
			panic(err)
		}
		for _, val := range lst {
			if val.IsDir() {
				dc++
				dw(path.Join(dir, val.Name()))
			} else {
				fc++
			}
		}
	}
	dw(root)
	return dc, fc
}

func filepathWalker(root string, err error) (int, int) {
	var dc, fc int
	filepath.Walk(root, func(wPath string, info os.FileInfo, err error) error {
		if wPath == root {
			return nil
		}

		if info.IsDir() {
			dc++
			return nil
		}

		if wPath != root {
			fc++
		}
		return nil
	})
	return dc, fc
}

func chuckNorris(root string, err error) (int, int) {
	var dc, fc int
	filepath.WalkDir(root, func(wPath string, info fs.DirEntry, err error) error {
		if wPath == root {
			return nil
		}

		if info.IsDir() {
			dc++
			return nil
		}

		if wPath != root {
			fc++
		}
		return nil
	})
	return dc, fc
}
