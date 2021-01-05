# epub

Golang epub reader library

## Summary

* Written on pure Go
* Require Go version >= 1.14

## Install

```plaintext
go get github.com/wmentor/epub
```

## Usage

```golang
package main

import (
  "fmt"
  "os"

  "github.com/wmentor/epub"
)

func main() {

  err := epub.Reader("./data/test.epub", func(chapter string, chapterHTML []byte) bool {
    fmt.Println(chapter)
  })
  if err != nil {
    panic(err)
  }

  // print epub to Stdout as text
  epub.ToTxt("./data/test.epub", os.Stdout)
}
```
