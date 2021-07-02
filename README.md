[![Go Report Card](https://goreportcard.com/badge/github.com/askgitdev/treequery)](https://goreportcard.com/report/github.com/askgitdev/treequery)
[![BuildStatus](https://github.com/askgitdev/askgit/workflows/tests/badge.svg)](https://github.com/askgitdev/askgit/actions?workflow=ci)

## treequery

`treequery` is a CLI for executing [Tree-sitter queries](https://tree-sitter.github.io/tree-sitter/using-parsers#query-syntax) on source code files.
It uses [`enry`](https://github.com/go-enry/go-enry) to detect a language and apply the right Tree-sitter [parser](https://tree-sitter.github.io/tree-sitter/#available-parsers).
The default output includes a list of line number locations where there's a query match, followed by a snippet of the matching code.

### Getting Started

Running `make` in the root of this repo will produce a `tq` binary, which can be used as such:

```
> ./tq testdata/TriestBase.java "(method_declaration name: (identifier) @method_name)"
./treequery/testdata/TriestBase.java:20
handleEdge
./treequery/testdata/TriestBase.java:48
swapIn
./treequery/testdata/TriestBase.java:53
addEdge
./treequery/testdata/TriestBase.java:88
removeEdge
./treequery/testdata/TriestBase.java:119
getEstimate
```

The above example shows a way to select all method names in the `testdata/TriestBase.java` file.
To exclude file names and line locations, use the `-q` (quiet) flag:

```
handleEdge
swapIn
addEdge
removeEdge
getEstimate
```

### Supported Languages

- [ ] bash
- [ ] c
- [ ] cpp
- [ ] csharp
- [ ] css
- [ ] elm
- [x] golang
- [ ] html
- [x] java
- [x] javascript
- [ ] lua
- [ ] ocaml
- [ ] php
- [x] python
- [x] ruby
- [ ] rust
- [ ] scala
- [ ] svelte
- [ ] toml
- [ ] typescript
- [ ] yaml
