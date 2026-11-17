// moment.js locale configuration
// locale : Hungarian [hu]
// author : Adam Brunner : https://github.com/adambrunner
// author : Peter Viszt  : https://github.com/passatgt

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("hu", func() {
		moment.DefineLocale("hu", js.Obj(
			"months", []string{"január", "február", "március", "április", "május", "június", "július", "augusztus", "szeptember", "október", "november", "december"},
			"monthsShort", []string{"jan.", "feb.", "márc.", "ápr.", "máj.", "jún.", "júl.", "aug.", "szept.", "okt.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"vasárnap", "hétfő", "kedd", "szerda", "csütörtök", "péntek", "szombat"},
			"weekdaysShort", []string{"vas", "hét", "kedd", "sze", "csüt", "pén", "szo"},
			"weekdaysMin", []string{"v", "h", "k", "sze", "cs", "p", "szo"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "YYYY.MM.DD.",
				"LL", "YYYY. MMMM D.",
				"LLL", "YYYY. MMMM D. H:mm",
				"LLLL", "YYYY. MMMM D., dddd H:mm",
			),
			"meridiemParse", jsre.MustCompile("de|du", "i"),
			"isPM", huIsPM,
			"meridiem", huMeridiem,
			"calendar", js.Obj(
				"sameDay", "[ma] LT[-kor]",
				"nextDay", "[holnap] LT[-kor]",
				"nextWeek", huCalendarNextWeek,
				"lastDay", "[tegnap] LT[-kor]",
				"lastWeek", huCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s múlva",
				"past", "%s",
				"s", huRelativeTimeSeconds,
				"ss", huRelativeTimeSeconds,
				"m", huRelativeTimeSeconds,
				"mm", huRelativeTimeSeconds,
				"h", huRelativeTimeSeconds,
				"hh", huRelativeTimeSeconds,
				"d", huRelativeTimeSeconds,
				"dd", huRelativeTimeSeconds,
				"M", huRelativeTimeSeconds,
				"MM", huRelativeTimeSeconds,
				"y", huRelativeTimeSeconds,
				"yy", huRelativeTimeSeconds,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}\\.", ""),
			"ordinal", "%d.",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var huWeekEndings = []string{"vasárnap", "hétfőn", "kedden", "szerdán", "csütörtökön", "pénteken", "szombaton"}

// huRelativeTimeSeconds is translate of hu.
func huRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	plain := isFuture || withoutSuffix
	pick := func(a, b string) string {
		if plain {
			return a
		}
		return b
	}
	n := num(number)
	switch key {
	case "s":
		return pick("néhány másodperc", "néhány másodperce")
	case "ss":
		return n + pick(" másodperc", " másodperce")
	case "m":
		return "egy" + pick(" perc", " perce")
	case "mm":
		return n + pick(" perc", " perce")
	case "h":
		return "egy" + pick(" óra", " órája")
	case "hh":
		return n + pick(" óra", " órája")
	case "d":
		return "egy" + pick(" nap", " napja")
	case "dd":
		return n + pick(" nap", " napja")
	case "M":
		return "egy" + pick(" hónap", " hónapja")
	case "MM":
		return n + pick(" hónap", " hónapja")
	case "y":
		return "egy" + pick(" év", " éve")
	case "yy":
		return n + pick(" év", " éve")
	}
	return ""
}

// huWeek is the calendar text for a weekday of this or last week.
func huWeek(m *moment.Moment, isFuture bool) string {
	prefix := "[múlt] "
	if isFuture {
		prefix = ""
	}
	ending := "undefined"
	if d := m.Day(); d >= 0 && d < 7 {
		ending = huWeekEndings[int(d)]
	}
	return prefix + "[" + ending + "] LT[-kor]"
}

// huCalendarNextWeek is calendar.nextWeek of hu.
func huCalendarNextWeek(m, _ *moment.Moment) string { return huWeek(m, true) }

// huCalendarLastWeek is calendar.lastWeek of hu.
func huCalendarLastWeek(m, _ *moment.Moment) string { return huWeek(m, false) }

// huIsPM is the isPM callback of hu.
func huIsPM(_ *moment.Locale, input string) bool { return js.ToLower(js.CharAt16(input, 1)) == "u" }

// huMeridiem is the meridiem callback of hu.
func huMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours < 12 {
		if isLower {
			return "de"
		}
		return "DE"
	}
	if isLower {
		return "du"
	}
	return "DU"
}
