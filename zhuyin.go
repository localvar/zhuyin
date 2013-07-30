package zhuyin

var (
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

		// "一": "y", "ㄨ": "w",
	}

	pinyinTones = [6][5]rune{
		{'a', 'ā', 'á', 'ǎ', 'à'},
		{'o', 'ō', 'ó', 'ǒ', 'ò'},
		{'e', 'ē', 'é', 'ě', 'è'},
		{'i', 'ī', 'í', 'ǐ', 'ì'},
		{'u', 'ū', 'ú', 'ǔ', 'ù'},
		{'ü', 'ǖ', 'ǘ', 'ǚ', 'ǜ'},
	}

	zhuyinTones = [5]string{"˙", "", "ˊ", "ˇ", "ˋ"}
)

func toneChar(c byte, tone byte) rune {
	r := rune(c)
	if r == 'v' {
		r = 'ü'
	}
	for _, t := range pinyinTones {
		if r == t[0] {
			return t[tone]
		}
	}
	return r
}

func toneRhymes(s string, tone byte) string {
	if len(s) == 1 {
		return string(toneChar(s[0], tone))
	}

	a, b := s[0], s[1]
	if a == 'a' || ((a == 'o' || a == 'e') && b != 'a') {
		return string(toneChar(a, tone)) + s[1:]
	}
	return string(a) + string(toneChar(b, tone)) + s[2:]
}

func isRhymes(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'v'
}

func isConsonant(b byte) bool {
	return b >= 'a' && b <= 'z' && !isRhymes(b)
}

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

func encodePinyin(consonant string, rhymes string, tone byte) string {
	if len(rhymes) == 0 {
		return ""
	}

	if len(consonant) > 0 {
		if _, ok := map_p2z[consonant]; !ok {
			return ""
		}

		if rhymes[0] == 'v' {
			c := consonant[0]
			if c == 'j' || c == 'q' || c == 'x' || c == 'y' {
				rhymes = "u" + rhymes[1:]
			}
		}
	}

	if _, ok := map_p2z[rhymes]; !ok {
		return ""
	}

	if rhymes = toneRhymes(rhymes, tone); len(rhymes) == 0 {
		return ""
	}

	if rhymes[0] == 'v' {
		rhymes = "ü" + rhymes[1:]
	}

	return consonant + rhymes
}

func EncodePinyin(s string) string {
	if s == "e5" {
		return "ê"
	}
	return encodePinyin(split(s))
}

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

func decodePinyin(s string) (string, string, byte) {
	var consonant, rhymes string

	for i := 0; i < len(s); i++ {
		c := s[i]
		if !isConsonant(c) {
			consonant, rhymes = s[:i], s[i:]
			break
		}
	}
	if len(consonant) > 0 {
		if _, ok := map_p2z[consonant]; !ok {
			return "", "", 0
		}
	}

	rhymes, tone := decodeRhymes(rhymes)
	if len(rhymes) == 0 {
		return "", "", 0
	}

	if len(consonant) > 0 && rhymes[0] == 'u' {
		c := consonant[0]
		if c == 'j' || c == 'q' || c == 'x' || c == 'y' {
			rhymes = "v" + rhymes[1:]
		}
	}

	if _, ok := map_p2z[rhymes]; !ok {
		return "", "", 0
	}

	return consonant, rhymes, tone
}

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

	if len(consonant) > 0 {
		if l, ok := map_p2z[consonant]; ok {
			consonant = l
		} else {
			return ""
		}
	}

	if len(rhymes) > 0 {
		if l, ok := map_p2z[rhymes]; ok {
			rhymes = l
		} else {
			return ""
		}
	}

	return consonant + rhymes + zhuyinTones[tone]
}

func EncodeZhuyin(s string) string {
	if s == "e5" {
		return "ㄝ"
	}
	return encodeZhuyin(split(s))
}

func decodeZhuyin(s string) (string, string, byte) {
	var consonant, rhymes string
	var tone byte = 1

split_input:
	for i, ch := range s {
		if v, ok := map_z2p[string(ch)]; ok {
			if i == 0 && isConsonant(v[0]) {
				consonant = v
			} else {
				rhymes = rhymes + string(ch)
			}
			continue
		}

		if i == 0 {
			return "", "", 0
		}

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
		if consonant == "zh" || consonant == "ch" || consonant == "sh" ||
			consonant == "r" || consonant == "z" || consonant == "c" ||
			consonant == "s" {
			rhymes = "i"
		}
		return consonant, rhymes, tone
	}

	rhymes, ok := map_z2p[rhymes]
	if (!ok) || (!isRhymes(rhymes[0])) {
		return "", "", 0
	}
	if len(consonant) == 0 {
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
		}
	}

	return consonant, rhymes, tone
}

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

func PinyinToZhuyin(s string) string {
	return encodeZhuyin(decodePinyin(s))
}

func ZhuyinToPinyin(s string) string {
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
