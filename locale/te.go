// moment.js locale configuration
// locale : Telugu [te]
// author : Krishna Chaitanya Thota : https://github.com/kcthota

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("te", func() {
		moment.DefineLocale("te", js.Obj(
			"months", []string{"జనవరి", "ఫిబ్రవరి", "మార్చి", "ఏప్రిల్", "మే", "జూన్", "జులై", "ఆగస్టు", "సెప్టెంబర్", "అక్టోబర్", "నవంబర్", "డిసెంబర్"},
			"monthsShort", []string{"జన.", "ఫిబ్ర.", "మార్చి", "ఏప్రి.", "మే", "జూన్", "జులై", "ఆగ.", "సెప్.", "అక్టో.", "నవ.", "డిసె."},
			"monthsParseExact", true,
			"weekdays", []string{"ఆదివారం", "సోమవారం", "మంగళవారం", "బుధవారం", "గురువారం", "శుక్రవారం", "శనివారం"},
			"weekdaysShort", []string{"ఆది", "సోమ", "మంగళ", "బుధ", "గురు", "శుక్ర", "శని"},
			"weekdaysMin", []string{"ఆ", "సో", "మం", "బు", "గు", "శు", "శ"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm",
				"LTS", "A h:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm",
				"LLLL", "dddd, D MMMM YYYY, A h:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[నేడు] LT",
				"nextDay", "[రేపు] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[నిన్న] LT",
				"lastWeek", "[గత] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s లో",
				"past", "%s క్రితం",
				"s", "కొన్ని క్షణాలు",
				"ss", "%d సెకన్లు",
				"m", "ఒక నిమిషం",
				"mm", "%d నిమిషాలు",
				"h", "ఒక గంట",
				"hh", "%d గంటలు",
				"d", "ఒక రోజు",
				"dd", "%d రోజులు",
				"M", "ఒక నెల",
				"MM", "%d నెలలు",
				"y", "ఒక సంవత్సరం",
				"yy", "%d సంవత్సరాలు",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}వ", ""),
			"ordinal", "%dవ",
			"meridiemParse", jsre.MustCompile("రాత్రి|ఉదయం|మధ్యాహ్నం|సాయంత్రం", ""),
			"meridiemHour", teMeridiemHour,
			"meridiem", teMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

// teMeridiemHour is the meridiemHour callback of te.
func teMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return dayPartHour(hour, meridiem, "రాత్రి", "ఉదయం", "మధ్యాహ్నం", "సాయంత్రం", 10)
}

// teMeridiem is the meridiem callback of te.
func teMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "రాత్రి"
	case hour < 10:
		return "ఉదయం"
	case hour < 17:
		return "మధ్యాహ్నం"
	case hour < 20:
		return "సాయంత్రం"
	}
	return "రాత్రి"
}
