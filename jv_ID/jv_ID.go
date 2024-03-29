package jv_ID

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type jv_ID struct {
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

// New returns a new instance of translator for the 'jv_ID' locale
func New() locales.Translator {
	return &jv_ID{
		locale:                 "jv_ID",
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
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
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
		timezones:              map[string]string{"JST": "Wektu Standar Jepang", "PST": "Wektu Standar Pasifik", "AKDT": "Wektu Ketigo Alaska", "HNPM": "Wektu Standar Santa Pierre lan Miquelon", "HADT": "Wektu Ketigo Hawaii-Aleutian", "GYT": "Wektu Guyana", "HNOG": "Wektu Standar Grinland Kulon", "HEOG": "Wektu Ketigo Grinland Kulon", "ART": "Wektu Standar Argentina", "ACST": "Wektu Standar Australia Tengah", "WIB": "Wektu Indonesia sisih Kulon", "LHST": "Wektu Standar Lord Howe", "ChST": "Wektu Standar Chamorro", "CLST": "Wektu Ketigo Chili", "HEEG": "Wektu Ketigo Grinland Wetan", "NZDT": "Wektu Ketigo Selandia Anyar", "MYT": "Wektu Malaysia", "COST": "Wektu Ketigo Kolombia", "ADT": "Wektu Ketigo Atlantik", "AWST": "Wektu Standar Australia sisih Kulon", "UYST": "Wektu Ketigo Uruguay", "WART": "Wektu Standar Argentina sisih Kulon", "ARST": "Wektu Ketigo Argentina", "HNCU": "Wektu Standar Kuba", "CST": "Wektu Standar Tengah", "HEPMX": "Wektu Ketigo Pasifik Meksiko", "HKST": "Wektu Ketigo Hong Kong", "TMST": "Wektu Ketigo Turkmenistan", "WARST": "Wektu Ketigo Argentina sisih Kulon", "CHADT": "Wektu Ketigo Chatham", "EDT": "Wektu Ketigo sisih Wetah", "GMT": "Wektu Rerata Greenwich", "HNEG": "Wektu Standar Grinland Wetan", "AKST": "Wektu Standar Alaska", "ACDT": "Wektu Ketigo Australia Tengah", "WEZ": "Wektu Standar Eropa sisih Kulon", "COT": "Wektu Standar Kolombia", "MST": "MST", "AEST": "Wektu Standar Australia sisih Wetan", "HAT": "Wektu Ketigo Newfoundland", "JDT": "Wektu Ketigo Jepang", "IST": "Wektu Standar India", "LHDT": "Wektu Ketigo Lord Howe", "HAST": "Wektu Standar Hawaii-Aleutian", "HKT": "Wektu Standar Hong Kong", "OEZ": "Wektu Standar Eropa sisih Wetan", "VET": "Wektu Venezuela", "HECU": "Wektu Ketigo Kuba", "BT": "Wektu Bhutan", "WAST": "Wektu Ketigo Afrika Kulon", "AWDT": "Wektu Ketigo Australia sisih Kulon", "SGT": "Wektu Standar Singapura", "WITA": "Wektu Indonesia Tengah", "HENOMX": "Wektu Ketigo Meksiko Lor-Kulon", "HNT": "Wektu Standar Newfoundland", "MEZ": "Wektu Standar Eropa Tengah", "CHAST": "Wektu Standar Chatham", "∅∅∅": "Wektu Ketigo Brasilia", "HNPMX": "Wektu Standar Pasifik Meksiko", "MDT": "MDT", "ECT": "Wektu Ekuador", "EST": "Wektu Standar sisih Wetan", "BOT": "Wektu Bolivia", "MESZ": "Wektu Ketigo Eropa Tengah", "CDT": "Wektu Ketigo Tengah", "UYT": "Wektu Standar Uruguay", "ACWST": "Wektu Standar Australia Tengah sisih Kulon", "WESZ": "Wektu Ketigo Eropa sisih Kulon", "OESZ": "Wektu Ketigo Eropa sisih Wetan", "PDT": "Wektu Ketigo Pasifik", "GFT": "Wektu Guiana Prancis", "ACWDT": "Wektu Ketigo Australia Tengah sisih Kulon", "SRT": "Wektu Suriname", "AST": "Wektu Standar Atlantik", "WAT": "Wektu Standar Afrika Kulon", "TMT": "Wektu Standar Turkmenistan", "NZST": "Wektu Standar Selandia Anyar", "HNNOMX": "Wektu Standar Meksiko Lor-Kulon", "CLT": "Wektu Standar Chili", "AEDT": "Wektu Ketigo Australia sisih Wetan", "EAT": "Wektu Afrika Wetan", "CAT": "Wektu Afrika Tengah", "WIT": "Wektu Indonesia sisih Wetan", "SAST": "Wektu Standar Afrika Kidul", "HEPM": "Wektu Ketigo Santa Pierre lan Miquelon"},
	}
}

// Locale returns the current translators string locale
func (jv *jv_ID) Locale() string {
	return jv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'jv_ID'
func (jv *jv_ID) PluralsCardinal() []locales.PluralRule {
	return jv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'jv_ID'
func (jv *jv_ID) PluralsOrdinal() []locales.PluralRule {
	return jv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'jv_ID'
func (jv *jv_ID) PluralsRange() []locales.PluralRule {
	return jv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'jv_ID'
func (jv *jv_ID) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'jv_ID'
func (jv *jv_ID) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'jv_ID'
func (jv *jv_ID) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (jv *jv_ID) MonthAbbreviated(month time.Month) string {
	return jv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (jv *jv_ID) MonthsAbbreviated() []string {
	return jv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (jv *jv_ID) MonthNarrow(month time.Month) string {
	return jv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (jv *jv_ID) MonthsNarrow() []string {
	return jv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (jv *jv_ID) MonthWide(month time.Month) string {
	return jv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (jv *jv_ID) MonthsWide() []string {
	return jv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (jv *jv_ID) WeekdayAbbreviated(weekday time.Weekday) string {
	return jv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (jv *jv_ID) WeekdaysAbbreviated() []string {
	return jv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (jv *jv_ID) WeekdayNarrow(weekday time.Weekday) string {
	return jv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (jv *jv_ID) WeekdaysNarrow() []string {
	return jv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (jv *jv_ID) WeekdayShort(weekday time.Weekday) string {
	return jv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (jv *jv_ID) WeekdaysShort() []string {
	return jv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (jv *jv_ID) WeekdayWide(weekday time.Weekday) string {
	return jv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (jv *jv_ID) WeekdaysWide() []string {
	return jv.daysWide
}

// Decimal returns the decimal point of number
func (jv *jv_ID) Decimal() string {
	return jv.decimal
}

// Group returns the group of number
func (jv *jv_ID) Group() string {
	return jv.group
}

// Group returns the minus sign of number
func (jv *jv_ID) Minus() string {
	return jv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'jv_ID' and handles both Whole and Real numbers based on 'v'
func (jv *jv_ID) FmtNumber(num float64, v uint64) string {

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

// FmtPercent returns 'num' with digits/precision of 'v' for 'jv_ID' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (jv *jv_ID) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'jv_ID'
func (jv *jv_ID) FmtCurrency(num float64, v uint64, currency currency.Type) string {

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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'jv_ID'
// in accounting notation.
func (jv *jv_ID) FmtAccounting(num float64, v uint64, currency currency.Type) string {

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

// FmtDateShort returns the short date representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtDateLong(t time.Time) string {

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

// FmtDateFull returns the full date representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtDateFull(t time.Time) string {

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

// FmtTimeShort returns the short time representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtTimeShort(t time.Time) string {

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

// FmtTimeMedium returns the medium time representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtTimeMedium(t time.Time) string {

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

// FmtTimeLong returns the long time representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtTimeLong(t time.Time) string {

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

// FmtTimeFull returns the full time representation of 't' for 'jv_ID'
func (jv *jv_ID) FmtTimeFull(t time.Time) string {

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
