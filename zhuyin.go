// Copyright (c) 2013 Localvar. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zhuyin

var (
	// map of Pinyin to Zhuyin
	map_p2z = map[string]string{
		// Consonant
		"b": "ㄅ", "p": "ㄆ", "m": "ㄇ", "f": "ㄈ",
		"d": "ㄉ", "t": "ㄊ", "n": "ㄋ", "l": "ㄌ",
		"g": "ㄍ", "k": "ㄎ", "h": "ㄏ",
		"j": "ㄐ", "q": "ㄑ", "x": "ㄒ",
		"zh": "ㄓ", "ch": "ㄔ", "sh": "ㄕ", "r": "ㄖ",
		"z": "ㄗ", "c": "ㄘ", "s": "ㄙ",
		"y": "一", "w": "ㄨ",
		// Rhymes
		"a": "ㄚ", "o": "ㄛ", "e": "ㄜ",
		"i": "一", "u": "ㄨ", "v": "ㄩ",
		"ai": "ㄞ", "ei": "ㄟ", "ui": "ㄨㄟ",
		"ao": "ㄠ", "ou": "ㄡ", "iu": "一ㄡ",
		"an": "ㄢ", "en": "ㄣ", "in": "一ㄣ",
		"ang": "ㄤ", "eng": "ㄥ", "ing": "一ㄥ",
		"ong": "ㄨㄥ", "ie": "一ㄝ", "er": "ㄦ",
		"ue": "ㄩㄝ", "ve": "ㄩㄝ", // 'ue' is same as 've', for typo
		"un": "ㄨㄣ", "vn": "ㄩㄣ", "ia": "一ㄚ",
		"ua": "ㄨㄚ", "uan": "ㄨㄢ", "van": "ㄩㄢ",
		"uai": "ㄨㄞ", "uo": "ㄨㄛ", "iong": "ㄩㄥ",
		"iang": "一ㄤ", "uang": "ㄨㄤ", "ian": "一ㄢ",
		"iao": "一ㄠ",
	}

	// map of Zhuyin to Pinyin
	map_z2p = map[string]string{
		// Consonant
		"ㄅ": "b", "ㄆ": "p", "ㄇ": "m", "ㄈ": "f",
		"ㄉ": "d", "ㄊ": "t", "ㄋ": "n", "ㄌ": "l",
		"ㄍ": "g", "ㄎ": "k", "ㄏ": "h",
		"ㄐ": "j", "ㄑ": "q", "ㄒ": "x",
		"ㄓ": "zh", "ㄔ": "ch", "ㄕ": "sh", "ㄖ": "r",
		"ㄗ": "z", "ㄘ": "c", "ㄙ": "s",

		// Rhymes
		"ㄚ": "a", "ㄛ": "o", "ㄜ": "e", "ㄝ": "e",
		"一": "i", "ㄨ": "u", "ㄩ": "v",
		"ㄞ": "ai", "ㄟ": "ei", "ㄦ": "er",
		"ㄠ": "ao", "ㄡ": "ou",
		"ㄢ": "an", "ㄣ": "en",
		"ㄤ": "ang", "ㄥ": "eng",

		"ㄨㄥ": "ong", "一ㄝ": "ie",
		"一ㄡ": "iu", "一ㄣ": "in", "一ㄥ": "ing",
		"ㄩㄝ": "ve",
		"ㄨㄣ": "un", "ㄩㄣ": "vn", "一ㄚ": "ia",
		"ㄨㄚ": "ua", "ㄨㄢ": "uan", "ㄩㄢ": "van",
		"ㄨㄞ": "uai", "ㄨㄛ": "uo", "ㄩㄥ": "iong",
		"一ㄤ": "iang", "ㄨㄤ": "uang", "一ㄢ": "ian",
		"一ㄠ": "iao", "ㄨㄟ": "ui",

		// 'y' and 'w' is not included because '一' and 'ㄨ' are already
		// mapped to 'i' and 'u'
		// "一": "y", "ㄨ": "w",
	}

	// tonal marks for Pinyin
	pinyinTones = [6][5]rune{
		{'a', 'ā', 'á', 'ǎ', 'à'},
		{'o', 'ō', 'ó', 'ǒ', 'ò'},
		{'e', 'ē', 'é', 'ě', 'è'},
		{'i', 'ī', 'í', 'ǐ', 'ì'},
		{'u', 'ū', 'ú', 'ǔ', 'ù'},
		{'ü', 'ǖ', 'ǘ', 'ǚ', 'ǜ'},
	}

	// tonal marks for Zhuyin
	zhuyinTones = [5]string{"˙", "", "ˊ", "ˇ", "ˋ"}
)

