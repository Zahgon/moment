// moment.js locale configuration
// locale : Northern Kurdish [ku-kmr]
// authors : Mazlum Özdogan : https://github.com/mergehez

package locale

import (
	"strings"

	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ku-kmr", func() {
		moment.DefineLocale("ku-kmr", js.Obj(
			"months", []string{"Rêbendan", "Sibat", "Adar", "Nîsan", "Gulan", "Hezîran", "Tîrmeh", "Tebax", "Îlon", "Cotmeh", "Mijdar", "Berfanbar"},
			"monthsShort", []string{"Rêb", "Sib", "Ada", "Nîs", "Gul", "Hez", "Tîr", "Teb", "Îlo", "Cot", "Mij", "Ber"},
			"monthsParseExact", true,
			"weekdays", []string{"Yekşem", "Duşem", "Sêşem", "Çarşem", "Pêncşem", "În", "Şemî"},
			"weekdaysShort", []string{"Yek", "Du", "Sê", "Çar", "Pên", "În", "Şem"},
			"weekdaysMin", []string{"Ye", "Du", "Sê", "Ça", "Pê", "În", "Şe"},
			"meridiem", kuKmrMeridiem,
			"meridiemParse", jsre.MustCompile("bn|BN|pn|PN", ""),
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "Do MMMM[a] YYYY[an]",
				"LLL", "Do MMMM[a] YYYY[an] HH:mm",
				"LLLL", "dddd, Do MMMM[a] YYYY[an] HH:mm",
				"ll", "Do MMM[.] YYYY[an]",
				"lll", "Do MMM[.] YYYY[an] HH:mm",
				"llll", "ddd[.], Do MMM[.] YYYY[an] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Îro di saet] LT [de]",
				"nextDay", "[Sibê di saet] LT [de]",
				"nextWeek", "dddd [di saet] LT [de]",
				"lastDay", "[Duh di saet] LT [de]",
				"lastWeek", "dddd[a borî di saet] LT [de]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "di %s de",
				"past", "berî %s",
				"s", kuKmrRelativeTimeSeconds,
				"ss", kuKmrRelativeTimeSeconds,
				"m", kuKmrRelativeTimeSeconds,
				"mm", kuKmrRelativeTimeSeconds,
				"h", kuKmrRelativeTimeSeconds,
				"hh", kuKmrRelativeTimeSeconds,
				"d", kuKmrRelativeTimeSeconds,
				"dd", kuKmrRelativeTimeSeconds,
				"w", kuKmrRelativeTimeSeconds,
				"ww", kuKmrRelativeTimeSeconds,
				"M", kuKmrRelativeTimeSeconds,
				"MM", kuKmrRelativeTimeSeconds,
				"y", kuKmrRelativeTimeSeconds,
				"yy", kuKmrRelativeTimeSeconds,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(?:yê|ê|\\.)", ""),
			"ordinal", kuKmrOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// kuKmrMeridiem is the meridiem callback of ku-kmr.
func kuKmrMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours < 12 {
		if isLower {
			return "bn"
		}
		return "BN"
	}
	if isLower {
		return "pn"
	}
	return "PN"
}

// kuKmrRelativeTimeSeconds is processRelativeTime of ku-kmr.
func kuKmrRelativeTimeSeconds(number float64, withoutSuffix bool, key string, _ bool) string {
	n := num(number)
	format := map[string][2]string{
		"s": {"çend sanîye", "çend sanîyeyan"}, "ss": {n + " sanîye", n + " sanîyeyan"},
		"m": {"deqîqeyek", "deqîqeyekê"}, "mm": {n + " deqîqe", n + " deqîqeyan"},
		"h": {"saetek", "saetekê"}, "hh": {n + " saet", n + " saetan"},
		"d": {"rojek", "rojekê"}, "dd": {n + " roj", n + " rojan"},
		"w": {"hefteyek", "hefteyekê"}, "ww": {n + " hefte", n + " hefteyan"},
		"M": {"mehek", "mehekê"}, "MM": {n + " meh", n + " mehan"},
		"y": {"salek", "salekê"}, "yy": {n + " sal", n + " salan"},
	}
	if withoutSuffix {
		return format[key][0]
	}
	return format[key][1]
}

// kuKmrEzafe is the ezafe suffix of a number: "yê" after 2, 3, 50, 70 and 80 (but not 12 or 13), "ê" otherwise.
func kuKmrEzafe(n string) string {
	length := js.Len16(n)
	l := js.Substring16(n, length-1, length)
	ll := ""
	if length > 1 {
		ll = js.Substring16(n, length-2, length)
	}
	twelveOrThirteen := js.StringToNumber(ll) == 12 || js.StringToNumber(ll) == 13
	if !twelveOrThirteen && (l == "2" || l == "3" || ll == "50" || l == "70" || l == "80") {
		return "yê"
	}
	return "ê"
}

// kuKmrOrdinal is the ordinal callback of ku-kmr.
func kuKmrOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	p := js.ToLower(period)
	if strings.Contains(p, "w") || strings.Contains(p, "m") {
		return str(number) + "."
	}
	return str(number) + kuKmrEzafe(str(number))
}
