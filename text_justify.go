package livetest

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	line := []string{}
	lineLength := 0

	for i := 0; i < len(words); i++ {
		if lineLength+len(line)+len(words[i]) > maxWidth {
			spaces := maxWidth - lineLength
			if len(line) == 1 {
				result = append(result, line[0]+strings.Repeat(" ", spaces))
			} else {
				for j := 0; j < spaces; j++ {
					line[j%(len(line)-1)] += " "
				}
				result = append(result, strings.Join(line, ""))
			}

			line = []string{}
			lineLength = 0
		}

		line = append(line, words[i])
		lineLength += len(words[i])
	}

	lastLine := strings.Join(line, " ")
	lastLine += strings.Repeat(" ", maxWidth-len(lastLine))
	result = append(result, lastLine)

	return result
}
