package livetest

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0)
	for _, s := range strings.Split(path, "/") {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if s != "." && s != "" {
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}
