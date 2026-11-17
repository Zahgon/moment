// moment.js locale configuration
// locale : Urdu [ur]
// author : Sawood Alam : https://github.com/ibnesayeed
// author : Zack : https://github.com/ZackVision

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ur", func() {
		moment.DefineLocale("ur", js.Obj(
			"months", []string{"جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
			"monthsShort", []string{"جنوری", "فروری", "مارچ", "اپریل", "مئی", "جون", "جولائی", "اگست", "ستمبر", "اکتوبر", "نومبر", "دسمبر"},
			"weekdays", []string{"اتوار", "پیر", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
			"weekdaysShort", []string{"اتوار", "پیر", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
			"weekdaysMin", []string{"اتوار", "پیر", "منگل", "بدھ", "جمعرات", "جمعہ", "ہفتہ"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd، D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("صبح|شام", ""),
			"isPM", urIsPM,
			"meridiem", urMeridiem,
			"calendar", js.Obj(
				"sameDay", "[آج بوقت] LT",
				"nextDay", "[کل بوقت] LT",
				"nextWeek", "dddd [بوقت] LT",
				"lastDay", "[گذشتہ روز بوقت] LT",
				"lastWeek", "[گذشتہ] dddd [بوقت] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s بعد",
				"past", "%s قبل",
				"s", "چند سیکنڈ",
				"ss", "%d سیکنڈ",
				"m", "ایک منٹ",
				"mm", "%d منٹ",
				"h", "ایک گھنٹہ",
				"hh", "%d گھنٹے",
				"d", "ایک دن",
				"dd", "%d دن",
				"M", "ایک ماہ",
				"MM", "%d ماہ",
				"y", "ایک سال",
				"yy", "%d سال",
			),
			"preparse", urPreparse,
			"postformat", urPostformat,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// urIsPM is the isPM callback of ur.
func urIsPM(_ *moment.Locale, input string) bool { return input == "شام" }

// urMeridiem is the meridiem callback of ur.
func urMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "صبح"
	}
	return "شام"
}

var (
	urPreparse   = dvPreparse
	urPostformat = dvPostformat
)
