// moment.js locale configuration
// locale : Nepalese [ne]
// author : suvash : https://github.com/suvash

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ne", func() {
		moment.DefineLocale("ne", js.Obj(
			"months", []string{"जनवरी", "फेब्रुवरी", "मार्च", "अप्रिल", "मई", "जुन", "जुलाई", "अगष्ट", "सेप्टेम्बर", "अक्टोबर", "नोभेम्बर", "डिसेम्बर"},
			"monthsShort", []string{"जन.", "फेब्रु.", "मार्च", "अप्रि.", "मई", "जुन", "जुलाई.", "अग.", "सेप्ट.", "अक्टो.", "नोभे.", "डिसे."},
			"monthsParseExact", true,
			"weekdays", []string{"आइतबार", "सोमबार", "मङ्गलबार", "बुधबार", "बिहिबार", "शुक्रबार", "शनिबार"},
			"weekdaysShort", []string{"आइत.", "सोम.", "मङ्गल.", "बुध.", "बिहि.", "शुक्र.", "शनि."},
			"weekdaysMin", []string{"आ.", "सो.", "मं.", "बु.", "बि.", "शु.", "श."},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "Aको h:mm बजे",
				"LTS", "Aको h:mm:ss बजे",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, Aको h:mm बजे",
				"LLLL", "dddd, D MMMM YYYY, Aको h:mm बजे",
			),
			"preparse", nePreparse,
			"postformat", nePostformat,
			"meridiemParse", jsre.MustCompile("राति|बिहान|दिउँसो|साँझ", ""),
			"meridiemHour", neMeridiemHour,
			"meridiem", neMeridiem,
			"calendar", js.Obj(
				"sameDay", "[आज] LT",
				"nextDay", "[भोलि] LT",
				"nextWeek", "[आउँदो] dddd[,] LT",
				"lastDay", "[हिजो] LT",
				"lastWeek", "[गएको] dddd[,] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%sमा",
				"past", "%s अगाडि",
				"s", "केही क्षण",
				"ss", "%d सेकेण्ड",
				"m", "एक मिनेट",
				"mm", "%d मिनेट",
				"h", "एक घण्टा",
				"hh", "%d घण्टा",
				"d", "एक दिन",
				"dd", "%d दिन",
				"M", "एक महिना",
				"MM", "%d महिना",
				"y", "एक बर्ष",
				"yy", "%d बर्ष",
			),
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	nePreparse   = hiPreparse
	nePostformat = hiPostformat
)

// neMeridiemHour is the meridiemHour callback of ne.
func neMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return dayPartHour(hour, meridiem, "राति", "बिहान", "दिउँसो", "साँझ", 10)
}

// neMeridiem is the meridiem callback of ne.
func neMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 3:
		return "राति"
	case hour < 12:
		return "बिहान"
	case hour < 16:
		return "दिउँसो"
	case hour < 20:
		return "साँझ"
	}
	return "राति"
}
