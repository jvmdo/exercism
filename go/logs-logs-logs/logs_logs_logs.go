package logs

import "strings"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	var app string

	for _, char := range log {
		switch {
			case char == '\u2757': // ❗
				app = "recommendation"
			case char == '\U0001F50D': // 🔍
				app = "search"
			case char == '\u2600': // ☀
				app = "weather"
			default:
				app = "default"
		}

		if app != "default" {
			break
		}
	}

	return app

	// Or...
	// appMap := map[rune]string{
	// 	'\u2757': "recommendation",
	// 	'\U0001F50D': "search",
	// 	'\u2600': "weather",
	// }

	// for _, char := range log {
	// 	if value, exists := appMap[char]; exists {
	// 		return value
	// 	}
	// }

	// return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string = log

	for _, char := range log {
		if char == oldRune {
			newLog = strings.Replace(log, string(oldRune), string(newRune), -1)
		}
	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
