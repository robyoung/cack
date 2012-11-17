package main

import (
	"bufio"
	"path/filepath"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: cack [pattern] [paths..]\n")
		os.Exit(1)
	}
	pattern := regexp.MustCompile(os.Args[1])
	paths := os.Args[2:]
	if len(paths) == 0 {
		paths = append(paths, ".")
	}

	for _, path := range paths {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("Failed to open file: %v\n", path)
			} else {
				reader := bufio.NewReaderSize(file, 4092)
				for i := 1; ; i++ {
					line, err := reader.ReadString('\n')
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
