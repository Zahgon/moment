// moment.js locale configuration
// locale : Filipino [fil]
// author : Dan Hagman : https://github.com/hagmandan
// author : Matthew Co : https://github.com/matthewdeeco

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fil", func() {
		moment.DefineLocale("fil", js.Obj(
			"months", []string{"Enero", "Pebrero", "Marso", "Abril", "Mayo", "Hunyo", "Hulyo", "Agosto", "Setyembre", "Oktubre", "Nobyembre", "Disyembre"},
			"monthsShort", []string{"Ene", "Peb", "Mar", "Abr", "May", "Hun", "Hul", "Ago", "Set", "Okt", "Nob", "Dis"},
			"weekdays", []string{"Linggo", "Lunes", "Martes", "Miyerkules", "Huwebes", "Biyernes", "Sabado"},
			"weekdaysShort", []string{"Lin", "Lun", "Mar", "Miy", "Huw", "Biy", "Sab"},
			"weekdaysMin", []string{"Li", "Lu", "Ma", "Mi", "Hu", "Bi", "Sab"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "MM/D/YYYY",
				"LL", "MMMM D, YYYY",
				"LLL", "MMMM D, YYYY HH:mm",
				"LLLL", "dddd, MMMM DD, YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "LT [ngayong araw]",
				"nextDay", "[Bukas ng] LT",
				"nextWeek", "LT [sa susunod na] dddd",
				"lastDay", "LT [kahapon]",
				"lastWeek", "LT [noong nakaraang] dddd",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "sa loob ng %s",
				"past", "%s ang nakalipas",
				"s", "ilang segundo",
				"ss", "%d segundo",
				"m", "isang minuto",
				"mm", "%d minuto",
				"h", "isang oras",
				"hh", "%d oras",
				"d", "isang araw",
				"dd", "%d araw",
				"M", "isang buwan",
				"MM", "%d buwan",
				"y", "isang taon",
				"yy", "%d taon",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}", ""),
			"ordinal", filOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// filOrdinal is the ordinal callback of fil: the argument itself, unchanged.
func filOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value { return number }
