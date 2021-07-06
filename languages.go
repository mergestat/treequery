package main

import (
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
	"github.com/smacker/go-tree-sitter/svelte"
	"github.com/smacker/go-tree-sitter/toml"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
	"github.com/smacker/go-tree-sitter/yaml"
)

// getTSLanguageFromEnry retrieves the tree sitter language from a language name string (defined by the enry package)
func getTSLanguageFromEnry(lang string) *sitter.Language {
	switch lang {
	case "Shell":
		return bash.GetLanguage()
	case "C":
		return c.GetLanguage()
	case "C++":
		return cpp.GetLanguage()
	case "C#":
		return csharp.GetLanguage()
	case "CSS":
		return css.GetLanguage()
	case "Elm":
		return elm.GetLanguage()
	case "Go":
		return golang.GetLanguage()
	case "HTML":
		return html.GetLanguage()
	case "Java":
		return java.GetLanguage()
	case "JavaScript":
		return javascript.GetLanguage()
	case "Lua":
		return lua.GetLanguage()
	case "OCaml":
		return ocaml.GetLanguage()
	case "Python":
		return python.GetLanguage()
	case "PHP":
		return php.GetLanguage()
	case "Ruby":
		return ruby.GetLanguage()
	case "Rust":
		return rust.GetLanguage()
	case "Scala":
		return scala.GetLanguage()
	case "Svelte":
		return svelte.GetLanguage()
	case "TOML":
		return toml.GetLanguage()
	case "TypeScript":
		return typescript.GetLanguage()
	case "YAML":
		return yaml.GetLanguage()
	default:
		return nil
	}
}
