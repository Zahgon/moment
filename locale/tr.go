// moment.js locale configuration
// locale : Turkish [tr]
// authors : Erhan Gundogan : https://github.com/erhangundogan,
//           Burak Yiğit Kaya: https://github.com/BYK

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("tr", func() {
		moment.DefineLocale("tr", js.Obj(
			"months", []string{"Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık"},
			"monthsShort", []string{"Oca", "Şub", "Mar", "Nis", "May", "Haz", "Tem", "Ağu", "Eyl", "Eki", "Kas", "Ara"},
			"weekdays", []string{"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi"},
			"weekdaysShort", []string{"Paz", "Pzt", "Sal", "Çar", "Per", "Cum", "Cmt"},
			"weekdaysMin", []string{"Pz", "Pt", "Sa", "Ça", "Pe", "Cu", "Ct"},
			"meridiem", trMeridiem,
			"meridiemParse", jsre.MustCompile("öö|ÖÖ|ös|ÖS", ""),
			"isPM", trIsPM,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[bugün saat] LT",
				"nextDay", "[yarın saat] LT",
				"nextWeek", "[gelecek] dddd [saat] LT",
				"lastDay", "[dün] LT",
				"lastWeek", "[geçen] dddd [saat] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s sonra",
				"past", "%s önce",
				"s", "birkaç saniye",
				"ss", "%d saniye",
				"m", "bir dakika",
				"mm", "%d dakika",
				"h", "bir saat",
				"hh", "%d saat",
				"d", "bir gün",
				"dd", "%d gün",
				"w", "bir hafta",
				"ww", "%d hafta",
				"M", "bir ay",
				"MM", "%d ay",
				"y", "bir yıl",
				"yy", "%d yıl",
			),
			"ordinal", trOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var trSuffixes = map[string]string{"1": "'inci", "5": "'inci", "8": "'inci", "70": "'inci", "80": "'inci",
	"2": "'nci", "7": "'nci", "20": "'nci", "50": "'nci", "3": "'üncü", "4": "'üncü", "100": "'üncü", "6": "'ncı",
	"9": "'uncu", "10": "'uncu", "30": "'uncu", "60": "'ıncı", "90": "'ıncı"}

// trMeridiem is the meridiem callback of tr.
func trMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours < 12 {
		if isLower {
			return "öö"
		}
		return "ÖÖ"
	}
	if isLower {
		return "ös"
	}
	return "ÖS"
}

// trIsPM is the isPM callback of tr.
func trIsPM(_ *moment.Locale, input string) bool { return input == "ös" || input == "ÖS" }

// trOrdinal is the ordinal callback of tr.
func trOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	return oghuzOrdinal(trSuffixes, "'ıncı", number, period)
}