// toneChar returns the tonal marks for char 'c' of 'tone' in Pinyin
func getTonalMark(c byte, tone byte) string {
	r := rune(c)
	if r == 'v' {
		r = 'ü'
	}
	for _, t := range pinyinTones {
		if r == t[0] {
			return string(t[tone])
		}
	}
	panic("IMPOSSIBLE: should not run to here.")
}

// toneRhymes returns the toned rhymes in Pinyin
func toneRhymes(s string, tone byte) string {
	// if only one character, tone this character
	if len(s) == 1 {
		return getTonalMark(s[0], tone)
	}

	a, b := s[0], s[1]

	// Tone the 1st character, if:
	// * the first character is 'a'
	// * there's no 'a' and the 1st character is 'o' or 'e'
	// * the 2nd character is not rhymes
	if a == 'a' || ((a == 'o' || a == 'e') && b != 'a') || !isRhymes(b) {
		return getTonalMark(a, tone) + s[1:]
	}

	// tone the 2nd character otherwise
	return string(a) + getTonalMark(b, tone) + s[2:]
}

// return true if the input character is rhymes, otherwise return false
func isRhymes(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'v'
}

// return true if the input character is consonant, otherwise return false
func isConsonant(b byte) bool {
	return b >= 'a' && b <= 'z' && !isRhymes(b)
}

// split the input string into consonant, rhymes and tone
// for example: 'zhang1' will be split to consonant 'zh', rhymes 'ang' and
//              tone '1'
// returns an empty rhymes in case an error
func split(s string) (string, string, byte) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if !isConsonant(c) {
			break
		}
	}
	consonant := s[:i]

	for ; i < len(s); i++ {
		c := s[i]
		if c < 'a' || c > 'z' {
			break
		}
	}
	rhymes := s[len(consonant):i]

	// rhymes could not be empty, and the length of tone is at most 1
	if len(rhymes) == 0 || len(s)-i > 2 {
		return "", "", 0
	}

	var tone byte
	if i < len(s) {
		tone = s[i] - '0'
		if tone < 0 || tone > 4 {
			return "", "", 0
		}
	}

	return consonant, rhymes, tone
}

// encodePinyin encode the input consonant, rhymes and tone into Pinyin
// for example: encodePinyin("zh", "ang", 1) outputs 'zhāng'
// return an empty string in case an error
func encodePinyin(consonant string, rhymes string, tone byte) string {
	// rhymes could not be empty and the maximum value of tone is 4
	if len(rhymes) == 0 || tone > 4 {
		return ""
	}

	if len(consonant) > 0 {
		// is it an valid consonant?
		if !isConsonant(consonant[0]) {
			return ""
		}
		if _, ok := map_p2z[consonant]; !ok {
			return ""
		}

		// convert rhymes 'ü' to 'u' if consonant is 'j', 'q', 'x' or 'y'
		if rhymes[0] == 'v' {
			c := consonant[0]
			if c == 'j' || c == 'q' || c == 'x' || c == 'y' {
				rhymes = "u" + rhymes[1:]
			}
		}
	}

	// is it an valid rhymes?
	if !isRhymes(rhymes[0]) {
		return ""
	}
	if _, ok := map_p2z[rhymes]; !ok {
		return ""
	}

	// tone the rhymes and convert 'v' to 'ü'
	rhymes = toneRhymes(rhymes, tone)
	if rhymes[0] == 'v' {
		rhymes = "ü" + rhymes[1:]
	}

	return consonant + rhymes
}

// EncodePinyin encode the input string into Pinyin
// for example: EncodePinyin("zhang1") outputs 'zhāng'
// return an empty string in case an error
func EncodePinyin(s string) string {
	// the special case
	if s == "e5" {
		return "ê"
	}
	return encodePinyin(split(s))
}

