// moment.js locale configuration
// locale : Amharic (Ethiopia) [am-et]
// author : Tekle Ayele : https://github.com/tekleayele

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("am-et", func() {
		moment.DefineLocale("am-et", js.Obj(
			"months", []string{"ጃንዩወሪ", "ፌብሩወሪ", "ማርች", "ኤፕሪል", "ሜይ", "ጁን", "ጁላይ", "ኦገስት", "ሴፕቴምበር", "ኦክቶበር", "ኖቬምበር", "ዲሴምበር"},
			"monthsShort", []string{"ጃን", "ፌብ", "ማር", "ኤፕር", "ሜይ", "ጁን", "ጁላይ", "ኦገ", "ሴፕ", "ኦክቶ", "ኖቬ", "ዲሴ"},
			"weekdays", []string{"እሑድ", "ሰኞ", "ማክሰኞ", "እሮብ", "ሀሙስ", "ዓርብ", "ቅዳሜ"},
			"weekdaysShort", []string{"እሑ", "ሰኞ", "ማክ", "እሮ", "ሀሙ", "ዓር", "ቅዳ"},
			"weekdaysMin", []string{"እሁ", "ሰ", "ማ", "እ", "ሀ", "ዓ", "ቅ"},
			"meridiemParse", jsre.MustCompile("ጥዋት|ከሰዓት", ""),
			"isPM", amEtIsPM,
			"meridiem", amEtMeridiem,
			"longDateFormat", js.Obj(
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY h:mm A",
				"LLLL", "dddd, D MMMM YYYY h:mm A",
			),
			"calendar", js.Obj(
				"sameDay", "[ዛሬ በ] LT",
				"nextDay", "[ነገ በ] LT",
				"nextWeek", "dddd [በ] LT",
				"lastDay", "[ትናንትና በ] LT",
				"lastWeek", "[ያለፈው ሳምንት] dddd [በ] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s ውስጥ",
				"past", "%s በፊት",
				"s", "ጥቂት ሰከንዶች",
				"ss", "%d ሰከንዶች",
				"m", "አንድ ደቂቃ",
				"mm", "%d ደቂቃዎች",
				"h", "አንድ ሰአት",
				"hh", "%d ሰአታት",
				"d", "አንድ ቀን",
				"dd", "%d ቀናት",
				"M", "አንድ ወር",
				"MM", "%d ወራት",
				"y", "አንድ ዓመት",
				"yy", "%d ዓመታት",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(ኛ)", ""),
			"ordinal", amEtOrdinal,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

// amEtIsPM is the isPM callback of am-et.
func amEtIsPM(_ *moment.Locale, input string) bool { return input == "ከሰዓት" }

// amEtMeridiem is the meridiem callback of am-et.
func amEtMeridiem(_ *moment.Locale, hours, _ float64, _ bool) string {
	if hours < 12 {
		return "ጥዋት"
	}
	return "ከሰዓት"
}

// amEtOrdinal is the ordinal callback of am-et.
func amEtOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return str(number) + "ኛ"
}
