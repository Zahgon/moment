// moment.js locale configuration
// locale : Chuvash [cv]
// author : Anatoly Mironov : https://github.com/mirontoli

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("cv", func() {
		moment.DefineLocale("cv", js.Obj(
			"months", []string{"кӑрлач", "нарӑс", "пуш", "ака", "май", "ҫӗртме", "утӑ", "ҫурла", "авӑн", "юпа", "чӳк", "раштав"},
			"monthsShort", []string{"кӑр", "нар", "пуш", "ака", "май", "ҫӗр", "утӑ", "ҫур", "авн", "юпа", "чӳк", "раш"},
			"weekdays", []string{"вырсарникун", "тунтикун", "ытларикун", "юнкун", "кӗҫнерникун", "эрнекун", "шӑматкун"},
			"weekdaysShort", []string{"выр", "тун", "ытл", "юн", "кӗҫ", "эрн", "шӑм"},
			"weekdaysMin", []string{"вр", "тн", "ыт", "юн", "кҫ", "эр", "шм"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD-MM-YYYY",
				"LL", "YYYY [ҫулхи] MMMM [уйӑхӗн] D[-мӗшӗ]",
				"LLL", "YYYY [ҫулхи] MMMM [уйӑхӗн] D[-мӗшӗ], HH:mm",
				"LLLL", "dddd, YYYY [ҫулхи] MMMM [уйӑхӗн] D[-мӗшӗ], HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Паян] LT [сехетре]",
				"nextDay", "[Ыран] LT [сехетре]",
				"lastDay", "[Ӗнер] LT [сехетре]",
				"nextWeek", "[Ҫитес] dddd LT [сехетре]",
				"lastWeek", "[Иртнӗ] dddd LT [сехетре]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", cvRelativeTimeFuture,
				"past", "%s каялла",
				"s", "пӗр-ик ҫеккунт",
				"ss", "%d ҫеккунт",
				"m", "пӗр минут",
				"mm", "%d минут",
				"h", "пӗр сехет",
				"hh", "%d сехет",
				"d", "пӗр кун",
				"dd", "%d кун",
				"M", "пӗр уйӑх",
				"MM", "%d уйӑх",
				"y", "пӗр ҫул",
				"yy", "%d ҫул",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-мӗш", ""),
			"ordinal", "%d-мӗш",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var (
	cvHours = jsre.MustCompile(`сехет$`, "i")
	cvYears = jsre.MustCompile(`ҫул$`, "i")
)

// cvRelativeTimeFuture is relativeTime.future of cv.
func cvRelativeTimeFuture(output string) string {
	switch {
	case cvHours.Exec(output) != nil:
		return output + "рен"
	case cvYears.Exec(output) != nil:
		return output + "тан"
	}
	return output + "ран"
}
