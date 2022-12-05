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

func CopyMap(original map[int][]string) map[int][]string {
	newMap := map[int][]string{}
	for i, row := range original {
		var newRow []string
		newRow = append(newRow, row...)

		newMap[i] = newRow
	}
	return newMap
}
