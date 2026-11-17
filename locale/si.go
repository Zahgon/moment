// moment.js locale configuration
// locale : Sinhalese [si]
// author : Sampath Sitinamaluwa : https://github.com/sampathsris

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("si", func() {
		moment.DefineLocale("si", js.Obj(
			"months", []string{"ජනවාරි", "පෙබරවාරි", "මාර්තු", "අප්‍රේල්", "මැයි", "ජූනි", "ජූලි", "අගෝස්තු", "සැප්තැම්බර්", "ඔක්තෝබර්", "නොවැම්බර්", "දෙසැම්බර්"},
			"monthsShort", []string{"ජන", "පෙබ", "මාර්", "අප්", "මැයි", "ජූනි", "ජූලි", "අගෝ", "සැප්", "ඔක්", "නොවැ", "දෙසැ"},
			"weekdays", []string{"ඉරිදා", "සඳුදා", "අඟහරුවාදා", "බදාදා", "බ්‍රහස්පතින්දා", "සිකුරාදා", "සෙනසුරාදා"},
			"weekdaysShort", []string{"ඉරි", "සඳු", "අඟ", "බදා", "බ්‍රහ", "සිකු", "සෙන"},
			"weekdaysMin", []string{"ඉ", "ස", "අ", "බ", "බ්‍ර", "සි", "සෙ"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "a h:mm",
				"LTS", "a h:mm:ss",
				"L", "YYYY/MM/DD",
				"LL", "YYYY MMMM D",
				"LLL", "YYYY MMMM D, a h:mm",
				"LLLL", "YYYY MMMM D [වැනි] dddd, a h:mm:ss",
			),
			"calendar", js.Obj(
				"sameDay", "[අද] LT[ට]",
				"nextDay", "[හෙට] LT[ට]",
				"nextWeek", "dddd LT[ට]",
				"lastDay", "[ඊයේ] LT[ට]",
				"lastWeek", "[පසුගිය] dddd LT[ට]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%sකින්",
				"past", "%sකට පෙර",
				"s", "තත්පර කිහිපය",
				"ss", "තත්පර %d",
				"m", "මිනිත්තුව",
				"mm", "මිනිත්තු %d",
				"h", "පැය",
				"hh", "පැය %d",
				"d", "දිනය",
				"dd", "දින %d",
				"M", "මාසය",
				"MM", "මාස %d",
				"y", "වසර",
				"yy", "වසර %d",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2} වැනි", ""),
			"ordinal", siOrdinal,
			"meridiemParse", jsre.MustCompile("පෙර වරු|පස් වරු|පෙ.ව|ප.ව.", ""),
			"isPM", siIsPM,
			"meridiem", siMeridiem,
		))
	})
}

// siOrdinal is the ordinal callback of si.
func siOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return str(number) + " වැනි"
}

// siIsPM is the isPM callback of si.
func siIsPM(_ *moment.Locale, input string) bool {
	return input == "ප.ව." || input == "පස් වරු"
}

// siMeridiem is the meridiem callback of si.
func siMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours > 11 {
		if isLower {
			return "ප.ව."
		}
		return "පස් වරු"
	}
	if isLower {
		return "පෙ.ව."
	}
	return "පෙර වරු"
}
