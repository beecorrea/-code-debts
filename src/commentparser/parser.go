package commentparser

import (
	"bufio"
	"log"

	"github.com/beecorrea/debt-checker/src/utils/fileutils"
)

type CommentParser struct {
	Path string 
	Reader *bufio.Scanner
}

func NewParser(p string) CommentParser{
	f, err := fileutils.OpenFile(p)
	
	if err != nil {
		log.Panicf("error opening file %v", err)
	}

	scanner := bufio.NewScanner(f)
	
	return CommentParser{
		Path: p,
		Reader: scanner,
	}
}

func (p CommentParser) GetCommentsFromFile() []string{
	comments := make([]string, 0)
	// Read lines
	for p.Reader.Scan() {
		text := p.Reader.Text()
		if IsComment(text) {
			comments = append(comments, text)
		}
	}

	return comments
}