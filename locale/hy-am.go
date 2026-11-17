// moment.js locale configuration
// locale : Armenian [hy-am]
// author : Armendarabyan : https://github.com/armendarabyan

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("hy-am", func() {
		moment.DefineLocale("hy-am", js.Obj(
			"months", js.Obj(
				"format", []string{"հունվարի", "փետրվարի", "մարտի", "ապրիլի", "մայիսի", "հունիսի", "հուլիսի", "օգոստոսի", "սեպտեմբերի", "հոկտեմբերի", "նոյեմբերի", "դեկտեմբերի"},
				"standalone", []string{"հունվար", "փետրվար", "մարտ", "ապրիլ", "մայիս", "հունիս", "հուլիս", "օգոստոս", "սեպտեմբեր", "հոկտեմբեր", "նոյեմբեր", "դեկտեմբեր"},
			),
			"monthsShort", []string{"հնվ", "փտր", "մրտ", "ապր", "մյս", "հնս", "հլս", "օգս", "սպտ", "հկտ", "նմբ", "դկտ"},
			"weekdays", []string{"կիրակի", "երկուշաբթի", "երեքշաբթի", "չորեքշաբթի", "հինգշաբթի", "ուրբաթ", "շաբաթ"},
			"weekdaysShort", []string{"կրկ", "երկ", "երք", "չրք", "հնգ", "ուրբ", "շբթ"},
			"weekdaysMin", []string{"կրկ", "երկ", "երք", "չրք", "հնգ", "ուրբ", "շբթ"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY թ.",
				"LLL", "D MMMM YYYY թ., HH:mm",
				"LLLL", "dddd, D MMMM YYYY թ., HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[այսօր] LT",
				"nextDay", "[վաղը] LT",
				"lastDay", "[երեկ] LT",
				"nextWeek", hyAmCalendarNextWeek,
				"lastWeek", hyAmCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s հետո",
				"past", "%s առաջ",
				"s", "մի քանի վայրկյան",
				"ss", "%d վայրկյան",
				"m", "րոպե",
				"mm", "%d րոպե",
				"h", "ժամ",
				"hh", "%d ժամ",
				"d", "օր",
				"dd", "%d օր",
				"M", "ամիս",
				"MM", "%d ամիս",
				"y", "տարի",
				"yy", "%d տարի",
			),
			"meridiemParse", jsre.MustCompile("գիշերվա|առավոտվա|ցերեկվա|երեկոյան", ""),
			"isPM", hyAmIsPM,
			"meridiem", hyAmMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}|\\d{1,2}-(ին|րդ)", ""),
			"ordinal", hyAmOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var hyAmPM = jsre.MustCompile(`^(ցերեկվա|երեկոյան)$`, "")

// hyAmCalendarNextWeek is calendar.nextWeek of hy-am.
func hyAmCalendarNextWeek(_, _ *moment.Moment) string { return "dddd [օրը ժամը] LT" }

// hyAmCalendarLastWeek is calendar.lastWeek of hy-am.
func hyAmCalendarLastWeek(_, _ *moment.Moment) string {
	return "[անցած] dddd [օրը ժամը] LT"
}

// hyAmIsPM is the isPM callback of hy-am.
func hyAmIsPM(_ *moment.Locale, input string) bool { return hyAmPM.Test(input) }

// hyAmMeridiem is the meridiem callback of hy-am.
func hyAmMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "գիշերվա"
	case hour < 12:
		return "առավոտվա"
	case hour < 17:
		return "ցերեկվա"
	}
	return "երեկոյան"
}

// hyAmOrdinal is the ordinal callback of hy-am.
func hyAmOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "DDD", "w", "W", "DDDo":
		if sameNumber(number, 1) {
			return str(number) + "-ին"
		}
		return str(number) + "-րդ"
	}
	return number
}
