package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		fmt.Printf("Too few arguments.")
		return
	}
	kw := os.Args[1] // keyword for path
	paths := filepath.SplitList(os.Getenv("PATH"))

	for i, p := range paths {
		if strings.Contains(strings.ToLower(p), strings.ToLower(kw)) {
			fmt.Printf("#%-2d : %q\n", i, p)
		}
	}
}
