// moment.js locale configuration
// locale : English (Canada) [en-ca]
// author : Jonathan Abourbih : https://github.com/jonbca

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("en-ca", func() {
		moment.DefineLocale("en-ca", js.Obj(
			"months", []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
			"monthsShort", []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
			"weekdays", []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
			"weekdaysShort", []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
			"weekdaysMin", []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"},
			"longDateFormat", js.Obj(
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "YYYY-MM-DD",
				"LL", "MMMM D, YYYY",
				"LLL", "MMMM D, YYYY h:mm A",
				"LLLL", "dddd, MMMM D, YYYY h:mm A",
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
			"ordinal", enCaOrdinal,
		))
	})
}

var enCaOrdinal = englishOrdinal
