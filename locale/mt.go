// moment.js locale configuration
// locale : Maltese (Malta) [mt]
// author : Alessandro Maruccia : https://github.com/alesma

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("mt", func() {
		moment.DefineLocale("mt", js.Obj(
			"months", []string{"Jannar", "Frar", "Marzu", "April", "Mejju", "Ġunju", "Lulju", "Awwissu", "Settembru", "Ottubru", "Novembru", "Diċembru"},
			"monthsShort", []string{"Jan", "Fra", "Mar", "Apr", "Mej", "Ġun", "Lul", "Aww", "Set", "Ott", "Nov", "Diċ"},
			"weekdays", []string{"Il-Ħadd", "It-Tnejn", "It-Tlieta", "L-Erbgħa", "Il-Ħamis", "Il-Ġimgħa", "Is-Sibt"},
			"weekdaysShort", []string{"Ħad", "Tne", "Tli", "Erb", "Ħam", "Ġim", "Sib"},
			"weekdaysMin", []string{"Ħa", "Tn", "Tl", "Er", "Ħa", "Ġi", "Si"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Illum fil-]LT",
				"nextDay", "[Għada fil-]LT",
				"nextWeek", "dddd [fil-]LT",
				"lastDay", "[Il-bieraħ fil-]LT",
				"lastWeek", "dddd [li għadda] [fil-]LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "f’ %s",
				"past", "%s ilu",
				"s", "ftit sekondi",
				"ss", "%d sekondi",
				"m", "minuta",
				"mm", "%d minuti",
				"h", "siegħa",
				"hh", "%d siegħat",
				"d", "ġurnata",
				"dd", "%d ġranet",
				"M", "xahar",
				"MM", "%d xhur",
				"y", "sena",
				"yy", "%d sni",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}º", ""),
			"ordinal", "%dº",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}
