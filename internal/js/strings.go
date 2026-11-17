package js

import (
	"strings"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
)

// ToString is the String(value) conversion.
func ToString(v any) string {
	switch x := v.(type) {
	case nil:
		return "null"
	case Undefined:
		return "undefined"
	case bool:
		if x {
			return "true"
		}
		return "false"
	case float64:
		return NumberToString(x)
	case string:
		return x
	case []string:
		return strings.Join(x, ",")
	case []any:
		parts := make([]string, len(x))
		for i, e := range x {
			if !IsNullish(e) {
				parts[i] = ToString(e)
			}
		}
		return strings.Join(parts, ",")
	}
	return ToString(ToPrimitive(v, true))
}

// IsSpace reports whether r is a JavaScript WhiteSpace or LineTerminator character (what trim and Number()
// remove).
func IsSpace(r rune) bool {
	switch r {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0xA0, 0x1680, 0x2028, 0x2029, 0x202F, 0x205F, 0x3000, 0xFEFF:
		return true
	}
	return r >= 0x2000 && r <= 0x200A
}

// TrimSpace is String.prototype.trim.
func TrimSpace(s string) string {
	return strings.TrimFunc(s, IsSpace)
}

// TrimLeftSpace is String.prototype.trimStart.
func TrimLeftSpace(s string) string {
	return strings.TrimLeftFunc(s, IsSpace)
}

