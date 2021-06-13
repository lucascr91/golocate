package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var commands []string = os.Args
	if len(commands) == 2 {
		var pattern string = os.Args[1]
		home, _ := os.UserHomeDir()
		files, _ := os.ReadDir(home)
		for _, match := range wildCard(home, files, pattern) {
			fmt.Println(match)
		}
	}
}

func wildCard(home string, files []os.DirEntry, pattern string) []string {
	var result []string
	for _, file := range files {
		if file.IsDir() {
			matches, _ := filepath.Glob(filepath.Join(home, file.Name(), pattern))
			for _, match := range matches {
				result = append(result, match)
			}
		}
	}
	return result
}
