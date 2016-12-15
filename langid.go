package main

import (
	"fmt"
)

/*
#include "langid_go.h"
*/
import "C"

func DetectLanguage(data string) string {
	lang := C.detect_language(C.CString(data))
	return C.GoString(lang)
}

func main() {
	fmt.Println(DetectLanguage("hello"))
}
