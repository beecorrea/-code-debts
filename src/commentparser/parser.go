package commentparser

import (
	"bufio"
)

type CommentParser struct {
	Path string 
	Reader *bufio.Scanner
}

func NewParser(p string, r *bufio.Scanner) CommentParser{
	return CommentParser{
		Path: p,
		Reader: r,
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