// moment.js locale configuration
// locale : Kyrgyz [ky]
// author : Chyngyz Arystan uulu : https://github.com/chyngyz

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ky", func() {
		moment.DefineLocale("ky", js.Obj(
			"months", []string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
			"monthsShort", []string{"янв", "фев", "март", "апр", "май", "июнь", "июль", "авг", "сен", "окт", "ноя", "дек"},
			"weekdays", []string{"Жекшемби", "Дүйшөмбү", "Шейшемби", "Шаршемби", "Бейшемби", "Жума", "Ишемби"},
			"weekdaysShort", []string{"Жек", "Дүй", "Шей", "Шар", "Бей", "Жум", "Ише"},
			"weekdaysMin", []string{"Жк", "Дй", "Шй", "Шр", "Бй", "Жм", "Иш"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Бүгүн саат] LT",
				"nextDay", "[Эртең саат] LT",
				"nextWeek", "dddd [саат] LT",
				"lastDay", "[Кечээ саат] LT",
				"lastWeek", "[Өткөн аптанын] dddd [күнү] [саат] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s ичинде",
				"past", "%s мурун",
				"s", "бирнече секунд",
				"ss", "%d секунд",
				"m", "бир мүнөт",
				"mm", "%d мүнөт",
				"h", "бир саат",
				"hh", "%d саат",
				"d", "бир күн",
				"dd", "%d күн",
				"M", "бир ай",
				"MM", "%d ай",
				"y", "бир жыл",
				"yy", "%d жыл",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(чи|чы|чү|чу)", ""),
			"ordinal", kyOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var kySuffixes = map[string]string{"0": "-чү", "1": "-чи", "2": "-чи", "3": "-чү", "4": "-чү", "5": "-чи", "6": "-чы",
	"7": "-чи", "8": "-чи", "9": "-чу", "10": "-чу", "20": "-чы", "30": "-чу", "40": "-чы", "50": "-чү", "60": "-чы",
	"70": "-чи", "80": "-чи", "90": "-чу", "100": "-чү"}

// kyOrdinal is the ordinal callback of ky.
func kyOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return turkicOrdinal(kySuffixes, number)
}
