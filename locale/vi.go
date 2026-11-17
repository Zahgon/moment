// moment.js locale configuration
// locale : Vietnamese [vi]
// author : Bang Nguyen : https://github.com/bangnk
// author : Chien Kira : https://github.com/chienkira

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("vi", func() {
		moment.DefineLocale("vi", js.Obj(
			"months", []string{"tháng 1", "tháng 2", "tháng 3", "tháng 4", "tháng 5", "tháng 6", "tháng 7", "tháng 8", "tháng 9", "tháng 10", "tháng 11", "tháng 12"},
			"monthsShort", []string{"Thg 01", "Thg 02", "Thg 03", "Thg 04", "Thg 05", "Thg 06", "Thg 07", "Thg 08", "Thg 09", "Thg 10", "Thg 11", "Thg 12"},
			"monthsParseExact", true,
			"weekdays", []string{"chủ nhật", "thứ hai", "thứ ba", "thứ tư", "thứ năm", "thứ sáu", "thứ bảy"},
			"weekdaysShort", []string{"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
			"weekdaysMin", []string{"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
			"weekdaysParseExact", true,
			"meridiemParse", jsre.MustCompile("sa|ch", "i"),
			"isPM", viIsPM,
			"meridiem", viMeridiem,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM [năm] YYYY",
				"LLL", "D MMMM [năm] YYYY HH:mm",
				"LLLL", "dddd, D MMMM [năm] YYYY HH:mm",
				"l", "DD/M/YYYY",
				"ll", "D MMM YYYY",
				"lll", "D MMM YYYY HH:mm",
				"llll", "ddd, D MMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Hôm nay lúc] LT",
				"nextDay", "[Ngày mai lúc] LT",
				"nextWeek", "dddd [tuần tới lúc] LT",
				"lastDay", "[Hôm qua lúc] LT",
				"lastWeek", "dddd [tuần trước lúc] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s tới",
				"past", "%s trước",
				"s", "vài giây",
				"ss", "%d giây",
				"m", "một phút",
				"mm", "%d phút",
				"h", "một giờ",
				"hh", "%d giờ",
				"d", "một ngày",
				"dd", "%d ngày",
				"w", "một tuần",
				"ww", "%d tuần",
				"M", "một tháng",
				"MM", "%d tháng",
				"y", "một năm",
				"yy", "%d năm",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}", ""),
			"ordinal", viOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var viPM = jsre.MustCompile(`^ch$`, "i")

// viIsPM is the isPM callback of vi.
func viIsPM(_ *moment.Locale, input string) bool { return viPM.Test(input) }

// viMeridiem is the meridiem callback of vi.
func viMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours < 12 {
		if isLower {
			return "sa"
		}
		return "SA"
	}
	if isLower {
		return "ch"
	}
	return "CH"
}

// viOrdinal is the ordinal callback of vi: the argument itself, unchanged.
func viOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value { return number }
