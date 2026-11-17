// moment.js locale configuration
// locale : Arabic (Libya) [ar-ly]
// author : Ali Hmer: https://github.com/kikoanis

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ar-ly", func() {
		moment.DefineLocale("ar-ly", js.Obj(
			"months", []string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
			"monthsShort", []string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
			"weekdays", []string{"الأحد", "الإثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
			"weekdaysShort", []string{"أحد", "إثنين", "ثلاثاء", "أربعاء", "خميس", "جمعة", "سبت"},
			"weekdaysMin", []string{"ح", "ن", "ث", "ر", "خ", "ج", "س"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "D/‏M/‏YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("ص|م", ""),
			"isPM", arLyIsPM,
			"meridiem", arLyMeridiem,
			"calendar", js.Obj(
				"sameDay", "[اليوم عند الساعة] LT",
				"nextDay", "[غدًا عند الساعة] LT",
				"nextWeek", "dddd [عند الساعة] LT",
				"lastDay", "[أمس عند الساعة] LT",
				"lastWeek", "dddd [عند الساعة] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "بعد %s",
				"past", "منذ %s",
				"s", arLyRelativeTimeSeconds,
				"ss", arLyRelativeTimeSecondsN,
				"m", arLyRelativeTimeMinute,
				"mm", arLyRelativeTimeMinutes,
				"h", arLyRelativeTimeHour,
				"hh", arLyRelativeTimeHours,
				"d", arLyRelativeTimeDay,
				"dd", arLyRelativeTimeDays,
				"M", arLyRelativeTimeMonth,
				"MM", arLyRelativeTimeMonths,
				"y", arLyRelativeTimeYear,
				"yy", arLyRelativeTimeYears,
			),
			"preparse", arLyPreparse,
			"postformat", arLyPostformat,
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}

var (
	arLyRelativeTimeSeconds  = arabicPluralize("s")
	arLyRelativeTimeSecondsN = arabicPluralize("s")
	arLyRelativeTimeMinute   = arabicPluralize("m")
	arLyRelativeTimeMinutes  = arabicPluralize("m")
	arLyRelativeTimeHour     = arabicPluralize("h")
	arLyRelativeTimeHours    = arabicPluralize("h")
	arLyRelativeTimeDay      = arabicPluralize("d")
	arLyRelativeTimeDays     = arabicPluralize("d")
	arLyRelativeTimeMonth    = arabicPluralize("M")
	arLyRelativeTimeMonths   = arabicPluralize("M")
	arLyRelativeTimeYear     = arabicPluralize("y")
	arLyRelativeTimeYears    = arabicPluralize("y")
)

var (
	arLyIsPM     = arabicIsPM
	arLyMeridiem = arabicMeridiem
)

// arLyPreparse reads the Arabic comma.
func arLyPreparse(_ *moment.Locale, s string) string { return arabicComma.Replace(s, ",") }

// arLyPostformat keeps ASCII digits (its digit table maps each digit to itself) and writes the Arabic comma.
func arLyPostformat(_ *moment.Locale, s string) string {
	return comma.Replace(mapMatches(asciiDigit, s, map[string]string{"0": "0", "1": "1", "2": "2", "3": "3",
		"4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9"}), "،")
}
