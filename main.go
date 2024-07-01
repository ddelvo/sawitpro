package main

import (
	"fmt"
	"os"
)

func main() {
	// This is just a placeholder to show how to print to stderr and return a non-zero exit code
	// Please replace the whole content of this function with your solution.
	fail := false
	if fail {
		// On input faulty, print "FAIL" to stderr and exit with status 1
		fmt.Fprintln(os.Stderr, "FAIL")
		os.Exit(1)
	}
	// On input correct, print the result to stdout and exit with status 0
	output := 8
	fmt.Printf("%d\n", output)
}
