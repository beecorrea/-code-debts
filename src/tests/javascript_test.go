package tests

import (
	"log"
	"testing"

	"github.com/beecorrea/debt-checker/src/commentparser"
)

func TestJavascriptFile(t *testing.T) {
	path := "mocks/javascript.js"
	parser := commentparser.NewParser(path)	

	debts := []string{"DEBT use hello as parameter", "DEBT return hello sentence"}

	comments := parser.GetCommentsFromFile()

	if len(comments) == 0 {
		log.Panicf("no comments in file")
	}

	for i, c := range comments {
		t.Run("should read comments from JS file", func (t *testing.T){
			debt := commentparser.GetCommentFromLine(c)
			
			if debt != debts[i] {
				t.Errorf("expected %v, received %v", debts[i], debt)
			}
		})
	}
}