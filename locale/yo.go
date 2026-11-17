// moment.js locale configuration
// locale : Yoruba Nigeria [yo]
// author : Atolagbe Abisoye : https://github.com/andela-batolagbe

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("yo", func() {
		moment.DefineLocale("yo", js.Obj(
			"months", []string{"Sẹ́rẹ́", "Èrèlè", "Ẹrẹ̀nà", "Ìgbé", "Èbibi", "Òkùdu", "Agẹmo", "Ògún", "Owewe", "Ọ̀wàrà", "Bélú", "Ọ̀pẹ̀̀"},
			"monthsShort", []string{"Sẹ́r", "Èrl", "Ẹrn", "Ìgb", "Èbi", "Òkù", "Agẹ", "Ògú", "Owe", "Ọ̀wà", "Bél", "Ọ̀pẹ̀̀"},
			"weekdays", []string{"Àìkú", "Ajé", "Ìsẹ́gun", "Ọjọ́rú", "Ọjọ́bọ", "Ẹtì", "Àbámẹ́ta"},
			"weekdaysShort", []string{"Àìk", "Ajé", "Ìsẹ́", "Ọjr", "Ọjb", "Ẹtì", "Àbá"},
			"weekdaysMin", []string{"Àì", "Aj", "Ìs", "Ọr", "Ọb", "Ẹt", "Àb"},
			"longDateFormat", js.Obj(
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY h:mm A",
				"LLLL", "dddd, D MMMM YYYY h:mm A",
			),
			"calendar", js.Obj(
				"sameDay", "[Ònì ni] LT",
				"nextDay", "[Ọ̀la ni] LT",
				"nextWeek", "dddd [Ọsẹ̀ tón'bọ] [ni] LT",
				"lastDay", "[Àna ni] LT",
				"lastWeek", "dddd [Ọsẹ̀ tólọ́] [ni] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "ní %s",
				"past", "%s kọjá",
				"s", "ìsẹjú aayá die",
				"ss", "aayá %d",
				"m", "ìsẹjú kan",
				"mm", "ìsẹjú %d",
				"h", "wákati kan",
				"hh", "wákati %d",
				"d", "ọjọ́ kan",
				"dd", "ọjọ́ %d",
				"M", "osù kan",
				"MM", "osù %d",
				"y", "ọdún kan",
				"yy", "ọdún %d",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("ọjọ́\\s\\d{1,2}", ""),
			"ordinal", "ọjọ́ %d",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}
