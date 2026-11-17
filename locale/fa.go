// moment.js locale configuration
// locale : Persian [fa]
// author : Ebrahim Byagowi : https://github.com/ebraminio

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fa", func() {
		moment.DefineLocale("fa", js.Obj(
			"months", []string{"ژانویه", "فوریه", "مارس", "آوریل", "مه", "ژوئن", "ژوئیه", "اوت", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"},
			"monthsShort", []string{"ژانویه", "فوریه", "مارس", "آوریل", "مه", "ژوئن", "ژوئیه", "اوت", "سپتامبر", "اکتبر", "نوامبر", "دسامبر"},
			"weekdays", []string{"یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه"},
			"weekdaysShort", []string{"یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه"},
			"weekdaysMin", []string{"ی", "د", "س", "چ", "پ", "ج", "ش"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("قبل از ظهر|بعد از ظهر", ""),
			"isPM", faIsPM,
			"meridiem", faMeridiem,
			"calendar", js.Obj(
				"sameDay", "[امروز ساعت] LT",
				"nextDay", "[فردا ساعت] LT",
				"nextWeek", "dddd [ساعت] LT",
				"lastDay", "[دیروز ساعت] LT",
				"lastWeek", "dddd [پیش] [ساعت] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "در %s",
				"past", "%s پیش",
				"s", "چند ثانیه",
				"ss", "%d ثانیه",
				"m", "یک دقیقه",
				"mm", "%d دقیقه",
				"h", "یک ساعت",
				"hh", "%d ساعت",
				"d", "یک روز",
				"dd", "%d روز",
				"M", "یک ماه",
				"MM", "%d ماه",
				"y", "یک سال",
				"yy", "%d سال",
			),
			"preparse", faPreparse,
			"postformat", faPostformat,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}م", ""),
			"ordinal", "%dم",
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}

var (
	faSymbols, faNumbers = digitTables("۰۱۲۳۴۵۶۷۸۹")
	faDigit              = jsre.MustCompile(`[۰-۹]`, "g")
	faPM                 = jsre.MustCompile(`بعد از ظهر`, "")
)

// faIsPM is the isPM callback of fa.
func faIsPM(_ *moment.Locale, input string) bool { return faPM.Test(input) }

// faMeridiem is the meridiem callback of fa.
func faMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "قبل از ظهر"
	}
	return "بعد از ظهر"
}

// faPreparse reads Persian digits and the Arabic comma.
func faPreparse(_ *moment.Locale, s string) string {
	return arabicComma.Replace(mapMatches(faDigit, s, faNumbers), ",")
}

// faPostformat writes Persian digits and the Arabic comma.
func faPostformat(_ *moment.Locale, s string) string {
	return comma.Replace(mapMatches(asciiDigit, s, faSymbols), "،")
}
