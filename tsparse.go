package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-enry/go-enry/v2"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/python"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// getTSLanguageFromEnry retrieves the tree sitter language from a language name string (defined by the enry package)
func getTSLanguageFromEnry(lang string) *sitter.Language {
	switch lang {
	case "JavaScript":
		return javascript.GetLanguage()
	case "Go":
		return golang.GetLanguage()
	case "Python":
		return python.GetLanguage()
	case "Java":
		return java.GetLanguage()
	default:
		return nil
	}
}

func main() {
	flag.Parse()
	path := flag.Arg(0)
	query := flag.Arg(1)

	absPath, err := filepath.Abs(path)
	handleErr(err)

	f, err := fileExists(absPath)
	handleErr(err)

	if !f {
		handleErr(errors.New("file does not exist"))
	}

	contents, err := ioutil.ReadFile(absPath)
	handleErr(err)

	lang := enry.GetLanguage(absPath, contents)
	if lang == "" {
		handleErr(errors.New("language could not be detected"))
	}

	language := getTSLanguageFromEnry(lang)
	parser := sitter.NewParser()
	parser.SetLanguage(language)

	tree := parser.Parse(nil, contents)
	n := tree.RootNode()

	if query == "" {
		// fmt.Println(absPath)
		// fmt.Println(n.Content(contents))
		fmt.Println(n)
	}

	q, err := sitter.NewQuery([]byte(query), language)
	if err != nil {
		handleErr(fmt.Errorf("problem with query: %w", err))
	}

	qc := sitter.NewQueryCursor()
	qc.Exec(q, n)

	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		for _, c := range m.Captures {
			fmt.Println(absPath + ":" + fmt.Sprintf("%d", c.Node.StartPoint().Row+1))
			fmt.Println(c.Node.Content(contents))
		}
	}
}
