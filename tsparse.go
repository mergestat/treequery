package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-enry/go-enry/v2"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	absPath, _ := filepath.Abs(path)
	fmt.Println("Path to file: " + absPath)
	f, err := exists(absPath)
	if !f {
		fmt.Println("There was either an error in the command line input or a faulty filepath. Please try again.")
		log.Fatal(err)
	}
	contents, _ := ioutil.ReadFile(absPath)
	lang := enry.GetLanguage(absPath, []byte(contents))
	fmt.Println("language: " + lang)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
