package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[[INF]*[ERR]*[TRC]*[DBG]*[WRN]*[FTL]*\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`(<[~=\-\*]*>)+`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count = count + 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+[[:alnum:]]+\s{1}`)
	var regexStr string
	var userName string
	var found bool
	for i := range lines {
		if re.MatchString(lines[i]) {
			regexStr = re.FindString(lines[i])
			_, userName, found = strings.Cut(regexStr, " ")
			if found {
				lines[i] = fmt.Sprintf("[USR] %s %s", strings.Trim(userName, " "), lines[i])
			}
		}
	}
	return lines
}