// UTF16 returns the UTF-16 code units of s.
func UTF16(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

// FromUTF16 converts code units back to a string; a lone surrogate becomes U+FFFD.
func FromUTF16(u []uint16) string {
	return string(utf16.Decode(u))
}

// Len16 is the length property of a string (UTF-16 code units).
func Len16(s string) int {
	n := 0
	for _, r := range s {
		if r >= 0x10000 {
			n += 2
		} else {
			n++
		}
	}
	return n
}

func clampIndex(i, n int) int {
	if i < 0 {
		return 0
	}
	if i > n {
		return n
	}
	return i
}

// Slice16 is String.prototype.slice(start, end) with end = Len16(s) when omitted by the caller.
func Slice16(s string, start, end int) string {
	u := UTF16(s)
	n := len(u)
	if start < 0 {
		start += n
	}
	if end < 0 {
		end += n
	}
	start, end = clampIndex(start, n), clampIndex(end, n)
	if start >= end {
		return ""
	}
	return FromUTF16(u[start:end])
}

// Substr16 is String.prototype.substr(start, length).
func Substr16(s string, start, length int) string {
	u := UTF16(s)
	n := len(u)
	if start < 0 {
		start += n
	}
	start = clampIndex(start, n)
	end := clampIndex(start+length, n)
	if length <= 0 || start >= end {
		return ""
	}
	return FromUTF16(u[start:end])
}

// Substring16 is String.prototype.substring(start, end).
func Substring16(s string, start, end int) string {
	n := Len16(s)
	start, end = clampIndex(start, n), clampIndex(end, n)
	if start > end {
		start, end = end, start
	}
	return Slice16(s, start, end)
}

// CharAt16 is String.prototype.charAt(i).
func CharAt16(s string, i int) string {
	u := UTF16(s)
	if i < 0 || i >= len(u) {
		return ""
	}
	return FromUTF16(u[i : i+1])
}

// IndexOf16 is String.prototype.indexOf(sub) in UTF-16 units.
func IndexOf16(s, sub string) int {
	i := strings.Index(s, sub)
	if i < 0 {
		return -1
	}
	return Len16(s[:i])
}

// ReplaceFirst is String.prototype.replace(pattern, replacement) with a string pattern: the first occurrence is
// replaced and the replacement's $$, $& , $` and $' patterns are expanded.
func ReplaceFirst(s, pattern, replacement string) string {
	i := strings.Index(s, pattern)
	if i < 0 {
		return s
	}
	var b strings.Builder
	for j := 0; j < len(replacement); j++ {
		c := replacement[j]
		if c != '$' || j+1 >= len(replacement) {
			b.WriteByte(c)
			continue
		}
		switch replacement[j+1] {
		case '$':
			b.WriteByte('$')
		case '&':
			b.WriteString(pattern)
		case '`':
			b.WriteString(s[:i])
		case '\'':
			b.WriteString(s[i+len(pattern):])
		default:
			b.WriteByte(c)
			continue
		}
		j++
	}
	return s[:i] + b.String() + s[i+len(pattern):]
}

// Split is String.prototype.split(separator) with a string separator; an empty separator splits into UTF-16
// code units.
func Split(s, sep string) []string {
	if sep == "" {
		u := UTF16(s)
		out := make([]string, len(u))
		for i := range u {
			out[i] = FromUTF16(u[i : i+1])
		}
		return out
	}
	return strings.Split(s, sep)
}

// upperSpecial holds the unconditional multi-character uppercase mappings of Unicode SpecialCasing.
var upperSpecial = map[rune]string{
	0x00DF: "SS", 0x0149: "\u02BCN", 0x01F0: "J\u030C", 0x0390: "\u0399\u0308\u0301", 0x03B0: "\u03A5\u0308\u0301",
	0x0587: "\u0535\u0552", 0x1E96: "H\u0331", 0x1E97: "T\u0308", 0x1E98: "W\u030A", 0x1E99: "Y\u030A",
	0x1E9A: "A\u02BE", 0x1F50: "\u03A5\u0313", 0x1F52: "\u03A5\u0313\u0300", 0x1F54: "\u03A5\u0313\u0301",
	0x1F56: "\u03A5\u0313\u0342", 0x1FB6: "\u0391\u0342", 0x1FC6: "\u0397\u0342", 0x1FD2: "\u0399\u0308\u0300",
	0x1FD3: "\u0399\u0308\u0301", 0x1FD6: "\u0399\u0342", 0x1FD7: "\u0399\u0308\u0342", 0x1FE2: "\u03A5\u0308\u0300",
	0x1FE3: "\u03A5\u0308\u0301", 0x1FE4: "\u03A1\u0313", 0x1FE6: "\u03A5\u0342", 0x1FE7: "\u03A5\u0308\u0342",
	0x1FF6: "\u03A9\u0342", 0x1FB3: "\u0391\u0399", 0x1FBC: "\u0391\u0399", 0x1FC3: "\u0397\u0399",
	0x1FCC: "\u0397\u0399", 0x1FF3: "\u03A9\u0399", 0x1FFC: "\u03A9\u0399", 0x1FB2: "\u1FBA\u0399",
	0x1FB4: "\u0386\u0399", 0x1FC2: "\u1FCA\u0399", 0x1FC4: "\u0389\u0399", 0x1FF2: "\u1FFA\u0399",
	0x1FF4: "\u038F\u0399", 0x1FB7: "\u0391\u0342\u0399", 0x1FC7: "\u0397\u0342\u0399", 0x1FF7: "\u03A9\u0342\u0399",
	0xFB00: "FF", 0xFB01: "FI", 0xFB02: "FL", 0xFB03: "FFI", 0xFB04: "FFL", 0xFB05: "ST", 0xFB06: "ST",
	0xFB13: "\u0544\u0546", 0xFB14: "\u0544\u0535", 0xFB15: "\u0544\u053B", 0xFB16: "\u054E\u0546",
	0xFB17: "\u0544\u053D",
}

func init() {
	// Greek letters with ypogegrammeni or prosgegrammeni uppercase to the capital without it plus IOTA.
	for _, base := range []rune{0x1F80, 0x1F90, 0x1FA0} {
		capital := map[rune]rune{0x1F80: 0x1F08, 0x1F90: 0x1F28, 0x1FA0: 0x1F68}[base]
		for i := rune(0); i < 8; i++ {
			upperSpecial[base+i] = string([]rune{capital + i, 0x0399})
			upperSpecial[base+8+i] = string([]rune{capital + i, 0x0399})
		}
	}
}

// ToUpper is String.prototype.toUpperCase (full Unicode case mapping).
func ToUpper(s string) string {
	var b strings.Builder
	for _, r := range s {
		if m, ok := upperSpecial[r]; ok {
			b.WriteString(m)
			continue
		}
		b.WriteRune(unicode.ToUpper(r))
	}
	return b.String()
}

// ToLower is String.prototype.toLowerCase (full Unicode case mapping, including the final sigma rule).
func ToLower(s string) string {
	rs := []rune(s)
	var b strings.Builder
	for i, r := range rs {
		switch r {
		case 0x0130:
			b.WriteString("i\u0307")
		case 0x03A3:
			if finalSigma(rs, i) {
				b.WriteRune(0x03C2)
			} else {
				b.WriteRune(0x03C3)
			}
		default:
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func isCased(r rune) bool {
	return unicode.IsUpper(r) || unicode.IsLower(r) || unicode.IsTitle(r) ||
		unicode.Is(unicode.Other_Lowercase, r) || unicode.Is(unicode.Other_Uppercase, r)
}

func isCaseIgnorable(r rune) bool {
	switch r {
	case '\'', '.', ':', '^', '`', 0x00A8, 0x00AD, 0x00AF, 0x00B4, 0x00B7, 0x00B8, 0x2018, 0x2019, 0x2024, 0x2027:
		return true
	}
	return unicode.In(r, unicode.Mn, unicode.Me, unicode.Cf, unicode.Lm, unicode.Sk)
}

// finalSigma applies the Final_Sigma condition to the capital sigma at rs[i].
func finalSigma(rs []rune, i int) bool {
	before := false
	for j := i - 1; j >= 0; j-- {
		if isCaseIgnorable(rs[j]) {
			continue
		}
		before = isCased(rs[j])
		break
	}
	if !before {
		return false
	}
	for j := i + 1; j < len(rs); j++ {
		if isCaseIgnorable(rs[j]) {
			continue
		}
		return !isCased(rs[j])
	}
	return true
}

// FirstRune returns the first character of s (as charAt(0) for BMP text) or "".
func FirstRune(s string) string {
	if s == "" {
		return ""
	}
	_, n := utf8.DecodeRuneInString(s)
	return s[:n]
}
