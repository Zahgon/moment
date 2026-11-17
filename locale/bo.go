// moment.js locale configuration
// locale : Tibetan [bo]
// author : Thupten N. Chakrishar : https://github.com/vajradog

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("bo", func() {
		moment.DefineLocale("bo", js.Obj(
			"months", []string{"ཟླ་བ་དང་པོ", "ཟླ་བ་གཉིས་པ", "ཟླ་བ་གསུམ་པ", "ཟླ་བ་བཞི་པ", "ཟླ་བ་ལྔ་པ", "ཟླ་བ་དྲུག་པ", "ཟླ་བ་བདུན་པ", "ཟླ་བ་བརྒྱད་པ", "ཟླ་བ་དགུ་པ", "ཟླ་བ་བཅུ་པ", "ཟླ་བ་བཅུ་གཅིག་པ", "ཟླ་བ་བཅུ་གཉིས་པ"},
			"monthsShort", []string{"ཟླ་1", "ཟླ་2", "ཟླ་3", "ཟླ་4", "ཟླ་5", "ཟླ་6", "ཟླ་7", "ཟླ་8", "ཟླ་9", "ཟླ་10", "ཟླ་11", "ཟླ་12"},
			"monthsShortRegex", jsre.MustCompile("^(ཟླ་\\d{1,2})", ""),
			"monthsParseExact", true,
			"weekdays", []string{"གཟའ་ཉི་མ་", "གཟའ་ཟླ་བ་", "གཟའ་མིག་དམར་", "གཟའ་ལྷག་པ་", "གཟའ་ཕུར་བུ", "གཟའ་པ་སངས་", "གཟའ་སྤེན་པ་"},
			"weekdaysShort", []string{"ཉི་མ་", "ཟླ་བ་", "མིག་དམར་", "ལྷག་པ་", "ཕུར་བུ", "པ་སངས་", "སྤེན་པ་"},
			"weekdaysMin", []string{"ཉི", "ཟླ", "མིག", "ལྷག", "ཕུར", "སངས", "སྤེན"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm",
				"LTS", "A h:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm",
				"LLLL", "dddd, D MMMM YYYY, A h:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[དི་རིང] LT",
				"nextDay", "[སང་ཉིན] LT",
				"nextWeek", "[བདུན་ཕྲག་རྗེས་མ], LT",
				"lastDay", "[ཁ་སང] LT",
				"lastWeek", "[བདུན་ཕྲག་མཐའ་མ] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s ལ་",
				"past", "%s སྔན་ལ",
				"s", "ལམ་སང",
				"ss", "%d སྐར་ཆ།",
				"m", "སྐར་མ་གཅིག",
				"mm", "%d སྐར་མ",
				"h", "ཆུ་ཚོད་གཅིག",
				"hh", "%d ཆུ་ཚོད",
				"d", "ཉིན་གཅིག",
				"dd", "%d ཉིན་",
				"M", "ཟླ་བ་གཅིག",
				"MM", "%d ཟླ་བ",
				"y", "ལོ་གཅིག",
				"yy", "%d ལོ",
			),
			"preparse", boPreparse,
			"postformat", boPostformat,
			"meridiemParse", jsre.MustCompile("མཚན་མོ|ཞོགས་ཀས|ཉིན་གུང|དགོང་དག|མཚན་མོ", ""),
			"meridiemHour", boMeridiemHour,
			"meridiem", boMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	boSymbols, boNumbers = digitTables("༠༡༢༣༤༥༦༧༨༩")
	boDigit              = jsre.MustCompile(`[༡༢༣༤༥༦༧༨༩༠]`, "g")
)

// boPreparse reads Tibetan digits.
func boPreparse(_ *moment.Locale, s string) string { return mapMatches(boDigit, s, boNumbers) }

// boPostformat writes Tibetan digits.
func boPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, boSymbols) }

// boMeridiemHour is the meridiemHour callback of bo.
func boMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	if (meridiem == "མཚན་མོ" && hourNumber(hour) >= 4) || (meridiem == "ཉིན་གུང" && hourNumber(hour) < 5) || meridiem == "དགོང་དག" {
		return plus12(hour)
	}
	return hour
}

// boMeridiem is the meridiem callback of bo.
func boMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "མཚན་མོ"
	case hour < 10:
		return "ཞོགས་ཀས"
	case hour < 17:
		return "ཉིན་གུང"
	case hour < 20:
		return "དགོང་དག"
	}
	return "མཚན་མོ"
}
