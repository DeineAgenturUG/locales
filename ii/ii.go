package ii

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type ii struct {
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

// New returns a new instance of translator for the 'ii' locale
func New() locales.Translator {
	return &ii{
		locale:                 "ii",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ꋍꆪ", "ꑍꆪ", "ꌕꆪ", "ꇖꆪ", "ꉬꆪ", "ꃘꆪ", "ꏃꆪ", "ꉆꆪ", "ꈬꆪ", "ꊰꆪ", "ꊰꊪꆪ", "ꊰꑋꆪ"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "ꋍꆪ", "ꑍꆪ", "ꌕꆪ", "ꇖꆪ", "ꉬꆪ", "ꃘꆪ", "ꏃꆪ", "ꉆꆪ", "ꈬꆪ", "ꊰꆪ", "ꊰꊪꆪ", "ꊰꑋꆪ"},
		daysAbbreviated:        []string{"ꑭꆏ", "ꆏꋍ", "ꆏꑍ", "ꆏꌕ", "ꆏꇖ", "ꆏꉬ", "ꆏꃘ"},
		daysNarrow:             []string{"ꆏ", "ꋍ", "ꑍ", "ꌕ", "ꇖ", "ꉬ", "ꃘ"},
		daysWide:               []string{"ꑭꆏꑍ", "ꆏꊂꋍ", "ꆏꊂꑍ", "ꆏꊂꌕ", "ꆏꊂꇖ", "ꆏꊂꉬ", "ꆏꊂꃘ"},
		periodsAbbreviated:     []string{"ꎸꄑ", "ꁯꋒ"},
		periodsWide:            []string{"ꎸꄑ", "ꁯꋒ"},
		erasAbbreviated:        []string{"ꃅꋊꂿ", "ꃅꋊꊂ"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"GFT": "GFT", "WAT": "WAT", "WAST": "WAST", "COT": "COT", "HKT": "HKT", "HNNOMX": "HNNOMX", "CST": "CST", "HAT": "HAT", "ART": "ART", "CHAST": "CHAST", "CDT": "CDT", "COST": "COST", "HEOG": "HEOG", "NZST": "NZST", "PDT": "PDT", "HEEG": "HEEG", "ARST": "ARST", "EST": "EST", "ChST": "ChST", "CLT": "CLT", "TMST": "TMST", "HNCU": "HNCU", "WIT": "WIT", "HNT": "HNT", "LHDT": "LHDT", "HEPM": "HEPM", "HAST": "HAST", "UYT": "UYT", "AST": "AST", "AEDT": "AEDT", "HNOG": "HNOG", "JST": "JST", "AKDT": "AKDT", "AWDT": "AWDT", "HADT": "HADT", "GMT": "GMT", "HKST": "HKST", "∅∅∅": "∅∅∅", "BOT": "BOT", "WIB": "WIB", "LHST": "LHST", "MDT": "MDT", "EAT": "EAT", "WART": "WART", "SRT": "SRT", "ECT": "ECT", "SGT": "SGT", "WEZ": "WEZ", "ACWST": "ACWST", "WESZ": "WESZ", "EDT": "EDT", "HNPMX": "HNPMX", "GYT": "GYT", "MST": "MST", "WITA": "WITA", "HNEG": "HNEG", "HEPMX": "HEPMX", "AEST": "AEST", "PST": "PST", "AKST": "AKST", "CLST": "CLST", "ADT": "ADT", "VET": "VET", "ACDT": "ACDT", "CHADT": "CHADT", "HECU": "HECU", "SAST": "SAST", "MEZ": "MEZ", "CAT": "CAT", "JDT": "JDT", "WARST": "WARST", "HENOMX": "HENOMX", "BT": "BT", "IST": "IST", "OEZ": "OEZ", "OESZ": "OESZ", "ACWDT": "ACWDT", "AWST": "AWST", "MESZ": "MESZ", "HNPM": "HNPM", "ACST": "ACST", "MYT": "MYT", "UYST": "UYST", "TMT": "TMT", "NZDT": "NZDT"},
	}
}

// Locale returns the current translators string locale
func (ii *ii) Locale() string {
	return ii.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ii'
func (ii *ii) PluralsCardinal() []locales.PluralRule {
	return ii.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ii'
func (ii *ii) PluralsOrdinal() []locales.PluralRule {
	return ii.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ii'
func (ii *ii) PluralsRange() []locales.PluralRule {
	return ii.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ii'
func (ii *ii) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ii'
func (ii *ii) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ii'
func (ii *ii) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ii *ii) MonthAbbreviated(month time.Month) string {
	return ii.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ii *ii) MonthsAbbreviated() []string {
	return ii.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ii *ii) MonthNarrow(month time.Month) string {
	return ii.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ii *ii) MonthsNarrow() []string {
	return ii.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ii *ii) MonthWide(month time.Month) string {
	return ii.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ii *ii) MonthsWide() []string {
	return ii.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ii *ii) WeekdayAbbreviated(weekday time.Weekday) string {
	return ii.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ii *ii) WeekdaysAbbreviated() []string {
	return ii.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ii *ii) WeekdayNarrow(weekday time.Weekday) string {
	return ii.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ii *ii) WeekdaysNarrow() []string {
	return ii.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ii *ii) WeekdayShort(weekday time.Weekday) string {
	return ii.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ii *ii) WeekdaysShort() []string {
	return ii.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ii *ii) WeekdayWide(weekday time.Weekday) string {
	return ii.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ii *ii) WeekdaysWide() []string {
	return ii.daysWide
}

// Decimal returns the decimal point of number
func (ii *ii) Decimal() string {
	return ii.decimal
}

// Group returns the group of number
func (ii *ii) Group() string {
	return ii.group
}

// Group returns the minus sign of number
func (ii *ii) Minus() string {
	return ii.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ii' and handles both Whole and Real numbers based on 'v'
func (ii *ii) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ii.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ii.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ii' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ii *ii) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ii.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ii.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ii'
func (ii *ii) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ii.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ii.group[0])
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

	for j := len(ii.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ii.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ii.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ii.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ii'
// in accounting notation.
func (ii *ii) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ii.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ii.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ii.group[0])
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

		for j := len(ii.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ii.currencyNegativePrefix[j])
		}

		b = append(b, ii.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ii.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ii.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ii.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ii'
func (ii *ii) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ii'
func (ii *ii) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ii.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ii'
func (ii *ii) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ii.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ii'
func (ii *ii) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, ii.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ii.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ii'
func (ii *ii) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ii'
func (ii *ii) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ii'
func (ii *ii) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ii'
func (ii *ii) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ii.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ii.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ii.periodsAbbreviated[0]...)
	} else {
		b = append(b, ii.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ii.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
