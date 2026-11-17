// moment.js locale configuration
// locale : Korean [ko]
// author : Kyungwook, Park : https://github.com/kyungw00k
// author : Jeeeyul Lee <jeeeyul@gmail.com>

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ko", func() {
		moment.DefineLocale("ko", js.Obj(
			"months", []string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
			"monthsShort", []string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
			"weekdays", []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"},
			"weekdaysShort", []string{"일", "월", "화", "수", "목", "금", "토"},
			"weekdaysMin", []string{"일", "월", "화", "수", "목", "금", "토"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm",
				"LTS", "A h:mm:ss",
				"L", "YYYY.MM.DD.",
				"LL", "YYYY년 MMMM D일",
				"LLL", "YYYY년 MMMM D일 A h:mm",
				"LLLL", "YYYY년 MMMM D일 dddd A h:mm",
				"l", "YYYY.MM.DD.",
				"ll", "YYYY년 MMMM D일",
				"lll", "YYYY년 MMMM D일 A h:mm",
				"llll", "YYYY년 MMMM D일 dddd A h:mm",
			),
			"calendar", js.Obj(
				"sameDay", "오늘 LT",
				"nextDay", "내일 LT",
				"nextWeek", "dddd LT",
				"lastDay", "어제 LT",
				"lastWeek", "지난주 dddd LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s 후",
				"past", "%s 전",
				"s", "몇 초",
				"ss", "%d초",
				"m", "1분",
				"mm", "%d분",
				"h", "한 시간",
				"hh", "%d시간",
				"d", "하루",
				"dd", "%d일",
				"M", "한 달",
				"MM", "%d달",
				"y", "일 년",
				"yy", "%d년",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(일|월|주)", ""),
			"ordinal", koOrdinal,
			"meridiemParse", jsre.MustCompile("오전|오후", ""),
			"isPM", koIsPM,
			"meridiem", koMeridiem,
		))
	})
}

// koOrdinal is the ordinal callback of ko.
func koOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	switch period {
	case "d", "D", "DDD":
		return str(number) + "일"
	case "M":
		return str(number) + "월"
	case "w", "W":
		return str(number) + "주"
	}
	return number
}

// koIsPM is the isPM callback of ko.
func koIsPM(_ *moment.Locale, token string) bool { return token == "오후" }

// koMeridiem is the meridiem callback of ko.
func koMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "오전"
	}
	return "오후"
}
