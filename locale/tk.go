// moment.js locale configuration
// locale : Turkmen [tk]
// author : Atamyrat Abdyrahmanov : https://github.com/atamyratabdy

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("tk", func() {
		moment.DefineLocale("tk", js.Obj(
			"months", []string{"Ýanwar", "Fewral", "Mart", "Aprel", "Maý", "Iýun", "Iýul", "Awgust", "Sentýabr", "Oktýabr", "Noýabr", "Dekabr"},
			"monthsShort", []string{"Ýan", "Few", "Mar", "Apr", "Maý", "Iýn", "Iýl", "Awg", "Sen", "Okt", "Noý", "Dek"},
			"weekdays", []string{"Ýekşenbe", "Duşenbe", "Sişenbe", "Çarşenbe", "Penşenbe", "Anna", "Şenbe"},
			"weekdaysShort", []string{"Ýek", "Duş", "Siş", "Çar", "Pen", "Ann", "Şen"},
			"weekdaysMin", []string{"Ýk", "Dş", "Sş", "Çr", "Pn", "An", "Şn"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[bugün sagat] LT",
				"nextDay", "[ertir sagat] LT",
				"nextWeek", "[indiki] dddd [sagat] LT",
				"lastDay", "[düýn] LT",
				"lastWeek", "[geçen] dddd [sagat] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s soň",
				"past", "%s öň",
				"s", "birnäçe sekunt",
				"m", "bir minut",
				"mm", "%d minut",
				"h", "bir sagat",
				"hh", "%d sagat",
				"d", "bir gün",
				"dd", "%d gün",
				"M", "bir aý",
				"MM", "%d aý",
				"y", "bir ýyl",
				"yy", "%d ýyl",
			),
			"ordinal", tkOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var tkSuffixes = map[string]string{"1": "'inji", "5": "'inji", "8": "'inji", "70": "'inji", "80": "'inji",
	"2": "'nji", "7": "'nji", "20": "'nji", "50": "'nji", "3": "'ünji", "4": "'ünji", "100": "'ünji", "6": "'njy",
	"9": "'unjy", "10": "'unjy", "30": "'unjy", "60": "'ynjy", "90": "'ynjy"}

// tkOrdinal is the ordinal callback of tk.
func tkOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	return oghuzOrdinal(tkSuffixes, "'unjy", number, period)
}

// oghuzOrdinal is the ordinal of tr and tk: days are the argument itself, zero has its own suffix, other numbers
// take the suffix of their last digit, tens or hundreds.
func oghuzOrdinal(suffixes map[string]string, zero string, number moment.Value, period string) moment.Value {
	switch period {
	case "d", "D", "Do", "DD":
		return number
	}
	if sameNumber(number, 0) {
		return str(number) + zero
	}
	return js.Add(number, digitSuffix(suffixes, js.ToNumber(number)))
}
