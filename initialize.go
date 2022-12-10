package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	day := "day" + args[0]
	path := "./" + day + "/"
	filename := day + ".go"
	testname := day + "_test.go"

	os.Mkdir("./"+day, os.ModePerm)
	ioutil.WriteFile(path+filename, []byte(content), fs.ModePerm)
	ioutil.WriteFile(path+testname, []byte(test), fs.ModePerm)

	fmt.Println(day)
}

const content = `package main

func main() {

}`

const test = `package main

import "testing"

TestFunc(t *testing.T)  {

}`
