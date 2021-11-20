package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/beecorrea/debt-checker/src/commentparser"
	"github.com/beecorrea/debt-checker/src/utils/fileutils"
)


func main(){
	path := "test1.txt"

	f, err := fileutils.OpenFile(path)
	
	if err != nil {
		log.Panicf("error opening file %v", err)
	}

	scanner := bufio.NewScanner(f)

	parser := commentparser.NewParser(path, scanner)

	comments := parser.GetCommentsFromFile()

	if len(comments) == 0 {
		log.Panicf("no comments in file")
	}

	for _, c := range comments {
		debt := commentparser.GetCommentFromLine(c)
		fmt.Printf("%v\n", debt)
	}
}