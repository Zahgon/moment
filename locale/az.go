// moment.js locale configuration
// locale : Azerbaijani [az]
// author : topchiyev : https://github.com/topchiyev

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("az", func() {
		moment.DefineLocale("az", js.Obj(
			"months", []string{"yanvar", "fevral", "mart", "aprel", "may", "iyun", "iyul", "avqust", "sentyabr", "oktyabr", "noyabr", "dekabr"},
			"monthsShort", []string{"yan", "fev", "mar", "apr", "may", "iyn", "iyl", "avq", "sen", "okt", "noy", "dek"},
			"weekdays", []string{"Bazar", "Bazar ertəsi", "Çərşənbə axşamı", "Çərşənbə", "Cümə axşamı", "Cümə", "Şənbə"},
			"weekdaysShort", []string{"Baz", "BzE", "ÇAx", "Çər", "CAx", "Cüm", "Şən"},
			"weekdaysMin", []string{"Bz", "BE", "ÇA", "Çə", "CA", "Cü", "Şə"},
			"weekdaysParseExact", true,
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
				"nextDay", "[sabah saat] LT",
				"nextWeek", "[gələn həftə] dddd [saat] LT",
				"lastDay", "[dünən] LT",
				"lastWeek", "[keçən həftə] dddd [saat] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s sonra",
				"past", "%s əvvəl",
				"s", "bir neçə saniyə",
				"ss", "%d saniyə",
				"m", "bir dəqiqə",
				"mm", "%d dəqiqə",
				"h", "bir saat",
				"hh", "%d saat",
				"d", "bir gün",
				"dd", "%d gün",
				"M", "bir ay",
				"MM", "%d ay",
				"y", "bir il",
				"yy", "%d il",
			),
			"meridiemParse", jsre.MustCompile("gecə|səhər|gündüz|axşam", ""),
			"isPM", azIsPM,
			"meridiem", azMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(ıncı|inci|nci|üncü|ncı|uncu)", ""),
			"ordinal", azOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var (
	azSuffixes = map[string]string{"1": "-inci", "5": "-inci", "8": "-inci", "70": "-inci", "80": "-inci",
		"2": "-nci", "7": "-nci", "20": "-nci", "50": "-nci", "3": "-üncü", "4": "-üncü", "100": "-üncü",
		"6": "-ncı", "9": "-uncu", "10": "-uncu", "30": "-uncu", "60": "-ıncı", "90": "-ıncı"}
	azPM = jsre.MustCompile(`^(gündüz|axşam)$`, "")
)

// azIsPM is the isPM callback of az.
func azIsPM(_ *moment.Locale, input string) bool { return azPM.Test(input) }

// azMeridiem is the meridiem callback of az.
func azMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "gecə"
	case hour < 12:
		return "səhər"
	case hour < 17:
		return "gündüz"
	}
	return "axşam"
}

// azOrdinal is the ordinal callback of az: the suffix of the last digit, the tens or the hundreds.
func azOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	if sameNumber(number, 0) {
		return str(number) + "-ıncı"
	}
	return js.Add(number, digitSuffix(azSuffixes, js.ToNumber(number)))
}
