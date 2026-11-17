// moment.js locale configuration
// locale : Spanish (Mexico) [es-mx]
// author : JC Franco : https://github.com/jcfranco

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("es-mx", func() {
		moment.DefineLocale("es-mx", js.Obj(
			"months", []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
			"monthsShort", esMxMonthsShort,
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
				"sameDay", esMxCalendarSameDay,
				"nextDay", esMxCalendarNextDay,
				"nextWeek", esMxCalendarNextWeek,
				"lastDay", esMxCalendarLastDay,
				"lastWeek", esMxCalendarLastWeek,
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
				"dow", 0.0,
				"doy", 4.0,
			),
			"invalidDate", "Fecha inválida",
		))
	})
}

var (
	esMxMonthsShort      = spanishMonthsShort
	esMxCalendarSameDay  = spanishCalendar("[hoy a la", "] LT")
	esMxCalendarNextDay  = spanishCalendar("[mañana a la", "] LT")
	esMxCalendarNextWeek = spanishCalendar("dddd [a la", "] LT")
	esMxCalendarLastDay  = spanishCalendar("[ayer a la", "] LT")
	esMxCalendarLastWeek = spanishCalendar("[el] dddd [pasado a la", "] LT")
)
