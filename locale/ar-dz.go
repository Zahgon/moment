// moment.js locale configuration
// locale : Arabic (Algeria) [ar-dz]
// author : Amine Roukh: https://github.com/Amine27
// author : Abdel Said: https://github.com/abdelsaid
// author : Ahmed Elkhatib
// author : forabi https://github.com/forabi
// author : Noureddine LOUAHEDJ : https://github.com/noureddinem

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ar-dz", func() {
		moment.DefineLocale("ar-dz", js.Obj(
			"months", []string{"جانفي", "فيفري", "مارس", "أفريل", "ماي", "جوان", "جويلية", "أوت", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
			"monthsShort", []string{"جانفي", "فيفري", "مارس", "أفريل", "ماي", "جوان", "جويلية", "أوت", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
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
			"isPM", arDzIsPM,
			"meridiem", arDzMeridiem,
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
				"s", arDzRelativeTimeSeconds,
				"ss", arDzRelativeTimeSecondsN,
				"m", arDzRelativeTimeMinute,
				"mm", arDzRelativeTimeMinutes,
				"h", arDzRelativeTimeHour,
				"hh", arDzRelativeTimeHours,
				"d", arDzRelativeTimeDay,
				"dd", arDzRelativeTimeDays,
				"M", arDzRelativeTimeMonth,
				"MM", arDzRelativeTimeMonths,
				"y", arDzRelativeTimeYear,
				"yy", arDzRelativeTimeYears,
			),
			"postformat", arDzPostformat,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 4.0,
			),
		))
	})
}

var (
	arDzRelativeTimeSeconds  = arabicPluralize("s")
	arDzRelativeTimeSecondsN = arabicPluralize("s")
	arDzRelativeTimeMinute   = arabicPluralize("m")
	arDzRelativeTimeMinutes  = arabicPluralize("m")
	arDzRelativeTimeHour     = arabicPluralize("h")
	arDzRelativeTimeHours    = arabicPluralize("h")
	arDzRelativeTimeDay      = arabicPluralize("d")
	arDzRelativeTimeDays     = arabicPluralize("d")
	arDzRelativeTimeMonth    = arabicPluralize("M")
	arDzRelativeTimeMonths   = arabicPluralize("M")
	arDzRelativeTimeYear     = arabicPluralize("y")
	arDzRelativeTimeYears    = arabicPluralize("y")
)

var (
	arDzIsPM     = arabicIsPM
	arDzMeridiem = arabicMeridiem
)

// arDzPostformat writes the Arabic comma.
func arDzPostformat(_ *moment.Locale, s string) string { return comma.Replace(s, "،") }
