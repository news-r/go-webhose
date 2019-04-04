# go-webhose

Access [webhose.io](https://webhose.io) API from golang.

## Install

```bash
go get github.com/JohnCoene/go-webhose
```

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