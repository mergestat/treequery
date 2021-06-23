package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	query := flag.Arg(1)
	absPath, _ := filepath.Abs(path)
	fmt.Println("Path to file: " + absPath)
	f, err := exists(absPath)
	if !f {
		fmt.Println("There was either an error in the command line input or a faulty filepath. Please try again.")
		log.Fatal(err)
	}
	filename := filepath.Base(absPath)

	contents, _ := ioutil.ReadFile(absPath)
	if len(contents) <= 0 {
		handleErr(errors.New("empty or faulty file input"))
	}
	lang := enry.GetLanguage(absPath, contents)
	if len(lang) <= 0 {
		handleErr(errors.New("language could not be detected"))
	}
	parser, grammar := getTSParser(lang)
	tree := parser.Parse(nil, contents)
	n := tree.RootNode()

	fmt.Println("language: " + lang)
	fmt.Println("AST:", n)

	fmt.Println("Root type:", n.Type())
	fmt.Println("Root children:", n.ChildCount())

	fmt.Println("\nFunctions in input:")
	q, errQuery := sitter.NewQuery([]byte(query), &grammar)
	handleErr(errQuery)
	qc := sitter.NewQueryCursor()
	qc.Exec(q, n)

	var codeElements []*sitter.Node
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			codeElements = append(codeElements, c.Node)
			fmt.Println("-", filename, ":", c.Node.StartPoint().Row, "-", c.Node.EndPoint().Row)
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
	switch strings.ToLower(lang) {
	case "javascript":
		grammar = javascript.GetLanguage()
	case "go":
		grammar = golang.GetLanguage()
	case "python":
		grammar = python.GetLanguage()
	case "java":
		grammar = java.GetLanguage()
	default:
		handleErr(errors.New("language not supported at this time"))
	}
	parser.SetLanguage(grammar)
	return *parser, *grammar
}

func funcName(content []byte, n *sitter.Node) string {
	if n == nil {
		return ""
	}

	return n.ChildByFieldName("name").Content(content)
}
