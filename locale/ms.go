// moment.js locale configuration
// locale : Malay [ms]
// author : Weldan Jamili : https://github.com/weldan

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ms", func() {
		moment.DefineLocale("ms", js.Obj(
			"months", []string{"Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"},
			"monthsShort", []string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ogs", "Sep", "Okt", "Nov", "Dis"},
			"weekdays", []string{"Ahad", "Isnin", "Selasa", "Rabu", "Khamis", "Jumaat", "Sabtu"},
			"weekdaysShort", []string{"Ahd", "Isn", "Sel", "Rab", "Kha", "Jum", "Sab"},
			"weekdaysMin", []string{"Ah", "Is", "Sl", "Rb", "Km", "Jm", "Sb"},
			"longDateFormat", js.Obj(
				"LT", "HH.mm",
				"LTS", "HH.mm.ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY [pukul] HH.mm",
				"LLLL", "dddd, D MMMM YYYY [pukul] HH.mm",
			),
			"meridiemParse", jsre.MustCompile("pagi|tengahari|petang|malam", ""),
			"meridiemHour", msMeridiemHour,
			"meridiem", msMeridiem,
			"calendar", js.Obj(
				"sameDay", "[Hari ini pukul] LT",
				"nextDay", "[Esok pukul] LT",
				"nextWeek", "dddd [pukul] LT",
				"lastDay", "[Kelmarin pukul] LT",
				"lastWeek", "dddd [lepas pukul] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "dalam %s",
				"past", "%s yang lepas",
				"s", "beberapa saat",
				"ss", "%d saat",
				"m", "seminit",
				"mm", "%d minit",
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
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var (
	msMeridiemHour = msMyMeridiemHour
	msMeridiem     = msMyMeridiem
)
