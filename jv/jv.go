package jv

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type jv struct {
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

// New returns a new instance of translator for the 'jv' locale
func New() locales.Translator {
	return &jv{
		locale:                 "jv",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "A$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CN¥", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
		daysAbbreviated:        []string{"Ahd", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
		daysNarrow:             []string{"A", "S", "S", "R", "K", "J", "S"},
		daysShort:              []string{"Ahd", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
		daysWide:               []string{"Ahad", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"},
		periodsAbbreviated:     []string{"Isuk", "Wengi"},
		periodsNarrow:          []string{"Isuk", "Wengi"},
		periodsWide:            []string{"Isuk", "Wengi"},
		erasAbbreviated:        []string{"SM", "M"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Sakdurunge Masehi", "Masehi"},
		timezones:              map[string]string{"MESZ": "Wektu Ketigo Eropa Tengah", "SRT": "Wektu Suriname", "HEEG": "Wektu Ketigo Grinland Wetan", "HNPM": "Wektu Standar Santa Pierre lan Miquelon", "HEPM": "Wektu Ketigo Santa Pierre lan Miquelon", "LHST": "Wektu Standar Lord Howe", "UYST": "Wektu Ketigo Uruguay", "WART": "Wektu Standar Argentina sisih Kulon", "HNCU": "Wektu Standar Kuba", "CST": "Wektu Standar Tengah", "WARST": "Wektu Ketigo Argentina sisih Kulon", "HNT": "Wektu Standar Newfoundland", "ARST": "Wektu Ketigo Argentina", "COST": "Wektu Ketigo Kolombia", "HEOG": "Wektu Ketigo Grinland Kulon", "JDT": "Wektu Ketigo Jepang", "HNNOMX": "Wektu Standar Meksiko Lor-Kulon", "SAST": "Wektu Standar Afrika Kidul", "∅∅∅": "Wektu Ketigo Brasilia", "HAST": "Wektu Standar Hawaii-Aleutian", "ADT": "Wektu Ketigo Atlantik", "WITA": "Wektu Indonesia Tengah", "AKDT": "Wektu Ketigo Alaska", "CAT": "Wektu Afrika Tengah", "WESZ": "Wektu Ketigo Eropa sisih Kulon", "GFT": "Wektu Guiana Prancis", "ACWST": "Wektu Standar Australia Tengah sisih Kulon", "CLT": "Wektu Standar Chili", "HNEG": "Wektu Standar Grinland Wetan", "NZDT": "Wektu Ketigo Selandia Anyar", "PST": "Wektu Standar Pasifik", "ACDT": "Wektu Ketigo Australia Tengah", "WIB": "Wektu Indonesia sisih Kulon", "ACWDT": "Wektu Ketigo Australia Tengah sisih Kulon", "MST": "MST", "AWDT": "Wektu Ketigo Australia sisih Kulon", "TMT": "Wektu Standar Turkmenistan", "WEZ": "Wektu Standar Eropa sisih Kulon", "GMT": "Wektu Rerata Greenwich", "EAT": "Wektu Afrika Wetan", "HAT": "Wektu Ketigo Newfoundland", "HENOMX": "Wektu Ketigo Meksiko Lor-Kulon", "ACST": "Wektu Standar Australia Tengah", "CHAST": "Wektu Standar Chatham", "CDT": "Wektu Ketigo Tengah", "HNOG": "Wektu Standar Grinland Kulon", "HKST": "Wektu Ketigo Hong Kong", "SGT": "Wektu Standar Singapura", "HEPMX": "Wektu Ketigo Pasifik Meksiko", "JST": "Wektu Standar Jepang", "PDT": "Wektu Ketigo Pasifik", "BOT": "Wektu Bolivia", "IST": "Wektu Standar India", "GYT": "Wektu Guyana", "ECT": "Wektu Ekuador", "AEST": "Wektu Standar Australia sisih Wetan", "AWST": "Wektu Standar Australia sisih Kulon", "OESZ": "Wektu Ketigo Eropa sisih Wetan", "AKST": "Wektu Standar Alaska", "WAT": "Wektu Standar Afrika Kulon", "WAST": "Wektu Ketigo Afrika Kulon", "AEDT": "Wektu Ketigo Australia sisih Wetan", "HECU": "Wektu Ketigo Kuba", "OEZ": "Wektu Standar Eropa sisih Wetan", "ART": "Wektu Standar Argentina", "EDT": "Wektu Ketigo sisih Wetah", "MYT": "Wektu Malaysia", "HADT": "Wektu Ketigo Hawaii-Aleutian", "UYT": "Wektu Standar Uruguay", "COT": "Wektu Standar Kolombia", "BT": "Wektu Bhutan", "CHADT": "Wektu Ketigo Chatham", "MDT": "MDT", "AST": "Wektu Standar Atlantik", "WIT": "Wektu Indonesia sisih Wetan", "MEZ": "Wektu Standar Eropa Tengah", "EST": "Wektu Standar sisih Wetan", "ChST": "Wektu Standar Chamorro", "CLST": "Wektu Ketigo Chili", "NZST": "Wektu Standar Selandia Anyar", "LHDT": "Wektu Ketigo Lord Howe", "HNPMX": "Wektu Standar Pasifik Meksiko", "HKT": "Wektu Standar Hong Kong", "TMST": "Wektu Ketigo Turkmenistan", "VET": "Wektu Venezuela"},
	}
}

// Locale returns the current translators string locale
func (jv *jv) Locale() string {
	return jv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'jv'
func (jv *jv) PluralsCardinal() []locales.PluralRule {
	return jv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'jv'
func (jv *jv) PluralsOrdinal() []locales.PluralRule {
	return jv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'jv'
func (jv *jv) PluralsRange() []locales.PluralRule {
	return jv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'jv'
func (jv *jv) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'jv'
func (jv *jv) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'jv'
func (jv *jv) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (jv *jv) MonthAbbreviated(month time.Month) string {
	return jv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (jv *jv) MonthsAbbreviated() []string {
	return jv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (jv *jv) MonthNarrow(month time.Month) string {
	return jv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (jv *jv) MonthsNarrow() []string {
	return jv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (jv *jv) MonthWide(month time.Month) string {
	return jv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (jv *jv) MonthsWide() []string {
	return jv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (jv *jv) WeekdayAbbreviated(weekday time.Weekday) string {
	return jv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (jv *jv) WeekdaysAbbreviated() []string {
	return jv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (jv *jv) WeekdayNarrow(weekday time.Weekday) string {
	return jv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (jv *jv) WeekdaysNarrow() []string {
	return jv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (jv *jv) WeekdayShort(weekday time.Weekday) string {
	return jv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (jv *jv) WeekdaysShort() []string {
	return jv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (jv *jv) WeekdayWide(weekday time.Weekday) string {
	return jv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (jv *jv) WeekdaysWide() []string {
	return jv.daysWide
}

// Decimal returns the decimal point of number
func (jv *jv) Decimal() string {
	return jv.decimal
}

// Group returns the group of number
func (jv *jv) Group() string {
	return jv.group
}

// Group returns the minus sign of number
func (jv *jv) Minus() string {
	return jv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'jv' and handles both Whole and Real numbers based on 'v'
func (jv *jv) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, jv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'jv' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (jv *jv) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, jv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, jv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'jv'
func (jv *jv) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := jv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
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

	for j := len(jv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, jv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, jv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, jv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'jv'
// in accounting notation.
func (jv *jv) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := jv.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
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

		for j := len(jv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, jv.currencyNegativePrefix[j])
		}

		b = append(b, jv.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(jv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, jv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, jv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'jv'
func (jv *jv) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'jv'
func (jv *jv) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'jv'
func (jv *jv) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'jv'
func (jv *jv) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, jv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'jv'
func (jv *jv) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'jv'
func (jv *jv) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'jv'
func (jv *jv) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'jv'
func (jv *jv) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := jv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
