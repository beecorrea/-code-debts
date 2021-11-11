package commentparser

import (
	"strings"
)

const (
	COMMENT = "//"
)

func GetCommentFromLine(line string) string {
	comment := SearchComment(line) 
	return comment
}

func SearchComment(line string) string{
	lineArray := ClearSpacesFromLine(line)
	if(lineArray[0] == COMMENT){
		return BuildComment(lineArray)
	}
	
	return ""	
}

func BuildComment(words []string) string {
	return strings.Join(words, " ")
}

func ClearSpacesFromLine(line string) []string{
	return strings.Split(line, " ")
}
