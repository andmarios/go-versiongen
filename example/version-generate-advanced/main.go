package main

import (
	"log"
	"time"

	"github.com/andmarios/go-versiongen"
)

func main() {
	// versiongen has some knobs you can turn and also supports a custom
	// path for the version file.

	// Set dirty string to +, so when you have uncommited changes, instead
	// of something like v0.1-1-g79d2733-dirty you will get v0.1-1-g79d2733+
	versiongen.DirtyString = "+"

	// Set Timeout to 30 second if your repo is so big, git operation take
	// too long to complete:
	versiongen.Timeout = 30 * time.Second

	// Set files to ignore when looking for uncommited changes:
	versiongen.IgnoreFiles = []string{"version.go", "autogen_template.go"}

	// Create version file to a custom path. We always start
	// from the directory we run go generate on:
	err = versiongen.CreateFile("version-generate/softversion.go")
	if err != nil {
		log.Fatalln(err)
	}
}
