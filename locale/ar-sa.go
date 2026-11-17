// moment.js locale configuration
// locale : Arabic (Saudi Arabia) [ar-sa]
// author : Suhail Alkowaileet : https://github.com/xsoh

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ar-sa", func() {
		moment.DefineLocale("ar-sa", js.Obj(
			"months", []string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
			"monthsShort", []string{"يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
			"weekdays", []string{"الأحد", "الإثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
			"weekdaysShort", []string{"أحد", "إثنين", "ثلاثاء", "أربعاء", "خميس", "جمعة", "سبت"},
			"weekdaysMin", []string{"ح", "ن", "ث", "ر", "خ", "ج", "س"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("ص|م", ""),
			"isPM", arSaIsPM,
			"meridiem", arSaMeridiem,
			"calendar", js.Obj(
				"sameDay", "[اليوم على الساعة] LT",
				"nextDay", "[غدا على الساعة] LT",
				"nextWeek", "dddd [على الساعة] LT",
				"lastDay", "[أمس على الساعة] LT",
				"lastWeek", "dddd [على الساعة] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "في %s",
				"past", "منذ %s",
				"s", "ثوان",
				"ss", "%d ثانية",
				"m", "دقيقة",
				"mm", "%d دقائق",
				"h", "ساعة",
				"hh", "%d ساعات",
				"d", "يوم",
				"dd", "%d أيام",
				"M", "شهر",
				"MM", "%d أشهر",
				"y", "سنة",
				"yy", "%d سنوات",
			),
			"preparse", arSaPreparse,
			"postformat", arSaPostformat,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	arSaIsPM       = arabicIsPM
	arSaMeridiem   = arabicMeridiem
	arSaPreparse   = arabicDigitsPreparse
	arSaPostformat = arabicDigitsPostformat
)
