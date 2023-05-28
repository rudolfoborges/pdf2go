# pdf2go

A simple Golang library for converting PDF to text and html.

[![Go Report Card](https://goreportcard.com/badge/github.com/rudolfoborges/pdf2go)](https://goreportcard.com/report/github.com/rudolfoborges/pdf2go)
[![GoDoc](https://godoc.org/github.com/rudolfoborges/pdf2go?status.svg)](https://godoc.org/github.com/rudolfoborges/pdf2go)
[![Build Status](https://travis-ci.org/rudolfoborges/pdf2go.svg?branch=main)](https://travis-ci.org/rudolfoborges/pdf2go)
[![codecov](https://codecov.io/gh/rudolfoborges/pdf2go/branch/main/graph/badge.svg)](https://codecov.io/gh/rudolfoborges/pdf2go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/rudolfoborges/pdf2go
```

## Poppler Installation

Needs [poppler](https://poppler.freedesktop.org/) to be installed on your system.

### Mac OS

```bash
brew install poppler
```

### Ubuntu

```bash
sudo apt-get install libpoppler-cpp-dev
```

### Centos

```bash
sudo yum install poppler-utils
```

### Windows

Download the latest version from [poppler](http://blog.alivate.com.au/poppler-windows/)

## Usage

### Extract Text

```go
package main

import (
    "fmt"
    "github.com/rudolfoborges/pdf2go"
)

func main() {
    pdf, err := pdf2go.New("path/to/file.pdf", pdf2go.Config{
        LogLevel: pdf2go.LogLevelError,
    })

    if err != nil {
        panic(err)
    }

    text, err := pdf.Text()
    if err != nil {
        panic(err)
    }

    fmt.Println(text)

    pages, err := pdf.Pages()

    if err != nil {
        panic(err)
    }

    for _, page := range pages {
        fmt.Println(page.Text())
    }
}
```

### Extract Html

```go
package main

import (
    "fmt"
    "github.com/rudolfoborges/pdf2go"
)

func main() {
    pdf, err := pdf2go.New("path/to/file.pdf", pdf2go.Config{
        LogLevel: pdf2go.LogLevelError,
    })

    if err != nil {
        panic(err)
    }

    html, err := pdf.Html()
    if err != nil {
        panic(err)
    }

    fmt.Println(html)

    pages, err := pdf.Pages()

    if err != nil {
        panic(err)
    }

    for _, page := range pages {
        fmt.Println(page.Html())
    }
}
```

### More examples on [examples](https://github.com/rudolfoborges/pdf2go/tree/main/examples) folder

## Next Steps

-   [ ] Add image extraction
-   [ ] Extract image from specific page
