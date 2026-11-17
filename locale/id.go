// moment.js locale configuration
// locale : Indonesian [id]
// author : Mohammad Satrio Utomo : https://github.com/tyok
// reference: http://id.wikisource.org/wiki/Pedoman_Umum_Ejaan_Bahasa_Indonesia_yang_Disempurnakan

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

// The month parse lists of id set only August (index 7); the other months keep the patterns built from their names,
// as the sparse arrays of the original do.
func init() {
	moment.RegisterLocale("id", func() {
		moment.DefineLocale("id", js.Obj(
			"months", []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
			"monthsShort", []string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"},
			"monthsParse", []any{js.Undef, js.Undef, js.Undef, js.Undef, js.Undef, js.Undef, js.Undef,
				jsre.MustCompile("^agustus|^agu|^agt", "i")},
			"longMonthsParse", []any{js.Undef, js.Undef, js.Undef, js.Undef, js.Undef, js.Undef, js.Undef,
				jsre.MustCompile("^agustus$", "i")},
			"shortMonthsParse", []any{js.Undef, js.Undef, js.Undef, js.Undef, js.Undef, js.Undef, js.Undef,
				jsre.MustCompile("^agu$|^agt$", "i")},
			"weekdays", []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"},
			"weekdaysShort", []string{"Min", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
			"weekdaysMin", []string{"Mg", "Sn", "Sl", "Rb", "Km", "Jm", "Sb"},
			"longDateFormat", js.Obj(
				"LT", "HH.mm",
				"LTS", "HH.mm.ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY [pukul] HH.mm",
				"LLLL", "dddd, D MMMM YYYY [pukul] HH.mm",
			),
			"meridiemParse", jsre.MustCompile("pagi|siang|sore|malam", ""),
			"meridiemHour", idMeridiemHour,
			"meridiem", idMeridiem,
			"calendar", js.Obj(
				"sameDay", "[Hari ini pukul] LT",
				"nextDay", "[Besok pukul] LT",
				"nextWeek", "dddd [pukul] LT",
				"lastDay", "[Kemarin pukul] LT",
				"lastWeek", "dddd [lalu pukul] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "dalam %s",
				"past", "%s yang lalu",
				"s", "beberapa detik",
				"ss", "%d detik",
				"m", "semenit",
				"mm", "%d menit",
				"h", "sejam",
				"hh", "%d jam",
				"d", "sehari",
				"dd", "%d hari",
				"M", "sebulan",
				"MM", "%d bulan",
				"y", "setahun",
				"yy", "%d tahun",
			),
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

// idMeridiemHour is the meridiemHour callback of id.
func idMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return malayMeridiemHour(hour, meridiem, "pagi", "siang", "sore", "malam")
}

// malayMeridiemHour is the meridiemHour of id, ms, ms-my and jv for their morning, midday, afternoon and night
// words.
func malayMeridiemHour(hour any, meridiem, morning, midday, afternoon, night string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case morning:
		return hour
	case midday:
		if hourNumber(hour) >= 11 {
			return hour
		}
		return plus12(hour)
	case afternoon, night:
		return plus12(hour)
	}
	return js.Undef
}

// idMeridiem is the meridiem callback of id.
func idMeridiem(_ *moment.Locale, hours, _ float64, _ bool) string {
	switch {
	case hours < 11:
		return "pagi"
	case hours < 15:
		return "siang"
	case hours < 19:
		return "sore"
	}
	return "malam"
}
