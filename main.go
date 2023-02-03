// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()

	if err != nil {
		fail(err)
	}
	res := wordcount(src)
	fmt.Println(res)
}

// match returns true if src matches pattern,
// false otherwise.
func wordcount(src string) int {
	arr := strings.Split(src, " ")
	return len(arr)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	flag.Parse()
	myVal := strings.Join(flag.Args(), " ")
	if myVal == "" {
		return myVal, errors.New("missing string to match")
	}
	return myVal, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
