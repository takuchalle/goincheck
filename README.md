# goincheck
[coincheck exchange client API](https://coincheck.com/ja/documents/exchange/api) for golang

[![Build Status](https://travis-ci.org/takuyaohashi/goincheck.svg?branch=master)](https://travis-ci.org/takuyaohashi/goincheck)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/takuyaohashi/goincheck/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/takuyaohashi/goincheck

# Installation

```
$ go get github.com/takuyaohashi/goincheck
```

# Usage Example
```go
package main

import (
    "fmt"

    "github.com/takuyaohashi/goincheck"
)

const (
    accessKey       = "hoge"
    secretAccessKey = "huga"
)

func main() {
    client, _ := goincheck.NewClient(accessKey, secretAccessKey)
    tikcer, _ := client.GetTicker()
    fmt.Printf("Tikcer = %+v\n", tikcer)
}
```

For detail, please check ``sample/cmd/goincheck`` directory or [GoDoc](http://godoc.org/github.com/takuyaohashi/goincheck).

# License
MIT
# Author
Takuya OHASHI
