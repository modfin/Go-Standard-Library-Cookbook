package main

import (
	"runtime"
	"runtime/debug"

	"github.com/rs/zerolog/log" // Suck in a new logger, just to have a dependency to log..
)

const info = `
Application %s starting.
The binary was build by GO: %s`

func dumpModule(m *debug.Module, indent ...string) {
	if m == nil {
		return
	}
	var pf string
	if len(indent) > 0 {
		pf = indent[0]
	}
	log.Printf("  %sPath: %s", pf, m.Path)
	log.Printf("  %sVersion: %s", pf, m.Version)
	log.Printf("  %sChecksum: %s", pf, m.Sum)
	dumpModule(m.Replace, pf+"  ")
}

func main() {
	log.Printf(info, "Example", runtime.Version())

	// Add ReadBuildInfo (added in GO 1.12) to log output
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Print("Failed to read BuildInfo")
	}
	log.Printf("Main package path: %s", bi.Path)
	log.Print("Main module:")
	dumpModule(&bi.Main)
	for i, dep := range bi.Deps {
		log.Printf("Dependency #%d", i)
		dumpModule(dep)
	}
}
