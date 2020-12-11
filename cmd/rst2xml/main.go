package main

import (
	"os"

	pygolin_runtime "github.com/pygolin/runtime"
	_ "github.com/pygolin/stdlib"

	mod "github.com/pygolin/pypi-docutils/cmd/rst2xml/__main__"
)

func main() {
	pygolin_runtime.ImportModule(pygolin_runtime.NewRootFrame(), "traceback")
	os.Exit(pygolin_runtime.RunMain(mod.Code))
}
