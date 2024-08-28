package main

import (
	"strings"
)

func simplifyPath(path string) string {
	fields := strings.Split(path, "/")
	s := []string{}
	for _, f := range fields {
		if f == "" || f == "." {
			continue
		}
		switch f {
		case "..":
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		default:
			s = append(s, f)
		}
	}
	builder := strings.Builder{}
	for i := 0; i < len(s); i++ {
		builder.WriteByte('/')
		builder.WriteString(s[i])
	}
	r := builder.String()
	if r == "" {
		return "/"
	}
	return r
}
