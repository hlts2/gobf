# gobf

[![Go Report Card](https://goreportcard.com/badge/github.com/hlts2/gobf)](https://goreportcard.com/report/github.com/hlts2/gobf)
[![GoDoc](http://godoc.org/github.com/hlts2/gobf?status.svg)](http://godoc.org/github.com/hlts2/gobf)

gobf is a simple library that generates brute force string patterns.


## Requirement

Go 1.15

## Installing

```
go get github.com/hlts2/gobf
```

## Example

```go
package main

import (
	"context"
	"log"

	"github.com/hlts2/gobf"
)

func main() {
	bf, err := gobf.New(
		gobf.WithNumber(true),
		gobf.WithUpper(true),
		gobf.WithLower(true),
		gobf.WithSize(4),
		gobf.WithConcrencyLimit(1000000),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("start to search pattern: s3Df")

	err = bf.Do(context.Background(), func(pattern string) {
		if pattern == "s3Df" {
			log.Printf("find: %s\n", pattern)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

### Option

You can use options to change the type of string generated.

```go
// Generate a pattern of four-letter lowercase and number combinations.
bf, err := gobf.New(
  gobf.WithNumber(true),
  // gobf.WithUpper(true),
  gobf.WithLower(true),
  gobf.WithSize(4),
)
```


## Contribution
1. Fork it ( https://github.com/hlts2/gobf/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## Author
[hlts2](https://github.com/hlts2)
