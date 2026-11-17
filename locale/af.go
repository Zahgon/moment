// moment.js locale configuration
// locale : Afrikaans [af]
// author : Werner Mollentze : https://github.com/wernerm

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("af", func() {
		moment.DefineLocale("af", js.Obj(
			"months", []string{"Januarie", "Februarie", "Maart", "April", "Mei", "Junie", "Julie", "Augustus", "September", "Oktober", "November", "Desember"},
			"monthsShort", []string{"Jan", "Feb", "Mrt", "Apr", "Mei", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Des"},
			"weekdays", []string{"Sondag", "Maandag", "Dinsdag", "Woensdag", "Donderdag", "Vrydag", "Saterdag"},
			"weekdaysShort", []string{"Son", "Maa", "Din", "Woe", "Don", "Vry", "Sat"},
			"weekdaysMin", []string{"So", "Ma", "Di", "Wo", "Do", "Vr", "Sa"},
			"meridiemParse", jsre.MustCompile("vm|nm", "i"),
			"isPM", afIsPM,
			"meridiem", afMeridiem,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Vandag om] LT",
				"nextDay", "[Môre om] LT",
				"nextWeek", "dddd [om] LT",
				"lastDay", "[Gister om] LT",
				"lastWeek", "[Laas] dddd [om] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "oor %s",
				"past", "%s gelede",
				"s", "'n paar sekondes",
				"ss", "%d sekondes",
				"m", "'n minuut",
				"mm", "%d minute",
				"h", "'n uur",
				"hh", "%d ure",
				"d", "'n dag",
				"dd", "%d dae",
				"M", "'n maand",
				"MM", "%d maande",
				"y", "'n jaar",
				"yy", "%d jaar",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(ste|de)", ""),
			"ordinal", afOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var afPM = jsre.MustCompile(`^nm$`, "i")

// afIsPM is the isPM callback of af.
func afIsPM(_ *moment.Locale, input string) bool { return afPM.Test(input) }

// afMeridiem is the meridiem callback of af.
func afMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours < 12 {
		if isLower {
			return "vm"
		}
		return "VM"
	}
	if isLower {
		return "nm"
	}
	return "NM"
}

// afOrdinal is the ordinal callback of af.
func afOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	if sameNumber(number, 1) || sameNumber(number, 8) || js.ToNumber(number) >= 20 {
		return str(number) + "ste"
	}
	return str(number) + "de"
}
