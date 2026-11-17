// moment.js locale configuration
// locale : Chinese (Hong Kong) [zh-hk]
// author : Ben : https://github.com/ben-lin
// author : Chris Lam : https://github.com/hehachris
// author : Konstantin : https://github.com/skfd
// author : Anthony : https://github.com/anthonylau

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("zh-hk", func() {
		moment.DefineLocale("zh-hk", js.Obj(
			"months", []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
			"monthsShort", []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
			"weekdays", []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
			"weekdaysShort", []string{"週日", "週一", "週二", "週三", "週四", "週五", "週六"},
			"weekdaysMin", []string{"日", "一", "二", "三", "四", "五", "六"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY/MM/DD",
				"LL", "YYYY年M月D日",
				"LLL", "YYYY年M月D日 HH:mm",
				"LLLL", "YYYY年M月D日dddd HH:mm",
				"l", "YYYY/M/D",
				"ll", "YYYY年M月D日",
				"lll", "YYYY年M月D日 HH:mm",
				"llll", "YYYY年M月D日dddd HH:mm",
			),
			"meridiemParse", jsre.MustCompile("凌晨|早上|上午|中午|下午|晚上", ""),
			"meridiemHour", zhHkMeridiemHour,
			"meridiem", zhHkMeridiem,
			"calendar", js.Obj(
				"sameDay", "[今天]LT",
				"nextDay", "[明天]LT",
				"nextWeek", "[下]ddddLT",
				"lastDay", "[昨天]LT",
				"lastWeek", "[上]ddddLT",
				"sameElse", "L",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(日|月|週)", ""),
			"ordinal", zhHkOrdinal,
			"relativeTime", js.Obj(
				"future", "%s後",
				"past", "%s前",
				"s", "幾秒",
				"ss", "%d 秒",
				"m", "1 分鐘",
				"mm", "%d 分鐘",
				"h", "1 小時",
				"hh", "%d 小時",
				"d", "1 天",
				"dd", "%d 天",
				"M", "1 個月",
				"MM", "%d 個月",
				"y", "1 年",
				"yy", "%d 年",
			),
		))
	})
}

// zhHkMeridiemHour is the meridiemHour callback of zh-hk (and zh-mo, zh-tw).
func zhHkMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "凌晨", "早上", "上午":
		return hour
	case "中午":
		if hourNumber(hour) >= 11 {
			return hour
		}
		return plus12(hour)
	case "下午", "晚上":
		return plus12(hour)
	}
	return js.Undef
}

// zhHkMeridiem is the meridiem callback of zh-hk.
func zhHkMeridiem(_ *moment.Locale, hour, minute float64, _ bool) string {
	hm := float64(hour*100) + minute
	switch {
	case hm < 600:
		return "凌晨"
	case hm < 900:
		return "早上"
	case hm < 1200:
		return "上午"
	case hm == 1200:
		return "中午"
	case hm < 1800:
		return "下午"
	}
	return "晚上"
}

// zhHkOrdinal is the ordinal callback of zh-hk (and zh-mo, zh-tw).
func zhHkOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	return chineseOrdinal(number, period, "週")
}
