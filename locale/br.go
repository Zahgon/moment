// moment.js locale configuration
// locale : Breton [br]
// author : Jean-Baptiste Le Duigou : https://github.com/jbleduigou

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("br", func() {
		moment.DefineLocale("br", js.Obj(
			"months", []string{"Genver", "Cʼhwevrer", "Meurzh", "Ebrel", "Mae", "Mezheven", "Gouere", "Eost", "Gwengolo", "Here", "Du", "Kerzu"},
			"monthsShort", []string{"Gen", "Cʼhwe", "Meu", "Ebr", "Mae", "Eve", "Gou", "Eos", "Gwe", "Her", "Du", "Ker"},
			"weekdays", []string{"Sul", "Lun", "Meurzh", "Mercʼher", "Yaou", "Gwener", "Sadorn"},
			"weekdaysShort", []string{"Sul", "Lun", "Meu", "Mer", "Yao", "Gwe", "Sad"},
			"weekdaysMin", []string{"Su", "Lu", "Me", "Mer", "Ya", "Gw", "Sa"},
			"weekdaysParse", []*jsre.Regexp{
				jsre.MustCompile("^Su", "i"),
				jsre.MustCompile("^Lu", "i"),
				jsre.MustCompile("^Me([^r]|$)", "i"),
				jsre.MustCompile("^Mer", "i"),
				jsre.MustCompile("^Ya", "i"),
				jsre.MustCompile("^Gw", "i"),
				jsre.MustCompile("^Sa", "i"),
			},
			"fullWeekdaysParse", []*jsre.Regexp{
				jsre.MustCompile("^sul", "i"),
				jsre.MustCompile("^lun", "i"),
				jsre.MustCompile("^meurzh", "i"),
				jsre.MustCompile("^merc[ʼ\\']her", "i"),
				jsre.MustCompile("^yaou", "i"),
				jsre.MustCompile("^gwener", "i"),
				jsre.MustCompile("^sadorn", "i"),
			},
			"shortWeekdaysParse", []*jsre.Regexp{
				jsre.MustCompile("^Sul", "i"),
				jsre.MustCompile("^Lun", "i"),
				jsre.MustCompile("^Meu", "i"),
				jsre.MustCompile("^Mer", "i"),
				jsre.MustCompile("^Yao", "i"),
				jsre.MustCompile("^Gwe", "i"),
				jsre.MustCompile("^Sad", "i"),
			},
			"minWeekdaysParse", []*jsre.Regexp{
				jsre.MustCompile("^Su", "i"),
				jsre.MustCompile("^Lu", "i"),
				jsre.MustCompile("^Me([^r]|$)", "i"),
				jsre.MustCompile("^Mer", "i"),
				jsre.MustCompile("^Ya", "i"),
				jsre.MustCompile("^Gw", "i"),
				jsre.MustCompile("^Sa", "i"),
			},
			"monthsRegex", jsre.MustCompile("^(genver|c[ʼ\\']hwevrer|meurzh|ebrel|mae|mezheven|gouere|eost|gwengolo|here|du|kerzu|gen|c[ʼ\\']hwe|meu|ebr|mae|eve|gou|eos|gwe|her|du|ker)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(genver|c[ʼ\\']hwevrer|meurzh|ebrel|mae|mezheven|gouere|eost|gwengolo|here|du|kerzu|gen|c[ʼ\\']hwe|meu|ebr|mae|eve|gou|eos|gwe|her|du|ker)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(genver|c[ʼ\\']hwevrer|meurzh|ebrel|mae|mezheven|gouere|eost|gwengolo|here|du|kerzu)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(gen|c[ʼ\\']hwe|meu|ebr|mae|eve|gou|eos|gwe|her|du|ker)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^gen", "i"),
				jsre.MustCompile("^c[ʼ\\']hwe", "i"),
				jsre.MustCompile("^meu", "i"),
				jsre.MustCompile("^ebr", "i"),
				jsre.MustCompile("^mae", "i"),
				jsre.MustCompile("^(mez|eve)", "i"),
				jsre.MustCompile("^gou", "i"),
				jsre.MustCompile("^eos", "i"),
				jsre.MustCompile("^gwe", "i"),
				jsre.MustCompile("^her", "i"),
				jsre.MustCompile("^du", "i"),
				jsre.MustCompile("^ker", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^gen", "i"),
				jsre.MustCompile("^c[ʼ\\']hwe", "i"),
				jsre.MustCompile("^meu", "i"),
				jsre.MustCompile("^ebr", "i"),
				jsre.MustCompile("^mae", "i"),
				jsre.MustCompile("^(mez|eve)", "i"),
				jsre.MustCompile("^gou", "i"),
				jsre.MustCompile("^eos", "i"),
				jsre.MustCompile("^gwe", "i"),
				jsre.MustCompile("^her", "i"),
				jsre.MustCompile("^du", "i"),
				jsre.MustCompile("^ker", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^gen", "i"),
				jsre.MustCompile("^c[ʼ\\']hwe", "i"),
				jsre.MustCompile("^meu", "i"),
				jsre.MustCompile("^ebr", "i"),
				jsre.MustCompile("^mae", "i"),
				jsre.MustCompile("^(mez|eve)", "i"),
				jsre.MustCompile("^gou", "i"),
				jsre.MustCompile("^eos", "i"),
				jsre.MustCompile("^gwe", "i"),
				jsre.MustCompile("^her", "i"),
				jsre.MustCompile("^du", "i"),
				jsre.MustCompile("^ker", "i"),
			},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D [a viz] MMMM YYYY",
				"LLL", "D [a viz] MMMM YYYY HH:mm",
				"LLLL", "dddd, D [a viz] MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Hiziv da] LT",
				"nextDay", "[Warcʼhoazh da] LT",
				"nextWeek", "dddd [da] LT",
				"lastDay", "[Decʼh da] LT",
				"lastWeek", "dddd [paset da] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "a-benn %s",
				"past", "%s ʼzo",
				"s", "un nebeud segondennoù",
				"ss", "%d eilenn",
				"m", "ur vunutenn",
				"mm", brRelativeTimeMinutes,
				"h", "un eur",
				"hh", "%d eur",
				"d", "un devezh",
				"dd", brRelativeTimeMinutes,
				"M", "ur miz",
				"MM", brRelativeTimeMinutes,
				"y", "ur bloaz",
				"yy", brRelativeTimeYears,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(añ|vet)", ""),
			"ordinal", brOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
			"meridiemParse", jsre.MustCompile("a.m.|g.m.", ""),
			"isPM", brIsPM,
			"meridiem", brMeridiem,
		))
	})
}

