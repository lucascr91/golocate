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
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}
		for _, match := range basicSearch(home, pattern) {
			fmt.Println(match)
		}
	}
}

func basicSearch(dir string, pattern string) []string {
	var result []string
	// get files in current directory
	topMatches, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		fmt.Println(err)
	}
	for _, match := range topMatches {
		result = append(result, match)
	}
	// get files in sub-directories
	var folders []string = listFolders(dir)
	if len(folders)>0 {
		for _, folder := range folders {
			for _, subfile := range basicSearch(folder, pattern) {
				result = append(result, subfile)
			}
		}
	}

	return result
}

func listFolders(dir string) []string {
	var dirs []string		
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		if file.IsDir() {
			dirs=append(dirs, filepath.Join(dir, file.Name()))
		}
	}
	return dirs
}
