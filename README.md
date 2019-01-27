# | extended terminal format |

[![Go version](https://img.shields.io/badge/go_version-1.11-blue.svg)](https://golang.org/doc/go1.11)
[![Go Report Card](https://goreportcard.com/badge/git.xdrm.io/go/clifmt)](https://goreportcard.com/report/git.xdrm.io/go/clifmt)
[![Go doc](https://godoc.org/git.xdrm.io/go/clifmt?status.svg)](https://godoc.org/git.xdrm.io/go/clifmt)
[![buddy branch](https://app.buddy.works/xdrmbracketsdev/clifmt/repository/branch/master/badge.svg?token=33f90a953be299fc439c760e2eab36c708f8ea1f87f1159dd77924589b364b2d "buddy branch")](https://app.buddy.works/xdrmbracketsdev/clifmt/repository/branch/master)


Simple utility written in `go` that extends the standard `fmt.Sprintf` and `fmt.Printf` functions. It allows you to set foreground/background color, **bold**, <u>underlined</u> and _italic_ text.

<!-- toc -->

- [(1) Format](#1-format)
  * [[Colorization]](#colorization)
    + [Base format](#base-format)
        * [Example](#example)
    + [Foreground only](#foreground-only)
        * [Example](#example-1)
    + [Background only](#background-only)
        * [Example](#example-2)
  * [[Markdown-like format]](#markdown-like-format)
    + [Bold format](#bold-format)
        * [Example](#example-3)
    + [Italic format](#italic-format)
        * [Example](#example-4)
    + [Underline format](#underline-format)
        * [Example](#example-5)
    + [Hyperlink format](#hyperlink-format)
        * [Example](#example-6)
- [(2) Screenshot](#2-screenshot)
        * [Colorizing format example :](#colorizing-format-example-)
        * [Markdown-like format example](#markdown-like-format-example)
- [(3) Incoming features](#3-incoming-features)

<!-- tocstop -->

----

## (1) Format



### [Colorization]

#### Base format

```go
${<target text>}(<fg>:<bg>)
```

- `<target text>` is the text that will be colorized.

- `<fg>` is the name of the foreground color (*c.f. [color list](https://git.xdrm.io/go/clifmt/src/master/colors.go#L15)*), or an hex color (*e.g.`#00ffaa`, `#0fa`*).

- `<bg>` is the name of the background color with the same syntax as for the foreground.



###### Example

```go
clifmt.Printf("normal text ${red text over black}(red:#000) normal text")
```

> Note that it is not recommended to nest the different coloring formats.



#### Foreground only

```go
${<target text>}(<fg>)
```

- `<target text>` is the text that will be colorized.
- `<fg>` is the name of the foreground color.



###### Example

```go
clifmt.Printf("normal text ${red text}(red) normal text")
```



#### Background only

```go
${<target text>}(:<bg>)
```

- `<target text>` is the text that will be colorized.
- `<bg>` is the name of the background color.



###### Example

```go
clifmt.Printf("normal text ${text over red}(#ff0000) normal text")
```



### [Markdown-like format]



#### Bold format

```go
**<target text>**
```

- `<target text>` is the text that will be bold.

###### Example

```go
clifmt.Printf("normal text **bold text** normal text")
```



#### Italic format

```go
*<target text>*
```

- `<target text>` is the text that will be italic.

###### Example

```go
clifmt.Printf("normal text *italic text* normal text")
```



#### Underline format

```go
_<target text>_
```

- `<target text>` is the text that will be underlined.

###### Example

```go
clifmt.Printf("normal text _underline text_ normal text")
```





#### Hyperlink format

```go
[<target text>](<target url>)
```

- `<target text>` is the text that will be displayed.
- `<target url>` is the url the hyperlink links to.

###### Example

```go
clifmt.Printf("normal text [hyper](link) normal text")
```



----

## (2) Screenshot



###### Colorizing format example :

![colorizing example](https://0x0.st/sCPc.png)



###### Markdown-like format example

![markdown-like example](https://0x0.st/sC9F.png)



----

## (3) Incoming features

- [x] **markdown-like formatting** - bold, italic, underlined, (strike-through?)
- [ ] **global alignment** - align text dynamically
- [ ] **progress-line** - redrawing format to show, for instance an animated progress bar on the same line

