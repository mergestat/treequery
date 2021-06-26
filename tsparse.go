package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

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
	query := flag.Arg(1)
	absPath, _ := filepath.Abs(path)
	fmt.Println("Path to file: " + absPath)
	f, _ := exists(absPath)
	if !f {
		handleErr(errors.New("there was either an error in the command line input or a faulty filepath"))
	}

	contents, readErr := ioutil.ReadFile(absPath)
	handleErr(readErr)
	lang := enry.GetLanguage(absPath, contents)
	fmt.Println("language: " + lang)
	if lang == "" {
		handleErr(errors.New("language could not be detected"))
	}
	parser, grammar := getTSParser(lang)
	tree := parser.Parse(nil, contents)
	n := tree.RootNode()

	fmt.Println("AST:", n)

	fmt.Println("Root type:", n.Type())
	fmt.Println("Root children:", n.ChildCount())

	fmt.Println("\nFunctions in input:")
	q, errQuery := sitter.NewQuery([]byte(query), &grammar)
	handleErr(errQuery)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, n)

	var funcs []*sitter.Node
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			funcs = append(funcs, c.Node)
			fmt.Println("-", absPath+":"+strconv.FormatUint(uint64(c.Node.StartPoint().Row), 10), "-", c.Node.EndPoint().Row)
			fmt.Println(c.Node.Content(contents))
		}
	}
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

func getTSParser(lang string) (sitter.Parser, sitter.Language) {
	parser := sitter.NewParser()
	grammar := new(sitter.Language)
	switch lang {
	case "JavaScript":
		grammar = javascript.GetLanguage()
	case "Go":
		grammar = golang.GetLanguage()
	case "Python":
		grammar = python.GetLanguage()
	case "Java":
		grammar = java.GetLanguage()
	default:
		handleErr(errors.New("language not supported at this time"))
	}
	parser.SetLanguage(grammar)
	return *parser, *grammar
}
