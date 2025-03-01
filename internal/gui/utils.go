package gui

import (
	"regexp"
	"strings"
	"time"
	"unicode"
)

func GetTimestamp() (ts string) {
	currentTime := time.Now()

	return currentTime.Format("2006.01.02 15:04:05")
}

// tk.GetOpenFile's output is inconsistent, extractFilePaths attempts to correct that
func extractFilePaths(pathStr []string) []string {
	var paths []string

	if len(pathStr) == 1 {
		input := pathStr[0]

		re := regexp.MustCompile(`\{([^}]+)\}`)
		matches := re.FindAllStringSubmatch(input, -1)
		for _, match := range matches {
			paths = append(paths, match[1])
			input = strings.Replace(input, match[0], "", 1)
		}

		remainingTokens := strings.Fields(input)
		var currentPath string
		for _, token := range remainingTokens {
			if isPathStart(token) {
				if currentPath != "" {
					paths = append(paths, currentPath)
				}
				currentPath = token
			} else {
				if currentPath == "" {
					currentPath = token
				} else {
					currentPath += " " + token
				}
			}
		}
		if currentPath != "" {
			paths = append(paths, currentPath)
		}

	} else {
		paths = pathStr
	}

	return paths

}

func isPathStart(token string) bool {
	if strings.HasPrefix(token, "/") {
		return true
	}
	if len(token) >= 2 && token[1] == ':' && unicode.IsLetter(rune(token[0])) {
		return true
	}
	return false
}
