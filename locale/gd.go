// moment.js locale configuration
// locale : Scottish Gaelic [gd]
// author : Jon Ashdown : https://github.com/jonashdown

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("gd", func() {
		moment.DefineLocale("gd", js.Obj(
			"months", []string{"Am Faoilleach", "An Gearran", "Am Màrt", "An Giblean", "An Cèitean", "An t-Ògmhios", "An t-Iuchar", "An Lùnastal", "An t-Sultain", "An Dàmhair", "An t-Samhain", "An Dùbhlachd"},
			"monthsShort", []string{"Faoi", "Gear", "Màrt", "Gibl", "Cèit", "Ògmh", "Iuch", "Lùn", "Sult", "Dàmh", "Samh", "Dùbh"},
			"monthsParseExact", true,
			"weekdays", []string{"Didòmhnaich", "Diluain", "Dimàirt", "Diciadain", "Diardaoin", "Dihaoine", "Disathairne"},
			"weekdaysShort", []string{"Did", "Dil", "Dim", "Dic", "Dia", "Dih", "Dis"},
			"weekdaysMin", []string{"Dò", "Lu", "Mà", "Ci", "Ar", "Ha", "Sa"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[An-diugh aig] LT",
				"nextDay", "[A-màireach aig] LT",
				"nextWeek", "dddd [aig] LT",
				"lastDay", "[An-dè aig] LT",
				"lastWeek", "dddd [seo chaidh] [aig] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "ann an %s",
				"past", "bho chionn %s",
				"s", "beagan diogan",
				"ss", "%d diogan",
				"m", "mionaid",
				"mm", "%d mionaidean",
				"h", "uair",
				"hh", "%d uairean",
				"d", "latha",
				"dd", "%d latha",
				"M", "mìos",
				"MM", "%d mìosan",
				"y", "bliadhna",
				"yy", "%d bliadhna",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(d|na|mh)", ""),
			"ordinal", gdOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var gdOrdinal = gaOrdinal
