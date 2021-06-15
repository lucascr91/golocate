package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var commands []string = os.Args
	if len(commands) == 1 {
		log.Fatal("Please, add a desired pattern and/or a flag")
	}
	var pattern string = os.Args[1]
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	if len(commands) == 2 {
		for _, match := range basicSearch(home, pattern) {
			fmt.Println(match)
		}
	} else if len(commands) == 3 {
		switch commands[2] {
		case "-r":
			for _, match := range filterRegex(basicSearch(home, "*"), pattern) {
				fmt.Println(match)
			}
		case "-c":
			for _, match := range filterContains(basicSearch(home, "*"), pattern) {
				fmt.Println(match)
			}
		}
	} else {
		fmt.Println("Please, add a desired pattern and/or a flag")
	}
}

func basicSearch(dir string, pattern string) []string {
	var result []string
	// get files in current directory
	matches, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		fmt.Println(err)
	}
	for _, match := range matches {
		result = append(result, match)
	}
	// get files in sub-directories
	var folders []string = listFolders(dir)
	if len(folders) > 0 {
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
			dirs = append(dirs, filepath.Join(dir, file.Name()))
		}
	}
	return dirs
}

func filterRegex(files []string, pattern string) []string {
	pattern = strings.Replace(pattern, "'", "", -1)
	r, _ := regexp.Compile(pattern)
	var result []string
	for _, file := range files {
		var a []string = strings.Split(file, "/")
		var chunk string = a[len(a)-1]
		if r.MatchString(chunk) {
			result = append(result, file)
		}
	}
	return result
}

func filterContains(files []string, pattern string) []string {
	var result []string
	for _, file := range files {
		var a []string = strings.Split(file, "/")
		var chunk string = a[len(a)-1]
		if strings.Contains(chunk, pattern) {
			result = append(result, file)
		}
	}
	return result
}
