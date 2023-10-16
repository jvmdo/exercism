package parsinglogfiles

import "regexp"
import "fmt"

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	if err != nil {
		return false
	}

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`<[*~=-]*>+`)

	if err != nil {
		return nil
	}

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`(?i)".*?(password).*"`)

	if err != nil {
		return -1
	}

	var count int

	for _, line := range lines {
		if len(re.FindStringSubmatch(line)) > 0 {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`end-of-line\d+`)

	if err != nil {
		return text
	}

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User\s+(\w+)`)

	if err != nil {
		return nil
	}

	for i, line := range lines {
		if (re.MatchString(line)) {
			lines[i] = fmt.Sprintf("[USR] %s %s", re.FindStringSubmatch(line)[1], line)
		}
	}

	return lines
}
