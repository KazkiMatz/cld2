package cld2

import (
	"testing"
)

var testData = [...]struct {
	Code string
	Text string
}{
	{"en", "Hey, this is an english sentence"},
	{"es", "¿Tienes planes para mantener activamente este proyecto? ¿Puedo ayudar de alguna manera?"},
	{"ko", "할머니들을 만나고 이야기한 기회가 있다. 선물도 했음."},
	{"vi", "Quản lý marketing và bán hàng"},
	{"th", "ระดับปริญญาตรี มหาวิทยาลัยกรุงเทพ คณะบริหารธุรกิจ"},
	{"ja", "Goにby-reference rangeが無いと不平を言っているわけではなく、rangeが非自明であることに対して不平を言っているのです。"},
}

var testDataMixed = [...]struct {
	Code1 string
	Code2 string
	Text  string
}{
	{"en", "vi", "Quản lý marketing và bán hàng. To create a Go slice backed by a C array, one needs to acquire this length at runtime and use a type conversion to a pointer to a very big array and then slice it to the length that you want. Xóa việc đã xem để làm mới gợi ý việc làm."},
	{"ko", "vi", "할머니들을 만나고 이야기한 기회가 있다. 선물도 했음. Quản lý marketing và bán hàng"},
	{"vi", "", "Quản lý marketing và bán hàng."},
	{"th", "vi", "ระดับปริญญาตรี มหาวิทยาลัยกรุงเทพ คณะบริหารธุรกิจ. Quản lý marketing và bán hàng."},
	{"ja", "vi", "Goにby-reference rangeが無いと不平を言っているわけではなく、rangeが非自明であることに対して不平を言っているのです。Quản lý marketing và bán hàng."},
}

func TestDetect(t *testing.T) {
	for _, item := range testData {
		infoSet := Detect(item.Text)
		primaryInfo := infoSet[0]
		if primaryInfo.Code != item.Code || primaryInfo.Percent < 75 {
			t.Errorf("Language: want \"%s\" above 75%%, got %s (%d%%)", item.Code, primaryInfo.Code, primaryInfo.Percent)
		}
	}
}

func TestDetectWithMixedData(t *testing.T) {
	for _, item := range testDataMixed {
		infoSet := Detect(item.Text)

		primaryInfo := infoSet[0]
		if primaryInfo.Code != item.Code1 || primaryInfo.Percent < 60 {
			t.Errorf("Primary language: want \"%s\" above 60%%, got %s (%d%%)", item.Code1, primaryInfo.Code, primaryInfo.Percent)
		}

		if len(infoSet) < 2 {
			if item.Code2 != "" {
				t.Errorf("Secondary language \"%s\" was not detected (primary: \"%s\")", item.Code2, item.Code1)
			}
			continue
		}

		secondaryInfo := infoSet[1]
		if secondaryInfo.Code != item.Code2 {
			t.Errorf("Secondary language: want \"%s\", got %s", item.Code2, secondaryInfo.Code)
		}
	}
}
