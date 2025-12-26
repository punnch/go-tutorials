package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func cathPanic() {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			fmt.Println(err)
		} else {
			panic(r)
		}
	}
}

func scanDirectory(path string) {
	fmt.Println(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fullpath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(fullpath)
		} else {
			fmt.Println(fullpath)
		}
	}
}

func main() {
	defer cathPanic()
	scanDirectory(".")
}
