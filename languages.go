package main

import (
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/ruby"
)

// getTSLanguageFromEnry retrieves the tree sitter language from a language name string (defined by the enry package)
func getTSLanguageFromEnry(lang string) *sitter.Language {
	switch lang {
	case "Go":
		return golang.GetLanguage()
	case "Java":
		return java.GetLanguage()
	case "JavaScript":
		return javascript.GetLanguage()
	case "Python":
		return python.GetLanguage()
	case "Ruby":
		return ruby.GetLanguage()
	default:
		return nil
	}
}
