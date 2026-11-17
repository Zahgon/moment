// moment.js locale configuration
// locale : Spanish [es]
// author : Julio Napurí : https://github.com/julionc

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("es", func() {
		moment.DefineLocale("es", js.Obj(
			"months", []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
			"monthsShort", esMonthsShort,
			"monthsRegex", jsre.MustCompile("^(enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre|ene\\.?|feb\\.?|mar\\.?|abr\\.?|may\\.?|jun\\.?|jul\\.?|ago\\.?|sep\\.?|oct\\.?|nov\\.?|dic\\.?)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre|ene\\.?|feb\\.?|mar\\.?|abr\\.?|may\\.?|jun\\.?|jul\\.?|ago\\.?|sep\\.?|oct\\.?|nov\\.?|dic\\.?)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(enero|febrero|marzo|abril|mayo|junio|julio|agosto|septiembre|octubre|noviembre|diciembre)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(ene\\.?|feb\\.?|mar\\.?|abr\\.?|may\\.?|jun\\.?|jul\\.?|ago\\.?|sep\\.?|oct\\.?|nov\\.?|dic\\.?)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^ene", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^mar", "i"),
				jsre.MustCompile("^abr", "i"),
				jsre.MustCompile("^may", "i"),
				jsre.MustCompile("^jun", "i"),
				jsre.MustCompile("^jul", "i"),
				jsre.MustCompile("^ago", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^oct", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dic", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^ene", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^mar", "i"),
				jsre.MustCompile("^abr", "i"),
				jsre.MustCompile("^may", "i"),
				jsre.MustCompile("^jun", "i"),
				jsre.MustCompile("^jul", "i"),
				jsre.MustCompile("^ago", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^oct", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dic", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^ene", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^mar", "i"),
				jsre.MustCompile("^abr", "i"),
				jsre.MustCompile("^may", "i"),
				jsre.MustCompile("^jun", "i"),
				jsre.MustCompile("^jul", "i"),
				jsre.MustCompile("^ago", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^oct", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dic", "i"),
			},
			"weekdays", []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
			"weekdaysShort", []string{"dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."},
			"weekdaysMin", []string{"do", "lu", "ma", "mi", "ju", "vi", "sá"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D [de] MMMM [de] YYYY",
				"LLL", "D [de] MMMM [de] YYYY H:mm",
				"LLLL", "dddd, D [de] MMMM [de] YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", esCalendarSameDay,
				"nextDay", esCalendarNextDay,
				"nextWeek", esCalendarNextWeek,
				"lastDay", esCalendarLastDay,
				"lastWeek", esCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "en %s",
				"past", "hace %s",
				"s", "unos segundos",
				"ss", "%d segundos",
				"m", "un minuto",
				"mm", "%d minutos",
				"h", "una hora",
				"hh", "%d horas",
				"d", "un día",
				"dd", "%d días",
				"w", "una semana",
				"ww", "%d semanas",
				"M", "un mes",
				"MM", "%d meses",
				"y", "un año",
				"yy", "%d años",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}º", ""),
			"ordinal", "%dº",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
			"invalidDate", "Fecha inválida",
		))
	})
}

var (
	esMonthsShort      = spanishMonthsShort
	esCalendarSameDay  = spanishCalendar("[hoy a la", "] LT")
	esCalendarNextDay  = spanishCalendar("[mañana a la", "] LT")
	esCalendarNextWeek = spanishCalendar("dddd [a la", "] LT")
	esCalendarLastDay  = spanishCalendar("[ayer a la", "] LT")
	esCalendarLastWeek = spanishCalendar("[el] dddd [pasado a la", "] LT")
)

var (
	spanishMonthsShortDot = []string{"ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.",
		"nov.", "dic."}
	spanishMonthsShortPlain = []string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov",
		"dic"}
	spanishDashFormat = jsre.MustCompile(`-MMM-`, "")
)

// spanishMonthsShort drops the abbreviation dots in a -MMM- format (es, es-do, es-mx, es-us).
func spanishMonthsShort(_ *moment.Locale, m *moment.Moment, format any) any {
	if m == nil {
		return spanishMonthsShortDot
	}
	list := spanishMonthsShortDot
	if spanishDashFormat.Test(js.ToString(format)) {
		list = spanishMonthsShortPlain
	}
	if i := m.Month(); i >= 0 && int(i) < len(list) {
		return list[int(i)]
	}
	return js.Undef
}

// spanishCalendar builds a calendar callback: "a la" becomes "a las" except at one o'clock.
func spanishCalendar(before, after string) func(*moment.Moment, *moment.Moment) string {
	return func(m, _ *moment.Moment) string {
		plural := ""
		if m.Hours() != 1 {
			plural = "s"
		}
		return before + plural + after
	}
}
