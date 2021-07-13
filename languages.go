package main

import (
	_ "embed"

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

var (
	shellQuery      string
	cQuery          string
	cppQuery        string
	chashQuery      string
	cssQuery        string
	elmQuery        string
	goQuery         string
	htmlQuery       string
	javaQuery       string
	jsQuery         string
	luaQuery        string
	ocamlQuery      string
	pythonQuery     string
	phpQuery        string
	rubyQuery       string
	rustQuery       string
	scalaQuery      string
	tomlQuery       string
	typescriptQuery string
	yamlQuery       string
)

// getTSLanguageFromEnry retrieves the tree sitter language from a language name string (defined by the enry package)
func getTSLanguageFromEnry(lang string) (*sitter.Language, string) {
	switch lang {
	case "Shell":
		return bash.GetLanguage(), shellQuery
	case "C":
		return c.GetLanguage(), cQuery
	case "C++":
		return cpp.GetLanguage(), cppQuery
	case "C#":
		return csharp.GetLanguage(), chashQuery
	case "CSS":
		return css.GetLanguage(), cssQuery
	case "Elm":
		return elm.GetLanguage(), elmQuery
	case "Go":
		return golang.GetLanguage(), goQuery
	case "HTML":
		return html.GetLanguage(), htmlQuery
	case "Java":
		return java.GetLanguage(), javaQuery
	case "JavaScript":
		return javascript.GetLanguage(), jsQuery
	case "Lua":
		return lua.GetLanguage(), luaQuery
	case "OCaml":
		return ocaml.GetLanguage(), ocamlQuery
	case "Python":
		return python.GetLanguage(), pythonQuery
	case "PHP":
		return php.GetLanguage(), phpQuery
	case "Ruby":
		return ruby.GetLanguage(), rubyQuery
	case "Rust":
		return rust.GetLanguage(), rustQuery
	case "Scala":
		return scala.GetLanguage(), scalaQuery
	//case "Svelte":
	//return svelte.GetLanguage()
	case "TOML":
		return toml.GetLanguage(), tomlQuery
	case "TypeScript":
		return typescript.GetLanguage(), typescriptQuery
	case "YAML":
		return yaml.GetLanguage(), yamlQuery
	default:
		return nil, ""
	}
}
