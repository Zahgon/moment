package moment

import (
	"math"
	"strings"

	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

var (
	match1                 = jsre.MustCompile(`\d`, "")
	match2                 = jsre.MustCompile(`\d\d`, "")
	match3                 = jsre.MustCompile(`\d{3}`, "")
	match4                 = jsre.MustCompile(`\d{4}`, "")
	match6                 = jsre.MustCompile(`[+-]?\d{6}`, "")
	match1to2              = jsre.MustCompile(`\d\d?`, "")
	match3to4              = jsre.MustCompile(`\d\d\d\d?`, "")
	match5to6              = jsre.MustCompile(`\d\d\d\d\d\d?`, "")
	match1to3              = jsre.MustCompile(`\d{1,3}`, "")
	match1to4              = jsre.MustCompile(`\d{1,4}`, "")
	match1to6              = jsre.MustCompile(`[+-]?\d{1,6}`, "")
	matchUnsigned          = jsre.MustCompile(`\d+`, "")
	matchSigned            = jsre.MustCompile(`[+-]?\d+`, "")
	matchOffset            = jsre.MustCompile(`Z|[+-]\d\d:?\d\d`, "gi")
	matchShortOffset       = jsre.MustCompile(`Z|[+-]\d\d(?::?\d\d)?`, "gi")
	matchTimestamp         = jsre.MustCompile(`[+-]?\d+(\.\d{1,3})?`, "")
	match1to2NoLeadingZero = jsre.MustCompile(`^[1-9]\d?`, "")
	match1to2HasZero       = jsre.MustCompile(`^([1-9]\d|\d)`, "")
	matchWord              = jsre.MustCompile(`[0-9]{0,256}['a-z\u00A0-\u05FF\u0700-\uD7FF\uF900-\uFDCF\uFDF0-\uFF07`+
		`\uFF10-\uFFEF]{1,256}|[\u0600-\u06FF\/]{1,256}(\s*?[\u0600-\u06FF]{1,256}){1,2}`, "i")
)

type tokenFunc func(m *Moment, format string) Value

type regexFunc func(strict bool, locale *Locale) *jsre.Regexp

type parseFunc func(input string, config *Moment, token string)

var (
	formatTokenFunctions = map[string]tokenFunc{}
	parseRegexes         = map[string]regexFunc{}
	parseTokens          = map[string]parseFunc{}
)

// addFormatToken registers a formatting token with its zero-padded and ordinal variants.
func addFormatToken(token string, padded string, length int, forceSign bool, ordinal string, fn tokenFunc) {
	if token != "" {
		formatTokenFunctions[token] = fn
	}
	if padded != "" {
		formatTokenFunctions[padded] = func(m *Moment, format string) Value {
			return zeroFill(js.ToNumber(fn(m, format)), length, forceSign)
		}
	}
	if ordinal != "" {
		formatTokenFunctions[ordinal] = func(m *Moment, format string) Value {
			return m.LocaleData().Ordinal(fn(m, format), token)
		}
	}
}

func num(f func(m *Moment) float64) tokenFunc {
	return func(m *Moment, _ string) Value { return f(m) }
}

func addRegexToken(token string, regex, strictRegex *jsre.Regexp) {
	parseRegexes[token] = func(strict bool, _ *Locale) *jsre.Regexp {
		if strict && strictRegex != nil {
			return strictRegex
		}
		return regex
	}
}

func addParseToken(tokens []string, fn parseFunc) {
	for _, t := range tokens {
		parseTokens[t] = fn
	}
}

func addArrayToken(tokens []string, index int) {
	addParseToken(tokens, func(input string, config *Moment, _ string) {
		config.aSet(index, toInt(input))
	})
}

func addWeekParseToken(tokens []string, fn func(input string, week *js.Object, config *Moment, token string)) {
	addParseToken(tokens, func(input string, config *Moment, token string) {
		if config.w == nil {
			config.w = js.NewObject()
		}
		fn(input, config.w, config, token)
	})
}

func addTimeToArrayFromToken(token, input string, config *Moment) {
	if fn, ok := parseTokens[token]; ok {
		fn(input, config, token)
	}
}

// zeroFill pads the absolute value of number with zeros to length, as moment's zeroFill does.
func zeroFill(number float64, length int, forceSign bool) string {
	abs := js.NumberToString(math.Abs(number))
	fill := length - js.Len16(abs)
	sign := ""
	if number >= 0 {
		if forceSign {
			sign = "+"
		}
	} else {
		sign = "-"
	}
	return sign + js.Substr16(js.NumberToString(math.Pow(10, math.Max(0, float64(fill)))), 1, math.MaxInt32) + abs
}

func hFormat(m *Moment) float64 {
	if h := math.Mod(m.Hours(), 12); h != 0 && !math.IsNaN(h) {
		return h
	}
	return 12
}

func kFormat(m *Moment) float64 {
	if h := m.Hours(); h != 0 && !math.IsNaN(h) {
		return h
	}
	return 24
}

func offsetToken(separator string) tokenFunc {
	return func(m *Moment, _ string) Value {
		offset := js.ToNumber(m.UTCOffset())
		sign := "+"
		if offset < 0 {
			offset = -offset
			sign = "-"
		}
		return sign + zeroFill(js.ToInt32(offset/60), 2, false) + separator + zeroFill(math.Mod(js.ToInt32(offset), 60), 2, false)
	}
}

func init() {
	registerFormatTokens()
	registerParseTokens()
}

func registerFormatTokens() {
	addFormatToken("M", "MM", 2, false, "Mo", num(func(m *Moment) float64 { return m.Month() + 1 }))
	addFormatToken("MMM", "", 0, false, "", func(m *Moment, f string) Value { return m.LocaleData().MonthsShort(m, f) })
	addFormatToken("MMMM", "", 0, false, "", func(m *Moment, f string) Value { return m.LocaleData().Months(m, f) })
	addFormatToken("Q", "", 0, false, "Qo", num((*Moment).Quarter))
	addFormatToken("D", "DD", 2, false, "Do", num((*Moment).Date))
	addFormatToken("DDD", "DDDD", 3, false, "DDDo", num((*Moment).DayOfYear))
	addFormatToken("d", "", 0, false, "do", num((*Moment).Day))
	addFormatToken("dd", "", 0, false, "", func(m *Moment, f string) Value { return m.LocaleData().WeekdaysMin(m) })
	addFormatToken("ddd", "", 0, false, "", func(m *Moment, f string) Value {
		return m.LocaleData().WeekdaysShort(m)
	})
	addFormatToken("dddd", "", 0, false, "", func(m *Moment, f string) Value { return m.LocaleData().Weekdays(m, f) })
	addFormatToken("e", "", 0, false, "", num((*Moment).Weekday))
	addFormatToken("E", "", 0, false, "", num((*Moment).ISOWeekday))
	addFormatToken("eHHmm", "", 0, false, "", func(m *Moment, _ string) Value {
		return js.NumberToString(m.Weekday()) + zeroFill(m.Hours(), 2, false) + zeroFill(m.Minutes(), 2, false)
	})
	addFormatToken("w", "ww", 2, false, "wo", num((*Moment).Week))
	addFormatToken("W", "WW", 2, false, "Wo", num((*Moment).ISOWeek))
	addFormatToken("", "gg", 2, false, "", num(func(m *Moment) float64 { return math.Mod(m.WeekYear(), 100) }))
	addFormatToken("", "GG", 2, false, "", num(func(m *Moment) float64 { return math.Mod(m.ISOWeekYear(), 100) }))
	for _, t := range []string{"gggg", "ggggg"} {
		addFormatToken("", t, len(t), false, "", num((*Moment).WeekYear))
	}
	for _, t := range []string{"GGGG", "GGGGG"} {
		addFormatToken("", t, len(t), false, "", num((*Moment).ISOWeekYear))
	}
	addFormatToken("Y", "", 0, false, "", func(m *Moment, _ string) Value {
		y := m.Year()
		if y <= 9999 {
			return zeroFill(y, 4, false)
		}
		return "+" + js.NumberToString(y)
	})
	addFormatToken("", "YY", 2, false, "", num(func(m *Moment) float64 { return math.Mod(m.Year(), 100) }))
	addFormatToken("", "YYYY", 4, false, "", num((*Moment).Year))
	addFormatToken("", "YYYYY", 5, false, "", num((*Moment).Year))
	addFormatToken("", "YYYYYY", 6, true, "", num((*Moment).Year))
	addFormatToken("H", "HH", 2, false, "", num((*Moment).Hours))
	addFormatToken("h", "hh", 2, false, "", num(hFormat))
	addFormatToken("k", "kk", 2, false, "", num(kFormat))
	addFormatToken("hmm", "", 0, false, "", func(m *Moment, _ string) Value {
		return js.NumberToString(hFormat(m)) + zeroFill(m.Minutes(), 2, false)
	})
	addFormatToken("hmmss", "", 0, false, "", func(m *Moment, _ string) Value {
		return js.NumberToString(hFormat(m)) + zeroFill(m.Minutes(), 2, false) + zeroFill(m.Seconds(), 2, false)
	})
	addFormatToken("Hmm", "", 0, false, "", func(m *Moment, _ string) Value {
		return js.NumberToString(m.Hours()) + zeroFill(m.Minutes(), 2, false)
	})
	addFormatToken("Hmmss", "", 0, false, "", func(m *Moment, _ string) Value {
		return js.NumberToString(m.Hours()) + zeroFill(m.Minutes(), 2, false) + zeroFill(m.Seconds(), 2, false)
	})
	addFormatToken("a", "", 0, false, "", func(m *Moment, _ string) Value {
		return m.LocaleData().Meridiem(m.Hours(), m.Minutes(), true)
	})
	addFormatToken("A", "", 0, false, "", func(m *Moment, _ string) Value {
		return m.LocaleData().Meridiem(m.Hours(), m.Minutes(), false)
	})
	addFormatToken("m", "mm", 2, false, "", num((*Moment).Minutes))
	addFormatToken("s", "ss", 2, false, "", num((*Moment).Seconds))
	addFormatToken("S", "", 0, false, "", num(func(m *Moment) float64 { return js.ToInt32(m.Milliseconds() / 100) }))
	addFormatToken("", "SS", 2, false, "", num(func(m *Moment) float64 { return js.ToInt32(m.Milliseconds() / 10) }))
	addFormatToken("", "SSS", 3, false, "", num((*Moment).Milliseconds))
	scale := 10.0
	for n := 4; n <= 9; n++ {
		factor := scale
		addFormatToken("", strings.Repeat("S", n), n, false, "", num(func(m *Moment) float64 {
			return m.Milliseconds() * factor
		}))
		scale *= 10
	}
	addFormatToken("X", "", 0, false, "", num((*Moment).Unix))
	addFormatToken("x", "", 0, false, "", num((*Moment).ValueOf))
	addFormatToken("z", "", 0, false, "", func(m *Moment, _ string) Value { return m.ZoneAbbr() })
	addFormatToken("zz", "", 0, false, "", func(m *Moment, _ string) Value { return m.ZoneName() })
	addFormatToken("Z", "", 0, false, "", offsetToken(":"))
	addFormatToken("ZZ", "", 0, false, "", offsetToken(""))
	for _, t := range []string{"N", "NN", "NNN"} {
		addFormatToken(t, "", 0, false, "", func(m *Moment, _ string) Value { return m.EraAbbr() })
	}
	addFormatToken("NNNN", "", 0, false, "", func(m *Moment, _ string) Value { return m.EraName() })
	addFormatToken("NNNNN", "", 0, false, "", func(m *Moment, _ string) Value { return m.EraNarrow() })
	addFormatToken("y", "y", 1, false, "yo", num((*Moment).EraYear))
	for i, t := range []string{"yy", "yyy", "yyyy"} {
		addFormatToken("y", t, i+2, false, "", num((*Moment).EraYear))
	}
}

func registerParseTokens() {
	addRegexToken("M", match1to2, match1to2NoLeadingZero)
	addRegexToken("MM", match1to2, match2)
	parseRegexes["MMM"] = func(strict bool, l *Locale) *jsre.Regexp { return l.MonthsShortRegex(strict) }
	parseRegexes["MMMM"] = func(strict bool, l *Locale) *jsre.Regexp { return l.MonthsRegex(strict) }
	addParseToken([]string{"M", "MM"}, func(input string, config *Moment, _ string) {
		config.aSet(monthIndex, toInt(input)-1)
	})
	addParseToken([]string{"MMM", "MMMM"}, func(input string, config *Moment, token string) {
		month := config.locale.MonthsParse(input, token, js.Truthy(config.strict))
		if !js.IsNullish(month) {
			config.aSet(monthIndex, month)
		} else {
			getParsingFlags(config).Set("invalidMonth", input)
		}
	})
	addRegexToken("Q", match1, nil)
	addParseToken([]string{"Q"}, func(input string, config *Moment, _ string) {
		config.aSet(monthIndex, (toInt(input)-1)*3)
	})
	addRegexToken("D", match1to2, match1to2NoLeadingZero)
	addRegexToken("DD", match1to2, match2)
	parseRegexes["Do"] = func(strict bool, l *Locale) *jsre.Regexp {
		if strict {
			if re, ok := l.own.Get("_dayOfMonthOrdinalParse").(*jsre.Regexp); ok {
				return re
			}
			re, _ := l.own.Get("_ordinalParse").(*jsre.Regexp)
			return re
		}
		re, _ := l.own.Get("_dayOfMonthOrdinalParseLenient").(*jsre.Regexp)
		return re
	}
	addArrayToken([]string{"D", "DD"}, dateIndex)
	addParseToken([]string{"Do"}, func(input string, config *Moment, _ string) {
		m := match1to2.Match(input)
		if m == nil {
			panic(typeError("Cannot read properties of null (reading '0')"))
		}
		config.aSet(dateIndex, toInt(m.Groups[0]))
	})
	addRegexToken("DDD", match1to3, nil)
	addRegexToken("DDDD", match3, nil)
	addParseToken([]string{"DDD", "DDDD"}, func(input string, config *Moment, _ string) {
		config.dayOfYear = toInt(input)
	})
	for _, t := range []string{"d", "e", "E"} {
		addRegexToken(t, match1to2, nil)
	}
	addRegexToken("eHHmm", match5to6, nil)
	parseRegexes["dd"] = func(strict bool, l *Locale) *jsre.Regexp { return l.WeekdaysMinRegex(strict) }
	parseRegexes["ddd"] = func(strict bool, l *Locale) *jsre.Regexp { return l.WeekdaysShortRegex(strict) }
	parseRegexes["dddd"] = func(strict bool, l *Locale) *jsre.Regexp { return l.WeekdaysRegex(strict) }
	addWeekParseToken([]string{"dd", "ddd", "dddd"}, func(input string, week *js.Object, config *Moment, token string) {
		weekday := config.locale.WeekdaysParse(input, token, js.Truthy(config.strict))
		if !js.IsNullish(weekday) {
			week.Set("d", weekday)
		} else {
			getParsingFlags(config).Set("invalidWeekday", input)
		}
	})
	addWeekParseToken([]string{"d", "e", "E"}, func(input string, week *js.Object, _ *Moment, token string) {
		week.Set(token, toInt(input))
	})
	addWeekParseToken([]string{"eHHmm"}, func(input string, week *js.Object, config *Moment, _ string) {
		end := js.Len16(input) - 4
		week.Set("e", toInt(js.Substr16(input, 0, end)))
		config.aSet(hourIndex, toInt(js.Substr16(input, end, 2)))
		config.aSet(minuteIndex, toInt(js.Substr16(input, end+2, math.MaxInt32)))
	})
	for _, t := range []string{"w", "W"} {
		addRegexToken(t, match1to2, match1to2NoLeadingZero)
		addRegexToken(t+t, match1to2, match2)
	}
	addWeekParseToken([]string{"w", "ww", "W", "WW"}, func(input string, week *js.Object, _ *Moment, token string) {
		week.Set(token[:1], toInt(input))
	})
	for _, p := range []string{"G", "g"} {
		addRegexToken(p, matchSigned, nil)
		addRegexToken(p+p, match1to2, match2)
		addRegexToken(strings.Repeat(p, 4), match1to4, match4)
		addRegexToken(strings.Repeat(p, 5), match1to6, match6)
	}
	addWeekParseToken([]string{"gggg", "ggggg", "GGGG", "GGGGG"}, func(input string, week *js.Object, _ *Moment, token string) {
		week.Set(token[:2], toInt(input))
	})
	addWeekParseToken([]string{"gg", "GG"}, func(input string, week *js.Object, _ *Moment, token string) {
		week.Set(token, ParseTwoDigitYear(input))
	})
	addRegexToken("Y", matchSigned, nil)
	addRegexToken("YY", match1to2, match2)
	addRegexToken("YYYY", match1to4, match4)
	addRegexToken("YYYYY", match1to6, match6)
	addRegexToken("YYYYYY", match1to6, match6)
	addArrayToken([]string{"YYYYY", "YYYYYY"}, yearIndex)
	addParseToken([]string{"YYYY"}, func(input string, config *Moment, _ string) {
		if js.Len16(input) == 2 {
			config.aSet(yearIndex, ParseTwoDigitYear(input))
		} else {
			config.aSet(yearIndex, toInt(input))
		}
	})
	addParseToken([]string{"YY"}, func(input string, config *Moment, _ string) {
		config.aSet(yearIndex, ParseTwoDigitYear(input))
	})
	addParseToken([]string{"Y"}, func(input string, config *Moment, _ string) {
		config.aSet(yearIndex, js.ParseInt(input))
	})
	meridiemRegex := func(_ bool, l *Locale) *jsre.Regexp {
		re, _ := l.own.Get("_meridiemParse").(*jsre.Regexp)
		return re
	}
	parseRegexes["a"] = meridiemRegex
	parseRegexes["A"] = meridiemRegex
	addRegexToken("H", match1to2, match1to2HasZero)
	addRegexToken("h", match1to2, match1to2NoLeadingZero)
	addRegexToken("k", match1to2, match1to2NoLeadingZero)
	for _, t := range []string{"HH", "hh", "kk"} {
		addRegexToken(t, match1to2, match2)
	}
	addRegexToken("hmm", match3to4, nil)
	addRegexToken("hmmss", match5to6, nil)
	addRegexToken("Hmm", match3to4, nil)
	addRegexToken("Hmmss", match5to6, nil)
	addArrayToken([]string{"H", "HH"}, hourIndex)
	addParseToken([]string{"k", "kk"}, func(input string, config *Moment, _ string) {
		k := toInt(input)
		if k == 24 {
			k = 0
		}
		config.aSet(hourIndex, k)
	})
	addParseToken([]string{"a", "A"}, func(input string, config *Moment, _ string) {
		config.isPm = config.locale.IsPM(input)
		config.meridiem = input
	})
	addParseToken([]string{"h", "hh"}, func(input string, config *Moment, _ string) {
		config.aSet(hourIndex, toInt(input))
		getParsingFlags(config).Set("bigHour", true)
	})
	hourMinute := func(bigHour bool) parseFunc {
		return func(input string, config *Moment, _ string) {
			pos := js.Len16(input) - 2
			config.aSet(hourIndex, toInt(js.Substr16(input, 0, pos)))
			config.aSet(minuteIndex, toInt(js.Substr16(input, pos, math.MaxInt32)))
			if bigHour {
				getParsingFlags(config).Set("bigHour", true)
			}
		}
	}
	hourMinuteSecond := func(bigHour bool) parseFunc {
		return func(input string, config *Moment, _ string) {
			n := js.Len16(input)
			config.aSet(hourIndex, toInt(js.Substr16(input, 0, n-4)))
			config.aSet(minuteIndex, toInt(js.Substr16(input, n-4, 2)))
			config.aSet(secondIndex, toInt(js.Substr16(input, n-2, math.MaxInt32)))
			if bigHour {
				getParsingFlags(config).Set("bigHour", true)
			}
		}
	}
	addParseToken([]string{"hmm"}, hourMinute(true))
	addParseToken([]string{"hmmss"}, hourMinuteSecond(true))
	addParseToken([]string{"Hmm"}, hourMinute(false))
	addParseToken([]string{"Hmmss"}, hourMinuteSecond(false))
	addRegexToken("m", match1to2, match1to2HasZero)
	addRegexToken("mm", match1to2, match2)
	addArrayToken([]string{"m", "mm"}, minuteIndex)
	addRegexToken("s", match1to2, match1to2HasZero)
	addRegexToken("ss", match1to2, match2)
	addArrayToken([]string{"s", "ss"}, secondIndex)
	addRegexToken("S", match1to3, match1)
	addRegexToken("SS", match1to3, match2)
	addRegexToken("SSS", match1to3, match3)
	for n := 4; n <= 9; n++ {
		addRegexToken(strings.Repeat("S", n), matchUnsigned, nil)
	}
	for n := 1; n <= 9; n++ {
		addParseToken([]string{strings.Repeat("S", n)}, func(input string, config *Moment, _ string) {
			config.aSet(millisecondIndex, toInt(js.StringToNumber("0."+input)*1000))
		})
	}
	addRegexToken("x", matchSigned, nil)
	addRegexToken("X", matchTimestamp, nil)
	addParseToken([]string{"X"}, func(input string, config *Moment, _ string) {
		config.d = js.NewDate(js.ParseFloat(input) * 1000)
	})
	addParseToken([]string{"x"}, func(input string, config *Moment, _ string) {
		config.d = js.NewDate(toInt(input))
	})
	addRegexToken("Z", matchShortOffset, nil)
	addRegexToken("ZZ", matchShortOffset, nil)
	addParseToken([]string{"Z", "ZZ"}, func(input string, config *Moment, _ string) {
		offset := offsetFromString(matchShortOffset, input)
		config.useUTC = true
		config.tzm = offset
		if offset == nil {
			getParsingFlags(config).Set("invalidOffset", input)
		}
	})
	for _, t := range []string{"N", "NN", "NNN"} {
		parseRegexes[t] = func(strict bool, l *Locale) *jsre.Regexp { return l.ErasAbbrRegex(strict) }
	}
	parseRegexes["NNNN"] = func(strict bool, l *Locale) *jsre.Regexp { return l.ErasNameRegex(strict) }
	parseRegexes["NNNNN"] = func(strict bool, l *Locale) *jsre.Regexp { return l.ErasNarrowRegex(strict) }
	addParseToken([]string{"N", "NN", "NNN", "NNNN", "NNNNN"}, func(input string, config *Moment, token string) {
		era := config.locale.ErasParse(input, token, js.Truthy(config.strict))
		if js.Truthy(era) {
			getParsingFlags(config).Set("era", era)
		} else {
			getParsingFlags(config).Set("invalidEra", input)
		}
	})
	for _, t := range []string{"y", "yy", "yyy", "yyyy"} {
		addRegexToken(t, matchUnsigned, nil)
	}
	parseRegexes["yo"] = func(_ bool, l *Locale) *jsre.Regexp {
		if re, ok := l.own.Get("_eraYearOrdinalRegex").(*jsre.Regexp); ok {
			return re
		}
		return matchUnsigned
	}
	addArrayToken([]string{"y", "yy", "yyy", "yyyy"}, yearIndex)
	addParseToken([]string{"yo"}, func(input string, config *Moment, _ string) {
		var match *jsre.Match
		if re, ok := config.locale.own.Get("_eraYearOrdinalRegex").(*jsre.Regexp); ok {
			match = re.Match(input)
		}
		if f, ok := config.locale.own.Get("eraYearOrdinalParse").(func(*Locale, string, *jsre.Match) float64); ok {
			config.aSet(yearIndex, f(config.locale, input, match))
		} else {
			config.aSet(yearIndex, js.ParseInt(input))
		}
	})
}
