package ar_DJ

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type ar_DJ struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ar_DJ' locale
func New() locales.Translator {
	return &ar_DJ{
		locale:                 "ar_DJ",
		pluralsCardinal:        []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{1, 4, 5, 6},
		decimal:                "٫",
		group:                  "٬",
		minus:                  "؜-",
		percent:                "٪؜",
		perMille:               "؉",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "Fdj", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		monthsNarrow:           []string{"", "ي", "ف", "م", "أ", "و", "ن", "ل", "غ", "س", "ك", "ب", "د"},
		monthsWide:             []string{"", "يناير", "فبراير", "مارس", "أبريل", "مايو", "يونيو", "يوليو", "أغسطس", "سبتمبر", "أكتوبر", "نوفمبر", "ديسمبر"},
		daysAbbreviated:        []string{"الأحد", "الاثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
		daysNarrow:             []string{"ح", "ن", "ث", "ر", "خ", "ج", "س"},
		daysShort:              []string{"أحد", "إثنين", "ثلاثاء", "أربعاء", "خميس", "جمعة", "سبت"},
		daysWide:               []string{"الأحد", "الاثنين", "الثلاثاء", "الأربعاء", "الخميس", "الجمعة", "السبت"},
		periodsAbbreviated:     []string{"ص", "م"},
		periodsNarrow:          []string{"ص", "م"},
		periodsWide:            []string{"ص", "م"},
		erasAbbreviated:        []string{"ق.م", "م"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"قبل الميلاد", "ميلادي"},
		timezones:              map[string]string{"ChST": "توقيت تشامورو", "SAST": "توقيت جنوب أفريقيا", "MYT": "توقيت ماليزيا", "HAST": "توقيت هاواي ألوتيان الرسمي", "WAT": "توقيت غرب أفريقيا الرسمي", "WIT": "توقيت شرق إندونيسيا", "HNCU": "توقيت كوبا الرسمي", "HAT": "توقيت نيوفاوندلاند الصيفي", "ART": "توقيت الأرجنتين الرسمي", "GFT": "توقيت غويانا الفرنسية", "GYT": "توقيت غيانا", "ECT": "توقيت الإكوادور", "HNOG": "توقيت غرب غرينلاند الرسمي", "EST": "التوقيت الرسمي الشرقي لأمريكا الشمالية", "LHDT": "التوقيت الصيفي للورد هاو", "PST": "توقيت المحيط الهادي الرسمي", "AWST": "توقيت غرب أستراليا الرسمي", "WITA": "توقيت وسط إندونيسيا", "BOT": "توقيت بوليفيا", "HNEG": "توقيت شرق غرينلاند الرسمي", "CAT": "توقيت وسط أفريقيا", "CHAST": "توقيت تشاتام الرسمي", "EDT": "التوقيت الصيفي الشرقي لأمريكا الشمالية", "WAST": "توقيت غرب أفريقيا الصيفي", "HNT": "توقيت نيوفاوندلاند الرسمي", "BT": "توقيت بوتان", "WIB": "توقيت غرب إندونيسيا", "HNPMX": "توقيت المحيط الهادي الرسمي للمكسيك", "HEPMX": "توقيت المحيط الهادي الصيفي للمكسيك", "GMT": "توقيت غرينتش", "PDT": "توقيت المحيط الهادي الصيفي", "ARST": "توقيت الأرجنتين الصيفي", "HNPM": "توقيت سانت بيير وميكولون الرسمي", "IST": "توقيت الهند", "HADT": "توقيت هاواي ألوتيان الصيفي", "CLST": "توقيت تشيلي الصيفي", "WART": "توقيت غرب الأرجنتين الرسمي", "VET": "توقيت فنزويلا", "MESZ": "توقيت وسط أوروبا الصيفي", "ACST": "توقيت وسط أستراليا الرسمي", "CHADT": "توقيت تشاتام الصيفي", "HKT": "توقيت هونغ كونغ الرسمي", "JST": "توقيت اليابان الرسمي", "OEZ": "توقيت شرق أوروبا الرسمي", "MEZ": "توقيت وسط أوروبا الرسمي", "ACDT": "توقيت وسط أستراليا الصيفي", "COST": "توقيت كولومبيا الصيفي", "AWDT": "توقيت غرب أستراليا الصيفي", "TMT": "توقيت تركمانستان الرسمي", "HKST": "توقيت هونغ كونغ الصيفي", "JDT": "توقيت اليابان الصيفي", "HENOMX": "التوقيت الصيفي لشمال غرب المكسيك", "AKDT": "توقيت ألاسكا الصيفي", "CST": "التوقيت الرسمي المركزي لأمريكا الشمالية", "SRT": "توقيت سورينام", "AEDT": "توقيت شرق أستراليا الصيفي", "HEOG": "توقيت غرب غرينلاند الصيفي", "MDT": "MDT", "NZDT": "توقيت نيوزيلندا الصيفي", "AKST": "التوقيت الرسمي لألاسكا", "SGT": "توقيت سنغافورة", "HEPM": "توقيت سانت بيير وميكولون الصيفي", "∅∅∅": "توقيت برازيليا الصيفي", "ACWDT": "توقيت غرب وسط أستراليا الصيفي", "ADT": "التوقيت الصيفي الأطلسي", "AEST": "توقيت شرق أستراليا الرسمي", "WARST": "توقيت غرب الأرجنتين الصيفي", "NZST": "توقيت نيوزيلندا الرسمي", "HECU": "توقيت كوبا الصيفي", "CDT": "التوقيت الصيفي المركزي لأمريكا الشمالية", "WEZ": "توقيت غرب أوروبا الرسمي", "UYT": "توقيت أوروغواي الرسمي", "AST": "التوقيت الرسمي الأطلسي", "ACWST": "توقيت غرب وسط أستراليا الرسمي", "COT": "توقيت كولومبيا الرسمي", "TMST": "توقيت تركمانستان الصيفي", "EAT": "توقيت شرق أفريقيا", "CLT": "توقيت تشيلي الرسمي", "OESZ": "توقيت شرق أوروبا الصيفي", "HEEG": "توقيت شرق غرينلاند الصيفي", "HNNOMX": "التوقيت الرسمي لشمال غرب المكسيك", "WESZ": "توقيت غرب أوروبا الصيفي", "LHST": "توقيت لورد هاو الرسمي", "UYST": "توقيت أوروغواي الصيفي", "MST": "MST"},
	}
}

// Locale returns the current translators string locale
func (ar *ar_DJ) Locale() string {
	return ar.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ar_DJ'
func (ar *ar_DJ) PluralsCardinal() []locales.PluralRule {
	return ar.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ar_DJ'
func (ar *ar_DJ) PluralsOrdinal() []locales.PluralRule {
	return ar.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ar_DJ'
func (ar *ar_DJ) PluralsRange() []locales.PluralRule {
	return ar.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ar_DJ'
func (ar *ar_DJ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if nMod100 >= 3 && nMod100 <= 10 {
		return locales.PluralRuleFew
	} else if nMod100 >= 11 && nMod100 <= 99 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ar_DJ'
func (ar *ar_DJ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ar_DJ'
func (ar *ar_DJ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ar.CardinalPluralRule(num1, v1)
	end := ar.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleZero
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleTwo {
		return locales.PluralRuleZero
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleTwo {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ar *ar_DJ) MonthAbbreviated(month time.Month) string {
	return ar.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ar *ar_DJ) MonthsAbbreviated() []string {
	return ar.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ar *ar_DJ) MonthNarrow(month time.Month) string {
	return ar.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ar *ar_DJ) MonthsNarrow() []string {
	return ar.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ar *ar_DJ) MonthWide(month time.Month) string {
	return ar.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ar *ar_DJ) MonthsWide() []string {
	return ar.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ar *ar_DJ) WeekdayAbbreviated(weekday time.Weekday) string {
	return ar.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ar *ar_DJ) WeekdaysAbbreviated() []string {
	return ar.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ar *ar_DJ) WeekdayNarrow(weekday time.Weekday) string {
	return ar.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ar *ar_DJ) WeekdaysNarrow() []string {
	return ar.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ar *ar_DJ) WeekdayShort(weekday time.Weekday) string {
	return ar.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ar *ar_DJ) WeekdaysShort() []string {
	return ar.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ar *ar_DJ) WeekdayWide(weekday time.Weekday) string {
	return ar.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ar *ar_DJ) WeekdaysWide() []string {
	return ar.daysWide
}

// Decimal returns the decimal point of number
func (ar *ar_DJ) Decimal() string {
	return ar.decimal
}

// Group returns the group of number
func (ar *ar_DJ) Group() string {
	return ar.group
}

// Group returns the minus sign of number
func (ar *ar_DJ) Minus() string {
	return ar.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ar_DJ' and handles both Whole and Real numbers based on 'v'
func (ar *ar_DJ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ar.group) - 1; j >= 0; j-- {
					b = append(b, ar.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ar_DJ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ar *ar_DJ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 9
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ar.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ar_DJ'
func (ar *ar_DJ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ar.currencies[currency]
	l := len(s) + len(symbol) + 7 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ar.group) - 1; j >= 0; j-- {
					b = append(b, ar.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ar.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ar.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ar_DJ'
// in accounting notation.
func (ar *ar_DJ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ar.currencies[currency]
	l := len(s) + len(symbol) + 7 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			for j := len(ar.decimal) - 1; j >= 0; j-- {
				b = append(b, ar.decimal[j])
			}
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(ar.group) - 1; j >= 0; j-- {
					b = append(b, ar.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ar.minus) - 1; j >= 0; j-- {
			b = append(b, ar.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ar.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ar.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ar.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0xe2, 0x80, 0x8f, 0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ar.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ar.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xd8, 0x8c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ar.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ar_DJ'
func (ar *ar_DJ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ar.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ar.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ar.periodsAbbreviated[0]...)
	} else {
		b = append(b, ar.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ar.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
