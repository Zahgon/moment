// moment.js locale configuration
// locale : Bengali [bn]
// author : Kaushik Gandhi : https://github.com/kaushikgandhi

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("bn", func() {
		moment.DefineLocale("bn", js.Obj(
			"months", []string{"জানুয়ারি", "ফেব্রুয়ারি", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্টেম্বর", "অক্টোবর", "নভেম্বর", "ডিসেম্বর"},
			"monthsShort", []string{"জানু", "ফেব্রু", "মার্চ", "এপ্রিল", "মে", "জুন", "জুলাই", "আগস্ট", "সেপ্ট", "অক্টো", "নভে", "ডিসে"},
			"weekdays", []string{"রবিবার", "সোমবার", "মঙ্গলবার", "বুধবার", "বৃহস্পতিবার", "শুক্রবার", "শনিবার"},
			"weekdaysShort", []string{"রবি", "সোম", "মঙ্গল", "বুধ", "বৃহস্পতি", "শুক্র", "শনি"},
			"weekdaysMin", []string{"রবি", "সোম", "মঙ্গল", "বুধ", "বৃহ", "শুক্র", "শনি"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm সময়",
				"LTS", "A h:mm:ss সময়",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm সময়",
				"LLLL", "dddd, D MMMM YYYY, A h:mm সময়",
			),
			"calendar", js.Obj(
				"sameDay", "[আজ] LT",
				"nextDay", "[আগামীকাল] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[গতকাল] LT",
				"lastWeek", "[গত] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s পরে",
				"past", "%s আগে",
				"s", "কয়েক সেকেন্ড",
				"ss", "%d সেকেন্ড",
				"m", "এক মিনিট",
				"mm", "%d মিনিট",
				"h", "এক ঘন্টা",
				"hh", "%d ঘন্টা",
				"d", "এক দিন",
				"dd", "%d দিন",
				"M", "এক মাস",
				"MM", "%d মাস",
				"y", "এক বছর",
				"yy", "%d বছর",
			),
			"preparse", bnPreparse,
			"postformat", bnPostformat,
			"meridiemParse", jsre.MustCompile("রাত|সকাল|দুপুর|বিকাল|রাত", ""),
			"meridiemHour", bnMeridiemHour,
			"meridiem", bnMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	bengaliSymbols, bengaliNumbers = digitTables("০১২৩৪৫৬৭৮৯")
	bengaliDigit                   = jsre.MustCompile(`[১২৩৪৫৬৭৮৯০]`, "g")
	bnPreparse                     = bengaliPreparse
	bnPostformat                   = bengaliPostformat
)

// bengaliPreparse reads Bengali digits (bn and bn-bd).
func bengaliPreparse(_ *moment.Locale, s string) string {
	return mapMatches(bengaliDigit, s, bengaliNumbers)
}

// bengaliPostformat writes Bengali digits (bn and bn-bd).
func bengaliPostformat(_ *moment.Locale, s string) string {
	return mapMatches(asciiDigit, s, bengaliSymbols)
}

// bnMeridiemHour is the meridiemHour callback of bn.
func bnMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	if (meridiem == "রাত" && hourNumber(hour) >= 4) || (meridiem == "দুপুর" && hourNumber(hour) < 5) || meridiem == "বিকাল" {
		return plus12(hour)
	}
	return hour
}

// bnMeridiem is the meridiem callback of bn.
func bnMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "রাত"
	case hour < 10:
		return "সকাল"
	case hour < 17:
		return "দুপুর"
	case hour < 20:
		return "বিকাল"
	}
	return "রাত"
}
