package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-enry/go-enry/v2"
	sitter "github.com/smacker/go-tree-sitter"
)

var (
	noFileNames bool
	queryFile   string
	langFlag    string
)

var (
	ErrLangNotDetected  = errors.New("could not detect language")
	ErrLangNotSupported = errors.New("language is not supported")
)

func init() {
	flag.BoolVar(&noFileNames, "q", false, `"quiet" mode excludes file names from output`)
	flag.StringVar(&queryFile, "f", "", "query can be extracted from filepath")
	flag.StringVar(&langFlag, "l", "", "language can be given by user")
	flag.Parse()
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// printForFile prints output to stdout for a single file
func printForFile(path, captureName string) error {
	// this will handle relative paths as well (joins with cwd)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	contents, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	var lang string
	if langFlag != "" {
		lang = langFlag
	} else {
		lang = enry.GetLanguage(absPath, contents)
		if lang == "" {
			return ErrLangNotDetected
		}
	}

	language, query, err := getTSLanguageFromEnry(lang)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}

	if queryFile != "" {
		queryFilePath, err := filepath.Abs(queryFile)
		if err != nil {
			return err
		}
		queryContent, err := ioutil.ReadFile(queryFilePath)
		if err != nil {
			return err
		}
		query = string(queryContent)
	}

	parser := sitter.NewParser()
	parser.SetLanguage(language)

	tree := parser.Parse(nil, contents)
	n := tree.RootNode()

	q, err := sitter.NewQuery([]byte(query), language)
	if err != nil {
		return fmt.Errorf("problem with query: %w", err)
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
			if q.CaptureNameForId(c.Index) == captureName {
				if !noFileNames {
					fmt.Printf("%s:%d:%d\n", absPath, c.Node.StartPoint().Row+1, c.Node.StartPoint().Column+1)
				}
				fmt.Println(c.Node.Content(contents))
			}
		}
	}

	return nil
}

func main() {
	var path string
	var captureName string

	switch flag.NArg() {
	case 0:
		handleErr(fmt.Errorf("must supply a query capture"))
	case 1:
		path = "."
		captureName = flag.Arg(0)
	case 2:
		path = flag.Arg(0)
		captureName = flag.Arg(1)
	}

	absPath, err := filepath.Abs(path)
	handleErr(err)

	err = filepath.WalkDir(absPath, func(path string, d fs.DirEntry, err error) error {
		// skip hidden directories
		// TODO skip paths specified in .gitignore too?
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return fs.SkipDir
		}

		if !d.IsDir() {
			err := printForFile(path, captureName)
			if err != nil {
				if errors.Is(err, ErrLangNotDetected) || errors.Is(err, ErrLangNotSupported) {
					return nil
				}
				return err
			}
		}
		return nil
	})
	handleErr(err)
}
