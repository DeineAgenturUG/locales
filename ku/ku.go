package ku

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type ku struct {
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
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'ku' locale
func New() locales.Translator {
	return &ku{
		locale:                 "ku",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "₺", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "rêb", "reş", "ada", "avr", "gul", "pûş", "tîr", "gel", "rez", "kew", "ser", "ber"},
		monthsNarrow:           []string{"", "R", "R", "A", "A", "G", "P", "T", "G", "R", "K", "S", "B"},
		monthsWide:             []string{"", "rêbendanê", "reşemiyê", "adarê", "avrêlê", "gulanê", "pûşperê", "tîrmehê", "gelawêjê", "rezberê", "kewçêrê", "sermawezê", "berfanbarê"},
		daysAbbreviated:        []string{"yş", "dş", "sş", "çş", "pş", "în", "ş"},
		daysNarrow:             []string{"Y", "D", "S", "Ç", "P", "Î", "Ş"},
		daysShort:              []string{"yş", "dş", "sş", "çş", "pş", "în", "ş"},
		daysWide:               []string{"yekşem", "duşem", "sêşem", "çarşem", "pêncşem", "în", "şemî"},
		periodsAbbreviated:     []string{"BN", "PN"},
		periodsWide:            []string{"BN", "PN"},
		erasAbbreviated:        []string{"BZ", "PZ"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"berî zayînê", "piştî zayînê"},
		timezones:              map[string]string{"AEST": "AEST", "TMT": "TMT", "NZST": "NZST", "VET": "VET", "EDT": "EDT", "HAST": "HAST", "ACWST": "ACWST", "HKST": "HKST", "WIT": "WIT", "HNEG": "HNEG", "MYT": "MYT", "HNCU": "HNCU", "HECU": "HECU", "ARST": "ARST", "CHAST": "CHAST", "WESZ": "WESZ", "TMST": "TMST", "HEEG": "HEEG", "HNOG": "HNOG", "HENOMX": "HENOMX", "BT": "BT", "ACST": "ACST", "COT": "COT", "ECT": "ECT", "HNNOMX": "HNNOMX", "SAST": "SAST", "EST": "EST", "COST": "COST", "HEPMX": "HEPMX", "SRT": "SRT", "AWST": "AWST", "OEZ": "OEZ", "WITA": "WITA", "WAST": "WAST", "CLT": "CLT", "GMT": "GMT", "CAT": "CAT", "ACDT": "ACDT", "WAT": "WAT", "HEOG": "HEOG", "SGT": "SGT", "HNT": "HNT", "ADT": "ADT", "PDT": "PDT", "WIB": "WIB", "HADT": "HADT", "GFT": "GFT", "MST": "MST", "EAT": "EAT", "JDT": "JDT", "NZDT": "NZDT", "∅∅∅": "∅∅∅", "HNPMX": "HNPMX", "OESZ": "OESZ", "HEPM": "HEPM", "CST": "CST", "ChST": "ChST", "UYT": "UYT", "AEDT": "AEDT", "JST": "JST", "WARST": "WARST", "AKDT": "AKDT", "CLST": "CLST", "GYT": "GYT", "PST": "PST", "ART": "ART", "HNPM": "HNPM", "CHADT": "CHADT", "IST": "IST", "UYST": "UYST", "AST": "AST", "AWDT": "AWDT", "HAT": "HAT", "AKST": "AKST", "MESZ": "MESZ", "WEZ": "WEZ", "LHST": "LHST", "CDT": "CDT", "LHDT": "LHDT", "ACWDT": "ACWDT", "MDT": "MDT", "HKT": "HKT", "WART": "WART", "BOT": "BOT", "MEZ": "MEZ"},
	}
}

// Locale returns the current translators string locale
func (ku *ku) Locale() string {
	return ku.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ku'
func (ku *ku) PluralsCardinal() []locales.PluralRule {
	return ku.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ku'
func (ku *ku) PluralsOrdinal() []locales.PluralRule {
	return ku.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ku'
func (ku *ku) PluralsRange() []locales.PluralRule {
	return ku.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ku'
func (ku *ku) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ku'
func (ku *ku) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ku'
func (ku *ku) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ku *ku) MonthAbbreviated(month time.Month) string {
	return ku.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ku *ku) MonthsAbbreviated() []string {
	return ku.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ku *ku) MonthNarrow(month time.Month) string {
	return ku.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ku *ku) MonthsNarrow() []string {
	return ku.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ku *ku) MonthWide(month time.Month) string {
	return ku.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ku *ku) MonthsWide() []string {
	return ku.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ku *ku) WeekdayAbbreviated(weekday time.Weekday) string {
	return ku.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ku *ku) WeekdaysAbbreviated() []string {
	return ku.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ku *ku) WeekdayNarrow(weekday time.Weekday) string {
	return ku.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ku *ku) WeekdaysNarrow() []string {
	return ku.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ku *ku) WeekdayShort(weekday time.Weekday) string {
	return ku.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ku *ku) WeekdaysShort() []string {
	return ku.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ku *ku) WeekdayWide(weekday time.Weekday) string {
	return ku.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ku *ku) WeekdaysWide() []string {
	return ku.daysWide
}

// Decimal returns the decimal point of number
func (ku *ku) Decimal() string {
	return ku.decimal
}

// Group returns the group of number
func (ku *ku) Group() string {
	return ku.group
}

// Group returns the minus sign of number
func (ku *ku) Minus() string {
	return ku.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ku' and handles both Whole and Real numbers based on 'v'
func (ku *ku) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ku' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ku *ku) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.minus[0])
	}

	b = append(b, ku.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ku'
func (ku *ku) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ku.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ku.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ku.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ku.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ku'
// in accounting notation.
func (ku *ku) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ku.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ku.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ku.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ku.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ku.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ku.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ku.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ku'
func (ku *ku) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ku'
func (ku *ku) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ku'
func (ku *ku) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ku'
func (ku *ku) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ku'
func (ku *ku) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ku'
func (ku *ku) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ku'
func (ku *ku) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ku'
func (ku *ku) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	return string(b)
}
