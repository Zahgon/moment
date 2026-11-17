// moment.js locale configuration
// locale : Arabic [ar]
// author : Abdel Said: https://github.com/abdelsaid
// author : Ahmed Elkhatib
// author : forabi https://github.com/forabi

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ar", func() {
		moment.DefineLocale("ar", js.Obj(
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
			"isPM", arIsPM,
			"meridiem", arMeridiem,
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
				"s", arRelativeTimeSeconds,
				"ss", arRelativeTimeSecondsN,
				"m", arRelativeTimeMinute,
				"mm", arRelativeTimeMinutes,
				"h", arRelativeTimeHour,
				"hh", arRelativeTimeHours,
				"d", arRelativeTimeDay,
				"dd", arRelativeTimeDays,
				"M", arRelativeTimeMonth,
				"MM", arRelativeTimeMonths,
				"y", arRelativeTimeYear,
				"yy", arRelativeTimeYears,
			),
			"preparse", arPreparse,
			"postformat", arPostformat,
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}

var (
	arRelativeTimeSeconds  = arabicPluralize("s")
	arRelativeTimeSecondsN = arabicPluralize("s")
	arRelativeTimeMinute   = arabicPluralize("m")
	arRelativeTimeMinutes  = arabicPluralize("m")
	arRelativeTimeHour     = arabicPluralize("h")
	arRelativeTimeHours    = arabicPluralize("h")
	arRelativeTimeDay      = arabicPluralize("d")
	arRelativeTimeDays     = arabicPluralize("d")
	arRelativeTimeMonth    = arabicPluralize("M")
	arRelativeTimeMonths   = arabicPluralize("M")
	arRelativeTimeYear     = arabicPluralize("y")
	arRelativeTimeYears    = arabicPluralize("y")
)

var (
	arIsPM       = arabicIsPM
	arMeridiem   = arabicMeridiem
	arPreparse   = arabicDigitsPreparse
	arPostformat = arabicDigitsPostformat
)
