// moment.js locale configuration
// locale : Cambodian [km]
// author : Kruy Vanna : https://github.com/kruyvanna

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("km", func() {
		moment.DefineLocale("km", js.Obj(
			"months", []string{"មករា", "កុម្ភៈ", "មីនា", "មេសា", "ឧសភា", "មិថុនា", "កក្កដា", "សីហា", "កញ្ញា", "តុលា", "វិច្ឆិកា", "ធ្នូ"},
			"monthsShort", []string{"មករា", "កុម្ភៈ", "មីនា", "មេសា", "ឧសភា", "មិថុនា", "កក្កដា", "សីហា", "កញ្ញា", "តុលា", "វិច្ឆិកា", "ធ្នូ"},
			"weekdays", []string{"អាទិត្យ", "ច័ន្ទ", "អង្គារ", "ពុធ", "ព្រហស្បតិ៍", "សុក្រ", "សៅរ៍"},
			"weekdaysShort", []string{"អា", "ច", "អ", "ព", "ព្រ", "សុ", "ស"},
			"weekdaysMin", []string{"អា", "ច", "អ", "ព", "ព្រ", "សុ", "ស"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("ព្រឹក|ល្ងាច", ""),
			"isPM", kmIsPM,
			"meridiem", kmMeridiem,
			"calendar", js.Obj(
				"sameDay", "[ថ្ងៃនេះ ម៉ោង] LT",
				"nextDay", "[ស្អែក ម៉ោង] LT",
				"nextWeek", "dddd [ម៉ោង] LT",
				"lastDay", "[ម្សិលមិញ ម៉ោង] LT",
				"lastWeek", "dddd [សប្តាហ៍មុន] [ម៉ោង] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%sទៀត",
				"past", "%sមុន",
				"s", "ប៉ុន្មានវិនាទី",
				"ss", "%d វិនាទី",
				"m", "មួយនាទី",
				"mm", "%d នាទី",
				"h", "មួយម៉ោង",
				"hh", "%d ម៉ោង",
				"d", "មួយថ្ងៃ",
				"dd", "%d ថ្ងៃ",
				"M", "មួយខែ",
				"MM", "%d ខែ",
				"y", "មួយឆ្នាំ",
				"yy", "%d ឆ្នាំ",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("ទី\\d{1,2}", ""),
			"ordinal", "ទី%d",
			"preparse", kmPreparse,
			"postformat", kmPostformat,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var (
	kmSymbols, kmNumbers = digitTables("០១២៣៤៥៦៧៨៩")
	kmDigit              = jsre.MustCompile(`[១២៣៤៥៦៧៨៩០]`, "g")
)

// kmIsPM is the isPM callback of km.
func kmIsPM(_ *moment.Locale, input string) bool { return input == "ល្ងាច" }

// kmMeridiem is the meridiem callback of km.
func kmMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "ព្រឹក"
	}
	return "ល្ងាច"
}

// kmPreparse reads Khmer digits.
func kmPreparse(_ *moment.Locale, s string) string { return mapMatches(kmDigit, s, kmNumbers) }

// kmPostformat writes Khmer digits.
func kmPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, kmSymbols) }
