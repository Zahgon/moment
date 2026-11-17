// moment.js locale configuration
// locale : Konkani Devanagari script [gom-deva]
// author : The Discoverer : https://github.com/WikiDiscoverer

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("gom-deva", func() {
		moment.DefineLocale("gom-deva", js.Obj(
			"months", js.Obj(
				"standalone", []string{"जानेवारी", "फेब्रुवारी", "मार्च", "एप्रील", "मे", "जून", "जुलय", "ऑगस्ट", "सप्टेंबर", "ऑक्टोबर", "नोव्हेंबर", "डिसेंबर"},
				"format", []string{"जानेवारीच्या", "फेब्रुवारीच्या", "मार्चाच्या", "एप्रीलाच्या", "मेयाच्या", "जूनाच्या", "जुलयाच्या", "ऑगस्टाच्या", "सप्टेंबराच्या", "ऑक्टोबराच्या", "नोव्हेंबराच्या", "डिसेंबराच्या"},
				"isFormat", jsre.MustCompile("MMMM(\\s)+D[oD]?", ""),
			),
			"monthsShort", []string{"जाने.", "फेब्रु.", "मार्च", "एप्री.", "मे", "जून", "जुल.", "ऑग.", "सप्टें.", "ऑक्टो.", "नोव्हें.", "डिसें."},
			"monthsParseExact", true,
			"weekdays", []string{"आयतार", "सोमार", "मंगळार", "बुधवार", "बिरेस्तार", "सुक्रार", "शेनवार"},
			"weekdaysShort", []string{"आयत.", "सोम.", "मंगळ.", "बुध.", "ब्रेस्त.", "सुक्र.", "शेन."},
			"weekdaysMin", []string{"आ", "सो", "मं", "बु", "ब्रे", "सु", "शे"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "A h:mm [वाजतां]",
				"LTS", "A h:mm:ss [वाजतां]",
				"L", "DD-MM-YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY A h:mm [वाजतां]",
				"LLLL", "dddd, MMMM Do, YYYY, A h:mm [वाजतां]",
				"llll", "ddd, D MMM YYYY, A h:mm [वाजतां]",
			),
			"calendar", js.Obj(
				"sameDay", "[आयज] LT",
				"nextDay", "[फाल्यां] LT",
				"nextWeek", "[फुडलो] dddd[,] LT",
				"lastDay", "[काल] LT",
				"lastWeek", "[फाटलो] dddd[,] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s",
				"past", "%s आदीं",
				"s", gomDevaRelativeTimeSeconds,
				"ss", gomDevaRelativeTimeSeconds,
				"m", gomDevaRelativeTimeSeconds,
				"mm", gomDevaRelativeTimeSeconds,
				"h", gomDevaRelativeTimeSeconds,
				"hh", gomDevaRelativeTimeSeconds,
				"d", gomDevaRelativeTimeSeconds,
				"dd", gomDevaRelativeTimeSeconds,
				"M", gomDevaRelativeTimeSeconds,
				"MM", gomDevaRelativeTimeSeconds,
				"y", gomDevaRelativeTimeSeconds,
				"yy", gomDevaRelativeTimeSeconds,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(वेर)", ""),
			"ordinal", gomDevaOrdinal,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 3.0,
			),
			"meridiemParse", jsre.MustCompile("राती|सकाळीं|दनपारां|सांजे", ""),
			"meridiemHour", gomDevaMeridiemHour,
			"meridiem", gomDevaMeridiem,
		))
	})
}

// gomDevaRelativeTimeSeconds is processRelativeTime of gom-deva.
func gomDevaRelativeTimeSeconds(number float64, _ bool, key string, isFuture bool) string {
	n := num(number)
	format := map[string][2]string{
		"s": {"थोडया सॅकंडांनी", "थोडे सॅकंड"}, "ss": {n + " सॅकंडांनी", n + " सॅकंड"},
		"m": {"एका मिणटान", "एक मिनूट"}, "mm": {n + " मिणटांनी", n + " मिणटां"},
		"h": {"एका वरान", "एक वर"}, "hh": {n + " वरांनी", n + " वरां"},
		"d": {"एका दिसान", "एक दीस"}, "dd": {n + " दिसांनी", n + " दीस"},
		"M": {"एका म्हयन्यान", "एक म्हयनो"}, "MM": {n + " म्हयन्यानी", n + " म्हयने"},
		"y": {"एका वर्सान", "एक वर्स"}, "yy": {n + " वर्सांनी", n + " वर्सां"},
	}
	if isFuture {
		return format[key][0]
	}
	return format[key][1]
}

// gomDevaOrdinal is the ordinal callback of gom-deva.
func gomDevaOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	if period == "D" {
		return str(number) + "वेर"
	}
	return number
}

// gomDevaMeridiemHour is the meridiemHour callback of gom-deva.
func gomDevaMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return konkaniMeridiemHour(hour, meridiem, "राती", "सकाळीं", "दनपारां", "सांजे")
}

// konkaniMeridiemHour is the meridiemHour of gom-deva and gom-latn for their night, morning, afternoon and
// evening words.
func konkaniMeridiemHour(hour any, meridiem, night, morning, afternoon, evening string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case night:
		if hourNumber(hour) < 4 {
			return hour
		}
		return plus12(hour)
	case morning:
		return hour
	case afternoon:
		if hourNumber(hour) > 12 {
			return hour
		}
		return plus12(hour)
	case evening:
		return plus12(hour)
	}
	return js.Undef
}

// gomDevaMeridiem is the meridiem callback of gom-deva.
func gomDevaMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "राती"
	case hour < 12:
		return "सकाळीं"
	case hour < 16:
		return "दनपारां"
	case hour < 20:
		return "सांजे"
	}
	return "राती"
}
