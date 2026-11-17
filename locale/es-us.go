// moment.js locale configuration
// locale : Spanish (United States) [es-us]
// author : bustta : https://github.com/bustta
// author : chrisrodz : https://github.com/chrisrodz

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("es-us", func() {
		moment.DefineLocale("es-us", js.Obj(
			"months", []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
			"monthsShort", esUsMonthsShort,
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
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "MM/DD/YYYY",
				"LL", "D [de] MMMM [de] YYYY",
				"LLL", "D [de] MMMM [de] YYYY h:mm A",
				"LLLL", "dddd, D [de] MMMM [de] YYYY h:mm A",
			),
			"calendar", js.Obj(
				"sameDay", esUsCalendarSameDay,
				"nextDay", esUsCalendarNextDay,
				"nextWeek", esUsCalendarNextWeek,
				"lastDay", esUsCalendarLastDay,
				"lastWeek", esUsCalendarLastWeek,
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
				"doy", 6.0,
			),
		))
	})
}

var (
	esUsMonthsShort      = spanishMonthsShort
	esUsCalendarSameDay  = spanishCalendar("[hoy a la", "] LT")
	esUsCalendarNextDay  = spanishCalendar("[mañana a la", "] LT")
	esUsCalendarNextWeek = spanishCalendar("dddd [a la", "] LT")
	esUsCalendarLastDay  = spanishCalendar("[ayer a la", "] LT")
	esUsCalendarLastWeek = spanishCalendar("[el] dddd [pasado a la", "] LT")
)
