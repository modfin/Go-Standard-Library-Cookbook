package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

var (
	errLog  *log.Logger
	binary  string
	appPath string
	fMode   *string
	logf    func(string, ...interface{})
	data    = []string{"Mechanic: Somebody set up us the bomb", "Operator: Main screen turn on", "CATS: All your base are belong to us", "CATS: You have no chance to survive make your time", "Captain: Move 'ZIG'", "Captain: For great justice"}
)

func init() {
	errLog = log.New(os.Stderr, "", log.Ltime|log.Lshortfile)
	logf = log.Printf
}

func getApp() {
	var err error
	binary, err = os.Executable()
	if err != nil {
		errLog.Fatalf("Error: %v", err)
	}
	appPath, err = filepath.EvalSymlinks(filepath.Dir(binary))
	if err != nil {
		errLog.Fatalf("Error: %v", err)
	}
}

func setupFlags() {
	fMode = flag.String("mode", "", "")
	flag.Parse()
}

func parent() {
	getApp()
	os.Setenv("PPID", strconv.Itoa(os.Getpid()))
	logf("Binary: %s", binary)
	logf("AppPath: %s", appPath)
	ctx, cancel := context.WithCancel(context.Background())

	smurf := exec.CommandContext(ctx, binary, "-mode", "smurf")
	chInp, _ := smurf.StdinPipe()
	chOut, _ := smurf.StdoutPipe()
	s := bufio.NewScanner(chOut)
	smurf.Start()
	for _, str := range data {
		chInp.Write([]byte(str + "\n"))
		s.Scan()
		logf("%s", s.Text())
	}
	cancel()
}

func checkChildEnv() {
	if os.Getenv("PPID") != strconv.Itoa(os.Getppid()) {
		errLog.Fatal("Should only be run from within itself")
	}
}

func smurf() {
	checkChildEnv()
	inp := bufio.NewScanner(os.Stdin)
	for inp.Scan() {
		fmt.Printf("%s SMURF!\n", inp.Text())
	}
}

func main() {
	setupFlags()
	switch *fMode {
	case "":
		parent()
	case "smurf":
		smurf()
	}
}
