// moment.js locale configuration
// locale : English (New Zealand) [en-nz]
// author : Luke McGregor : https://github.com/lukemcgregor

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("en-nz", func() {
		moment.DefineLocale("en-nz", js.Obj(
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
			"ordinal", enNzOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var enNzOrdinal = englishOrdinal
