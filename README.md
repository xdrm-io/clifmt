# | extended terminal format |

[![Go version](https://img.shields.io/badge/go_version-1.11-blue.svg)](https://golang.org/doc/go1.11)
[![Go Report Card](https://goreportcard.com/badge/git.xdrm.io/go/clifmt)](https://goreportcard.com/report/git.xdrm.io/go/clifmt)
[![Go doc](https://godoc.org/git.xdrm.io/go/clifmt?status.svg)](https://godoc.org/git.xdrm.io/go/clifmt)


Simple utility written in `go` that extends the standard `fmt.Sprintf` and `fmt.Printf` functions. It allows you to set foreground/background color, **bold**, <u>underlined</u> and _italic_ text.



----

## (1) Format



##### Base format

```go
${<target text>}(<fg>:<bg>)
```

- `<target text>` is the text that will be colorized.
- `<fg>` is the name of the foreground color (*c.f. [color list](https://git.xdrm.io/go/clifmt/src/master/colors.go#L15)*), or an hex color (*e.g.`#00ffaa`, `#0fa`*).
- `<bg>` is the name of the background color with the same syntax as for the foreground.

> Note that it is not recommended to nest the different coloring formats.



##### Foreground only

```go
${<target text>}(<fg>)
```

- `<target text>` is the text that will be colorized.
- `<fg>` is the name of the foreground color.



##### Background only

```go
${<target text>}(:<bg>)
```

- `<target text>` is the text that will be colorized.
- `<bg>` is the name of the background color.





----

## (2) Screenshot

![default screenshot](https://0x0.st/sCPc.png)



----

## (3) Incoming features

- [ ] **markdown-like formatting** - bold, italic, underlined, (strike-through?)
- [ ] **global alignment** - align text dynamically
- [ ] **progress-line** - redrawing format to show, for instance an animated progress bar on the same line

