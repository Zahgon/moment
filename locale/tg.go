// moment.js locale configuration
// locale : Tajik [tg]
// author : Orif N. Jr. : https://github.com/orif-jr

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("tg", func() {
		moment.DefineLocale("tg", js.Obj(
			"months", js.Obj(
				"format", []string{"январи", "феврали", "марти", "апрели", "майи", "июни", "июли", "августи", "сентябри", "октябри", "ноябри", "декабри"},
				"standalone", []string{"январ", "феврал", "март", "апрел", "май", "июн", "июл", "август", "сентябр", "октябр", "ноябр", "декабр"},
			),
			"monthsShort", []string{"янв", "фев", "мар", "апр", "май", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
			"weekdays", []string{"якшанбе", "душанбе", "сешанбе", "чоршанбе", "панҷшанбе", "ҷумъа", "шанбе"},
			"weekdaysShort", []string{"яшб", "дшб", "сшб", "чшб", "пшб", "ҷум", "шнб"},
			"weekdaysMin", []string{"яш", "дш", "сш", "чш", "пш", "ҷм", "шб"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Имрӯз соати] LT",
				"nextDay", "[Фардо соати] LT",
				"lastDay", "[Дирӯз соати] LT",
				"nextWeek", "dddd[и] [ҳафтаи оянда соати] LT",
				"lastWeek", "dddd[и] [ҳафтаи гузашта соати] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "баъди %s",
				"past", "%s пеш",
				"s", "якчанд сония",
				"m", "як дақиқа",
				"mm", "%d дақиқа",
				"h", "як соат",
				"hh", "%d соат",
				"d", "як рӯз",
				"dd", "%d рӯз",
				"M", "як моҳ",
				"MM", "%d моҳ",
				"y", "як сол",
				"yy", "%d сол",
			),
			"meridiemParse", jsre.MustCompile("шаб|субҳ|рӯз|бегоҳ", ""),
			"meridiemHour", tgMeridiemHour,
			"meridiem", tgMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(ум|юм)", ""),
			"ordinal", tgOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var tgSuffixes = map[string]string{"0": "-ум", "1": "-ум", "2": "-юм", "3": "-юм", "4": "-ум", "5": "-ум", "6": "-ум",
	"7": "-ум", "8": "-ум", "9": "-ум", "10": "-ум", "12": "-ум", "13": "-ум", "20": "-ум", "30": "-юм", "40": "-ум",
	"50": "-ум", "60": "-ум", "70": "-ум", "80": "-ум", "90": "-ум", "100": "-ум"}

// tgMeridiemHour is the meridiemHour callback of tg.
func tgMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return dayPartHour(hour, meridiem, "шаб", "субҳ", "рӯз", "бегоҳ", 11)
}

// tgMeridiem is the meridiem callback of tg.
func tgMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "шаб"
	case hour < 11:
		return "субҳ"
	case hour < 16:
		return "рӯз"
	case hour < 19:
		return "бегоҳ"
	}
	return "шаб"
}

// tgOrdinal is the ordinal callback of tg.
func tgOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return turkicOrdinal(tgSuffixes, number)
}
