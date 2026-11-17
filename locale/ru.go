// moment.js locale configuration
// locale : Russian [ru]
// author : Viktorminator : https://github.com/Viktorminator
// author : Menelion Elensúle : https://github.com/Oire
// author : Коренберг Марк : https://github.com/socketpair

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ru", func() {
		moment.DefineLocale("ru", js.Obj(
			"months", js.Obj(
				"format", []string{"января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "ноября", "декабря"},
				"standalone", []string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"},
			),
			"monthsShort", js.Obj(
				"format", []string{"янв.", "февр.", "мар.", "апр.", "мая", "июня", "июля", "авг.", "сент.", "окт.", "нояб.", "дек."},
				"standalone", []string{"янв.", "февр.", "март", "апр.", "май", "июнь", "июль", "авг.", "сент.", "окт.", "нояб.", "дек."},
			),
			"weekdays", js.Obj(
				"standalone", []string{"воскресенье", "понедельник", "вторник", "среда", "четверг", "пятница", "суббота"},
				"format", []string{"воскресенье", "понедельник", "вторник", "среду", "четверг", "пятницу", "субботу"},
				"isFormat", jsre.MustCompile("\\[ ?[Вв] ?(?:прошлую|следующую|эту)? ?] ?dddd", ""),
			),
			"weekdaysShort", []string{"вс", "пн", "вт", "ср", "чт", "пт", "сб"},
			"weekdaysMin", []string{"вс", "пн", "вт", "ср", "чт", "пт", "сб"},
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^янв", "i"),
				jsre.MustCompile("^фев", "i"),
				jsre.MustCompile("^мар", "i"),
				jsre.MustCompile("^апр", "i"),
				jsre.MustCompile("^ма[йя]", "i"),
				jsre.MustCompile("^июн", "i"),
				jsre.MustCompile("^июл", "i"),
				jsre.MustCompile("^авг", "i"),
				jsre.MustCompile("^сен", "i"),
				jsre.MustCompile("^окт", "i"),
				jsre.MustCompile("^ноя", "i"),
				jsre.MustCompile("^дек", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^янв", "i"),
				jsre.MustCompile("^фев", "i"),
				jsre.MustCompile("^мар", "i"),
				jsre.MustCompile("^апр", "i"),
				jsre.MustCompile("^ма[йя]", "i"),
				jsre.MustCompile("^июн", "i"),
				jsre.MustCompile("^июл", "i"),
				jsre.MustCompile("^авг", "i"),
				jsre.MustCompile("^сен", "i"),
				jsre.MustCompile("^окт", "i"),
				jsre.MustCompile("^ноя", "i"),
				jsre.MustCompile("^дек", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^янв", "i"),
				jsre.MustCompile("^фев", "i"),
				jsre.MustCompile("^мар", "i"),
				jsre.MustCompile("^апр", "i"),
				jsre.MustCompile("^ма[йя]", "i"),
				jsre.MustCompile("^июн", "i"),
				jsre.MustCompile("^июл", "i"),
				jsre.MustCompile("^авг", "i"),
				jsre.MustCompile("^сен", "i"),
				jsre.MustCompile("^окт", "i"),
				jsre.MustCompile("^ноя", "i"),
				jsre.MustCompile("^дек", "i"),
			},
			"monthsRegex", jsre.MustCompile("^(январ[ья]|янв\\.?|феврал[ья]|февр?\\.?|марта?|мар\\.?|апрел[ья]|апр\\.?|ма[йя]|июн[ья]|июн\\.?|июл[ья]|июл\\.?|августа?|авг\\.?|сентябр[ья]|сент?\\.?|октябр[ья]|окт\\.?|ноябр[ья]|нояб?\\.?|декабр[ья]|дек\\.?)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(январ[ья]|янв\\.?|феврал[ья]|февр?\\.?|марта?|мар\\.?|апрел[ья]|апр\\.?|ма[йя]|июн[ья]|июн\\.?|июл[ья]|июл\\.?|августа?|авг\\.?|сентябр[ья]|сент?\\.?|октябр[ья]|окт\\.?|ноябр[ья]|нояб?\\.?|декабр[ья]|дек\\.?)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(январ[яь]|феврал[яь]|марта?|апрел[яь]|ма[яй]|июн[яь]|июл[яь]|августа?|сентябр[яь]|октябр[яь]|ноябр[яь]|декабр[яь])", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(янв\\.|февр?\\.|мар[т.]|апр\\.|ма[яй]|июн[ья.]|июл[ья.]|авг\\.|сент?\\.|окт\\.|нояб?\\.|дек\\.)", "i"),
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY г.",
				"LLL", "D MMMM YYYY г., H:mm",
				"LLLL", "dddd, D MMMM YYYY г., H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Сегодня, в] LT",
				"nextDay", "[Завтра, в] LT",
				"lastDay", "[Вчера, в] LT",
				"nextWeek", ruCalendarNextWeek,
				"lastWeek", ruCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "через %s",
				"past", "%s назад",
				"s", "несколько секунд",
				"ss", ruRelativeTimeSecondsN,
				"m", ruRelativeTimeSecondsN,
				"mm", ruRelativeTimeSecondsN,
				"h", "час",
				"hh", ruRelativeTimeSecondsN,
				"d", "день",
				"dd", ruRelativeTimeSecondsN,
				"w", "неделя",
				"ww", ruRelativeTimeSecondsN,
				"M", "месяц",
				"MM", ruRelativeTimeSecondsN,
				"y", "год",
				"yy", ruRelativeTimeSecondsN,
			),
			"meridiemParse", jsre.MustCompile("ночи|утра|дня|вечера", "i"),
			"isPM", ruIsPM,
			"meridiem", ruMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(й|го|я)", ""),
			"ordinal", ruOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var ruPM = jsre.MustCompile(`^(дня|вечера)$`, "")

// ruWeekCalendar is the calendar text for a weekday of the next or last week.
func ruWeekCalendar(m, now *moment.Moment, words [3]string) string {
	if now.Week() != m.Week() {
		switch m.Day() {
		case 0:
			return words[0]
		case 1, 2, 4:
			return words[1]
		case 3, 5, 6:
			return words[2]
		}
		return ""
	}
	if m.Day() == 2 {
		return "[Во] dddd, [в] LT"
	}
	return "[В] dddd, [в] LT"
}

// ruCalendarNextWeek is calendar.nextWeek of ru.
func ruCalendarNextWeek(m, now *moment.Moment) string {
	return ruWeekCalendar(m, now, [3]string{"[В следующее] dddd, [в] LT", "[В следующий] dddd, [в] LT",
		"[В следующую] dddd, [в] LT"})
}

// ruCalendarLastWeek is calendar.lastWeek of ru.
func ruCalendarLastWeek(m, now *moment.Moment) string {
	return ruWeekCalendar(m, now, [3]string{"[В прошлое] dddd, [в] LT", "[В прошлый] dddd, [в] LT",
		"[В прошлую] dddd, [в] LT"})
}

// ruRelativeTimeSecondsN is relativeTimeWithPlural of ru.
func ruRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, _ bool) string {
	pick := func(without, with string) string {
		if withoutSuffix {
			return without
		}
		return with
	}
	format := map[string]string{
		"ss": pick("секунда_секунды_секунд", "секунду_секунды_секунд"),
		"mm": pick("минута_минуты_минут", "минуту_минуты_минут"),
		"hh": "час_часа_часов", "dd": "день_дня_дней", "ww": "неделя_недели_недель", "MM": "месяц_месяца_месяцев",
		"yy": "год_года_лет",
	}
	if key == "m" {
		return pick("минута", "минуту")
	}
	return num(number) + " " + slavicPlural(js.Split(format[key], "_"), number)
}

// ruIsPM is the isPM callback of ru.
func ruIsPM(_ *moment.Locale, input string) bool { return ruPM.Test(input) }

// ruMeridiem is the meridiem callback of ru.
func ruMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "ночи"
	case hour < 12:
		return "утра"
	case hour < 17:
		return "дня"
	}
	return "вечера"
}

// ruOrdinal is the ordinal callback of ru.
func ruOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "M", "d", "DDD":
		return str(number) + "-й"
	case "D":
		return str(number) + "-го"
	case "w", "W":
		return str(number) + "-я"
	}
	return number
}
