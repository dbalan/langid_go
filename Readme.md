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
## Updating foreign code and models.
All C files except `langid_go.{c,h}` are copied from `langid.c` repo, including
`model.c` which contains a generated model. To update these files remove them

from repo, and pass LDFLAG: `-lliblangid` to cgo (here: 
https://github.com/cliqz-oss/langid_go/blob/master/langid_go.go#L3
)


## Bugs and todos
1. Loading your own model is not supported.
2. No interface to get the probability values out.


tags: golang, go, langid.go

