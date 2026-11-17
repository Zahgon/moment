// moment.js locale configuration
// locale : Greek [el]
// author : Aggelos Karalias : https://github.com/mehiel

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("el", func() {
		moment.DefineLocale("el", js.Obj(
			"monthsNominativeEl", []string{"Ιανουάριος", "Φεβρουάριος", "Μάρτιος", "Απρίλιος", "Μάιος", "Ιούνιος", "Ιούλιος", "Αύγουστος", "Σεπτέμβριος", "Οκτώβριος", "Νοέμβριος", "Δεκέμβριος"},
			"monthsGenitiveEl", []string{"Ιανουαρίου", "Φεβρουαρίου", "Μαρτίου", "Απριλίου", "Μαΐου", "Ιουνίου", "Ιουλίου", "Αυγούστου", "Σεπτεμβρίου", "Οκτωβρίου", "Νοεμβρίου", "Δεκεμβρίου"},
			"months", elMonths,
			"monthsShort", []string{"Ιαν", "Φεβ", "Μαρ", "Απρ", "Μαϊ", "Ιουν", "Ιουλ", "Αυγ", "Σεπ", "Οκτ", "Νοε", "Δεκ"},
			"weekdays", []string{"Κυριακή", "Δευτέρα", "Τρίτη", "Τετάρτη", "Πέμπτη", "Παρασκευή", "Σάββατο"},
			"weekdaysShort", []string{"Κυρ", "Δευ", "Τρι", "Τετ", "Πεμ", "Παρ", "Σαβ"},
			"weekdaysMin", []string{"Κυ", "Δε", "Τρ", "Τε", "Πε", "Πα", "Σα"},
			"meridiem", elMeridiem,
			"isPM", elIsPM,
			"meridiemParse", jsre.MustCompile("[ΠΜ]\\.?Μ?\\.?", "i"),
			"longDateFormat", js.Obj(
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY h:mm A",
				"LLLL", "dddd, D MMMM YYYY h:mm A",
			),
			"calendarEl", js.Obj(
				"sameDay", "[Σήμερα {}] LT",
				"nextDay", "[Αύριο {}] LT",
				"nextWeek", "dddd [{}] LT",
				"lastDay", "[Χθες {}] LT",
				"lastWeek", elCalendarElLastWeek,
				"sameElse", "L",
			),
			"calendar", elCalendar,
			"relativeTime", js.Obj(
				"future", "σε %s",
				"past", "%s πριν",
				"s", "λίγα δευτερόλεπτα",
				"ss", "%d δευτερόλεπτα",
				"m", "ένα λεπτό",
				"mm", "%d λεπτά",
				"h", "μία ώρα",
				"hh", "%d ώρες",
				"d", "μία μέρα",
				"dd", "%d μέρες",
				"M", "ένας μήνας",
				"MM", "%d μήνες",
				"y", "ένας χρόνος",
				"yy", "%d χρόνια",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}η", ""),
			"ordinal", "%dη",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var elDayBefore = jsre.MustCompile(`D`, "")

// elMonths returns the genitive month name when a day number precedes MMMM in the format.
func elMonths(l *moment.Locale, m *moment.Moment, format any) any {
	if m == nil {
		return l.Prop("monthsNominativeEl")
	}
	list := "monthsNominativeEl"
	if f, ok := format.(string); ok && elDayBefore.Test(js.Substring16(f, 0, js.IndexOf16(f, "MMMM"))) {
		list = "monthsGenitiveEl"
	}
	names, _ := l.Prop(list).([]string)
	if i := m.Month(); i >= 0 && int(i) < len(names) {
		return names[int(i)]
	}
	return js.Undef
}

// elMeridiem is the meridiem callback of el.
func elMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours > 11 {
		if isLower {
			return "μμ"
		}
		return "ΜΜ"
	}
	if isLower {
		return "πμ"
	}
	return "ΠΜ"
}

// elIsPM is the isPM callback of el.
func elIsPM(_ *moment.Locale, input string) bool { return js.CharAt16(js.ToLower(input), 0) == "μ" }

// elCalendarElLastWeek is calendarEl.lastWeek of el.
func elCalendarElLastWeek(m, _ *moment.Moment) string {
	if m.Day() == 6 {
		return "[το προηγούμενο] dddd [{}] LT"
	}
	return "[την προηγούμενη] dddd [{}] LT"
}

// elCalendar fills the article placeholder of the calendar format: "στη" at one o'clock, "στις" otherwise.
func elCalendar(l *moment.Locale, key string, mom *moment.Moment, _ *moment.Moment) string {
	formats, _ := l.Prop("calendarEl").(*js.Object)
	var output string
	if formats != nil {
		switch v := formats.Get(key).(type) {
		case string:
			output = v
		case func(*moment.Moment, *moment.Moment) string:
			output = v(mom, nil)
		}
	}
	hours := js.NaN
	if mom != nil {
		hours = mom.Hours()
	}
	article := "στις"
	if js.Mod(hours, 12) == 1 {
		article = "στη"
	}
	return js.ReplaceFirst(output, "{}", article)
}
