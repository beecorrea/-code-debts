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
	if cmtStart := hasComment(lineArray); cmtStart != -1{ 
		return buildComment(lineArray[cmtStart:])
	}
	
	return ""	
}

func buildComment(words []string) string {
	return strings.Join(words, " ")
}

func clearSpacesFromLine(line string) []string{
	return strings.Split(line, " ")
}

func hasComment(line []string) int{
	for i, word := range line{
		if word == COMMENT && line[i+1] == DEBT {
			return i+1
		}
	}

	return 0
}
