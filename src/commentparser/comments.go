package commentparser

import (
	"strings"
)

const (
	COMMENT = "//"
	DEBT = "DEBT"
)

func GetCommentFromLine(line string) string {
	comment := searchComment(line) 
	return comment
}

func IsComment(line string) bool {
	return strings.Contains(line, COMMENT)
}

func searchComment(line string) string{
	lineArray := clearSpacesFromLine(line)
	
	if(IsComment(lineArray[0]) && lineArray[1] == DEBT){ 
		return buildComment(lineArray[1:])
	}
	
	return ""	
}

func buildComment(words []string) string {
	return strings.Join(words, " ")
}

func clearSpacesFromLine(line string) []string{
	return strings.Split(line, " ")
}

