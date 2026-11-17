// moment.js locale configuration
// locale : Mongolian [mn]
// author : Javkhlantugs Nyamdorj : https://github.com/javkhaanj7

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("mn", func() {
		moment.DefineLocale("mn", js.Obj(
			"months", []string{"Нэгдүгээр сар", "Хоёрдугаар сар", "Гуравдугаар сар", "Дөрөвдүгээр сар", "Тавдугаар сар", "Зургадугаар сар", "Долдугаар сар", "Наймдугаар сар", "Есдүгээр сар", "Аравдугаар сар", "Арван нэгдүгээр сар", "Арван хоёрдугаар сар"},
			"monthsShort", []string{"1 сар", "2 сар", "3 сар", "4 сар", "5 сар", "6 сар", "7 сар", "8 сар", "9 сар", "10 сар", "11 сар", "12 сар"},
			"monthsParseExact", true,
			"weekdays", []string{"Ням", "Даваа", "Мягмар", "Лхагва", "Пүрэв", "Баасан", "Бямба"},
			"weekdaysShort", []string{"Ням", "Дав", "Мяг", "Лха", "Пүр", "Баа", "Бям"},
			"weekdaysMin", []string{"Ня", "Да", "Мя", "Лх", "Пү", "Ба", "Бя"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY-MM-DD",
				"LL", "YYYY оны MMMMын D",
				"LLL", "YYYY оны MMMMын D HH:mm",
				"LLLL", "dddd, YYYY оны MMMMын D HH:mm",
			),
			"meridiemParse", jsre.MustCompile("ҮӨ|ҮХ", "i"),
			"isPM", mnIsPM,
			"meridiem", mnMeridiem,
			"calendar", js.Obj(
				"sameDay", "[Өнөөдөр] LT",
				"nextDay", "[Маргааш] LT",
				"nextWeek", "[Ирэх] dddd LT",
				"lastDay", "[Өчигдөр] LT",
				"lastWeek", "[Өнгөрсөн] dddd LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s дараа",
				"past", "%s өмнө",
				"s", mnRelativeTimeSeconds,
				"ss", mnRelativeTimeSeconds,
				"m", mnRelativeTimeSeconds,
				"mm", mnRelativeTimeSeconds,
				"h", mnRelativeTimeSeconds,
				"hh", mnRelativeTimeSeconds,
				"d", mnRelativeTimeSeconds,
				"dd", mnRelativeTimeSeconds,
				"M", mnRelativeTimeSeconds,
				"MM", mnRelativeTimeSeconds,
				"y", mnRelativeTimeSeconds,
				"yy", mnRelativeTimeSeconds,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2} өдөр", ""),
			"ordinal", mnOrdinal,
		))
	})
}

// mnIsPM is the isPM callback of mn.
func mnIsPM(_ *moment.Locale, input string) bool { return input == "ҮХ" }

// mnMeridiem is the meridiem callback of mn.
func mnMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "ҮӨ"
	}
	return "ҮХ"
}

// mnRelativeTimeSeconds is translate of mn.
func mnRelativeTimeSeconds(number float64, withoutSuffix bool, key string, _ bool) string {
	pick := func(a, b string) string {
		if withoutSuffix {
			return a
		}
		return b
	}
	switch key {
	case "s":
		return pick("хэдхэн секунд", "хэдхэн секундын")
	case "ss":
		return num(number) + pick(" секунд", " секундын")
	case "m", "mm":
		return num(number) + pick(" минут", " минутын")
	case "h", "hh":
		return num(number) + pick(" цаг", " цагийн")
	case "d", "dd":
		return num(number) + pick(" өдөр", " өдрийн")
	case "M", "MM":
		return num(number) + pick(" сар", " сарын")
	case "y", "yy":
		return num(number) + pick(" жил", " жилийн")
	}
	return num(number)
}

// mnOrdinal is the ordinal callback of mn.
func mnOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "d", "D", "DDD":
		return str(number) + " өдөр"
	}
	return number
}
