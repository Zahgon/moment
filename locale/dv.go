// moment.js locale configuration
// locale : Maldivian [dv]
// author : Jawish Hameed : https://github.com/jawish

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("dv", func() {
		moment.DefineLocale("dv", js.Obj(
			"months", []string{"ޖެނުއަރީ", "ފެބްރުއަރީ", "މާރިޗު", "އޭޕްރީލު", "މޭ", "ޖޫން", "ޖުލައި", "އޯގަސްޓު", "ސެޕްޓެމްބަރު", "އޮކްޓޯބަރު", "ނޮވެމްބަރު", "ޑިސެމްބަރު"},
			"monthsShort", []string{"ޖެނުއަރީ", "ފެބްރުއަރީ", "މާރިޗު", "އޭޕްރީލު", "މޭ", "ޖޫން", "ޖުލައި", "އޯގަސްޓު", "ސެޕްޓެމްބަރު", "އޮކްޓޯބަރު", "ނޮވެމްބަރު", "ޑިސެމްބަރު"},
			"weekdays", []string{"އާދިއްތަ", "ހޯމަ", "އަންގާރަ", "ބުދަ", "ބުރާސްފަތި", "ހުކުރު", "ހޮނިހިރު"},
			"weekdaysShort", []string{"އާދިއްތަ", "ހޯމަ", "އަންގާރަ", "ބުދަ", "ބުރާސްފަތި", "ހުކުރު", "ހޮނިހިރު"},
			"weekdaysMin", []string{"އާދި", "ހޯމަ", "އަން", "ބުދަ", "ބުރާ", "ހުކު", "ހޮނި"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "D/M/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("މކ|މފ", ""),
			"isPM", dvIsPM,
			"meridiem", dvMeridiem,
			"calendar", js.Obj(
				"sameDay", "[މިއަދު] LT",
				"nextDay", "[މާދަމާ] LT",
				"nextWeek", "dddd LT",
				"lastDay", "[އިއްޔެ] LT",
				"lastWeek", "[ފާއިތުވި] dddd LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "ތެރޭގައި %s",
				"past", "ކުރިން %s",
				"s", "ސިކުންތުކޮޅެއް",
				"ss", "d% ސިކުންތު",
				"m", "މިނިޓެއް",
				"mm", "މިނިޓު %d",
				"h", "ގަޑިއިރެއް",
				"hh", "ގަޑިއިރު %d",
				"d", "ދުވަހެއް",
				"dd", "ދުވަސް %d",
				"M", "މަހެއް",
				"MM", "މަސް %d",
				"y", "އަހަރެއް",
				"yy", "އަހަރު %d",
			),
			"preparse", dvPreparse,
			"postformat", dvPostformat,
			"week", js.Obj(
				"dow", 7.0,
				"doy", 12.0,
			),
		))
	})
}

// dvIsPM is the isPM callback of dv.
func dvIsPM(_ *moment.Locale, input string) bool { return input == "މފ" }

// dvMeridiem is the meridiem callback of dv.
func dvMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "މކ"
	}
	return "މފ"
}

// dvPreparse reads the Arabic comma.
func dvPreparse(_ *moment.Locale, s string) string { return arabicComma.Replace(s, ",") }

// dvPostformat writes the Arabic comma.
func dvPostformat(_ *moment.Locale, s string) string { return comma.Replace(s, "،") }
