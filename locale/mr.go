// moment.js locale configuration
// locale : Marathi [mr]
// author : Harshad Kale : https://github.com/kalehv
// author : Vivek Athalye : https://github.com/vnathalye

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("mr", func() {
		moment.DefineLocale("mr", js.Obj(
			"months", []string{"जानेवारी", "फेब्रुवारी", "मार्च", "एप्रिल", "मे", "जून", "जुलै", "ऑगस्ट", "सप्टेंबर", "ऑक्टोबर", "नोव्हेंबर", "डिसेंबर"},
			"monthsShort", []string{"जाने.", "फेब्रु.", "मार्च.", "एप्रि.", "मे.", "जून.", "जुलै.", "ऑग.", "सप्टें.", "ऑक्टो.", "नोव्हें.", "डिसें."},
			"monthsParseExact", true,
			"weekdays", []string{"रविवार", "सोमवार", "मंगळवार", "बुधवार", "गुरूवार", "शुक्रवार", "शनिवार"},
			"weekdaysShort", []string{"रवि", "सोम", "मंगळ", "बुध", "गुरू", "शुक्र", "शनि"},
			"weekdaysMin", []string{"र", "सो", "मं", "बु", "गु", "शु", "श"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm वाजता",
				"LTS", "A h:mm:ss वाजता",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm वाजता",
				"LLLL", "dddd, D MMMM YYYY, A h:mm वाजता",
			),
			"calendar", js.Obj(
				"sameDay", "[आज] LT",
				"nextDay", "[उद्या] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[काल] LT",
				"lastWeek", "[मागील] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%sमध्ये",
				"past", "%sपूर्वी",
				"s", mrRelativeTimeSeconds,
				"ss", mrRelativeTimeSeconds,
				"m", mrRelativeTimeSeconds,
				"mm", mrRelativeTimeSeconds,
				"h", mrRelativeTimeSeconds,
				"hh", mrRelativeTimeSeconds,
				"d", mrRelativeTimeSeconds,
				"dd", mrRelativeTimeSeconds,
				"M", mrRelativeTimeSeconds,
				"MM", mrRelativeTimeSeconds,
				"y", mrRelativeTimeSeconds,
				"yy", mrRelativeTimeSeconds,
			),
			"preparse", mrPreparse,
			"postformat", mrPostformat,
			"meridiemParse", jsre.MustCompile("पहाटे|सकाळी|दुपारी|सायंकाळी|रात्री", ""),
			"meridiemHour", mrMeridiemHour,
			"meridiem", mrMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var mrForms = map[string][2]string{
	"s": {"काही सेकंद", "काही सेकंदां"}, "ss": {"%d सेकंद", "%d सेकंदां"}, "m": {"एक मिनिट", "एका मिनिटा"},
	"mm": {"%d मिनिटे", "%d मिनिटां"}, "h": {"एक तास", "एका तासा"}, "hh": {"%d तास", "%d तासां"},
	"d": {"एक दिवस", "एका दिवसा"}, "dd": {"%d दिवस", "%d दिवसां"}, "M": {"एक महिना", "एका महिन्या"},
	"MM": {"%d महिने", "%d महिन्यां"}, "y": {"एक वर्ष", "एका वर्षा"}, "yy": {"%d वर्षे", "%d वर्षां"},
}

// mrRelativeTimeSeconds is relativeTimeMr of mr.
func mrRelativeTimeSeconds(number float64, withoutSuffix bool, key string, _ bool) string {
	forms := mrForms[key]
	output := forms[1]
	if withoutSuffix {
		output = forms[0]
	}
	return replaceCount(output, number)
}

var (
	mrPreparse   = hiPreparse
	mrPostformat = hiPostformat
)

// mrMeridiemHour is the meridiemHour callback of mr.
func mrMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "पहाटे", "सकाळी":
		return hour
	case "दुपारी", "सायंकाळी", "रात्री":
		if hourNumber(hour) >= 12 {
			return hour
		}
		return plus12(hour)
	}
	return js.Undef
}

// mrMeridiem is the meridiem callback of mr.
func mrMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour >= 0 && hour < 6:
		return "पहाटे"
	case hour < 12:
		return "सकाळी"
	case hour < 17:
		return "दुपारी"
	case hour < 20:
		return "सायंकाळी"
	}
	return "रात्री"
}
