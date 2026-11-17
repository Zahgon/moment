// moment.js locale configuration
// locale : Georgian [ka]
// author : Irakli Janiashvili : https://github.com/IrakliJani

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ka", func() {
		moment.DefineLocale("ka", js.Obj(
			"months", []string{"იანვარი", "თებერვალი", "მარტი", "აპრილი", "მაისი", "ივნისი", "ივლისი", "აგვისტო", "სექტემბერი", "ოქტომბერი", "ნოემბერი", "დეკემბერი"},
			"monthsShort", []string{"იან", "თებ", "მარ", "აპრ", "მაი", "ივნ", "ივლ", "აგვ", "სექ", "ოქტ", "ნოე", "დეკ"},
			"weekdays", js.Obj(
				"standalone", []string{"კვირა", "ორშაბათი", "სამშაბათი", "ოთხშაბათი", "ხუთშაბათი", "პარასკევი", "შაბათი"},
				"format", []string{"კვირას", "ორშაბათს", "სამშაბათს", "ოთხშაბათს", "ხუთშაბათს", "პარასკევს", "შაბათს"},
				"isFormat", jsre.MustCompile("(წინა|შემდეგ)", ""),
			),
			"weekdaysShort", []string{"კვი", "ორშ", "სამ", "ოთხ", "ხუთ", "პარ", "შაბ"},
			"weekdaysMin", []string{"კვ", "ორ", "სა", "ოთ", "ხუ", "პა", "შა"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[დღეს] LT[-ზე]",
				"nextDay", "[ხვალ] LT[-ზე]",
				"lastDay", "[გუშინ] LT[-ზე]",
				"nextWeek", "[შემდეგ] dddd LT[-ზე]",
				"lastWeek", "[წინა] dddd LT-ზე",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", kaRelativeTimeFuture,
				"past", kaRelativeTimePast,
				"s", "რამდენიმე წამი",
				"ss", "%d წამი",
				"m", "წუთი",
				"mm", "%d წუთი",
				"h", "საათი",
				"hh", "%d საათი",
				"d", "დღე",
				"dd", "%d დღე",
				"M", "თვე",
				"MM", "%d თვე",
				"y", "წელი",
				"yy", "%d წელი",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("0|1-ლი|მე-\\d{1,2}|\\d{1,2}-ე", ""),
			"ordinal", kaOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var (
	kaFutureUnit = jsre.MustCompile(`(წამ|წუთ|საათ|წელ|დღ|თვ)(ი|ე)`, "")
	kaPastUnit   = jsre.MustCompile(`(წამი|წუთი|საათი|დღე|თვე)`, "")
	kaPastEnding = jsre.MustCompile(`(ი|ე)$`, "")
	kaYear       = jsre.MustCompile(`წელი`, "")
	kaYearEnding = jsre.MustCompile(`წელი$`, "")
)

// kaRelativeTimeFuture is relativeTime.future of ka.
func kaRelativeTimeFuture(s string) string {
	return kaFutureUnit.ReplaceFunc(s, func(m *jsre.Match) string {
		if m.Groups[2] == "ი" {
			return m.Groups[1] + "ში"
		}
		return m.Groups[1] + m.Groups[2] + "ში"
	})
}

// kaRelativeTimePast is relativeTime.past of ka.
func kaRelativeTimePast(s string) string {
	if kaPastUnit.Test(s) {
		return kaPastEnding.Replace(s, "ის წინ")
	}
	if kaYear.Test(s) {
		return kaYearEnding.Replace(s, "წლის წინ")
	}
	return s
}

// kaOrdinal is the ordinal callback of ka.
func kaOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	n := js.ToNumber(number)
	switch {
	case sameNumber(number, 0):
		return number
	case sameNumber(number, 1):
		return str(number) + "-ლი"
	case n < 20 || (n <= 100 && js.Mod(n, 20) == 0) || js.Mod(n, 100) == 0:
		return "მე-" + str(number)
	}
	return str(number) + "-ე"
}
