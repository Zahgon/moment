// moment.js locale configuration
// locale : Thai [th]
// author : Kridsada Thanabulpong : https://github.com/sirn

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("th", func() {
		moment.DefineLocale("th", js.Obj(
			"months", []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"},
			"monthsShort", []string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
			"monthsParseExact", true,
			"weekdays", []string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"},
			"weekdaysShort", []string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัส", "ศุกร์", "เสาร์"},
			"weekdaysMin", []string{"อา.", "จ.", "อ.", "พ.", "พฤ.", "ศ.", "ส."},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY เวลา H:mm",
				"LLLL", "วันddddที่ D MMMM YYYY เวลา H:mm",
			),
			"meridiemParse", jsre.MustCompile("ก่อนเที่ยง|หลังเที่ยง", ""),
			"isPM", thIsPM,
			"meridiem", thMeridiem,
			"calendar", js.Obj(
				"sameDay", "[วันนี้ เวลา] LT",
				"nextDay", "[พรุ่งนี้ เวลา] LT",
				"nextWeek", "dddd[หน้า เวลา] LT",
				"lastDay", "[เมื่อวานนี้ เวลา] LT",
				"lastWeek", "[วัน]dddd[ที่แล้ว เวลา] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "อีก %s",
				"past", "%sที่แล้ว",
				"s", "ไม่กี่วินาที",
				"ss", "%d วินาที",
				"m", "1 นาที",
				"mm", "%d นาที",
				"h", "1 ชั่วโมง",
				"hh", "%d ชั่วโมง",
				"d", "1 วัน",
				"dd", "%d วัน",
				"w", "1 สัปดาห์",
				"ww", "%d สัปดาห์",
				"M", "1 เดือน",
				"MM", "%d เดือน",
				"y", "1 ปี",
				"yy", "%d ปี",
			),
		))
	})
}

// thIsPM is the isPM callback of th.
func thIsPM(_ *moment.Locale, input string) bool { return input == "หลังเที่ยง" }

// thMeridiem is the meridiem callback of th.
func thMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "ก่อนเที่ยง"
	}
	return "หลังเที่ยง"
}
