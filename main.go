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
)

var (
	noFileNames bool
)

func init() {
	flag.BoolVar(&noFileNames, "q", false, `"quiet" excludes file names from output`)
	flag.Parse()
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	path := flag.Arg(0)
	query := flag.Arg(1)

	absPath, err := filepath.Abs(path)
	handleErr(err)

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
	defer q.Close()

	qc := sitter.NewQueryCursor()
	defer qc.Close()

	qc.Exec(q, n)

	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}

		// fmt.Println(q.CaptureNameForId(m.ID))
		for _, c := range m.Captures {
			if !noFileNames {
				fmt.Printf("%s:%d:%d\n", absPath, c.Node.StartPoint().Row+1, c.Node.StartPoint().Column+1)
			}
			fmt.Println(c.Node.Content(contents))
		}
	}
}
