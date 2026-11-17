package js

import (
	"math"
	"unicode/utf16"
)

// DateFields are the fields of a date string; Offset is the UTC offset in milliseconds when HasOffset.
type DateFields struct {
	Year, Month, Day, Hour, Minute, Second, Millisecond float64
	Offset                                              float64
	HasOffset                                           bool
}

const dpNone = math.MinInt32

type dpKind int

const (
	dpInvalid dpKind = iota
	dpEnd
	dpNumber
	dpSymbol
	dpWhite
	dpKeyword
	dpUnknown
)

type dpWordType int

const (
	dpNoWord dpWordType = iota
	dpMonthName
	dpAmPm
	dpZoneName
	dpTimeSeparator
)

type dpToken struct {
	kind     dpKind
	value    int
	length   int
	symbol   uint16
	wordType dpWordType
}

func (t dpToken) isNumber(length int) bool { return t.kind == dpNumber && t.length == length }
func (t dpToken) isSymbol(c byte) bool     { return t.kind == dpSymbol && t.symbol == uint16(c) }
func (t dpToken) isSign() bool             { return t.isSymbol('+') || t.isSymbol('-') }
func (t dpToken) isZ() bool {
	return t.kind == dpKeyword && t.wordType == dpZoneName && t.length == 1 && t.value == 0
}

var dpKeywords = []struct {
	prefix string
	kind   dpWordType
	value  int
}{
	{"jan", dpMonthName, 1}, {"feb", dpMonthName, 2}, {"mar", dpMonthName, 3}, {"apr", dpMonthName, 4},
	{"may", dpMonthName, 5}, {"jun", dpMonthName, 6}, {"jul", dpMonthName, 7}, {"aug", dpMonthName, 8},
	{"sep", dpMonthName, 9}, {"oct", dpMonthName, 10}, {"nov", dpMonthName, 11}, {"dec", dpMonthName, 12},
	{"am", dpAmPm, 0}, {"pm", dpAmPm, 12}, {"ut", dpZoneName, 0}, {"utc", dpZoneName, 0}, {"z", dpZoneName, 0},
	{"gmt", dpZoneName, 0}, {"cdt", dpZoneName, -5}, {"cst", dpZoneName, -6}, {"edt", dpZoneName, -4},
	{"est", dpZoneName, -5}, {"mdt", dpZoneName, -6}, {"mst", dpZoneName, -7}, {"pdt", dpZoneName, -7},
	{"pst", dpZoneName, -8}, {"t", dpTimeSeparator, 0},
}

// dpScanner splits a date string into numbers, symbols, words, white space and skipped characters.
type dpScanner struct {
	s    []uint16
	pos  int
	next dpToken
}

// dpIsSpace reports the white space of the date scanner: JavaScript white space and \n, \r, but not U+2028 or
// U+2029, which the scanner reads as word characters.
func dpIsSpace(c uint16) bool { return c != 0x2028 && c != 0x2029 && IsSpace(rune(c)) }

