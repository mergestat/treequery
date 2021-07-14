package main

import (
	"embed"
	"path"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/bash"
	"github.com/smacker/go-tree-sitter/c"
	"github.com/smacker/go-tree-sitter/cpp"
	"github.com/smacker/go-tree-sitter/csharp"
	"github.com/smacker/go-tree-sitter/css"
	"github.com/smacker/go-tree-sitter/elm"
	"github.com/smacker/go-tree-sitter/golang"
	"github.com/smacker/go-tree-sitter/html"
	"github.com/smacker/go-tree-sitter/java"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/smacker/go-tree-sitter/lua"
	"github.com/smacker/go-tree-sitter/ocaml"
	"github.com/smacker/go-tree-sitter/php"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/ruby"
	"github.com/smacker/go-tree-sitter/rust"
	"github.com/smacker/go-tree-sitter/scala"

	//"github.com/smacker/go-tree-sitter/svelte"
	"github.com/smacker/go-tree-sitter/toml"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
	"github.com/smacker/go-tree-sitter/yaml"
)

//go:embed queries
var f embed.FS

// getTSLanguageFromEnry retrieves the tree sitter language from a language name string (defined by the enry package)
func getTSLanguageFromEnry(lang string) (*sitter.Language, string, error) {
	var tsLang *sitter.Language
	switch lang {
	case "Shell":
		tsLang = bash.GetLanguage()
	case "C":
		tsLang = c.GetLanguage()
	case "C++":
		tsLang = cpp.GetLanguage()
	case "C#":
		tsLang = csharp.GetLanguage()
	case "CSS":
		tsLang = css.GetLanguage()
	case "Elm":
		tsLang = elm.GetLanguage()
	case "Go":
		tsLang = golang.GetLanguage()
	case "HTML":
		tsLang = html.GetLanguage()
	case "Java":
		tsLang = java.GetLanguage()
	case "JavaScript":
		tsLang = javascript.GetLanguage()
	case "Lua":
		tsLang = lua.GetLanguage()
	case "OCaml":
		tsLang = ocaml.GetLanguage()
	case "Python":
		tsLang = python.GetLanguage()
	case "PHP":
		tsLang = php.GetLanguage()
	case "Ruby":
		tsLang = ruby.GetLanguage()
	case "Rust":
		tsLang = rust.GetLanguage()
	case "Scala":
		tsLang = scala.GetLanguage()
	//case "Svelte":
	//return svelte.GetLanguage()
	case "TOML":
		tsLang = toml.GetLanguage()
	case "TypeScript":
		tsLang = typescript.GetLanguage()
	case "YAML":
		tsLang = yaml.GetLanguage()
	default:
		return nil, "", ErrLangNotSupported
	}

	query, err := f.ReadFile(path.Join("queries", lang, "main.scm"))
	if err != nil {
		return tsLang, "", err
	}

	return tsLang, string(query), nil
}
