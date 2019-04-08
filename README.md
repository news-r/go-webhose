[![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/JohnCoene/go-webhose/webhose) [![Go Report Card](https://goreportcard.com/badge/github.com/JohnCoene/go-webhose)](https://goreportcard.com/report/github.com/JohnCoene/go-webhose)

# go-webhose

Access [webhose.io](https://webhose.io) API from golang.

## Install

```bash
go get github.com/JohnCoene/go-webhose
```

## Documentation

Documentation is on [godoc](https://godoc.org/github.com/JohnCoene/go-webhose/webhose)

## Example

```go
package main

import (
	"fmt"
	"go-webhose/webhose"
)

func main() {
	client := &webhose.NewsClient{Token: "xxXXXX-XXxxXX-xXXxxxXX-xXXXx", Query: "go programming language"}

	news := webhose.GetArticles(client, 2) // client and number of pages of result

	fmt.Println(news.News)
}

```