func (sc *dpScanner) scan() dpToken {
	s := sc.s
	if sc.pos >= len(s) || s[sc.pos] == 0 {
		return dpToken{kind: dpEnd}
	}
	c := s[sc.pos]
	start := sc.pos
	switch {
	case c >= '0' && c <= '9':
		n := 0
		for sc.pos < len(s) && s[sc.pos] >= '0' && s[sc.pos] <= '9' {
			if n < math.MaxInt32/10-10 {
				n = n*10 + int(s[sc.pos]-'0')
			}
			sc.pos++
		}
		return dpToken{kind: dpNumber, value: n, length: sc.pos - start}
	case c == ':' || c == '-' || c == '+' || c == '.' || c == ')':
		sc.pos++
		return dpToken{kind: dpSymbol, symbol: c}
	case c >= 'A' && !dpIsSpace(c):
		var prefix []rune
		for sc.pos < len(s) && s[sc.pos] >= 'A' && !dpIsSpace(s[sc.pos]) {
			if len(prefix) < 3 {
				r := rune(s[sc.pos])
				if r >= 'A' && r <= 'Z' {
					r += 'a' - 'A'
				}
				prefix = append(prefix, r)
			}
			sc.pos++
		}
		length := sc.pos - start
		for _, k := range dpKeywords {
			if string(prefix) == k.prefix && (length <= 3 || k.kind == dpMonthName) {
				return dpToken{kind: dpKeyword, wordType: k.kind, value: k.value, length: length}
			}
		}
		return dpToken{kind: dpKeyword, wordType: dpNoWord, length: length}
	case dpIsSpace(c):
		for sc.pos < len(s) && dpIsSpace(s[sc.pos]) {
			sc.pos++
		}
		return dpToken{kind: dpWhite, length: sc.pos - start}
	case c == '(':
		balance := 0
		for {
			if sc.pos >= len(s) {
				break
			}
			switch s[sc.pos] {
			case ')':
				balance--
			case '(':
				balance++
			}
			sc.pos++
			if balance <= 0 || sc.pos >= len(s) || s[sc.pos] == 0 {
				break
			}
		}
		return dpToken{kind: dpUnknown}
	}
	sc.pos++
	return dpToken{kind: dpUnknown}
}

func (sc *dpScanner) take() dpToken {
	t := sc.next
	sc.next = sc.scan()
	return t
}

func (sc *dpScanner) skipSymbol(c byte) bool {
	if sc.next.isSymbol(c) {
		sc.take()
		return true
	}
	return false
}

type dpDay struct {
	comp       [3]int
	n          int
	namedMonth int
	iso        bool
}

func (d *dpDay) add(v int) bool {
	if d.n == 3 {
		return false
	}
	d.comp[d.n] = v
	d.n++
	return true
}

func dpBetween(x, lo, hi int) bool { return x >= lo && x <= hi }

func (d *dpDay) write(f *DateFields) bool {
	if d.n < 1 {
		return false
	}
	for d.n < 3 {
		d.comp[d.n] = 1
		d.n++
	}
	year, month, day := 0, dpNone, dpNone
	if d.namedMonth == dpNone {
		if d.iso || !dpBetween(d.comp[0], 1, 31) {
			year, month, day = d.comp[0], d.comp[1], d.comp[2]
		} else {
			month, day, year = d.comp[0], d.comp[1], d.comp[2]
		}
	} else {
		month = d.namedMonth
		if !dpBetween(d.comp[0], 1, 31) {
			year, day = d.comp[0], d.comp[1]
		} else {
			day, year = d.comp[0], d.comp[1]
		}
	}
	if !d.iso {
		if dpBetween(year, 0, 49) {
			year += 2000
		} else if dpBetween(year, 50, 99) {
			year += 1900
		}
	}
	if !dpBetween(month, 1, 12) || !dpBetween(day, 1, 31) || year < -(1<<30) || year > 1<<30-1 {
		return false
	}
	f.Year, f.Month, f.Day = float64(year), float64(month-1), float64(day)
	return true
}

type dpTime struct {
	comp       [4]int
	n          int
	hourOffset int
}

func (t *dpTime) empty() bool { return t.n == 0 }

func (t *dpTime) expecting(v int) bool {
	return (t.n == 1 && dpBetween(v, 0, 59)) || (t.n == 2 && dpBetween(v, 0, 59)) || (t.n == 3 && dpBetween(v, 0, 999))
}

func (t *dpTime) add(v int) bool {
	if t.n >= 4 {
		return false
	}
	t.comp[t.n] = v
	t.n++
	return true
}

func (t *dpTime) addFinal(v int) bool {
	if !t.add(v) {
		return false
	}
	for t.n < 4 {
		t.comp[t.n] = 0
		t.n++
	}
	return true
}

