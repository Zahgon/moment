// moment.js locale configuration
// locale : Lithuanian [lt]
// author : Mindaugas Mozūras : https://github.com/mmozuras

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("lt", func() {
		moment.DefineLocale("lt", js.Obj(
			"months", js.Obj(
				"format", []string{"sausio", "vasario", "kovo", "balandžio", "gegužės", "birželio", "liepos", "rugpjūčio", "rugsėjo", "spalio", "lapkričio", "gruodžio"},
				"standalone", []string{"sausis", "vasaris", "kovas", "balandis", "gegužė", "birželis", "liepa", "rugpjūtis", "rugsėjis", "spalis", "lapkritis", "gruodis"},
				"isFormat", jsre.MustCompile("D[oD]?(\\[[^\\[\\]]*\\]|\\s)+MMMM?|MMMM?(\\[[^\\[\\]]*\\]|\\s)+D[oD]?", ""),
			),
			"monthsShort", []string{"sau", "vas", "kov", "bal", "geg", "bir", "lie", "rgp", "rgs", "spa", "lap", "grd"},
			"weekdays", js.Obj(
				"format", []string{"sekmadienį", "pirmadienį", "antradienį", "trečiadienį", "ketvirtadienį", "penktadienį", "šeštadienį"},
				"standalone", []string{"sekmadienis", "pirmadienis", "antradienis", "trečiadienis", "ketvirtadienis", "penktadienis", "šeštadienis"},
				"isFormat", jsre.MustCompile("dddd HH:mm", ""),
			),
			"weekdaysShort", []string{"Sek", "Pir", "Ant", "Tre", "Ket", "Pen", "Šeš"},
			"weekdaysMin", []string{"S", "P", "A", "T", "K", "Pn", "Š"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY-MM-DD",
				"LL", "YYYY [m.] MMMM D [d.]",
				"LLL", "YYYY [m.] MMMM D [d.], HH:mm [val.]",
				"LLLL", "YYYY [m.] MMMM D [d.], dddd, HH:mm [val.]",
				"l", "YYYY-MM-DD",
				"ll", "YYYY [m.] MMMM D [d.]",
				"lll", "YYYY [m.] MMMM D [d.], HH:mm [val.]",
				"llll", "YYYY [m.] MMMM D [d.], ddd, HH:mm [val.]",
			),
			"calendar", js.Obj(
				"sameDay", "[Šiandien] LT",
				"nextDay", "[Rytoj] LT",
				"nextWeek", "dddd LT",
				"lastDay", "[Vakar] LT",
				"lastWeek", "[Praėjusį] dddd LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "po %s",
				"past", "prieš %s",
				"s", ltRelativeTimeSeconds,
				"ss", ltRelativeTimeSecondsN,
				"m", ltRelativeTimeMinute,
				"mm", ltRelativeTimeSecondsN,
				"h", ltRelativeTimeMinute,
				"hh", ltRelativeTimeSecondsN,
				"d", ltRelativeTimeMinute,
				"dd", ltRelativeTimeSecondsN,
				"M", ltRelativeTimeMinute,
				"MM", ltRelativeTimeSecondsN,
				"y", ltRelativeTimeMinute,
				"yy", ltRelativeTimeSecondsN,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-oji", ""),
			"ordinal", ltOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var ltUnits = map[string]string{
	"ss": "sekundė_sekundžių_sekundes", "m": "minutė_minutės_minutę", "mm": "minutės_minučių_minutes",
	"h": "valanda_valandos_valandą", "hh": "valandos_valandų_valandas", "d": "diena_dienos_dieną",
	"dd": "dienos_dienų_dienas", "M": "mėnuo_mėnesio_mėnesį", "MM": "mėnesiai_mėnesių_mėnesius",
	"y": "metai_metų_metus", "yy": "metai_metų_metus",
}

// ltForm returns form i of a unit's word list, "" when the unit has none.
func ltForm(key string, i int) string {
	forms := js.Split(ltUnits[key], "_")
	if len(forms) <= i {
		return ""
	}
	return forms[i]
}

// ltSpecial reports the numbers that take the genitive plural.
func ltSpecial(number float64) bool { return js.Mod(number, 10) == 0 || (number > 10 && number < 20) }

// ltRelativeTimeSeconds is translateSeconds of lt.
func ltRelativeTimeSeconds(_ float64, withoutSuffix bool, _ string, isFuture bool) string {
	switch {
	case withoutSuffix:
		return "kelios sekundės"
	case isFuture:
		return "kelių sekundžių"
	}
	return "kelias sekundes"
}

// ltRelativeTimeMinute is translateSingular of lt.
func ltRelativeTimeMinute(_ float64, withoutSuffix bool, key string, isFuture bool) string {
	switch {
	case withoutSuffix:
		return ltForm(key, 0)
	case isFuture:
		return ltForm(key, 1)
	}
	return ltForm(key, 2)
}

// ltRelativeTimeSecondsN is translate of lt.
func ltRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, isFuture bool) string {
	result := num(number) + " "
	switch {
	case number == 1:
		return result + ltRelativeTimeMinute(number, withoutSuffix, js.CharAt16(key, 0), isFuture)
	case withoutSuffix:
		if ltSpecial(number) {
			return result + ltForm(key, 1)
		}
		return result + ltForm(key, 0)
	case isFuture:
		return result + ltForm(key, 1)
	case ltSpecial(number):
		return result + ltForm(key, 1)
	}
	return result + ltForm(key, 2)
}

// ltOrdinal is the ordinal callback of lt.
func ltOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return str(number) + "-oji"
}
