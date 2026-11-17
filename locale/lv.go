// moment.js locale configuration
// locale : Latvian [lv]
// author : Kristaps Karlsons : https://github.com/skakri
// author : Jānis Elmeris : https://github.com/JanisE

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("lv", func() {
		moment.DefineLocale("lv", js.Obj(
			"months", []string{"janvāris", "februāris", "marts", "aprīlis", "maijs", "jūnijs", "jūlijs", "augusts", "septembris", "oktobris", "novembris", "decembris"},
			"monthsShort", []string{"jan", "feb", "mar", "apr", "mai", "jūn", "jūl", "aug", "sep", "okt", "nov", "dec"},
			"weekdays", []string{"svētdiena", "pirmdiena", "otrdiena", "trešdiena", "ceturtdiena", "piektdiena", "sestdiena"},
			"weekdaysShort", []string{"Sv", "P", "O", "T", "C", "Pk", "S"},
			"weekdaysMin", []string{"Sv", "P", "O", "T", "C", "Pk", "S"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY.",
				"LL", "YYYY. [gada] D. MMMM",
				"LLL", "YYYY. [gada] D. MMMM, HH:mm",
				"LLLL", "YYYY. [gada] D. MMMM, dddd, HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Šodien pulksten] LT",
				"nextDay", "[Rīt pulksten] LT",
				"nextWeek", "dddd [pulksten] LT",
				"lastDay", "[Vakar pulksten] LT",
				"lastWeek", "[Pagājušā] dddd [pulksten] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "pēc %s",
				"past", "pirms %s",
				"s", lvRelativeTimeSeconds,
				"ss", lvRelativeTimeSecondsN,
				"m", lvRelativeTimeMinute,
				"mm", lvRelativeTimeSecondsN,
				"h", lvRelativeTimeMinute,
				"hh", lvRelativeTimeSecondsN,
				"d", lvRelativeTimeMinute,
				"dd", lvRelativeTimeSecondsN,
				"M", lvRelativeTimeMinute,
				"MM", lvRelativeTimeSecondsN,
				"y", lvRelativeTimeMinute,
				"yy", lvRelativeTimeSecondsN,
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

var lvUnits = map[string][]string{
	"ss": {"sekundes", "sekundēm", "sekunde", "sekundes"}, "m": {"minūtes", "minūtēm", "minūte", "minūtes"},
	"mm": {"minūtes", "minūtēm", "minūte", "minūtes"}, "h": {"stundas", "stundām", "stunda", "stundas"},
	"hh": {"stundas", "stundām", "stunda", "stundas"}, "d": {"dienas", "dienām", "diena", "dienas"},
	"dd": {"dienas", "dienām", "diena", "dienas"}, "M": {"mēneša", "mēnešiem", "mēnesis", "mēneši"},
	"MM": {"mēneša", "mēnešiem", "mēnesis", "mēneši"}, "y": {"gada", "gadiem", "gads", "gadi"},
	"yy": {"gada", "gadiem", "gads", "gadi"},
}

// lvFormat picks the singular or plural form, nominative for a length of time and dative otherwise.
func lvFormat(forms []string, number float64, withoutSuffix bool) string {
	if len(forms) < 4 {
		return ""
	}
	singular := js.Mod(number, 10) == 1 && js.Mod(number, 100) != 11
	switch {
	case withoutSuffix && singular:
		return forms[2]
	case withoutSuffix:
		return forms[3]
	case singular:
		return forms[0]
	}
	return forms[1]
}

// lvRelativeTimeSeconds is relativeSeconds of lv.
func lvRelativeTimeSeconds(_ float64, withoutSuffix bool, _ string, _ bool) string {
	if withoutSuffix {
		return "dažas sekundes"
	}
	return "dažām sekundēm"
}

// lvRelativeTimeSecondsN is relativeTimeWithPlural of lv.
func lvRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, _ bool) string {
	return num(number) + " " + lvFormat(lvUnits[key], number, withoutSuffix)
}

// lvRelativeTimeMinute is relativeTimeWithSingular of lv.
func lvRelativeTimeMinute(number float64, withoutSuffix bool, key string, _ bool) string {
	return lvFormat(lvUnits[key], number, withoutSuffix)
}
