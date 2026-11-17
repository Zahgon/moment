// moment.js locale configuration
// locale : Japanese [ja]
// author : LI Long : https://github.com/baryon

package locale

import (
	"math"

	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ja", func() {
		moment.DefineLocale("ja", js.Obj(
			"eras", []any{
				js.Obj(
					"since", "2019-05-01",
					"offset", 1.0,
					"name", "令和",
					"narrow", "㋿",
					"abbr", "R",
				),
				js.Obj(
					"since", "1989-01-08",
					"until", "2019-04-30",
					"offset", 1.0,
					"name", "平成",
					"narrow", "㍻",
					"abbr", "H",
				),
				js.Obj(
					"since", "1926-12-25",
					"until", "1989-01-07",
					"offset", 1.0,
					"name", "昭和",
					"narrow", "㍼",
					"abbr", "S",
				),
				js.Obj(
					"since", "1912-07-30",
					"until", "1926-12-24",
					"offset", 1.0,
					"name", "大正",
					"narrow", "㍽",
					"abbr", "T",
				),
				js.Obj(
					"since", "1873-01-01",
					"until", "1912-07-29",
					"offset", 6.0,
					"name", "明治",
					"narrow", "㍾",
					"abbr", "M",
				),
				js.Obj(
					"since", "0001-01-01",
					"until", "1873-12-31",
					"offset", 1.0,
					"name", "西暦",
					"narrow", "AD",
					"abbr", "AD",
				),
				js.Obj(
					"since", "0000-12-31",
					"until", math.Inf(-1),
					"offset", 1.0,
					"name", "紀元前",
					"narrow", "BC",
					"abbr", "BC",
				),
			},
			"eraYearOrdinalRegex", jsre.MustCompile("(元|\\d+)年", ""),
			"eraYearOrdinalParse", jaEraYearOrdinalParse,
			"months", []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
			"monthsShort", []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
			"weekdays", []string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},
			"weekdaysShort", []string{"日", "月", "火", "水", "木", "金", "土"},
			"weekdaysMin", []string{"日", "月", "火", "水", "木", "金", "土"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY/MM/DD",
				"LL", "YYYY年M月D日",
				"LLL", "YYYY年M月D日 HH:mm",
				"LLLL", "YYYY年M月D日 dddd HH:mm",
				"l", "YYYY/MM/DD",
				"ll", "YYYY年M月D日",
				"lll", "YYYY年M月D日 HH:mm",
				"llll", "YYYY年M月D日(ddd) HH:mm",
			),
			"meridiemParse", jsre.MustCompile("午前|午後", "i"),
			"isPM", jaIsPM,
			"meridiem", jaMeridiem,
			"calendar", js.Obj(
				"sameDay", "[今日] LT",
				"nextDay", "[明日] LT",
				"nextWeek", jaCalendarNextWeek,
				"lastDay", "[昨日] LT",
				"lastWeek", jaCalendarLastWeek,
				"sameElse", "L",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}日", ""),
			"ordinal", jaOrdinal,
			"relativeTime", js.Obj(
				"future", "%s後",
				"past", "%s前",
				"s", "数秒",
				"ss", "%d秒",
				"m", "1分",
				"mm", "%d分",
				"h", "1時間",
				"hh", "%d時間",
				"d", "1日",
				"dd", "%d日",
				"M", "1ヶ月",
				"MM", "%dヶ月",
				"y", "1年",
				"yy", "%d年",
			),
		))
	})
}

// jaEraYearOrdinalParse reads the year of an era ("元" is the first year).
func jaEraYearOrdinalParse(_ *moment.Locale, input string, match *jsre.Match) float64 {
	if match == nil {
		return js.ParseInt(input)
	}
	group, defined := match.Groups[1], match.Defined[1]
	if defined && group == "元" {
		return 1
	}
	if !defined || group == "" {
		return js.ParseInt(input)
	}
	return js.ParseInt(group)
}

// jaIsPM is the isPM callback of ja.
func jaIsPM(_ *moment.Locale, input string) bool { return input == "午後" }

// jaMeridiem is the meridiem callback of ja.
func jaMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "午前"
	}
	return "午後"
}

// jaCalendarNextWeek is calendar.nextWeek of ja.
func jaCalendarNextWeek(m, now *moment.Moment) string {
	if now.Week() != m.Week() {
		return "[来週]dddd LT"
	}
	return "dddd LT"
}

// jaCalendarLastWeek is calendar.lastWeek of ja.
func jaCalendarLastWeek(m, now *moment.Moment) string {
	if m.Week() != now.Week() {
		return "[先週]dddd LT"
	}
	return "dddd LT"
}

// jaOrdinal is the ordinal callback of ja.
func jaOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "y":
		if sameNumber(number, 1) {
			return "元年"
		}
		return str(number) + "年"
	case "d", "D", "DDD":
		return str(number) + "日"
	}
	return number
}
