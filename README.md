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
	"github.com/wmentor/epub"
)

func main() {

  err := epub.Reader("./data/test.epub", func(chapter string, content []byte) bool {
    fmt.Println(chpater)
  })
  if err != nil {
    panic(err)
  }
}
```