func (t *dpTime) write(f *DateFields) bool {
	for t.n < 4 {
		t.comp[t.n] = 0
		t.n++
	}
	hour, minute, second, ms := t.comp[0], t.comp[1], t.comp[2], t.comp[3]
	if t.hourOffset != dpNone {
		if !dpBetween(hour, 0, 12) {
			return false
		}
		hour = hour%12 + t.hourOffset
	}
	if !dpBetween(hour, 0, 23) || !dpBetween(minute, 0, 59) || !dpBetween(second, 0, 59) || !dpBetween(ms, 0, 999) {
		if hour != 24 || minute != 0 || second != 0 || ms != 0 {
			return false
		}
	}
	f.Hour, f.Minute, f.Second, f.Millisecond = float64(hour), float64(minute), float64(second), float64(ms)
	return true
}

type dpZone struct {
	sign, hour, minute int
}

func (z *dpZone) set(hours int) {
	z.sign = 1
	if hours < 0 {
		z.sign = -1
	}
	z.hour = hours * z.sign
	z.minute = 0
}

func (z *dpZone) expecting(v int) bool {
	return z.hour != dpNone && z.minute == dpNone && dpBetween(v, 0, 59)
}

func (z *dpZone) isUTC() bool { return z.hour == 0 && z.minute == 0 }

func (z *dpZone) write(f *DateFields) bool {
	if z.sign == dpNone {
		return true
	}
	if z.hour == dpNone {
		z.hour = 0
	}
	if z.minute == dpNone {
		z.minute = 0
	}
	total := int64(z.hour)*3600 + int64(z.minute)*60
	if total > 1<<30-1 {
		return false
	}
	if z.sign < 0 {
		total = -total
	}
	f.Offset, f.HasOffset = float64(total)*1000, true
	return true
}

// dpMilliseconds keeps the first three significant digits of a fraction, given its digit count.
func dpMilliseconds(t dpToken) int {
	n, length := t.value, t.length
	switch {
	case length == 1:
		return n * 100
	case length == 2:
		return n * 10
	case length > 3:
		if length > 9 {
			length = 9
		}
		factor := 1
		for length > 3 {
			factor *= 10
			length--
		}
		return n / factor
	}
	return n
}

// dpISO parses the ECMAScript date-time string format; it returns the first token it did not handle, an
// invalid token when the string is malformed, or the end token when it consumed the whole string.
func dpISO(sc *dpScanner, day *dpDay, tm *dpTime, tz *dpZone) dpToken {
	if sc.next.isSign() {
		signTok := sc.take()
		if !sc.next.isNumber(6) {
			return signTok
		}
		sign := 1
		if signTok.symbol == '-' {
			sign = -1
		}
		year := sc.take().value
		if sign < 0 && year == 0 {
			return signTok
		}
		day.add(sign * year)
	} else if sc.next.isNumber(4) {
		day.add(sc.take().value)
	} else {
		return sc.take()
	}
	if sc.skipSymbol('-') {
		if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 1, 12) {
			return sc.take()
		}
		day.add(sc.take().value)
		if sc.skipSymbol('-') {
			if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 1, 31) {
				return sc.take()
			}
			day.add(sc.take().value)
		}
	}
	if !(sc.next.kind == dpKeyword && sc.next.wordType == dpTimeSeparator) {
		if sc.next.kind != dpEnd {
			return sc.take()
		}
	} else {
		sc.take()
		if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 0, 24) {
			return dpToken{kind: dpInvalid}
		}
		hour24 := sc.next.value == 24
		tm.add(sc.take().value)
		if !sc.skipSymbol(':') {
			return dpToken{kind: dpInvalid}
		}
		if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 0, 59) || (hour24 && sc.next.value > 0) {
			return dpToken{kind: dpInvalid}
		}
		tm.add(sc.take().value)
		if sc.skipSymbol(':') {
			if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 0, 59) || (hour24 && sc.next.value > 0) {
				return dpToken{kind: dpInvalid}
			}
			tm.add(sc.take().value)
			if sc.skipSymbol('.') {
				if sc.next.kind != dpNumber || (hour24 && sc.next.value > 0) {
					return dpToken{kind: dpInvalid}
				}
				tm.add(dpMilliseconds(sc.take()))
			}
		}
		if sc.next.isZ() {
			sc.take()
			tz.set(0)
		} else if sc.next.isSign() {
			tz.sign = 1
			if sc.take().symbol == '-' {
				tz.sign = -1
			}
			if sc.next.isNumber(4) {
				hm := sc.take().value
				if !dpBetween(hm/100, 0, 23) || !dpBetween(hm%100, 0, 59) {
					return dpToken{kind: dpInvalid}
				}
				tz.hour, tz.minute = hm/100, hm%100
			} else {
				if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 0, 23) {
					return dpToken{kind: dpInvalid}
				}
				tz.hour = sc.take().value
				if !sc.skipSymbol(':') {
					return dpToken{kind: dpInvalid}
				}
				if !sc.next.isNumber(2) || !dpBetween(sc.next.value, 0, 59) {
					return dpToken{kind: dpInvalid}
				}
				tz.minute = sc.take().value
			}
		}
		if sc.next.kind != dpEnd {
			return dpToken{kind: dpInvalid}
		}
	}
	if tz.hour == dpNone && tm.empty() {
		tz.set(0)
	}
	day.iso = true
	return dpToken{kind: dpEnd}
}

