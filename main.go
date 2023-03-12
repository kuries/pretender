package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"github.com/kuries/pretender/app"
)

func main() {
	exitcode.Exit(app.CLI(os.Args[1:]))
}
