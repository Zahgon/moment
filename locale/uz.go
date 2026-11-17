// moment.js locale configuration
// locale : Uzbek [uz]
// author : Sardor Muminov : https://github.com/muminoff

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("uz", func() {
		moment.DefineLocale("uz", js.Obj(
			"months", []string{"январ", "феврал", "март", "апрел", "май", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"},
			"monthsShort", []string{"янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
			"weekdays", []string{"Якшанба", "Душанба", "Сешанба", "Чоршанба", "Пайшанба", "Жума", "Шанба"},
			"weekdaysShort", []string{"Якш", "Душ", "Сеш", "Чор", "Пай", "Жум", "Шан"},
			"weekdaysMin", []string{"Як", "Ду", "Се", "Чо", "Па", "Жу", "Ша"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "D MMMM YYYY, dddd HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Бугун соат] LT [да]",
				"nextDay", "[Эртага] LT [да]",
				"nextWeek", "dddd [куни соат] LT [да]",
				"lastDay", "[Кеча соат] LT [да]",
				"lastWeek", "[Утган] dddd [куни соат] LT [да]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "Якин %s ичида",
				"past", "%s олдин",
				"s", "фурсат",
				"ss", "%d фурсат",
				"m", "бир дакика",
				"mm", "%d дакика",
				"h", "бир соат",
				"hh", "%d соат",
				"d", "бир кун",
				"dd", "%d кун",
				"M", "бир ой",
				"MM", "%d ой",
				"y", "бир йил",
				"yy", "%d йил",
			),
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}
