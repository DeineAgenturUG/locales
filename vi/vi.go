package vi

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type vi struct {
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

// New returns a new instance of translator for the 'vi' locale
func New() locales.Translator {
	return &vi{
		locale:                 "vi",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "thg 1", "thg 2", "thg 3", "thg 4", "thg 5", "thg 6", "thg 7", "thg 8", "thg 9", "thg 10", "thg 11", "thg 12"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "tháng 1", "tháng 2", "tháng 3", "tháng 4", "tháng 5", "tháng 6", "tháng 7", "tháng 8", "tháng 9", "tháng 10", "tháng 11", "tháng 12"},
		daysAbbreviated:        []string{"CN", "Th 2", "Th 3", "Th 4", "Th 5", "Th 6", "Th 7"},
		daysNarrow:             []string{"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
		daysShort:              []string{"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
		daysWide:               []string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"},
		periodsAbbreviated:     []string{"SA", "CH"},
		periodsNarrow:          []string{"s", "c"},
		periodsWide:            []string{"SA", "CH"},
		erasAbbreviated:        []string{"Trước CN", "sau CN"},
		erasNarrow:             []string{"tr. CN", "sau CN"},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACST": "Giờ Chuẩn Miền Trung Australia", "CLST": "Giờ Mùa Hè Chile", "HEOG": "Giờ Mùa Hè Miền Tây Greenland", "AWST": "Giờ Chuẩn Miền Tây Australia", "WITA": "Giờ Miền Trung Indonesia", "HNCU": "Giờ Chuẩn Cuba", "HNNOMX": "Giờ Chuẩn Tây Bắc Mexico", "HAT": "Giờ Mùa Hè Newfoundland", "HADT": "Giờ Mùa Hè Hawaii-Aleut", "SGT": "Giờ Singapore", "PST": "Giờ chuẩn Thái Bình Dương", "BOT": "Giờ Bolivia", "HNPMX": "Giờ Chuẩn Thái Bình Dương Mexico", "SRT": "Giờ Suriname", "TMT": "Giờ Chuẩn Turkmenistan", "AKST": "Giờ Chuẩn Alaska", "AST": "Giờ Chuẩn Đại Tây Dương", "MESZ": "Giờ mùa hè Trung Âu", "CHADT": "Giờ Mùa Hè Chatham", "UYT": "Giờ Chuẩn Uruguay", "WIB": "Giờ Miền Tây Indonesia", "AEST": "Giờ Chuẩn Miền Đông Australia", "WIT": "Giờ Miền Đông Indonesia", "HKT": "Giờ Chuẩn Hồng Kông", "JDT": "Giờ Mùa Hè Nhật Bản", "WARST": "Giờ mùa hè miền tây Argentina", "ART": "Giờ Chuẩn Argentina", "WEZ": "Giờ Chuẩn Tây Âu", "ACWDT": "Giờ Mùa Hè Miền Trung Tây Australia", "HNOG": "Giờ Chuẩn Miền Tây Greenland", "VET": "Giờ Venezuela", "CDT": "Giờ mùa hè miền Trung", "COT": "Giờ Chuẩn Colombia", "HEPMX": "Giờ Mùa Hè Thái Bình Dương Mexico", "MDT": "Giờ Mùa Hè Ma Cao", "OEZ": "Giờ chuẩn Đông Âu", "GMT": "Giờ Trung bình Greenwich", "MST": "Giờ Chuẩn Ma Cao", "AEDT": "Giờ Mùa Hè Miền Đông Australia", "AKDT": "Giờ Mùa Hè Alaska", "EST": "Giờ chuẩn miền Đông", "UYST": "Giờ Mùa Hè Uruguay", "GYT": "Giờ Guyana", "AWDT": "Giờ Mùa Hè Miền Tây Australia", "∅∅∅": "Giờ Mùa Hè Azores", "HECU": "Giờ Mùa Hè Cuba", "NZST": "Giờ Chuẩn New Zealand", "HEPM": "Giờ Mùa Hè Saint Pierre và Miquelon", "HAST": "Giờ Chuẩn Hawaii-Aleut", "OESZ": "Giờ mùa hè Đông Âu", "ACDT": "Giờ Mùa Hè Miền Trung Australia", "EDT": "Giờ mùa hè miền Đông", "GFT": "Giờ Guiana thuộc Pháp", "WAST": "Giờ Mùa Hè Tây Phi", "WART": "Giờ chuẩn miền tây Argentina", "HENOMX": "Giờ Mùa Hè Tây Bắc Mexico", "SAST": "Giờ Chuẩn Nam Phi", "HNT": "Giờ Chuẩn Newfoundland", "CAT": "Giờ Trung Phi", "WESZ": "Giờ mùa hè Tây Âu", "ChST": "Giờ Chamorro", "CLT": "Giờ Chuẩn Chile", "HKST": "Giờ Mùa Hè Hồng Kông", "HEEG": "Giờ Mùa Hè Miền Đông Greenland", "MEZ": "Giờ chuẩn Trung Âu", "WAT": "Giờ Chuẩn Tây Phi", "EAT": "Giờ Đông Phi", "HNEG": "Giờ Chuẩn Miền Đông Greenland", "ARST": "Giờ Mùa Hè Argentina", "BT": "Giờ Bhutan", "MYT": "Giờ Malaysia", "COST": "Giờ Mùa Hè Colombia", "ECT": "Giờ Ecuador", "JST": "Giờ Chuẩn Nhật Bản", "NZDT": "Giờ Mùa Hè New Zealand", "ACWST": "Giờ Chuẩn Miền Trung Tây Australia", "LHDT": "Giờ Mùa Hè Lord Howe", "LHST": "Giờ Chuẩn Lord Howe", "ADT": "Giờ mùa hè Đại Tây Dương", "TMST": "Giờ Mùa Hè Turkmenistan", "PDT": "Giờ mùa hè Thái Bình Dương", "CHAST": "Giờ Chuẩn Chatham", "HNPM": "Giờ Chuẩn St. Pierre và Miquelon", "IST": "Giờ Chuẩn Ấn Độ", "CST": "Giờ chuẩn miền Trung"},
	}
}

// Locale returns the current translators string locale
func (vi *vi) Locale() string {
	return vi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'vi'
func (vi *vi) PluralsCardinal() []locales.PluralRule {
	return vi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'vi'
func (vi *vi) PluralsOrdinal() []locales.PluralRule {
	return vi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'vi'
func (vi *vi) PluralsRange() []locales.PluralRule {
	return vi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'vi'
func (vi *vi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'vi'
func (vi *vi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'vi'
func (vi *vi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (vi *vi) MonthAbbreviated(month time.Month) string {
	return vi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (vi *vi) MonthsAbbreviated() []string {
	return vi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (vi *vi) MonthNarrow(month time.Month) string {
	return vi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (vi *vi) MonthsNarrow() []string {
	return vi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (vi *vi) MonthWide(month time.Month) string {
	return vi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (vi *vi) MonthsWide() []string {
	return vi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (vi *vi) WeekdayAbbreviated(weekday time.Weekday) string {
	return vi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (vi *vi) WeekdaysAbbreviated() []string {
	return vi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (vi *vi) WeekdayNarrow(weekday time.Weekday) string {
	return vi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (vi *vi) WeekdaysNarrow() []string {
	return vi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (vi *vi) WeekdayShort(weekday time.Weekday) string {
	return vi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (vi *vi) WeekdaysShort() []string {
	return vi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (vi *vi) WeekdayWide(weekday time.Weekday) string {
	return vi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (vi *vi) WeekdaysWide() []string {
	return vi.daysWide
}

// Decimal returns the decimal point of number
func (vi *vi) Decimal() string {
	return vi.decimal
}

// Group returns the group of number
func (vi *vi) Group() string {
	return vi.group
}

// Group returns the minus sign of number
func (vi *vi) Minus() string {
	return vi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'vi' and handles both Whole and Real numbers based on 'v'
func (vi *vi) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'vi' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (vi *vi) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, vi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'vi'
func (vi *vi) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, vi.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'vi'
// in accounting notation.
func (vi *vi) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, vi.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, vi.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, vi.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'vi'
func (vi *vi) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'vi'
func (vi *vi) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'vi'
func (vi *vi) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'vi'
func (vi *vi) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, vi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'vi'
func (vi *vi) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'vi'
func (vi *vi) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'vi'
func (vi *vi) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'vi'
func (vi *vi) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := vi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
