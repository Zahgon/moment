// moment.js locale configuration
// locale : Chinese (Taiwan) [zh-tw]
// author : Ben : https://github.com/ben-lin
// author : Chris Lam : https://github.com/hehachris

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("zh-tw", func() {
		moment.DefineLocale("zh-tw", js.Obj(
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
			"meridiemHour", zhTwMeridiemHour,
			"meridiem", zhTwMeridiem,
			"calendar", js.Obj(
				"sameDay", "[今天] LT",
				"nextDay", "[明天] LT",
				"nextWeek", "[下]dddd LT",
				"lastDay", "[昨天] LT",
				"lastWeek", "[上]dddd LT",
				"sameElse", "L",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(日|月|週)", ""),
			"ordinal", zhTwOrdinal,
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

var (
	zhTwMeridiemHour = zhHkMeridiemHour
	zhTwMeridiem     = zhCnMeridiem
	zhTwOrdinal      = zhHkOrdinal
)
