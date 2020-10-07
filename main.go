package main

import (
	"os"
	"strconv"
)

func main() {
	os.Stderr.WriteString("stderr\n")
	os.Stdout.WriteString("stdout\n")

	argsWithoutProg := os.Args[1:]
	retcode := 0
	if len(argsWithoutProg) != 0 {
		retcode, _ = strconv.Atoi(argsWithoutProg[0])
	}
	os.Exit(retcode)
}
