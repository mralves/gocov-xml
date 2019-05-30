package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlekSi/gocov-xml"
)

func main() {
	sourcePathPtr := flag.String(
		"source",
		"",
		"Absolute path to source. Defaults to current working directory.",
	)

	flag.Parse()

	// Parse the commandline arguments.
	var sourcePath string
	var err error
	if *sourcePathPtr != "" {
		sourcePath = *sourcePathPtr
		if !filepath.IsAbs(sourcePath) {
			panic(fmt.Sprintf("Source path is a relative path: %s", sourcePath))
		}
	} else {
		sourcePath, err = os.Getwd()
		if err != nil {
			panic(err)
		}
	}

	err = gocovxml.Parse(sourcePath)
	if err != nil {
		panic(err)
	}
}
