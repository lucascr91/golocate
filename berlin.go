package main

import (
	"fmt"
	"os"
)

func main() {
         // Open current directory

         f, err := os.Open(".")
         if err != nil {
                 panic(err)
         }

         // Get file names
         files, err := f.Readdirnames(0)
         if err != nil {
                 panic(err)
         }

         // show files
         for _, v := range files {
                 fmt.Println(v)
         }
}