// ParseDateString parses s as V8's Date.parse does and reports false when the result is NaN (Invalid Date).
// Month is 0-based. Years are full years (two-digit years already expanded as V8 does).
func ParseDateString(s string) (DateFields, bool) {
	sc := &dpScanner{s: utf16.Encode([]rune(s))}
	sc.next = sc.scan()
	day := &dpDay{namedMonth: dpNone}
	tm := &dpTime{hourOffset: dpNone}
	tz := &dpZone{sign: dpNone, hour: dpNone, minute: dpNone}
	var f DateFields
	tok := dpISO(sc, day, tm, tz)
	if tok.kind == dpInvalid {
		return f, false
	}
	hasNumber := day.n > 0
	for ; tok.kind != dpEnd; tok = sc.take() {
		switch {
		case tok.kind == dpNumber:
			hasNumber = true
			n := tok.value
			if sc.skipSymbol(':') {
				if sc.skipSymbol(':') {
					if !tm.empty() {
						return f, false
					}
					tm.add(n)
					tm.add(0)
				} else {
					if !tm.add(n) {
						return f, false
					}
					if sc.next.isSymbol('.') {
						sc.take()
					}
				}
			} else if sc.next.isSymbol('.') && tm.expecting(n) {
				sc.take()
				tm.add(n)
				if sc.next.kind != dpNumber {
					return f, false
				}
				tm.addFinal(dpMilliseconds(sc.take()))
			} else if tz.expecting(n) {
				tz.minute = n
			} else if tm.expecting(n) {
				tm.addFinal(n)
				p := sc.next
				if p.kind != dpEnd && p.kind != dpWhite && !p.isZ() && !p.isSign() {
					return f, false
				}
			} else {
				if !day.add(n) {
					return f, false
				}
				sc.skipSymbol('-')
			}
		case tok.kind == dpKeyword:
			switch {
			case tok.wordType == dpAmPm && !tm.empty():
				tm.hourOffset = tok.value
			case tok.wordType == dpMonthName:
				day.namedMonth = tok.value
				sc.skipSymbol('-')
			case tok.wordType == dpZoneName && hasNumber:
				tz.set(tok.value)
			default:
				if hasNumber || sc.next.kind == dpNumber {
					return f, false
				}
			}
		case tok.isSign() && (tz.isUTC() || !tm.empty()):
			tz.sign = 1
			if tok.symbol == '-' {
				tz.sign = -1
			}
			n, length := 0, 0
			if sc.next.kind == dpNumber {
				t := sc.take()
				n, length = t.value, t.length
			}
			hasNumber = true
			switch {
			case sc.next.isSymbol(':'):
				tz.hour, tz.minute = n, dpNone
			case length == 1 || length == 2:
				tz.hour, tz.minute = n, 0
			case length == 3 || length == 4:
				tz.hour, tz.minute = n/100, n%100
			default:
				return f, false
			}
		case (tok.isSign() || tok.isSymbol(')')) && hasNumber:
			return f, false
		}
	}
	ok := day.write(&f) && tm.write(&f) && tz.write(&f)
	return f, ok
}
