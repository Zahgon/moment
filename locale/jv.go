// moment.js locale configuration
// locale : Javanese [jv]
// author : Rony Lantip : https://github.com/lantip
// reference: http://jv.wikipedia.org/wiki/Basa_Jawa

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("jv", func() {
		moment.DefineLocale("jv", js.Obj(
			"months", []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "Nopember", "Desember"},
			"monthsShort", []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Ags", "Sep", "Okt", "Nop", "Des"},
			"weekdays", []string{"Minggu", "Senen", "Seloso", "Rebu", "Kemis", "Jemuwah", "Septu"},
			"weekdaysShort", []string{"Min", "Sen", "Sel", "Reb", "Kem", "Jem", "Sep"},
			"weekdaysMin", []string{"Mg", "Sn", "Sl", "Rb", "Km", "Jm", "Sp"},
			"longDateFormat", js.Obj(
				"LT", "HH.mm",
				"LTS", "HH.mm.ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY [pukul] HH.mm",
				"LLLL", "dddd, D MMMM YYYY [pukul] HH.mm",
			),
			"meridiemParse", jsre.MustCompile("enjing|siyang|sonten|ndalu", ""),
			"meridiemHour", jvMeridiemHour,
			"meridiem", jvMeridiem,
			"calendar", js.Obj(
				"sameDay", "[Dinten puniko pukul] LT",
				"nextDay", "[Mbenjang pukul] LT",
				"nextWeek", "dddd [pukul] LT",
				"lastDay", "[Kala wingi pukul] LT",
				"lastWeek", "dddd [kepengker pukul] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "wonten ing %s",
				"past", "%s ingkang kepengker",
				"s", "sawetawis detik",
				"ss", "%d detik",
				"m", "setunggal menit",
				"mm", "%d menit",
				"h", "setunggal jam",
				"hh", "%d jam",
				"d", "sedinten",
				"dd", "%d dinten",
				"M", "sewulan",
				"MM", "%d wulan",
				"y", "setaun",
				"yy", "%d taun",
			),
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

// jvMeridiemHour is the meridiemHour callback of jv.
func jvMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return malayMeridiemHour(hour, meridiem, "enjing", "siyang", "sonten", "ndalu")
}

// jvMeridiem is the meridiem callback of jv.
func jvMeridiem(_ *moment.Locale, hours, _ float64, _ bool) string {
	switch {
	case hours < 11:
		return "enjing"
	case hours < 15:
		return "siyang"
	case hours < 19:
		return "sonten"
	}
	return "ndalu"
}
