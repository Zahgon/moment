// moment.js locale configuration
// locale : Kurdish [ku]
// author : Shahram Mebashar : https://github.com/ShahramMebashar

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ku", func() {
		moment.DefineLocale("ku", js.Obj(
			"months", []string{"کانونی دووەم", "شوبات", "ئازار", "نیسان", "ئایار", "حوزەیران", "تەمموز", "ئاب", "ئەیلوول", "تشرینی یەكەم", "تشرینی دووەم", "كانونی یەکەم"},
			"monthsShort", []string{"کانونی دووەم", "شوبات", "ئازار", "نیسان", "ئایار", "حوزەیران", "تەمموز", "ئاب", "ئەیلوول", "تشرینی یەكەم", "تشرینی دووەم", "كانونی یەکەم"},
			"weekdays", []string{"یه‌كشه‌ممه‌", "دووشه‌ممه‌", "سێشه‌ممه‌", "چوارشه‌ممه‌", "پێنجشه‌ممه‌", "هه‌ینی", "شه‌ممه‌"},
			"weekdaysShort", []string{"یه‌كشه‌م", "دووشه‌م", "سێشه‌م", "چوارشه‌م", "پێنجشه‌م", "هه‌ینی", "شه‌ممه‌"},
			"weekdaysMin", []string{"ی", "د", "س", "چ", "پ", "ه", "ش"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("ئێواره‌|به‌یانی", ""),
			"isPM", kuIsPM,
			"meridiem", kuMeridiem,
			"calendar", js.Obj(
				"sameDay", "[ئه‌مرۆ كاتژمێر] LT",
				"nextDay", "[به‌یانی كاتژمێر] LT",
				"nextWeek", "dddd [كاتژمێر] LT",
				"lastDay", "[دوێنێ كاتژمێر] LT",
				"lastWeek", "dddd [كاتژمێر] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "له‌ %s",
				"past", "%s",
				"s", "چه‌ند چركه‌یه‌ك",
				"ss", "چركه‌ %d",
				"m", "یه‌ك خوله‌ك",
				"mm", "%d خوله‌ك",
				"h", "یه‌ك كاتژمێر",
				"hh", "%d كاتژمێر",
				"d", "یه‌ك ڕۆژ",
				"dd", "%d ڕۆژ",
				"M", "یه‌ك مانگ",
				"MM", "%d مانگ",
				"y", "یه‌ك ساڵ",
				"yy", "%d ساڵ",
			),
			"preparse", kuPreparse,
			"postformat", kuPostformat,
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}

var kuPM = jsre.MustCompile("ئێواره\u200c", "")

// kuIsPM is the isPM callback of ku.
func kuIsPM(_ *moment.Locale, input string) bool { return kuPM.Test(input) }

// kuMeridiem is the meridiem callback of ku.
func kuMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "به\u200cیانی"
	}
	return "ئێواره\u200c"
}

var (
	kuPreparse   = arabicDigitsPreparse
	kuPostformat = arabicDigitsPostformat
)
