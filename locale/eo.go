// moment.js locale configuration
// locale : Esperanto [eo]
// author : Colin Dean : https://github.com/colindean
// author : Mia Nordentoft Imperatori : https://github.com/miestasmia
// comment : miestasmia corrected the translation by colindean
// comment : Vivakvo corrected the translation by colindean and miestasmia

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("eo", func() {
		moment.DefineLocale("eo", js.Obj(
			"months", []string{"januaro", "februaro", "marto", "aprilo", "majo", "junio", "julio", "aŭgusto", "septembro", "oktobro", "novembro", "decembro"},
			"monthsShort", []string{"jan", "feb", "mart", "apr", "maj", "jun", "jul", "aŭg", "sept", "okt", "nov", "dec"},
			"weekdays", []string{"dimanĉo", "lundo", "mardo", "merkredo", "ĵaŭdo", "vendredo", "sabato"},
			"weekdaysShort", []string{"dim", "lun", "mard", "merk", "ĵaŭ", "ven", "sab"},
			"weekdaysMin", []string{"di", "lu", "ma", "me", "ĵa", "ve", "sa"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY-MM-DD",
				"LL", "[la] D[-an de] MMMM, YYYY",
				"LLL", "[la] D[-an de] MMMM, YYYY HH:mm",
				"LLLL", "dddd[n], [la] D[-an de] MMMM, YYYY HH:mm",
				"llll", "ddd, [la] D[-an de] MMM, YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("[ap]\\.t\\.m", "i"),
			"isPM", eoIsPM,
			"meridiem", eoMeridiem,
			"calendar", js.Obj(
				"sameDay", "[Hodiaŭ je] LT",
				"nextDay", "[Morgaŭ je] LT",
				"nextWeek", "dddd[n je] LT",
				"lastDay", "[Hieraŭ je] LT",
				"lastWeek", "[pasintan] dddd[n je] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "post %s",
				"past", "antaŭ %s",
				"s", "kelkaj sekundoj",
				"ss", "%d sekundoj",
				"m", "unu minuto",
				"mm", "%d minutoj",
				"h", "unu horo",
				"hh", "%d horoj",
				"d", "unu tago",
				"dd", "%d tagoj",
				"M", "unu monato",
				"MM", "%d monatoj",
				"y", "unu jaro",
				"yy", "%d jaroj",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}a", ""),
			"ordinal", "%da",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

// eoIsPM is the isPM callback of eo.
func eoIsPM(_ *moment.Locale, input string) bool { return js.ToLower(js.CharAt16(input, 0)) == "p" }

// eoMeridiem is the meridiem callback of eo.
func eoMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours > 11 {
		if isLower {
			return "p.t.m."
		}
		return "P.T.M."
	}
	if isLower {
		return "a.t.m."
	}
	return "A.T.M."
}
