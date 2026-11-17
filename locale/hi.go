// moment.js locale configuration
// locale : Hindi [hi]
// author : Mayank Singhal : https://github.com/mayanksinghal

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("hi", func() {
		moment.DefineLocale("hi", js.Obj(
			"months", js.Obj(
				"format", []string{"जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितम्बर", "अक्टूबर", "नवम्बर", "दिसम्बर"},
				"standalone", []string{"जनवरी", "फरवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्टूबर", "नवंबर", "दिसंबर"},
			),
			"monthsShort", []string{"जन.", "फ़र.", "मार्च", "अप्रै.", "मई", "जून", "जुल.", "अग.", "सित.", "अक्टू.", "नव.", "दिस."},
			"weekdays", []string{"रविवार", "सोमवार", "मंगलवार", "बुधवार", "गुरूवार", "शुक्रवार", "शनिवार"},
			"weekdaysShort", []string{"रवि", "सोम", "मंगल", "बुध", "गुरू", "शुक्र", "शनि"},
			"weekdaysMin", []string{"र", "सो", "मं", "बु", "गु", "शु", "श"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm बजे",
				"LTS", "A h:mm:ss बजे",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm बजे",
				"LLLL", "dddd, D MMMM YYYY, A h:mm बजे",
			),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^जन", "i"),
				jsre.MustCompile("^फ़र|फर", "i"),
				jsre.MustCompile("^मार्च", "i"),
				jsre.MustCompile("^अप्रै", "i"),
				jsre.MustCompile("^मई", "i"),
				jsre.MustCompile("^जून", "i"),
				jsre.MustCompile("^जुल", "i"),
				jsre.MustCompile("^अग", "i"),
				jsre.MustCompile("^सितं|सित", "i"),
				jsre.MustCompile("^अक्टू", "i"),
				jsre.MustCompile("^नव|नवं", "i"),
				jsre.MustCompile("^दिसं|दिस", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^जन", "i"),
				jsre.MustCompile("^फ़र|फर", "i"),
				jsre.MustCompile("^मार्च", "i"),
				jsre.MustCompile("^अप्रै", "i"),
				jsre.MustCompile("^मई", "i"),
				jsre.MustCompile("^जून", "i"),
				jsre.MustCompile("^जुल", "i"),
				jsre.MustCompile("^अग", "i"),
				jsre.MustCompile("^सितं|सित", "i"),
				jsre.MustCompile("^अक्टू", "i"),
				jsre.MustCompile("^नव|नवं", "i"),
				jsre.MustCompile("^दिसं|दिस", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^जन", "i"),
				jsre.MustCompile("^फ़र", "i"),
				jsre.MustCompile("^मार्च", "i"),
				jsre.MustCompile("^अप्रै", "i"),
				jsre.MustCompile("^मई", "i"),
				jsre.MustCompile("^जून", "i"),
				jsre.MustCompile("^जुल", "i"),
				jsre.MustCompile("^अग", "i"),
				jsre.MustCompile("^सित", "i"),
				jsre.MustCompile("^अक्टू", "i"),
				jsre.MustCompile("^नव", "i"),
				jsre.MustCompile("^दिस", "i"),
			},
			"monthsRegex", jsre.MustCompile("^(जनवरी|जन\\.?|फ़रवरी|फरवरी|फ़र\\.?|मार्च?|अप्रैल|अप्रै\\.?|मई?|जून?|जुलाई|जुल\\.?|अगस्त|अग\\.?|सितम्बर|सितंबर|सित\\.?|अक्टूबर|अक्टू\\.?|नवम्बर|नवंबर|नव\\.?|दिसम्बर|दिसंबर|दिस\\.?)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(जनवरी|जन\\.?|फ़रवरी|फरवरी|फ़र\\.?|मार्च?|अप्रैल|अप्रै\\.?|मई?|जून?|जुलाई|जुल\\.?|अगस्त|अग\\.?|सितम्बर|सितंबर|सित\\.?|अक्टूबर|अक्टू\\.?|नवम्बर|नवंबर|नव\\.?|दिसम्बर|दिसंबर|दिस\\.?)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(जनवरी?|फ़रवरी|फरवरी?|मार्च?|अप्रैल?|मई?|जून?|जुलाई?|अगस्त?|सितम्बर|सितंबर|सित?\\.?|अक्टूबर|अक्टू\\.?|नवम्बर|नवंबर?|दिसम्बर|दिसंबर?)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(जन\\.?|फ़र\\.?|मार्च?|अप्रै\\.?|मई?|जून?|जुल\\.?|अग\\.?|सित\\.?|अक्टू\\.?|नव\\.?|दिस\\.?)", "i"),
			"calendar", js.Obj(
				"sameDay", "[आज] LT",
				"nextDay", "[कल] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[कल] LT",
				"lastWeek", "[पिछले] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s में",
				"past", "%s पहले",
				"s", "कुछ ही क्षण",
				"ss", "%d सेकंड",
				"m", "एक मिनट",
				"mm", "%d मिनट",
				"h", "एक घंटा",
				"hh", "%d घंटे",
				"d", "एक दिन",
				"dd", "%d दिन",
				"M", "एक महीने",
				"MM", "%d महीने",
				"y", "एक वर्ष",
				"yy", "%d वर्ष",
			),
			"preparse", hiPreparse,
			"postformat", hiPostformat,
			"meridiemParse", jsre.MustCompile("रात|सुबह|दोपहर|शाम", ""),
			"meridiemHour", hiMeridiemHour,
			"meridiem", hiMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	hiSymbols, hiNumbers = digitTables("०१२३४५६७८९")
	hiDigit              = jsre.MustCompile(`[१२३४५६७८९०]`, "g")
)

// hiPreparse reads Devanagari digits.
func hiPreparse(_ *moment.Locale, s string) string { return mapMatches(hiDigit, s, hiNumbers) }

// hiPostformat writes Devanagari digits.
func hiPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, hiSymbols) }

// hiMeridiemHour is the meridiemHour callback of hi.
func hiMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return dayPartHour(hour, meridiem, "रात", "सुबह", "दोपहर", "शाम", 10)
}

// hiMeridiem is the meridiem callback of hi.
func hiMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "रात"
	case hour < 10:
		return "सुबह"
	case hour < 17:
		return "दोपहर"
	case hour < 20:
		return "शाम"
	}
	return "रात"
}
