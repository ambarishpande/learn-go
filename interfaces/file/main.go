package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide name of the file to be read.")
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	fmt.Println("Reading File: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
