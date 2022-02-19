package server

import (
	"errors"
	"regexp"
)

var tagPattern = regexp.MustCompile("\\$[{](\\w+)[}]")

func replaceAll(text string, values map[string]string) (string, error) {
	var output string
	parts := tagPattern.Split(text, -1)
	tags := tagPattern.FindAll([]byte(text), -1)

	output = parts[0]

	for index, tag := range tags {
		if value, ok := values[string(tag)]; ok {
			output += value + parts[index+1]
		} else {
			return "", errors.New("Tag `" + string(tag) + "` not found")
		}
	}

	return output, nil
}
