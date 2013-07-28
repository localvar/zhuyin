package zhuyin

import (
	"testing"
)

func TestEncodePinyin(t *testing.T) {
	cases := []string{
		// positive cases
		"e5", "ê",
		"ju0", "ju",
		"jv", "ju",
		"lv3", "lǚ",
		"lvan4", "lüàn", // not valid, for test only
		"zhuan4", "zhuàn",
		"zhao2", "zháo",
		"leng1", "lēng",
		"shui3", "shuǐ",
		"liu2", "liú",
		"an3", "ǎn",

		// negative cases
		"alkfj", "",
		"zhhhan3", "",
		"zhaang4", "",
	}

	t.Log("testing pinyin encoding...")
	for i := 0; i < len(cases); i += 2 {
		output := EncodePinyin(cases[i])
		if output != cases[i+1] {
			t.Errorf("input: %s   desired: %s   actural: %s", cases[i], cases[i+1], output)
		}
	}
}

func TestDecodePinyin(t *testing.T) {
	cases := []string{
		// positive cases
		"ê", "e5",
		"ju", "jv0",
		"lǚ", "lv3",
		"lüàn", "lvan4", // not valid, for test only
		"zhuàn", "zhuan4",
		"zháo", "zhao2",
		"lēng", "leng1",
		"shuǐ", "shui3",
		"liú", "liu2",
		"ǎn", "an3",

		// negative cases
		"alkfj", "",
		"zhhhan3", "",
		"zhaang4", "",
	}

	t.Log("testing pinyin decoding...")
	for i := 0; i < len(cases); i += 2 {
		output := DecodePinyin(cases[i])
		if output != cases[i+1] {
			t.Errorf("input: %s   desired: %s   actural: %s", cases[i], cases[i+1], output)
		}
	}
}

func TestEncodeZhuyin(t *testing.T) {
	cases := []string{
		// positive cases
		"e5", "ㄝ",
		"ju0", "ㄐㄩ˙",
		"jv", "ㄐㄩ˙",
		"lv3", "ㄌㄩˇ",
		"lvan4", "ㄌㄩㄢˋ", // not valid, for test only
		"zhuan4", "ㄓㄨㄢˋ",
		"zhao2", "ㄓㄠˊ",
		"zhi4", "ㄓˋ",
		"leng1", "ㄌㄥ",
		"shui3", "ㄕㄨㄟˇ",
		"liu2", "ㄌ一ㄡˊ",
		"an3", "ㄢˇ",
		"yu1", "ㄩ",
		"wu2", "ㄨˊ",
		"yve3", "ㄩㄝˇ",
		"yue4", "ㄩㄝˋ",

		// negative cases
		"alkfj", "",
		"zhhhan3", "",
		"zhaang4", "",
	}

	t.Log("testing zhuyin encoding...")
	for i := 0; i < len(cases); i += 2 {
		output := EncodeZhuyin(cases[i])
		if output != cases[i+1] {
			t.Errorf("input: %s   desired: %s   actural: %s", cases[i], cases[i+1], output)
		}
	}
}

func TestDecodeZhuyin(t *testing.T) {
}
