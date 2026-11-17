// moment.js locale configuration
// locale : Bengali (Bangladesh) [bn-bd]
// author : Asraf Hossain Patoary : https://github.com/ashwoolford

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("bn-bd", func() {
		moment.DefineLocale("bn-bd", js.Obj(
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
			"preparse", bnBdPreparse,
			"postformat", bnBdPostformat,
			"meridiemParse", jsre.MustCompile("রাত|ভোর|সকাল|দুপুর|বিকাল|সন্ধ্যা|রাত", ""),
			"meridiemHour", bnBdMeridiemHour,
			"meridiem", bnBdMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

// bnBdMeridiemHour is the meridiemHour callback of bn-bd.
func bnBdMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "রাত":
		if hourNumber(hour) < 4 {
			return hour
		}
		return plus12(hour)
	case "ভোর", "সকাল":
		return hour
	case "দুপুর":
		if hourNumber(hour) >= 3 {
			return hour
		}
		return plus12(hour)
	case "বিকাল", "সন্ধ্যা":
		return plus12(hour)
	}
	return js.Undef
}

// bnBdMeridiem is the meridiem callback of bn-bd.
func bnBdMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "রাত"
	case hour < 6:
		return "ভোর"
	case hour < 12:
		return "সকাল"
	case hour < 15:
		return "দুপুর"
	case hour < 18:
		return "বিকাল"
	case hour < 20:
		return "সন্ধ্যা"
	}
	return "রাত"
}

var (
	bnBdPreparse   = bengaliPreparse
	bnBdPostformat = bengaliPostformat
)