// decodeRhymes decode the input string into rhymes and tone
// for example: decodeRhymes("āng") outputs 'ang' and 1
// returns an empty rhymes in case an error
func decodeRhymes(s string) (string, byte) {
	var tone byte
	var rhymes string
	for _, ch := range s {
		for _, t := range pinyinTones {
			for j := 1; j < len(t); j++ {
				if ch == t[j] {
					ch = t[0]
					if tone > 0 {
						return "", 0
					}
					tone = byte(j)
				}
			}
		}
		if ch == 'ü' {
			ch = 'v'
		}
		rhymes = rhymes + string(ch)
	}

	return rhymes, tone
}

// decodePinyin decode the input string into consonant, rhymes and tone
// for example: decodePinyin("zhāng") outputs 'zh','ang' and 1
// return an empty rhymes in case an error
func decodePinyin(s string) (string, string, byte) {
	var consonant, rhymes string

	// split the input into consonant and rhymes(toned)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !isConsonant(c) {
			consonant, rhymes = s[:i], s[i:]
			break
		}
	}

	// is it an valid consonant?
	if len(consonant) > 0 {
		if _, ok := map_p2z[consonant]; !ok {
			return "", "", 0
		}
	}

	// decode the toned rhymes into rhymes and tone
	rhymes, tone := decodeRhymes(rhymes)
	if len(rhymes) == 0 {
		return "", "", 0
	}

	// convert 'u' to 'v' if consonant is 'j', 'q', 'x' or 'y'
	if len(consonant) > 0 && rhymes[0] == 'u' {
		c := consonant[0]
		if c == 'j' || c == 'q' || c == 'x' || c == 'y' {
			rhymes = "v" + rhymes[1:]
		}
	}

	// is it an valid rhymes?
	if _, ok := map_p2z[rhymes]; !ok {
		return "", "", 0
	}

	return consonant, rhymes, tone
}

// DecodePinyin decode the input Pinyin
// for example: DecodePinyin("zhāng") outputs 'zhang1'
// return an empty string in case an error
func DecodePinyin(s string) string {
	if s == "ê" {
		return "e5"
	}

	consonant, rhymes, tone := decodePinyin(s)
	if len(rhymes) == 0 {
		return ""
	}
	return consonant + rhymes + string(tone+'0')
}

// encodePinyin encode the input consonant, rhymes and tone into Zhuyin
// for example: encodePinyin("m", "in", 2) outputs 'ㄇ一ㄣˊ'
// return an empty string in case an error
func encodeZhuyin(consonant string, rhymes string, tone byte) string {
	if len(rhymes) == 0 {
		return ""
	}

	if rhymes[0] == 'u' && len(consonant) > 0 {
		c := consonant[0]
		if c == 'j' || c == 'q' || c == 'x' || c == 'y' {
			rhymes = "v" + rhymes[1:]
		}
	}

	// the special cases for 'Zheng3 Ti3 Ren4 Du2'
	if rhymes == "i" {
		if consonant == "zh" || consonant == "ch" || consonant == "sh" ||
			consonant == "r" || consonant == "z" || consonant == "c" ||
			consonant == "s" || consonant == "y" {
			rhymes = ""
		}
	} else if consonant == "w" {
		if rhymes == "u" {
			consonant = ""
		}
	} else if consonant == "y" {
		if rhymes == "v" || rhymes == "e" || rhymes == "ve" || rhymes == "in" ||
			rhymes == "van" || rhymes == "ing" || rhymes == "vn" {
			consonant = ""
		}
	}

	// consonant must be valid
	if len(consonant) > 0 {
		if l, ok := map_p2z[consonant]; ok {
			consonant = l
		} else {
			return ""
		}
	}

	// rhymes must be valid
	if len(rhymes) > 0 {
		if l, ok := map_p2z[rhymes]; ok {
			rhymes = l
		} else {
			return ""
		}
	}

	return consonant + rhymes + zhuyinTones[tone]
}

// encodeZhuyin encode the input string into Zhuyin
// for example: encodeZhuyin("min2") outputs 'ㄇ一ㄣˊ'
// return an empty string in case an error
func EncodeZhuyin(s string) string {
	if s == "e5" {
		return "ㄝ"
	}
	return encodeZhuyin(split(s))
}

