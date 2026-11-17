// moment.js locale configuration
// locale : Uzbek Latin [uz-latn]
// author : Rasulbek Mirzayev : github.com/Rasulbeeek

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("uz-latn", func() {
		moment.DefineLocale("uz-latn", js.Obj(
			"months", []string{"Yanvar", "Fevral", "Mart", "Aprel", "May", "Iyun", "Iyul", "Avgust", "Sentabr", "Oktabr", "Noyabr", "Dekabr"},
			"monthsShort", []string{"Yan", "Fev", "Mar", "Apr", "May", "Iyun", "Iyul", "Avg", "Sen", "Okt", "Noy", "Dek"},
			"weekdays", []string{"Yakshanba", "Dushanba", "Seshanba", "Chorshanba", "Payshanba", "Juma", "Shanba"},
			"weekdaysShort", []string{"Yak", "Dush", "Sesh", "Chor", "Pay", "Jum", "Shan"},
			"weekdaysMin", []string{"Ya", "Du", "Se", "Cho", "Pa", "Ju", "Sha"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "D MMMM YYYY, dddd HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Bugun soat] LT [da]",
				"nextDay", "[Ertaga] LT [da]",
				"nextWeek", "dddd [kuni soat] LT [da]",
				"lastDay", "[Kecha soat] LT [da]",
				"lastWeek", "[O'tgan] dddd [kuni soat] LT [da]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "Yaqin %s ichida",
				"past", "%s oldin",
				"s", "soniya",
				"ss", "%d soniya",
				"m", "bir daqiqa",
				"mm", "%d daqiqa",
				"h", "bir soat",
				"hh", "%d soat",
				"d", "bir kun",
				"dd", "%d kun",
				"M", "bir oy",
				"MM", "%d oy",
				"y", "bir yil",
				"yy", "%d yil",
			),
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}
