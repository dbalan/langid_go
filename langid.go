package language_detection

// #cgo LDFLAGS: -lprotobuf-c
// #include "langid_go.h"
import "C"

type LanguageIdentifier struct{}

func NewIdfr() *LanguageIdentifier {
	C.init()
	return &LanguageIdentifier{}
}

func (l *LanguageIdentifier) Destroy() {
	C.destroy()
}

func (l *LanguageIdentifier) DetectLanguage(data string) string {
	lang := C.detect_language(C.CString(data))
	return C.GoString(lang)
}

/*
func main() {
	li := NewIdfr()
	defer li.Destroy()

	val := []string{
		"Sɔsó egbexɔxɔ do amɛnyinyi to alɔkpa ɖekpo kpwimɛ",
		"മനുഷ്യാവകാശങ്ങള്‍",
	}

	for _, v := range val {
		fmt.Println(li.DetectLanguage(v))
	}
}
*/
