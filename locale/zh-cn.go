// moment.js locale configuration
// locale : Chinese (China) [zh-cn]
// author : suupic : https://github.com/suupic
// author : Zeno Zeng : https://github.com/zenozeng
// author : uu109 : https://github.com/uu109

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("zh-cn", func() {
		moment.DefineLocale("zh-cn", js.Obj(
			"months", []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
			"monthsShort", []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
			"weekdays", []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
			"weekdaysShort", []string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"},
			"weekdaysMin", []string{"日", "一", "二", "三", "四", "五", "六"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY/MM/DD",
				"LL", "YYYY年M月D日",
				"LLL", "YYYY年M月D日Ah点mm分",
				"LLLL", "YYYY年M月D日ddddAh点mm分",
				"l", "YYYY/M/D",
				"ll", "YYYY年M月D日",
				"lll", "YYYY年M月D日 HH:mm",
				"llll", "YYYY年M月D日dddd HH:mm",
			),
			"meridiemParse", jsre.MustCompile("凌晨|早上|上午|中午|下午|晚上", ""),
			"meridiemHour", zhCnMeridiemHour,
			"meridiem", zhCnMeridiem,
			"calendar", js.Obj(
				"sameDay", "[今天]LT",
				"nextDay", "[明天]LT",
				"nextWeek", zhCnCalendarNextWeek,
				"lastDay", "[昨天]LT",
				"lastWeek", zhCnCalendarLastWeek,
				"sameElse", "L",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(日|月|周)", ""),
			"ordinal", zhCnOrdinal,
			"relativeTime", js.Obj(
				"future", "%s后",
				"past", "%s前",
				"s", "几秒",
				"ss", "%d 秒",
				"m", "1 分钟",
				"mm", "%d 分钟",
				"h", "1 小时",
				"hh", "%d 小时",
				"d", "1 天",
				"dd", "%d 天",
				"w", "1 周",
				"ww", "%d 周",
				"M", "1 个月",
				"MM", "%d 个月",
				"y", "1 年",
				"yy", "%d 年",
			),
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// zhCnMeridiemHour is the meridiemHour callback of zh-cn.
func zhCnMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "凌晨", "早上", "上午":
		return hour
	case "下午", "晚上":
		return plus12(hour)
	}
	if hourNumber(hour) >= 11 {
		return hour
	}
	return plus12(hour)
}

// zhCnMeridiem is the meridiem callback of zh-cn (and zh-mo, zh-tw).
func zhCnMeridiem(_ *moment.Locale, hour, minute float64, _ bool) string {
	hm := float64(hour*100) + minute
	switch {
	case hm < 600:
		return "凌晨"
	case hm < 900:
		return "早上"
	case hm < 1130:
		return "上午"
	case hm < 1230:
		return "中午"
	case hm < 1800:
		return "下午"
	}
	return "晚上"
}

// zhCnCalendarNextWeek is calendar.nextWeek of zh-cn.
func zhCnCalendarNextWeek(m, now *moment.Moment) string {
	if now.Week() != m.Week() {
		return "[下]dddLT"
	}
	return "[本]dddLT"
}

// zhCnCalendarLastWeek is calendar.lastWeek of zh-cn.
func zhCnCalendarLastWeek(m, now *moment.Moment) string {
	if m.Week() != now.Week() {
		return "[上]dddLT"
	}
	return "[本]dddLT"
}

// zhCnOrdinal is the ordinal callback of zh-cn.
func zhCnOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	return chineseOrdinal(number, period, "周")
}

// chineseOrdinal is the ordinal of the Chinese locales: day, month and week suffixes, otherwise the argument itself.
func chineseOrdinal(number moment.Value, period, week string) moment.Value {
	switch period {
	case "d", "D", "DDD":
		return str(number) + "日"
	case "M":
		return str(number) + "月"
	case "w", "W":
		return str(number) + week
	}
	return number
}
