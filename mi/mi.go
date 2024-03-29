package mi

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type mi struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'mi' locale
func New() locales.Translator {
	return &mi{
		locale:                 "mi",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Kohi", "Hui", "Pou", "Pae", "Hara", "Pipi", "Hōngo", "Here", "Mahu", "Nuku", "Rangi", "Haki"},
		monthsNarrow:           []string{"", "K", "H", "P", "P", "H", "P", "H", "H", "M", "N", "R", "H"},
		monthsWide:             []string{"", "Kohitātea", "Huitanguru", "Poutūterangi", "Paengawhāwhā", "Haratua", "Pipiri", "Hōngongoi", "Hereturikōkā", "Mahuru", "Whiringa-ā-nuku", "Whiringa-ā-rangi", "Hakihea"},
		daysAbbreviated:        []string{"Tap", "Hin", "Tū", "Apa", "Par", "Mer", "Hor"},
		daysNarrow:             []string{"T", "H", "T", "A", "P", "M", "H"},
		daysShort:              []string{"Tap", "Hin", "Tū", "Apa", "Par", "Mer", "Hor"},
		daysWide:               []string{"Rātapu", "Rāhina", "Rātū", "Rāapa", "Rāpare", "Rāmere", "Rāhoroi"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"BCE", "CE"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"BCE", "CE"},
		timezones:              map[string]string{"BT": "BT", "LHST": "LHST", "NZDT": "NZDT", "HAT": "HAT", "HNEG": "HNEG", "HNPM": "HNPM", "HNOG": "HNOG", "JST": "JST", "CHADT": "CHADT", "TMST": "TMST", "HNNOMX": "HNNOMX", "SAST": "SAST", "HNT": "HNT", "ACST": "ACST", "HNPMX": "HNPMX", "MST": "MST", "MDT": "MDT", "ACDT": "ACDT", "HADT": "HADT", "IST": "IST", "EST": "Wā Arowhānui Rāwhiti", "LHDT": "LHDT", "AST": "Wā Arowhānui Ranatiki", "AEST": "AEST", "VET": "VET", "CHAST": "CHAST", "ChST": "ChST", "JDT": "JDT", "AKDT": "AKDT", "ARST": "ARST", "BOT": "BOT", "WIB": "WIB", "HKST": "HKST", "WITA": "WITA", "HECU": "HECU", "COT": "COT", "GYT": "GYT", "HAST": "HAST", "PDT": "Wā Awatea Kiwa", "ART": "ART", "WESZ": "Wā Raumati Uropi Uru", "ECT": "ECT", "HEOG": "HEOG", "EAT": "EAT", "AEDT": "AEDT", "HKT": "HKT", "WIT": "WIT", "HENOMX": "HENOMX", "WAT": "WAT", "COST": "COST", "GMT": "Wā Toharite Greenwich", "EDT": "Wā Awatea Rāwhiti", "UYT": "UYT", "UYST": "UYST", "TMT": "TMT", "MYT": "MYT", "WAST": "WAST", "WART": "WART", "WARST": "WARST", "OESZ": "Wā Raumati Uropi Rāwhiti", "∅∅∅": "∅∅∅", "AKST": "AKST", "MEZ": "Wā Arowhānui Uropi Waenga", "CST": "Wā Arowhānui Waenga", "ACWDT": "ACWDT", "CLST": "CLST", "SRT": "SRT", "CDT": "Wā Awatea Waenga", "HEEG": "HEEG", "MESZ": "Wā Raumati Uropi Waenga", "WEZ": "Wā Arowhānui Uropi Uru", "GFT": "GFT", "AWST": "AWST", "HNCU": "HNCU", "PST": "Wā Arowhānui Kiwa", "CLT": "CLT", "NZST": "NZST", "HEPM": "HEPM", "OEZ": "Wā Arowhānui Uropi Rāwhiti", "CAT": "CAT", "HEPMX": "HEPMX", "ADT": "Wā Awatea Ranatiki", "SGT": "SGT", "ACWST": "ACWST", "AWDT": "AWDT"},
	}
}

// Locale returns the current translators string locale
func (mi *mi) Locale() string {
	return mi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'mi'
func (mi *mi) PluralsCardinal() []locales.PluralRule {
	return mi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'mi'
func (mi *mi) PluralsOrdinal() []locales.PluralRule {
	return mi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'mi'
func (mi *mi) PluralsRange() []locales.PluralRule {
	return mi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'mi'
func (mi *mi) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'mi'
func (mi *mi) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'mi'
func (mi *mi) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (mi *mi) MonthAbbreviated(month time.Month) string {
	return mi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (mi *mi) MonthsAbbreviated() []string {
	return mi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (mi *mi) MonthNarrow(month time.Month) string {
	return mi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (mi *mi) MonthsNarrow() []string {
	return mi.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (mi *mi) MonthWide(month time.Month) string {
	return mi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (mi *mi) MonthsWide() []string {
	return mi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (mi *mi) WeekdayAbbreviated(weekday time.Weekday) string {
	return mi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (mi *mi) WeekdaysAbbreviated() []string {
	return mi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (mi *mi) WeekdayNarrow(weekday time.Weekday) string {
	return mi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (mi *mi) WeekdaysNarrow() []string {
	return mi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (mi *mi) WeekdayShort(weekday time.Weekday) string {
	return mi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (mi *mi) WeekdaysShort() []string {
	return mi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (mi *mi) WeekdayWide(weekday time.Weekday) string {
	return mi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (mi *mi) WeekdaysWide() []string {
	return mi.daysWide
}

// Decimal returns the decimal point of number
func (mi *mi) Decimal() string {
	return mi.decimal
}

// Group returns the group of number
func (mi *mi) Group() string {
	return mi.group
}

// Group returns the minus sign of number
func (mi *mi) Minus() string {
	return mi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'mi' and handles both Whole and Real numbers based on 'v'
func (mi *mi) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'mi' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (mi *mi) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, mi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'mi'
func (mi *mi) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(mi.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, mi.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, mi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'mi'
// in accounting notation.
func (mi *mi) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := mi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, mi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, mi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(mi.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, mi.currencyNegativePrefix[j])
		}

		b = append(b, mi.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(mi.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, mi.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, mi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'mi'
func (mi *mi) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'mi'
func (mi *mi) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'mi'
func (mi *mi) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'mi'
func (mi *mi) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, mi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, mi.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'mi'
func (mi *mi) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'mi'
func (mi *mi) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'mi'
func (mi *mi) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'mi'
func (mi *mi) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, mi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, mi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, mi.periodsAbbreviated[0]...)
	} else {
		b = append(b, mi.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := mi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
