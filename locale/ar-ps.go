// moment.js locale configuration
// locale : Arabic (Palestine) [ar-ps]
// author : Majd Al-Shihabi : https://github.com/majdal

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ar-ps", func() {
		moment.DefineLocale("ar-ps", js.Obj(
			"months", []string{"كانون الثاني", "شباط", "آذار", "نيسان", "أيّار", "حزيران", "تمّوز", "آب", "أيلول", "تشري الأوّل", "تشرين الثاني", "كانون الأوّل"},
			"monthsShort", []string{"ك٢", "شباط", "آذار", "نيسان", "أيّار", "حزيران", "تمّوز", "آب", "أيلول", "ت١", "ت٢", "ك١"},
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
			"isPM", arPsIsPM,
			"meridiem", arPsMeridiem,
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
			"preparse", arPsPreparse,
			"postformat", arPsPostformat,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	arPsIsPM       = arabicIsPM
	arPsMeridiem   = arabicMeridiem
	arPsPostformat = arabicDigitsPostformat
	arPsLateDigits = jsre.MustCompile(`[٣٤٥٦٧٨٩٠]`, "g")
	// arPsOneTwo matches 1 and 2 not followed (in the reversed text: not preceded) by ت or ك.
	arPsOneTwo = jsre.MustCompile(`[١٢](?![\u062a\u0643])`, "g")
)

// arPsPreparse reads Arabic-Indic digits except 1 and 2 directly before ت or ك, and the Arabic comma.
func arPsPreparse(_ *moment.Locale, s string) string {
	s = mapMatches(arPsLateDigits, s, arabicNumbers)
	s = reverseUnits(mapMatches(arPsOneTwo, reverseUnits(s), arabicNumbers))
	return arabicComma.Replace(s, ",")
}
