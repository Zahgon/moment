package moment

import (
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

// isoConstant and rfcConstant are moment.ISO_8601 and moment.RFC_2822, formats that select those parsers.
var (
	isoConstant = &js.Func{Name: "ISO_8601", Call: func(any, []any) any { return js.Undef }}
	rfcConstant = &js.Func{Name: "RFC_2822", Call: func(any, []any) any { return js.Undef }}
)

var (
	unescapeBracketsRegex = jsre.MustCompile(`\\(\[)|\\(\])|\[([^\]\[]*)\]|\\(.)`, "g")
	regexEscapeRegex      = jsre.MustCompile(`[-\/\\^$*+?.()|[\]{}]`, "g")
)

// regexEscape escapes the characters a RegExp source treats specially.
func regexEscape(s string) string {
	return regexEscapeRegex.Replace(s, `\$&`)
}

func unescapeFormat(s string) string {
	s = js.ReplaceFirst(s, `\`, "")
	s = unescapeBracketsRegex.ReplaceFunc(s, func(m *jsre.Match) string {
		for i := 1; i <= 4; i++ {
			if m.Defined[i] && m.Groups[i] != "" {
				return m.Groups[i]
			}
		}
		if m.Defined[4] {
			return m.Groups[4]
		}
		return "undefined"
	})
	return regexEscape(s)
}

func getParseRegexForToken(token string, config *Moment) *jsre.Regexp {
	if fn, ok := parseRegexes[token]; ok {
		return fn(js.Truthy(config.strict), config.locale)
	}
	re, err := jsre.Compile(unescapeFormat(token), "")
	if err != nil {
		return nil
	}
	return re
}

// firstMatch is String.prototype.match(re)[0] over UTF-16 text: the first match, or false.
func firstMatch(re *jsre.Regexp, s []uint16) ([]uint16, bool) {
	if re == nil {
		return nil, true
	}
	pos := re.ExecUTF16(s, 0)
	if pos == nil {
		return nil, false
	}
	return s[pos[0]:pos[1]], true
}

func indexUTF16(s, sub []uint16) int {
	for i := 0; i+len(sub) <= len(s); i++ {
		match := true
		for j := range sub {
			if s[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func (m *Moment) aGet(i int) Value {
	if i < len(m.a) {
		return m.a[i]
	}
	return js.Undef
}

func (m *Moment) aSet(i int, v Value) {
	for len(m.a) <= i {
		m.a = append(m.a, js.Undef)
	}
	m.a[i] = v
}

// configFromStringAndFormat parses config.i with the format config.f.
func configFromStringAndFormat(config *Moment) {
	if config.f == isoConstant {
		configFromISO(config)
		return
	}
	if config.f == rfcConstant {
		configFromRFC2822(config)
		return
	}
	config.a = []Value{}
	pf := getParsingFlags(config)
	pf.Set("empty", true)

	str := js.UTF16(js.ToString(config.i))
	stringLength := len(str)
	totalParsed := 0
	for _, token := range formatTokens(expandFormat(js.ToString(config.f), config.locale)) {
		parsed, matched := firstMatch(getParseRegexForToken(token, config), str)
		input := js.Undefined{}
		var parsedInput Value = input
		if matched {
			parsedInput = js.FromUTF16(parsed)
		}
		if len(parsed) > 0 {
			at := indexUTF16(str, parsed)
			if at > 0 {
				pushFlag(pf, "unusedInput", js.FromUTF16(str[:at]))
			}
			str = str[at+len(parsed):]
			totalParsed += len(parsed)
		}
		if _, ok := formatTokenFunctions[token]; ok {
			if len(parsed) > 0 {
				pf.Set("empty", false)
			} else {
				pushFlag(pf, "unusedTokens", token)
			}
			if s, ok := parsedInput.(string); ok {
				addTimeToArrayFromToken(token, s, config)
			}
		} else if js.Truthy(config.strict) && len(parsed) == 0 {
			pushFlag(pf, "unusedTokens", token)
		}
	}

	pf.Set("charsLeftOver", float64(stringLength-totalParsed))
	if len(str) > 0 {
		pushFlag(pf, "unusedInput", js.FromUTF16(str))
	}
	hour := config.aGet(hourIndex)
	if js.ToNumber(hour) <= 12 && js.StrictEquals(pf.Get("bigHour"), true) && js.ToNumber(hour) > 0 {
		pf.Set("bigHour", js.Undef)
	}
	parts := make([]any, len(config.a))
	copy(parts, config.a)
	pf.Set("parsedDateParts", parts)
	pf.Set("meridiem", config.meridiem)
	config.aSet(hourIndex, meridiemFixWrap(config.locale, config.aGet(hourIndex), config.meridiem))
	if era := pf.Get("era"); era != nil {
		config.aSet(yearIndex, config.locale.ErasConvertYear(era, config.aGet(yearIndex)))
	}
	configFromArray(config)
	checkOverflow(config)
}

func meridiemFixWrap(locale *Locale, hour, meridiem Value) Value {
	if js.IsNullish(meridiem) {
		return hour
	}
	if f, ok := locale.own.Get("meridiemHour").(func(*Locale, any, string) any); ok {
		return f(locale, hour, js.ToString(meridiem))
	}
	isPm := locale.IsPM(meridiem)
	if isPm && js.ToNumber(hour) < 12 {
		return js.Add(hour, 12.0)
	}
	if !isPm && js.StrictEquals(hour, 12.0) {
		return 0.0
	}
	return hour
}

// configFromStringAndArray parses with each format of config.f and keeps the best result.
func configFromStringAndArray(config *Moment) {
	formats := asArray(config.f)
	if len(formats) == 0 {
		getParsingFlags(config).Set("invalidFormat", true)
		config.d = js.NewDate(js.NaN)
		return
	}
	var best, temp *Moment
	scoreToBeat := js.NaN
	haveScore := false
	bestIsValid := false
	now := &nowCache{}
	for _, f := range formats {
		temp = copyConfig(newConfig(), config)
		if !js.IsNullish(config.useUTC) {
			temp.useUTC = config.useUTC
		}
		temp.defaultDatePartsNow = now
		temp.f = f
		configFromStringAndFormat(temp)
		validFound := isValidConfig(temp)
		pf := getParsingFlags(temp)
		unused, _ := pf.Get("unusedTokens").([]any)
		score := js.ToNumber(pf.Get("charsLeftOver")) + float64(len(unused))*10
		pf.Set("score", score)
		if !bestIsValid {
			if !haveScore || score < scoreToBeat || validFound {
				scoreToBeat, haveScore = score, true
				best = temp
				if validFound {
					bestIsValid = true
				}
			}
		} else if score < scoreToBeat {
			scoreToBeat = score
			best = temp
		}
	}
	if best == nil {
		best = temp
	}
	extendConfig(config, best)
}

// extendConfig copies every property the parse set on src onto config (moment's extend(config, tempConfig)).
func extendConfig(config, src *Moment) {
	*config = *src
}
