// moment.js locale configuration
// locale : Portuguese [pt]
// author : Jefferson : https://github.com/jalex79

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("pt", func() {
		moment.DefineLocale("pt", js.Obj(
			"months", []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
			"monthsShort", []string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"},
			"weekdays", []string{"Domingo", "Segunda-feira", "Terça-feira", "Quarta-feira", "Quinta-feira", "Sexta-feira", "Sábado"},
			"weekdaysShort", []string{"Dom", "Seg", "Ter", "Qua", "Qui", "Sex", "Sáb"},
			"weekdaysMin", []string{"Do", "2ª", "3ª", "4ª", "5ª", "6ª", "Sá"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D [de] MMMM [de] YYYY",
				"LLL", "D [de] MMMM [de] YYYY HH:mm",
				"LLLL", "dddd, D [de] MMMM [de] YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Hoje às] LT",
				"nextDay", "[Amanhã às] LT",
				"nextWeek", "dddd [às] LT",
				"lastDay", "[Ontem às] LT",
				"lastWeek", ptCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "em %s",
				"past", "há %s",
				"s", "segundos",
				"ss", "%d segundos",
				"m", "um minuto",
				"mm", "%d minutos",
				"h", "uma hora",
				"hh", "%d horas",
				"d", "um dia",
				"dd", "%d dias",
				"w", "uma semana",
				"ww", "%d semanas",
				"M", "um mês",
				"MM", "%d meses",
				"y", "um ano",
				"yy", "%d anos",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}º", ""),
			"ordinal", "%dº",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// ptCalendarLastWeek is calendar.lastWeek of pt.
func ptCalendarLastWeek(m, _ *moment.Moment) string {
	if d := m.Day(); d == 0 || d == 6 {
		return "[Último] dddd [às] LT"
	}
	return "[Última] dddd [às] LT"
}
