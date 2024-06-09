package site

import "strings"

func CheckPageDepth(url string, depth int) bool {
	items := strings.Split(url, "/")
	if len(items) >= depth {
		return false
	}
	return true
}
