// moment.js locale configuration
// locale : English (Australia) [en-au]
// author : Jared Morse : https://github.com/jarcoal

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("en-au", func() {
		moment.DefineLocale("en-au", js.Obj(
			"months", []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
			"monthsShort", []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
			"weekdays", []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			"weekdaysShort", []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
			"weekdaysMin", []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
			"longDateFormat", js.Obj(
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY h:mm A",
				"LLLL", "dddd, D MMMM YYYY h:mm A",
			),
			"calendar", js.Obj(
				"sameDay", "[Today at] LT",
				"nextDay", "[Tomorrow at] LT",
				"nextWeek", "dddd [at] LT",
				"lastDay", "[Yesterday at] LT",
				"lastWeek", "[Last] dddd [at] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "in %s",
				"past", "%s ago",
				"s", "a few seconds",
				"ss", "%d seconds",
				"m", "a minute",
				"mm", "%d minutes",
				"h", "an hour",
				"hh", "%d hours",
				"d", "a day",
				"dd", "%d days",
				"M", "a month",
				"MM", "%d months",
				"y", "a year",
				"yy", "%d years",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(st|nd|rd|th)", ""),
			"ordinal", enAuOrdinal,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 4.0,
			),
		))
	})
}

var enAuOrdinal = englishOrdinal

// englishOrdinal is the ordinal callback of the English locales: 1st, 2nd, 3rd, 4th, 11th to 13th.
func englishOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	n := js.ToNumber(number)
	b := js.Mod(n, 10)
	output := "th"
	if js.ToInt32(js.Mod(n, 100)/10) != 1 {
		switch b {
		case 1:
			output = "st"
		case 2:
			output = "nd"
		case 3:
			output = "rd"
		}
	}
	return str(number) + output
}
