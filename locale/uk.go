// moment.js locale configuration
// locale : Ukrainian [uk]
// author : zemlanin : https://github.com/zemlanin
// Author : Menelion Elensúle : https://github.com/Oire
// Author : Oleksandr Trukhnii : https://github.com/JohnJunior

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("uk", func() {
		moment.DefineLocale("uk", js.Obj(
			"months", js.Obj(
				"format", []string{"січня", "лютого", "березня", "квітня", "травня", "червня", "липня", "серпня", "вересня", "жовтня", "листопада", "грудня"},
				"standalone", []string{"січень", "лютий", "березень", "квітень", "травень", "червень", "липень", "серпень", "вересень", "жовтень", "листопад", "грудень"},
			),
			"monthsShort", []string{"січ", "лют", "бер", "квіт", "трав", "черв", "лип", "серп", "вер", "жовт", "лист", "груд"},
			"weekdays", ukWeekdays,
			"weekdaysShort", []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
			"weekdaysMin", []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY р.",
				"LLL", "D MMMM YYYY р., HH:mm",
				"LLLL", "dddd, D MMMM YYYY р., HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", ukCalendarSameDay,
				"nextDay", ukCalendarNextDay,
				"lastDay", ukCalendarLastDay,
				"nextWeek", ukCalendarNextWeek,
				"lastWeek", ukCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "за %s",
				"past", "%s тому",
				"s", "декілька секунд",
				"ss", ukRelativeTimeSecondsN,
				"m", ukRelativeTimeSecondsN,
				"mm", ukRelativeTimeSecondsN,
				"h", "годину",
				"hh", ukRelativeTimeSecondsN,
				"d", "день",
				"dd", ukRelativeTimeSecondsN,
				"M", "місяць",
				"MM", ukRelativeTimeSecondsN,
				"y", "рік",
				"yy", ukRelativeTimeSecondsN,
			),
			"meridiemParse", jsre.MustCompile("ночі|ранку|дня|вечора", ""),
			"isPM", ukIsPM,
			"meridiem", ukMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(й|го)", ""),
			"ordinal", ukOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var (
	ukWeekdayCases = map[string][]string{
		"nominative": {"неділя", "понеділок", "вівторок", "середа", "четвер", "пʼятниця", "субота"},
		"accusative": {"неділю", "понеділок", "вівторок", "середу", "четвер", "пʼятницю", "суботу"},
		"genitive":   {"неділі", "понеділка", "вівторка", "середи", "четверга", "пʼятниці", "суботи"},
	}
	ukAccusative = jsre.MustCompile(`(\[[ВвУу]\]) ?dddd`, "")
	ukGenitive   = jsre.MustCompile(`\[?(?:минулої|наступної)? ?\] ?dddd`, "")
	ukPM         = jsre.MustCompile(`^(дня|вечора)$`, "")
)

// ukWeekdays is weekdaysCaseReplace of uk: the accusative after "в"/"у", the genitive after "минулої" or
// "наступної", the nominative otherwise.
func ukWeekdays(_ *moment.Locale, m *moment.Moment, format any) any {
	nominative := ukWeekdayCases["nominative"]
	if b, ok := format.(bool); ok && b && m == nil {
		return append(append([]string{}, nominative[1:7]...), nominative[0])
	}
	if m == nil {
		return nominative
	}
	f := js.ToString(format)
	nounCase := "nominative"
	if ukAccusative.Test(f) {
		nounCase = "accusative"
	} else if ukGenitive.Test(f) {
		nounCase = "genitive"
	}
	if d := m.Day(); d >= 0 && d < 7 {
		return ukWeekdayCases[nounCase][int(d)]
	}
	return js.Undef
}

// ukHours appends "о" (or "об" at eleven) and the time to a calendar prefix.
func ukHours(prefix string) func(*moment.Moment, *moment.Moment) string {
	return func(m, _ *moment.Moment) string {
		suffix := ""
		if m.Hours() == 11 {
			suffix = "б"
		}
		return prefix + "о" + suffix + "] LT"
	}
}

var (
	ukCalendarSameDay  = ukHours("[Сьогодні ")
	ukCalendarNextDay  = ukHours("[Завтра ")
	ukCalendarLastDay  = ukHours("[Вчора ")
	ukCalendarNextWeek = ukHours("[У] dddd [")
)

// ukCalendarLastWeek is calendar.lastWeek of uk.
func ukCalendarLastWeek(m, now *moment.Moment) string {
	switch m.Day() {
	case 0, 3, 5, 6:
		return ukHours("[Минулої] dddd [")(m, now)
	case 1, 2, 4:
		return ukHours("[Минулого] dddd [")(m, now)
	}
	return ""
}

// ukRelativeTimeSecondsN is relativeTimeWithPlural of uk.
func ukRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, _ bool) string {
	pick := func(without, with string) string {
		if withoutSuffix {
			return without
		}
		return with
	}
	format := map[string]string{
		"ss": pick("секунда_секунди_секунд", "секунду_секунди_секунд"),
		"mm": pick("хвилина_хвилини_хвилин", "хвилину_хвилини_хвилин"),
		"hh": pick("година_години_годин", "годину_години_годин"),
		"dd": "день_дні_днів", "MM": "місяць_місяці_місяців", "yy": "рік_роки_років",
	}
	switch key {
	case "m":
		return pick("хвилина", "хвилину")
	case "h":
		return pick("година", "годину")
	}
	return num(number) + " " + slavicPlural(js.Split(format[key], "_"), number)
}

// ukIsPM is the isPM callback of uk.
func ukIsPM(_ *moment.Locale, input string) bool { return ukPM.Test(input) }

// ukMeridiem is the meridiem callback of uk.
func ukMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "ночі"
	case hour < 12:
		return "ранку"
	case hour < 17:
		return "дня"
	}
	return "вечора"
}

// ukOrdinal is the ordinal callback of uk.
func ukOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "M", "d", "DDD", "w", "W":
		return str(number) + "-й"
	case "D":
		return str(number) + "-го"
	}
	return number
}
