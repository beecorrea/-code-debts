package commentparser

import "bufio"

type CommentParser struct {
	Path string 
	Reader bufio.Scanner
}

func NewParser(p string, r bufio.Scanner) CommentParser{
	return CommentParser{
		Path: p,
		Reader: r,
	}
}

func (p CommentParser) GetCommentsFromFile(){

}