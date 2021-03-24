// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://code.google.com/p/cld2/.
package cld2

// #include <stdlib.h>
// #include "cld2.h"
import "C"
import "unsafe"

type LangInfo struct {
	Code       string
	Percent    int
	IsReliable bool
}

// Detect returns the language code for detected language
// in the given text.
func Detect(text string) []LangInfo {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	var set *C.LangInfo = C.DetectLang(cs, -1)
	defer C.free(unsafe.Pointer(set))

	size := 3
	infoSet := []LangInfo{}
	for _, res := range (*[1 << 30]C.LangInfo)(unsafe.Pointer(set))[:size:size] {
		code := C.GoString(res.code)
		if code == "un" {
			continue
		}

		infoSet = append(infoSet,
			LangInfo{
				Code:       code,
				Percent:    int(res.percent),
				IsReliable: (res.is_reliable == true),
			})
	}

	return infoSet
}
