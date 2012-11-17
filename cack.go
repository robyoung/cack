package main

import (
	"bufio"
	"path/filepath"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Verify we have enough arguments and exit if we don't
	if len(os.Args) < 2 {
		fmt.Printf("Usage: cack [pattern] [paths..]\n")
		os.Exit(1)
	}
	// Compile the first argument as a regular expression.
	// This will panic and exit if it can't be compiled.
	pattern := regexp.MustCompile(os.Args[1])

	// Store a slice of all subsequent arguments, adding a default
	// of the current directory if none are given.
	paths := os.Args[2:]
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	for _, path := range paths {
		// Walk the filesystem for each path, calling the annonymous
		// function provided for each subpath.
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			// Open the path
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("Failed to open file: %v\n", path)
			} else {
				// Read \n terminated lines from the path until there is an error
				reader := bufio.NewReaderSize(file, 4092)
				for i := 1; ; i++ {
					line, err := reader.ReadString('\n')
					// If a line matches the pattern then print the result.
					if pattern.MatchString(line) {
						fmt.Printf("%v:%v:%v", path, i, line)
					}
					if err != nil {
						break
					}
				}
			}
			return nil
		})
	}
}
