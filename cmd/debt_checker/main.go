package main

import (
	"fmt"
	"log"

	"github.com/beecorrea/debt-checker/src/commentparser"
)


func main(){
	path := "mocks/javascript.js"

	parser := commentparser.NewParser(path)
	
	comments := parser.GetCommentsFromFile()

	if len(comments) == 0 {
		log.Panicf("no comments in file")
	}

	for _, c := range comments {
		debt := commentparser.GetCommentFromLine(c)
		fmt.Printf("%v\n", debt)
	}
}