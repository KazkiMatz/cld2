package cld2

import (
	"testing"
)

var testData = [...]struct {
	lang string
	text string
}{
	{"en", "Hey, this is an english sentence"},
	{"es", "¿Tienes planes para mantener activamente este proyecto? ¿Puedo ayudar de alguna manera?"},
	{"ko", "할머니들을 만나고 이야기한 기회가 있다. 선물도 했음."},
	{"vi", "Quản lý marketing và bán hàng"},
	{"th", "ระดับปริญญาตรี มหาวิทยาลัยกรุงเทพ คณะบริหารธุรกิจ"},
	{"ja", "Goにby-reference rangeが無いと不平を言っているわけではなく、rangeが非自明であることに対して不平を言っているのです。"},
}

func TestDetect(t *testing.T) {
	for _, item := range testData {
		res := Detect(item.text)
		if res != item.lang {
			t.Errorf("Language: want \"%s\", got %s", item.lang, res)
		}
	}
}
