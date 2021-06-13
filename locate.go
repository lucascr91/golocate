package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	os.Chdir(filepath.Join(home, ""))
	matches, _:= filepath.Glob("*pdf")
	for _, match := range matches {
		fmt.Println(match)
	}

}