// decodeZhuyin decode the input string into consonant, rhymes and tone
// for example: decodeZhuyin("ㄇ一ㄣˊ") outputs 'm','in' and 2
// return an empty rhymes in case an error
func decodeZhuyin(s string) (string, string, byte) {
	var consonant, rhymes string
	var tone byte = 1

	// split input into consonant, rhymes and tone
split_input:
	for i, ch := range s {
		// if the character is consonant or rhymes
		if v, ok := map_z2p[string(ch)]; ok {
			// if it is the 1st character and it is consonant
			if i == 0 && isConsonant(v[0]) {
				consonant = v
			} else {
				// add it to rhymes, note, rhymes is still Zhuyin
				rhymes = rhymes + string(ch)
			}
			continue
		}

		// both consonant and rhymes is empty?
		if i == 0 {
			return "", "", 0
		}

		// try to find the tone
		st := s[i:]
		for j, t := range zhuyinTones {
			if st == t {
				tone = byte(j)
				break split_input
			}
		}
		return "", "", 0
	}

	if len(rhymes) == 0 {
		// if it is 'Zheng3 Ti3 Ren4 Du2', the rhymes should be 'i'
		if consonant == "zh" || consonant == "ch" || consonant == "sh" ||
			consonant == "r" || consonant == "z" || consonant == "c" ||
			consonant == "s" {
			rhymes = "i"
		}
		// rhymes will be empty if not 'Zheng3 Ti3 Ren4 Du2',
		// this is an error case, will be handled outside
		return consonant, rhymes, tone
	}

	// is it an valid rhymes?
	rhymes, ok := map_z2p[rhymes]
	if (!ok) || (!isRhymes(rhymes[0])) {
		return "", "", 0
	}

	if len(consonant) == 0 {
		// first:  check if it is 'Zheng3 Ti3 Ren4 Du2',
		// second: remove leading 'u' and set consonant to 'w', or
		//         remove leading 'i' and set consonant to 'y',
		// last:   check for the very special case,
		//         'ong' need to be converted to 'weng'
		if rhymes == "i" || rhymes == "v" || rhymes == "e" ||
			rhymes == "ve" || rhymes == "in" || rhymes == "van" ||
			rhymes == "ing" || rhymes == "vn" {
			consonant = "y"
		} else if rhymes == "u" {
			consonant = "w"
		} else if rhymes[0] == 'u' {
			consonant = "w"
			rhymes = rhymes[1:]
		} else if rhymes[0] == 'i' {
			consonant = "y"
			rhymes = rhymes[1:]
		} else if rhymes == "ong" {
			consonant = "w"
			rhymes = "eng"
		}
	}

	return consonant, rhymes, tone
}

// DecodePinyin decode the input Zhuyin
// for example: DecodeZhuyin("ㄇ一ㄣˊ") outputs 'min2'
// return an empty string in case an error
func DecodeZhuyin(s string) string {
	if s == "ㄝ" {
		return "e5"
	}

	consonant, rhymes, tone := decodeZhuyin(s)
	if len(rhymes) == 0 {
		return ""
	}
	return consonant + rhymes + string(tone+'0')
}

// PinyinToZhuyin converts the input Pinyin to Zhuyin
// for example: zhāng  -->  ㄓㄤ
func PinyinToZhuyin(s string) string {
	if s == "ê" {
		return "ㄝ"
	}
	return encodeZhuyin(decodePinyin(s))
}

// ZhuyinToPinyin converts the input Zhuyin to Pinyin
// for example: ㄓㄤ  -->  zhāng
func ZhuyinToPinyin(s string) string {
	if s == "ㄝ" {
		return "ê"
	}
	return encodePinyin(decodeZhuyin(s))
}

/*
type EncodeFunc func(p string) string

func encode(s string, fn EncodeFunc) []string {
	var result []string

	start := 0
	for i, c := range s {
		if c == ' ' || c == '\t' {
			if start < i {
				result = append(result, fn(s[start:i]))
			}
			start = i + 1
		} else if c >= '0' && c <= '5' {
			result = append(result, fn(s[start:i+1]))
			start = i + 1
		} else if c < 'a' || c > 'z' {

		}
	}

	if start < len(s) {
		result = append(result, fn(s[start:]))
	}

	return result
}
*/
