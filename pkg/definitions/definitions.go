package definitions

var UnihanReadingKeys = []string{
	"Cantonese",
	"Definition",
	"Fanqie",
	"Hangul",
	"HanyuPinlu",
	"HanyuPinyin",
	"Japanese",
	"JapaneseKun",
	"JapaneseOn",
	"Korean",
	"Mandarin",
	"SMSZD2003Readings",
	"Tang",
	"TGHZ2013",
	"Vietnamese",
	"XHC1983",
	"Zhuang",
}

type UnihanReadingKeyPosition int

const (
	Cantonese UnihanReadingKeyPosition = iota
	Definition
	Fanqie
	Hangul
	HanyuPinlu
	HanyuPinyin
	Japanese
	JapaneseKun
	JapaneseOn
	Korean
	Mandarin
	SMSZD2003Readings
	Tang
	TGHZ2013
	Vietnamese
	XHC1983
	Zhuang
	MAX // This should be the last constant
)

var UnihanReadingKeyMap = map[string]UnihanReadingKeyPosition{}

func init() {
	if len(UnihanReadingKeys) == 0 {
		panic("UnihanReadingKeys is empty")
	}
	if len(UnihanReadingKeys) != int(MAX) {
		panic("UnihanReadingKeys length does not match MAX")
	}

	for i, key := range UnihanReadingKeys {
		UnihanReadingKeyMap[key] = UnihanReadingKeyPosition(i)
	}
}

type UnihanReadings map[rune]UnihanReading

type UnihanReading [MAX]string
