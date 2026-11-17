// moment.js locale configuration
// locale : Belarusian [be]
// author : Dmitry Demidov : https://github.com/demidov91
// author: Praleska: http://praleska.pro/
// Author : Menelion Elensúle : https://github.com/Oire

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("be", func() {
		moment.DefineLocale("be", js.Obj(
			"months", js.Obj(
				"format", []string{"студзеня", "лютага", "сакавіка", "красавіка", "траўня", "чэрвеня", "ліпеня", "жніўня", "верасня", "кастрычніка", "лістапада", "снежня"},
				"standalone", []string{"студзень", "люты", "сакавік", "красавік", "травень", "чэрвень", "ліпень", "жнівень", "верасень", "кастрычнік", "лістапад", "снежань"},
			),
			"monthsShort", []string{"студ", "лют", "сак", "крас", "трав", "чэрв", "ліп", "жнів", "вер", "каст", "ліст", "снеж"},
			"weekdays", js.Obj(
				"format", []string{"нядзелю", "панядзелак", "аўторак", "сераду", "чацвер", "пятніцу", "суботу"},
				"standalone", []string{"нядзеля", "панядзелак", "аўторак", "серада", "чацвер", "пятніца", "субота"},
				"isFormat", jsre.MustCompile("\\[ ?[Ууў] ?(?:мінулую|наступную)? ?\\] ?dddd", ""),
			),
			"weekdaysShort", []string{"нд", "пн", "ат", "ср", "чц", "пт", "сб"},
			"weekdaysMin", []string{"нд", "пн", "ат", "ср", "чц", "пт", "сб"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY г.",
				"LLL", "D MMMM YYYY г., HH:mm",
				"LLLL", "dddd, D MMMM YYYY г., HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Сёння ў] LT",
				"nextDay", "[Заўтра ў] LT",
				"lastDay", "[Учора ў] LT",
				"nextWeek", beCalendarNextWeek,
				"lastWeek", beCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "праз %s",
				"past", "%s таму",
				"s", "некалькі секунд",
				"m", beRelativeTimeMinute,
				"mm", beRelativeTimeMinute,
				"h", beRelativeTimeMinute,
				"hh", beRelativeTimeMinute,
				"d", "дзень",
				"dd", beRelativeTimeMinute,
				"M", "месяц",
				"MM", beRelativeTimeMinute,
				"y", "год",
				"yy", beRelativeTimeMinute,
			),
			"meridiemParse", jsre.MustCompile("ночы|раніцы|дня|вечара", ""),
			"isPM", beIsPM,
			"meridiem", beMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(і|ы|га)", ""),
			"ordinal", beOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var bePM = jsre.MustCompile(`^(дня|вечара)$`, "")

// beRelativeTimeMinute is relativeTimeWithPlural of be.
func beRelativeTimeMinute(number float64, withoutSuffix bool, key string, _ bool) string {
	pick := func(without, with string) string {
		if withoutSuffix {
			return without
		}
		return with
	}
	format := map[string]string{
		"ss": pick("секунда_секунды_секунд", "секунду_секунды_секунд"),
		"mm": pick("хвіліна_хвіліны_хвілін", "хвіліну_хвіліны_хвілін"),
		"hh": pick("гадзіна_гадзіны_гадзін", "гадзіну_гадзіны_гадзін"),
		"dd": "дзень_дні_дзён", "MM": "месяц_месяцы_месяцаў", "yy": "год_гады_гадоў",
	}
	switch key {
	case "m":
		return pick("хвіліна", "хвіліну")
	case "h":
		return pick("гадзіна", "гадзіну")
	}
	return num(number) + " " + slavicPlural(js.Split(format[key], "_"), number)
}

// beCalendarNextWeek is calendar.nextWeek of be.
func beCalendarNextWeek(_, _ *moment.Moment) string { return "[У] dddd [ў] LT" }

// beCalendarLastWeek is calendar.lastWeek of be.
func beCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0, 3, 5, 6:
		return "[У мінулую] dddd [ў] LT"
	case 1, 2, 4:
		return "[У мінулы] dddd [ў] LT"
	}
	return ""
}

// beIsPM is the isPM callback of be.
func beIsPM(_ *moment.Locale, input string) bool { return bePM.Test(input) }

// beMeridiem is the meridiem callback of be.
func beMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "ночы"
	case hour < 12:
		return "раніцы"
	case hour < 17:
		return "дня"
	}
	return "вечара"
}

// beOrdinal is the ordinal callback of be.
func beOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	n := js.ToNumber(number)
	switch period {
	case "M", "d", "DDD", "w", "W":
		if (js.Mod(n, 10) == 2 || js.Mod(n, 10) == 3) && js.Mod(n, 100) != 12 && js.Mod(n, 100) != 13 {
			return str(number) + "-і"
		}
		return str(number) + "-ы"
	case "D":
		return str(number) + "-га"
	}
	return number
}
