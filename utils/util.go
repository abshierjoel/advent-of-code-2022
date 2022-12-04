package utils

import (
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) []string {
	file, _ := ioutil.ReadFile(filename)
	file_string := string(file)
	return strings.Split(file_string, "\n")
}
