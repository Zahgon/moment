// moment.js locale configuration
// locale : Pashto [ps]
// author : Hanif Hefaz : https://github.com/hanifhefaz

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ps", func() {
		moment.DefineLocale("ps", js.Obj(
			"months", js.Obj(
				"format", []string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
				"standalone", []string{"جنوري", "فېبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
				"isFormat", jsre.MustCompile("D[oD]?(\\[[^\\[\\]]*\\]|\\s|,)+MMMM?", ""),
			),
			"monthsShort", js.Obj(
				"format", []string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سېپتمبر", "اکتوبر", "نومبر", "دسمبر"},
				"standalone", []string{"جنوري", "فبروري", "مارچ", "اپریل", "مۍ", "جون", "جولای", "اګست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
			),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^جنوري$", "i"),
				jsre.MustCompile("^(فبروري|فېبروري)$", "i"),
				jsre.MustCompile("^مارچ$", "i"),
				jsre.MustCompile("^اپریل$", "i"),
				jsre.MustCompile("^مۍ$", "i"),
				jsre.MustCompile("^جون$", "i"),
				jsre.MustCompile("^جولای$", "i"),
				jsre.MustCompile("^اګست$", "i"),
				jsre.MustCompile("^(سپتمبر|سېپتمبر)$", "i"),
				jsre.MustCompile("^اکتوبر$", "i"),
				jsre.MustCompile("^نومبر$", "i"),
				jsre.MustCompile("^دسمبر$", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^جنوري$", "i"),
				jsre.MustCompile("^(فبروري|فېبروري)$", "i"),
				jsre.MustCompile("^مارچ$", "i"),
				jsre.MustCompile("^اپریل$", "i"),
				jsre.MustCompile("^مۍ$", "i"),
				jsre.MustCompile("^جون$", "i"),
				jsre.MustCompile("^جولای$", "i"),
				jsre.MustCompile("^اګست$", "i"),
				jsre.MustCompile("^(سپتمبر|سېپتمبر)$", "i"),
				jsre.MustCompile("^اکتوبر$", "i"),
				jsre.MustCompile("^نومبر$", "i"),
				jsre.MustCompile("^دسمبر$", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^جنوري$", "i"),
				jsre.MustCompile("^(فبروري|فېبروري)$", "i"),
				jsre.MustCompile("^مارچ$", "i"),
				jsre.MustCompile("^اپریل$", "i"),
				jsre.MustCompile("^مۍ$", "i"),
				jsre.MustCompile("^جون$", "i"),
				jsre.MustCompile("^جولای$", "i"),
				jsre.MustCompile("^اګست$", "i"),
				jsre.MustCompile("^(سپتمبر|سېپتمبر)$", "i"),
				jsre.MustCompile("^اکتوبر$", "i"),
				jsre.MustCompile("^نومبر$", "i"),
				jsre.MustCompile("^دسمبر$", "i"),
			},
			"monthsRegex", jsre.MustCompile("^(جنوري|فبروري|فېبروري|مارچ|اپریل|مۍ|جون|جولای|اګست|سپتمبر|سېپتمبر|اکتوبر|نومبر|دسمبر)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(جنوري|فبروري|فېبروري|مارچ|اپریل|مۍ|جون|جولای|اګست|سپتمبر|سېپتمبر|اکتوبر|نومبر|دسمبر)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(جنوري|فبروري|فېبروري|مارچ|اپریل|مۍ|جون|جولای|اګست|سپتمبر|سېپتمبر|اکتوبر|نومبر|دسمبر)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(جنوري|فبروري|فېبروري|مارچ|اپریل|مۍ|جون|جولای|اګست|سپتمبر|سېپتمبر|اکتوبر|نومبر|دسمبر)", "i"),
			"weekdays", []string{"یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنج شنبه", "جمعه", "شنبه"},
			"weekdaysShort", []string{"یکشنبه", "دوشنبه", "سه شنبه", "چهارشنبه", "پنج شنبه", "جمعه", "شنبه"},
			"weekdaysMin", []string{"ی", "د", "س", "چ", "پ", "ج", "ش"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "YYYY/M/D",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY H:mm",
				"LLLL", "dddd D, MMMM, YYYY [په] H:mm",
			),
			"meridiemParse", jsre.MustCompile("غ\\.م|غ\\.و\\.", ""),
			"isPM", psIsPM,
			"meridiem", psMeridiem,
			"calendar", js.Obj(
				"sameDay", "[نن په] LT",
				"nextDay", "[سبا په] LT",
				"nextWeek", "dddd [په] LT",
				"lastDay", "[پرون په] LT",
				"lastWeek", "dddd [په] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "په %s کې",
				"past", "%s مخکې",
				"s", psRelativeTimeSeconds,
				"ss", psRelativeTimeSeconds,
				"m", psRelativeTimeSeconds,
				"mm", psRelativeTimeSeconds,
				"h", psRelativeTimeSeconds,
				"hh", psRelativeTimeSeconds,
				"d", psRelativeTimeSeconds,
				"dd", psRelativeTimeSeconds,
				"M", psRelativeTimeSeconds,
				"MM", psRelativeTimeSeconds,
				"y", psRelativeTimeSeconds,
				"yy", psRelativeTimeSeconds,
			),
			"preparse", psPreparse,
			"postformat", psPostformat,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}", ""),
			"ordinal", "%d",
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}

var (
	psFuture = map[string]string{"s": "څو ثانيو", "ss": "%d ثانيو", "m": "يوه دقيقه", "mm": "%d دقيقو", "h": "يو ساعت",
		"hh": "%d ساعتو", "d": "يوه ورځ", "dd": "%d ورځو", "M": "يوه مياشت", "MM": "%d مياشتو", "y": "يو کال",
		"yy": "%d کالونو"}
	psOther = map[string]string{"s": "څو ثانيې", "ss": "%d ثانيې", "m": "يوه دقيقه", "mm": "%d دقيقې", "h": "يو ساعت",
		"hh": "%d ساعتونه", "d": "يوه ورځ", "dd": "%d ورځې", "M": "يوه مياشت", "MM": "%d مياشتې", "y": "يو کال",
		"yy": "%d کاله"}
	psPreparse   = faPreparse
	psPostformat = faPostformat
)

// psIsPM is the isPM callback of ps.
func psIsPM(_ *moment.Locale, input string) bool { return input == "غ.و." }

// psMeridiem is the meridiem callback of ps.
func psMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "غ.م"
	}
	return "غ.و."
}

// psRelativeTimeSeconds is relativeTime of ps: future forms only for an unsuffixed future.
func psRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	format := psOther
	if isFuture && !withoutSuffix {
		format = psFuture
	}
	return js.ReplaceFirst(format[key], "%d", num(number))
}
