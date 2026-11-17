// moment.js locale configuration
// locale : Portuguese (Brazil) [pt-br]
// author : Caio Ribeiro Pereira : https://github.com/caio-ribeiro-pereira

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("pt-br", func() {
		moment.DefineLocale("pt-br", js.Obj(
			"months", []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
			"monthsShort", []string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"},
			"weekdays", []string{"domingo", "segunda-feira", "terça-feira", "quarta-feira", "quinta-feira", "sexta-feira", "sábado"},
			"weekdaysShort", []string{"dom", "seg", "ter", "qua", "qui", "sex", "sáb"},
			"weekdaysMin", []string{"do", "2ª", "3ª", "4ª", "5ª", "6ª", "sá"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D [de] MMMM [de] YYYY",
				"LLL", "D [de] MMMM [de] YYYY [às] HH:mm",
				"LLLL", "dddd, D [de] MMMM [de] YYYY [às] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", ptBrCalendarSameDay,
				"nextDay", ptBrCalendarNextDay,
				"nextWeek", ptBrCalendarNextWeek,
				"lastDay", ptBrCalendarLastDay,
				"lastWeek", ptBrCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "em %s",
				"past", "há %s",
				"s", "poucos segundos",
				"ss", "%d segundos",
				"m", "um minuto",
				"mm", "%d minutos",
				"h", "uma hora",
				"hh", "%d horas",
				"d", "um dia",
				"dd", "%d dias",
				"M", "um mês",
				"MM", "%d meses",
				"y", "um ano",
				"yy", "%d anos",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}º", ""),
			"ordinal", "%dº",
			"invalidDate", "Data inválida",
		))
	})
}

// ptBrAt is "às", or "à" before two o'clock.
func ptBrAt(m *moment.Moment) string {
	if m.Hours() >= 2 {
		return "às"
	}
	return "à"
}

// ptBrCalendarSameDay is calendar.sameDay of pt-br.
func ptBrCalendarSameDay(m, _ *moment.Moment) string { return "[Hoje " + ptBrAt(m) + "] LT" }

// ptBrCalendarNextDay is calendar.nextDay of pt-br.
func ptBrCalendarNextDay(m, _ *moment.Moment) string { return "[Amanhã " + ptBrAt(m) + "] LT" }

// ptBrCalendarNextWeek is calendar.nextWeek of pt-br.
func ptBrCalendarNextWeek(m, _ *moment.Moment) string { return "dddd [" + ptBrAt(m) + "] LT" }

// ptBrCalendarLastDay is calendar.lastDay of pt-br.
func ptBrCalendarLastDay(m, _ *moment.Moment) string { return "[Ontem " + ptBrAt(m) + "] LT" }

// ptBrCalendarLastWeek is calendar.lastWeek of pt-br.
func ptBrCalendarLastWeek(m, _ *moment.Moment) string {
	if d := m.Day(); d == 0 || d == 6 {
		return "[Último] dddd [" + ptBrAt(m) + "] LT"
	}
	return "[Última] dddd [" + ptBrAt(m) + "] LT"
}
