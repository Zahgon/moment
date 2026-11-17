// moment.js locale configuration
// locale : Dutch [nl]
// author : Joris Röling : https://github.com/jorisroling
// author : Jacob Middag : https://github.com/middagj

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("nl", func() {
		moment.DefineLocale("nl", js.Obj(
			"months", []string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
			"monthsShort", nlMonthsShort,
			"monthsRegex", jsre.MustCompile("^(januari|februari|maart|april|mei|ju[nl]i|augustus|september|oktober|november|december|jan\\.?|feb\\.?|mrt\\.?|apr\\.?|ju[nl]\\.?|aug\\.?|sep\\.?|okt\\.?|nov\\.?|dec\\.?)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(januari|februari|maart|april|mei|ju[nl]i|augustus|september|oktober|november|december|jan\\.?|feb\\.?|mrt\\.?|apr\\.?|ju[nl]\\.?|aug\\.?|sep\\.?|okt\\.?|nov\\.?|dec\\.?)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(januari|februari|maart|april|mei|ju[nl]i|augustus|september|oktober|november|december)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(jan\\.?|feb\\.?|mrt\\.?|apr\\.?|mei|ju[nl]\\.?|aug\\.?|sep\\.?|okt\\.?|nov\\.?|dec\\.?)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jan", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^(maart|mrt\\.?)$", "i"),
				jsre.MustCompile("^apr", "i"),
				jsre.MustCompile("^mei$", "i"),
				jsre.MustCompile("^jun[i.]?$", "i"),
				jsre.MustCompile("^jul[i.]?$", "i"),
				jsre.MustCompile("^aug", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^okt", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dec", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jan", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^(maart|mrt\\.?)$", "i"),
				jsre.MustCompile("^apr", "i"),
				jsre.MustCompile("^mei$", "i"),
				jsre.MustCompile("^jun[i.]?$", "i"),
				jsre.MustCompile("^jul[i.]?$", "i"),
				jsre.MustCompile("^aug", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^okt", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dec", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jan", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^(maart|mrt\\.?)$", "i"),
				jsre.MustCompile("^apr", "i"),
				jsre.MustCompile("^mei$", "i"),
				jsre.MustCompile("^jun[i.]?$", "i"),
				jsre.MustCompile("^jul[i.]?$", "i"),
				jsre.MustCompile("^aug", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^okt", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dec", "i"),
			},
			"weekdays", []string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
			"weekdaysShort", []string{"zo.", "ma.", "di.", "wo.", "do.", "vr.", "za."},
			"weekdaysMin", []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD-MM-YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[vandaag om] LT",
				"nextDay", "[morgen om] LT",
				"nextWeek", "dddd [om] LT",
				"lastDay", "[gisteren om] LT",
				"lastWeek", "[afgelopen] dddd [om] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "over %s",
				"past", "%s geleden",
				"s", "een paar seconden",
				"ss", "%d seconden",
				"m", "één minuut",
				"mm", "%d minuten",
				"h", "één uur",
				"hh", "%d uur",
				"d", "één dag",
				"dd", "%d dagen",
				"w", "één week",
				"ww", "%d weken",
				"M", "één maand",
				"MM", "%d maanden",
				"y", "één jaar",
				"yy", "%d jaar",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(ste|de)", ""),
			"ordinal", nlOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var (
	nlMonthsShort          = dutchMonthsShort
	nlOrdinal              = dutchOrdinal
	dutchMonthsWithDots    = []string{"jan.", "feb.", "mrt.", "apr.", "mei", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."}
	dutchMonthsWithoutDots = []string{"jan", "feb", "mrt", "apr", "mei", "jun", "jul", "aug", "sep", "okt", "nov", "dec"}
)

// dutchMonthsShort is the monthsShort callback of nl and nl-be.
func dutchMonthsShort(_ *moment.Locale, m *moment.Moment, format any) any {
	return pickDottedMonth(m, format, dutchMonthsWithDots, dutchMonthsWithoutDots)
}

// dutchOrdinal is the ordinal callback of nl and nl-be.
func dutchOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	if sameNumber(number, 1) || sameNumber(number, 8) || js.ToNumber(number) >= 20 {
		return str(number) + "ste"
	}
	return str(number) + "de"
}
