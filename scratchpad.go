package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	dir := "documents"
	files, _ := ioutil.ReadDir(dir)
	fmt.Println(files)
	// Take a directory
	// Read all files in the directory. each file will be one document
	// tokenize each file and calculate tf for all

}
