# yamlfmt

[![Go build](https://github.com/stuart-warren/yamlfmt/workflows/Go/badge.svg)](https://github.com/stuart-warren/yamlfmt/actions)

based on gofmt, yamlfmt formats yaml files into a canonical format

* lists are not indented
* maps have sorted keys
* indent is 2 spaces
* documents always have separators `---`

```
$ yamlfmt --help
formats yaml files with 2 space indent, sorted dicts and non-indented lists
usage: yamlfmt [flags] [path ...]
  -d    display diffs instead of rewriting files
  -f    exit non zero if changes detected
  -l    list files whose formatting differs from yamlfmt's
  -s    sort maps & sequences, WARNING: This may break anchors & aliases
  -w    write result to (source) file instead of stdout
```

Without an explicit path, it processes the standard input. Given a file, it operates on that file; given a directory, it operates on all .yaml and .yml files in that directory, recursively.

By default, yamlfmt prints the reformatted sources to standard output.
