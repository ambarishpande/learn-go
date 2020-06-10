package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type docToCount map[string]int
type tokenToDocToCount map[string]docToCount

func main() {

	dir := "documents"
	files, _ := ioutil.ReadDir(dir)
	tm := tokenToDocToCount{}
	docTotalCount := map[string]int{}
	for _, file := range files {
		fptr, _ := os.Open(string(dir + "/" + file.Name()))
		content, _ := ioutil.ReadAll(fptr)
		tokens := strings.Split(string(content), " ")
		docTotalCount[file.Name()] = len(tokens)
		for _, token := range tokens {
			_, status := tm[token][file.Name()]
			if status {
				tm[token][file.Name()]++
			} else {
				_, st := tm[token]
				if !st {
					tm[token] = docToCount{}
				}
				tm[token][file.Name()] = 1
			}
		}
	}

	fmt.Println("Doc Total Count", docTotalCount)
	// Calculate TF IDF of each keyword in each document

	for term, dc := range tm {
		docCount := len(dc)
		for doc, count := range dc {

			tf := float64(float64(count) / float64(docTotalCount[doc]))

			idf := float64(1 / float64(docCount))

			fmt.Println(doc, term, count, docTotalCount[doc], docCount, tf, idf, float64(tf*idf))
		}
	}
	// Take a directory
	// Read all files in the directory. each file will be one document
	// tokenize each file and calculate tf for all

}
