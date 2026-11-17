// moment.js locale configuration
// locale : Burmese [my]
// author : Squar team, mysquar.com
// author : David Rossellat : https://github.com/gholadr
// author : Tin Aung Lin : https://github.com/thanyawzinmin

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("my", func() {
		moment.DefineLocale("my", js.Obj(
			"months", []string{"ဇန်နဝါရီ", "ဖေဖော်ဝါရီ", "မတ်", "ဧပြီ", "မေ", "ဇွန်", "ဇူလိုင်", "သြဂုတ်", "စက်တင်ဘာ", "အောက်တိုဘာ", "နိုဝင်ဘာ", "ဒီဇင်ဘာ"},
			"monthsShort", []string{"ဇန်", "ဖေ", "မတ်", "ပြီ", "မေ", "ဇွန်", "လိုင်", "သြ", "စက်", "အောက်", "နို", "ဒီ"},
			"weekdays", []string{"တနင်္ဂနွေ", "တနင်္လာ", "အင်္ဂါ", "ဗုဒ္ဓဟူး", "ကြာသပတေး", "သောကြာ", "စနေ"},
			"weekdaysShort", []string{"နွေ", "လာ", "ဂါ", "ဟူး", "ကြာ", "သော", "နေ"},
			"weekdaysMin", []string{"နွေ", "လာ", "ဂါ", "ဟူး", "ကြာ", "သော", "နေ"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[ယနေ.] LT [မှာ]",
				"nextDay", "[မနက်ဖြန်] LT [မှာ]",
				"nextWeek", "dddd LT [မှာ]",
				"lastDay", "[မနေ.က] LT [မှာ]",
				"lastWeek", "[ပြီးခဲ့သော] dddd LT [မှာ]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "လာမည့် %s မှာ",
				"past", "လွန်ခဲ့သော %s က",
				"s", "စက္ကန်.အနည်းငယ်",
				"ss", "%d စက္ကန့်",
				"m", "တစ်မိနစ်",
				"mm", "%d မိနစ်",
				"h", "တစ်နာရီ",
				"hh", "%d နာရီ",
				"d", "တစ်ရက်",
				"dd", "%d ရက်",
				"M", "တစ်လ",
				"MM", "%d လ",
				"y", "တစ်နှစ်",
				"yy", "%d နှစ်",
			),
			"preparse", myPreparse,
			"postformat", myPostformat,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var (
	mySymbols, myNumbers = digitTables("၀၁၂၃၄၅၆၇၈၉")
	myDigit              = jsre.MustCompile(`[၁၂၃၄၅၆၇၈၉၀]`, "g")
)

// myPreparse reads Myanmar digits.
func myPreparse(_ *moment.Locale, s string) string { return mapMatches(myDigit, s, myNumbers) }

// myPostformat writes Myanmar digits.
func myPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, mySymbols) }
