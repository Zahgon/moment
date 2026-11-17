// moment.js locale configuration
// locale : Tagalog (Philippines) [tl-ph]
// author : Dan Hagman : https://github.com/hagmandan

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("tl-ph", func() {
		moment.DefineLocale("tl-ph", js.Obj(
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
			"ordinal", tlPhOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// tlPhOrdinal is the ordinal callback of tl-ph: the argument itself, unchanged.
func tlPhOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value { return number }
