// moment.js locale configuration
// locale : Uyghur (China) [ug-cn]
// author: boyaq : https://github.com/boyaq

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ug-cn", func() {
		moment.DefineLocale("ug-cn", js.Obj(
			"months", []string{"يانۋار", "فېۋرال", "مارت", "ئاپرېل", "ماي", "ئىيۇن", "ئىيۇل", "ئاۋغۇست", "سېنتەبىر", "ئۆكتەبىر", "نويابىر", "دېكابىر"},
			"monthsShort", []string{"يانۋار", "فېۋرال", "مارت", "ئاپرېل", "ماي", "ئىيۇن", "ئىيۇل", "ئاۋغۇست", "سېنتەبىر", "ئۆكتەبىر", "نويابىر", "دېكابىر"},
			"weekdays", []string{"يەكشەنبە", "دۈشەنبە", "سەيشەنبە", "چارشەنبە", "پەيشەنبە", "جۈمە", "شەنبە"},
			"weekdaysShort", []string{"يە", "دۈ", "سە", "چا", "پە", "جۈ", "شە"},
			"weekdaysMin", []string{"يە", "دۈ", "سە", "چا", "پە", "جۈ", "شە"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY-MM-DD",
				"LL", "YYYY-يىلىM-ئاينىڭD-كۈنى",
				"LLL", "YYYY-يىلىM-ئاينىڭD-كۈنى، HH:mm",
				"LLLL", "dddd، YYYY-يىلىM-ئاينىڭD-كۈنى، HH:mm",
			),
			"meridiemParse", jsre.MustCompile("يېرىم كېچە|سەھەر|چۈشتىن بۇرۇن|چۈش|چۈشتىن كېيىن|كەچ", ""),
			"meridiemHour", ugCnMeridiemHour,
			"meridiem", ugCnMeridiem,
			"calendar", js.Obj(
				"sameDay", "[بۈگۈن سائەت] LT",
				"nextDay", "[ئەتە سائەت] LT",
				"nextWeek", "[كېلەركى] dddd [سائەت] LT",
				"lastDay", "[تۆنۈگۈن] LT",
				"lastWeek", "[ئالدىنقى] dddd [سائەت] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s كېيىن",
				"past", "%s بۇرۇن",
				"s", "نەچچە سېكونت",
				"ss", "%d سېكونت",
				"m", "بىر مىنۇت",
				"mm", "%d مىنۇت",
				"h", "بىر سائەت",
				"hh", "%d سائەت",
				"d", "بىر كۈن",
				"dd", "%d كۈن",
				"M", "بىر ئاي",
				"MM", "%d ئاي",
				"y", "بىر يىل",
				"yy", "%d يىل",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(-كۈنى|-ئاي|-ھەپتە)", ""),
			"ordinal", ugCnOrdinal,
			"preparse", ugCnPreparse,
			"postformat", ugCnPostformat,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

// ugCnMeridiemHour is the meridiemHour callback of ug-cn.
func ugCnMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "يېرىم كېچە", "سەھەر", "چۈشتىن بۇرۇن":
		return hour
	case "چۈشتىن كېيىن", "كەچ":
		return plus12(hour)
	}
	if hourNumber(hour) >= 11 {
		return hour
	}
	return plus12(hour)
}

// ugCnMeridiem is the meridiem callback of ug-cn.
func ugCnMeridiem(_ *moment.Locale, hour, minute float64, _ bool) string {
	hm := float64(hour*100) + minute
	switch {
	case hm < 600:
		return "يېرىم كېچە"
	case hm < 900:
		return "سەھەر"
	case hm < 1130:
		return "چۈشتىن بۇرۇن"
	case hm < 1230:
		return "چۈش"
	case hm < 1800:
		return "چۈشتىن كېيىن"
	}
	return "كەچ"
}

// ugCnOrdinal is the ordinal callback of ug-cn.
func ugCnOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "d", "D", "DDD":
		return str(number) + "-كۈنى"
	case "w", "W":
		return str(number) + "-ھەپتە"
	}
	return number
}

var (
	ugCnPreparse   = dvPreparse
	ugCnPostformat = dvPostformat
)
