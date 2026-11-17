// moment.js locale configuration
// locale : Kazakh [kk]
// authors : Nurlan Rakhimzhanov : https://github.com/nurlan

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("kk", func() {
		moment.DefineLocale("kk", js.Obj(
			"months", []string{"қаңтар", "ақпан", "наурыз", "сәуір", "мамыр", "маусым", "шілде", "тамыз", "қыркүйек", "қазан", "қараша", "желтоқсан"},
			"monthsShort", []string{"қаң", "ақп", "нау", "сәу", "мам", "мау", "шіл", "там", "қыр", "қаз", "қар", "жел"},
			"weekdays", []string{"жексенбі", "дүйсенбі", "сейсенбі", "сәрсенбі", "бейсенбі", "жұма", "сенбі"},
			"weekdaysShort", []string{"жек", "дүй", "сей", "сәр", "бей", "жұм", "сен"},
			"weekdaysMin", []string{"жк", "дй", "сй", "ср", "бй", "жм", "сн"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Бүгін сағат] LT",
				"nextDay", "[Ертең сағат] LT",
				"nextWeek", "dddd [сағат] LT",
				"lastDay", "[Кеше сағат] LT",
				"lastWeek", "[Өткен аптаның] dddd [сағат] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s ішінде",
				"past", "%s бұрын",
				"s", "бірнеше секунд",
				"ss", "%d секунд",
				"m", "бір минут",
				"mm", "%d минут",
				"h", "бір сағат",
				"hh", "%d сағат",
				"d", "бір күн",
				"dd", "%d күн",
				"M", "бір ай",
				"MM", "%d ай",
				"y", "бір жыл",
				"yy", "%d жыл",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(ші|шы)", ""),
			"ordinal", kkOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var kkSuffixes = map[string]string{"0": "-ші", "1": "-ші", "2": "-ші", "3": "-ші", "4": "-ші", "5": "-ші", "6": "-шы",
	"7": "-ші", "8": "-ші", "9": "-шы", "10": "-шы", "20": "-шы", "30": "-шы", "40": "-шы", "50": "-ші", "60": "-шы",
	"70": "-ші", "80": "-ші", "90": "-шы", "100": "-ші"}

// kkOrdinal is the ordinal callback of kk.
func kkOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return turkicOrdinal(kkSuffixes, number)
}

// turkicOrdinal appends the suffix of the number, its last digit or its hundreds (kk, ky and tg); with none, the
// JavaScript sum of the argument and undefined.
func turkicOrdinal(suffixes map[string]string, number moment.Value) moment.Value {
	n := js.ToNumber(number)
	b := "null"
	if n >= 100 {
		b = "100"
	}
	for _, key := range []string{str(number), num(js.Mod(n, 10)), b} {
		if s := suffixes[key]; s != "" {
			return str(number) + s
		}
	}
	return js.Add(number, js.Undef)
}
