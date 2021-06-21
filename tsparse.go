package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-enry/go-enry/v2"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
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
	contents = []byte(contents)
	lang := enry.GetLanguage(absPath, contents)
	fmt.Println("language: " + lang)

	parser := sitter.NewParser()
	parser.SetLanguage(javascript.GetLanguage())
	tree := parser.Parse(nil, contents)
	n := tree.RootNode()
<<<<<<< Updated upstream
	fmt.Println(n)
=======
	fmt.Println("language: " + lang)
	fmt.Println("AST:", n)

	fmt.Println("Root type:", n.Type())
	fmt.Println("Root children:", n.ChildCount())

	fmt.Println("\nFunctions in input:")
	q, _ := sitter.NewQuery([]byte("(function_declaration) @func"), java.GetLanguage())
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
			fmt.Println("-", funcName(contents, c.Node))
		}
	}
	fmt.Println(funcs)
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
=======

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getTSParser(lang string) sitter.Parser {
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
	return *parser
}

func funcName(content []byte, n *sitter.Node) string {
	if n == nil {
		return ""
	}

	if n.Type() != "function_declaration" {
		return ""
	}

	return n.ChildByFieldName("name").Content(content)
}
>>>>>>> Stashed changes
