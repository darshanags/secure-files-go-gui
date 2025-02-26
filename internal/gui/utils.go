package gui

import (
	"regexp"
	"strings"
	"time"
)

func GetTimestamp() (ts string) {
	currentTime := time.Now()

	return currentTime.Format("2006.01.02 15:04:05")
}

func extractFilePaths(pathStr []string) []string {

	re := regexp.MustCompile(`\{([^}]+)\}|([^ ]+)`)
	matches := re.FindAllString(pathStr[0], -1)

	filePaths := make([]string, 0, len(matches))
	for _, match := range matches {
		filePaths = append(filePaths, strings.Trim(match, "{} "))
	}

	return filePaths
}
