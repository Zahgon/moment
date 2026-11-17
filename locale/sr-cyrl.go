// moment.js locale configuration
// locale : Serbian Cyrillic [sr-cyrl]
// author : Milan Janačković<milanjanackovic@gmail.com> : https://github.com/milan-j
// author : Stefan Crnjaković <stefan@hotmail.rs> : https://github.com/crnjakovic

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sr-cyrl", func() {
		moment.DefineLocale("sr-cyrl", js.Obj(
			"months", []string{"јануар", "фебруар", "март", "април", "мај", "јун", "јул", "август", "септембар", "октобар", "новембар", "децембар"},
			"monthsShort", []string{"јан.", "феб.", "мар.", "апр.", "мај", "јун", "јул", "авг.", "сеп.", "окт.", "нов.", "дец."},
			"monthsParseExact", true,
			"weekdays", []string{"недеља", "понедељак", "уторак", "среда", "четвртак", "петак", "субота"},
			"weekdaysShort", []string{"нед.", "пон.", "уто.", "сре.", "чет.", "пет.", "суб."},
			"weekdaysMin", []string{"не", "по", "ут", "ср", "че", "пе", "су"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "D. M. YYYY.",
				"LL", "D. MMMM YYYY.",
				"LLL", "D. MMMM YYYY. H:mm",
				"LLLL", "dddd, D. MMMM YYYY. H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[данас у] LT",
				"nextDay", "[сутра у] LT",
				"nextWeek", srCyrlCalendarNextWeek,
				"lastDay", "[јуче у] LT",
				"lastWeek", srCyrlCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "за %s",
				"past", "пре %s",
				"s", "неколико секунди",
				"ss", srCyrlRelativeTimeSecondsN,
				"m", srCyrlRelativeTimeSecondsN,
				"mm", srCyrlRelativeTimeSecondsN,
				"h", srCyrlRelativeTimeSecondsN,
				"hh", srCyrlRelativeTimeSecondsN,
				"d", srCyrlRelativeTimeSecondsN,
				"dd", srCyrlRelativeTimeSecondsN,
				"M", srCyrlRelativeTimeSecondsN,
				"MM", srCyrlRelativeTimeSecondsN,
				"y", srCyrlRelativeTimeSecondsN,
				"yy", srCyrlRelativeTimeSecondsN,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}\\.", ""),
			"ordinal", "%d.",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var srCyrlWords = map[string][]string{
	"ss": {"секунда", "секунде", "секунди"}, "m": {"један минут", "једног минута"}, "mm": {"минут", "минута", "минута"},
	"h": {"један сат", "једног сата"}, "hh": {"сат", "сата", "сати"}, "d": {"један дан", "једног дана"},
	"dd": {"дан", "дана", "дана"}, "M": {"један месец", "једног месеца"}, "MM": {"месец", "месеца", "месеци"},
	"y": {"једну годину", "једне године"}, "yy": {"годину", "године", "година"},
}

var srCyrlLastWeekDays = []string{"[прошле] [недеље] [у] LT", "[прошлог] [понедељка] [у] LT",
	"[прошлог] [уторка] [у] LT", "[прошле] [среде] [у] LT", "[прошлог] [четвртка] [у] LT",
	"[прошлог] [петка] [у] LT", "[прошле] [суботе] [у] LT"}

// srCyrlCalendarNextWeek is calendar.nextWeek of sr-cyrl.
func srCyrlCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[у] [недељу] [у] LT"
	case 3:
		return "[у] [среду] [у] LT"
	case 6:
		return "[у] [суботу] [у] LT"
	case 1, 2, 4, 5:
		return "[у] dddd [у] LT"
	}
	return ""
}

// srCyrlCalendarLastWeek is calendar.lastWeek of sr-cyrl.
func srCyrlCalendarLastWeek(m, _ *moment.Moment) string {
	return serbianLastWeek(m, srCyrlLastWeekDays)
}

// srCyrlRelativeTimeSecondsN is translator.translate of sr-cyrl.
func srCyrlRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, isFuture bool) string {
	return serbianTranslate(srCyrlWords, "једна година", "годину", " година", number, withoutSuffix, key, isFuture)
}
