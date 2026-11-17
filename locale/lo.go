// moment.js locale configuration
// locale : Lao [lo]
// author : Ryan Hart : https://github.com/ryanhart2

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("lo", func() {
		moment.DefineLocale("lo", js.Obj(
			"months", []string{"ມັງກອນ", "ກຸມພາ", "ມີນາ", "ເມສາ", "ພຶດສະພາ", "ມິຖຸນາ", "ກໍລະກົດ", "ສິງຫາ", "ກັນຍາ", "ຕຸລາ", "ພະຈິກ", "ທັນວາ"},
			"monthsShort", []string{"ມັງກອນ", "ກຸມພາ", "ມີນາ", "ເມສາ", "ພຶດສະພາ", "ມິຖຸນາ", "ກໍລະກົດ", "ສິງຫາ", "ກັນຍາ", "ຕຸລາ", "ພະຈິກ", "ທັນວາ"},
			"weekdays", []string{"ອາທິດ", "ຈັນ", "ອັງຄານ", "ພຸດ", "ພະຫັດ", "ສຸກ", "ເສົາ"},
			"weekdaysShort", []string{"ທິດ", "ຈັນ", "ອັງຄານ", "ພຸດ", "ພະຫັດ", "ສຸກ", "ເສົາ"},
			"weekdaysMin", []string{"ທ", "ຈ", "ອຄ", "ພ", "ພຫ", "ສກ", "ສ"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "ວັນdddd D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("ຕອນເຊົ້າ|ຕອນແລງ", ""),
			"isPM", loIsPM,
			"meridiem", loMeridiem,
			"calendar", js.Obj(
				"sameDay", "[ມື້ນີ້ເວລາ] LT",
				"nextDay", "[ມື້ອື່ນເວລາ] LT",
				"nextWeek", "[ວັນ]dddd[ໜ້າເວລາ] LT",
				"lastDay", "[ມື້ວານນີ້ເວລາ] LT",
				"lastWeek", "[ວັນ]dddd[ແລ້ວນີ້ເວລາ] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "ອີກ %s",
				"past", "%sຜ່ານມາ",
				"s", "ບໍ່ເທົ່າໃດວິນາທີ",
				"ss", "%d ວິນາທີ",
				"m", "1 ນາທີ",
				"mm", "%d ນາທີ",
				"h", "1 ຊົ່ວໂມງ",
				"hh", "%d ຊົ່ວໂມງ",
				"d", "1 ມື້",
				"dd", "%d ມື້",
				"M", "1 ເດືອນ",
				"MM", "%d ເດືອນ",
				"y", "1 ປີ",
				"yy", "%d ປີ",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("(ທີ່)\\d{1,2}", ""),
			"ordinal", loOrdinal,
		))
	})
}

// loIsPM is the isPM callback of lo.
func loIsPM(_ *moment.Locale, input string) bool { return input == "ຕອນແລງ" }

// loMeridiem is the meridiem callback of lo.
func loMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "ຕອນເຊົ້າ"
	}
	return "ຕອນແລງ"
}

// loOrdinal is the ordinal callback of lo.
func loOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return "ທີ່" + str(number)
}
