# language detection

This is a go wrapper around [langid.c](https://github.com/saffsd/langid.c). The
repo include all the neccesory c code and the default trained model to be able
to use as is.

## Installation
Depends on protobuf-c library, in ubuntu this can be installed using
```
sudo apt-get install libprotobuf-c-dev
```

Install the go library
```
go get github.com/cliqz-oss/langid_go
```

## Usage
```golang
package main

import (
	"fmt"
	"github.com/cliqz-oss/langid_go"
)

func main() {
	li := langid_go.NewIdfr()
    
    // de-allocates the language model from memory.
	defer li.Destroy()
    
	fmt.Println(li.DetectLanguage("hello world"))
}

```

tags: golang, go, langid.go
