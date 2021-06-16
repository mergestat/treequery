package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-enry/go-enry/v2"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/python"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	absPath, _ := filepath.Abs(path)
	fmt.Println("Path to file: " + absPath)
	f, inputErr := exists(absPath)
	if !f {
		handleErr(inputErr)
	}

	contents, _ := ioutil.ReadFile(absPath)
	contents = []byte(contents)
	if len(contents) <= 0 {
		handleErr(errors.New("empty or faulty file input"))
	}
	lang := enry.GetLanguage(absPath, contents)
	if len(lang) <= 0 {
		handleErr(errors.New("language could not be detected"))
	}
	parser := sitter.NewParser()
	switch strings.ToLower(lang) {
	case "javascript":
		parser.SetLanguage(javascript.GetLanguage())
	case "go":
		parser.SetLanguage(golang.GetLanguage())
	case "python":
		parser.SetLanguage(python.GetLanguage())
	case "java":
		parser.SetLanguage(java.GetLanguage())
	default:
		handleErr(errors.New("language not supported at this time"))
	}

	tree := parser.Parse(nil, contents)
	n := tree.RootNode()
	fmt.Println("language: " + lang)
	fmt.Println(n)
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

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
