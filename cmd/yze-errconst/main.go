// Command yze-go-errconst runs the errconst analyzer as a standalone go/analysis
// checker (text, -json, and -fix output, and usable as a `go vet -vettool`).
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	errconst "github.com/gomatic/yze-errconst"
)

// run is the analysis entry point, indirected so the binary's wiring is testable
// without invoking the real driver (which loads packages and exits the process).
var run = singlechecker.Main

func main() { run(errconst.Analyzer) }
