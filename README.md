## treequery

`treequery` is a small CLI for executing [Tree-sitter queries](https://tree-sitter.github.io/tree-sitter/using-parsers#query-syntax) on source code files.
It uses [`enry`](https://github.com/go-enry/go-enry) to detect a language and apply the right Tree-sitter [parser](https://tree-sitter.github.io/tree-sitter/#available-parsers).


### Getting Started

Running `make` in the root of this repo will produce a `tq` binary, which can be used as such:

```
> ./tq testdata/TriestBase.java "(method_declaration name: (identifier) @method_name)"
/.../treequery/testdata/TriestBase.java:20
handleEdge
/.../treequery/testdata/TriestBase.java:48
swapIn
/.../treequery/testdata/TriestBase.java:53
addEdge
/.../treequery/testdata/TriestBase.java:88
removeEdge
/.../treequery/testdata/TriestBase.java:119
getEstimate
```

The above example shows a way to select all method names in the `testdata/TriestBase.java` file.
