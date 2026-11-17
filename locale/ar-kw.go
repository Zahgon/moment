// moment.js locale configuration
// locale : Arabic (Kuwait) [ar-kw]
// author : Nusret Parlak: https://github.com/nusretparlak

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("ar-kw", func() {
		moment.DefineLocale("ar-kw", js.Obj(
			"months", []string{"يناير", "فبراير", "مارس", "أبريل", "ماي", "يونيو", "يوليوز", "غشت", "شتنبر", "أكتوبر", "نونبر", "دجنبر"},
			"monthsShort", []string{"يناير", "فبراير", "مارس", "أبريل", "ماي", "يونيو", "يوليوز", "غشت", "شتنبر", "أكتوبر", "نونبر", "دجنبر"},
			"weekdays", []string{"الأحد", "الإتنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
			"weekdaysShort", []string{"احد", "اتنين", "ثلاثاء", "اربعاء", "خميس", "جمعة", "سبت"},
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
			"week", js.Obj(
				"dow", 0.0,
				"doy", 12.0,
			),
		))
	})
}
