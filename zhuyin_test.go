// Copyright (c) 2013 Localvar. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
		"yi2", "yí",
		"yuan2", "yuán",
		"yvan2", "yuán",
		"min2", "mín",

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
		"yí", "yi2",
		"yuán", "yvan2",

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
		"liu2", "ㄌㄧㄡˊ",
		"an3", "ㄢˇ",
		"yu1", "ㄩ",
		"wu2", "ㄨˊ",
		"yve3", "ㄩㄝˇ",
		"yue4", "ㄩㄝˋ",
		"yi2", "ㄧˊ",
		"yuan2", "ㄩㄢˊ",
		"yvan2", "ㄩㄢˊ",
		"min2", "ㄇㄧㄣˊ",

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
	cases := []string{
		// positive cases
		"ㄝ", "e5",
		"ㄐㄩ˙", "jv0",
		"ㄐㄩ˙", "jv0",
		"ㄌㄩˇ", "lv3",
		"ㄌㄩㄢˋ", "lvan4", // not valid, for test only
		"ㄓㄨㄢˋ", "zhuan4",
		"ㄓㄠˊ", "zhao2",
		"ㄓˋ", "zhi4",
		"ㄌㄥ", "leng1",
		"ㄕㄨㄟˇ", "shui3",
		"ㄌㄧㄡˊ", "liu2",
		"ㄢˇ", "an3",
		"ㄩ", "yv1",
		"ㄨˊ", "wu2",
		"ㄩㄝˇ", "yve3",
		"ㄩㄝˋ", "yve4",
		"ㄧˊ", "yi2",
		"ㄩㄢˊ", "yvan2",
		"ㄇㄧㄣˊ", "min2",
		"ㄨㄥˊ", "weng2",

		// negative cases
		"ㄩㄝㄝ", "",
		"ㄐˇ", "",
		"ㄕㄨㄕㄨ", "",
	}
	t.Log("testing zhuyin decoding...")
	for i := 0; i < len(cases); i += 2 {
		output := DecodeZhuyin(cases[i])
		if output != cases[i+1] {
			t.Errorf("input: %s   desired: %s   actural: %s", cases[i], cases[i+1], output)
		}
	}
}

func TestPinyinToZhuyin(t *testing.T) {
	cases := []string{
		"mín", "ㄇㄧㄣˊ",
		"zhāng", "ㄓㄤ",
		"wéng", "ㄨㄥˊ",
	}
	t.Log("testing pinyin to zhuyin...")
	for i := 0; i < len(cases); i += 2 {
		output := PinyinToZhuyin(cases[i])
		if output != cases[i+1] {
			t.Errorf("input: %s   desired: %s   actural: %s", cases[i], cases[i+1], output)
		}
	}
}

func TestZhuyinToPinyin(t *testing.T) {
	cases := []string{
		"ㄇㄧㄣˊ", "mín",
		"ㄓㄤ", "zhāng",
		"ㄨㄥˊ", "wéng",
	}

	t.Log("testing zhuyin to pinyin...")
	for i := 0; i < len(cases); i += 2 {
		output := ZhuyinToPinyin(cases[i])
		if output != cases[i+1] {
			t.Errorf("input: %s   desired: %s   actural: %s", cases[i], cases[i+1], output)
		}
	}
}