// brMutation softens the first consonant of text after the number 2.
func brMutation(text string, number float64) string {
	if number != 2 {
		return text
	}
	table := map[string]string{"m": "v", "b": "v", "d": "z"}
	first, ok := table[js.CharAt16(text, 0)]
	if !ok {
		return text
	}
	return first + js.Substring16(text, 1, js.Len16(text))
}

// brLastNumber is the last digit of number (for numbers above 9).
func brLastNumber(number float64) float64 {
	if number > 9 {
		return brLastNumber(js.Mod(number, 10))
	}
	return number
}

// brRelativeTimeMinutes is relativeTimeWithMutation of br.
func brRelativeTimeMinutes(number float64, _ bool, key string, _ bool) string {
	format := map[string]string{"mm": "munutenn", "MM": "miz", "dd": "devezh"}
	text, ok := format[key]
	if !ok {
		text = "undefined"
	}
	return num(number) + " " + brMutation(text, number)
}

// brRelativeTimeYears is specialMutationForYears of br.
func brRelativeTimeYears(number float64, _ bool, _ string, _ bool) string {
	switch brLastNumber(number) {
	case 1, 3, 4, 5, 9:
		return num(number) + " bloaz"
	}
	return num(number) + " vloaz"
}

// brOrdinal is the ordinal callback of br.
func brOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	if sameNumber(number, 1) {
		return str(number) + "añ"
	}
	return str(number) + "vet"
}

// brIsPM is the isPM callback of br.
func brIsPM(_ *moment.Locale, token string) bool { return token == "g.m." }

// brMeridiem is the meridiem callback of br.
func brMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "a.m."
	}
	return "g.m."
}
