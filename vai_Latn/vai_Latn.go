package vai_Latn

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type vai_Latn struct {
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

// New returns a new instance of translator for the 'vai_Latn' locale
func New() locales.Translator {
	return &vai_Latn{
		locale:                 "vai_Latn",
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
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "$", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "ꖨꖕꔞ", "ꕒꕡ", "ꕾꖺ", "ꖢꖕ", "ꖑꕱ", "ꖱꘋ", "ꖱꕞ", "ꗛꔕ", "ꕢꕌ", "ꕭꖃ", "ꔞꘋ", "ꖨꖕꗏ"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "luukao kemã", "ɓandaɓu", "vɔɔ", "fulu", "goo", "6", "7", "kɔnde", "saah", "galo", "kenpkato ɓololɔ", "luukao lɔma"},
		daysNarrow:             []string{"S", "M", "T", "W", "T", "F", "S"},
		daysWide:               []string{"lahadi", "tɛɛnɛɛ", "talata", "alaba", "aimisa", "aijima", "siɓiti"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"", ""},
		periodsWide:            []string{"", ""},
		erasAbbreviated:        []string{"BCE", "CE"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"OESZ": "OESZ", "WITA": "WITA", "SAST": "SAST", "MEZ": "MEZ", "SGT": "SGT", "TMST": "TMST", "TMT": "TMT", "WART": "WART", "WAST": "WAST", "AST": "AST", "HNOG": "HNOG", "AWST": "AWST", "WIT": "WIT", "AKDT": "AKDT", "MESZ": "MESZ", "ARST": "ARST", "CST": "CST", "WAT": "WAT", "AEDT": "AEDT", "HKST": "HKST", "OEZ": "OEZ", "CHAST": "CHAST", "CDT": "CDT", "UYT": "UYT", "ADT": "ADT", "HECU": "HECU", "CAT": "CAT", "NZST": "NZST", "EST": "EST", "CLT": "CLT", "MDT": "MDT", "EAT": "EAT", "JDT": "JDT", "WESZ": "WESZ", "UYST": "UYST", "COST": "COST", "ECT": "ECT", "AEST": "AEST", "VET": "VET", "HENOMX": "HENOMX", "ChST": "ChST", "GFT": "GFT", "ACWST": "ACWST", "HEPMX": "HEPMX", "BOT": "BOT", "HADT": "HADT", "HNPMX": "HNPMX", "HAT": "HAT", "AKST": "AKST", "HEPM": "HEPM", "LHST": "LHST", "HAST": "HAST", "COT": "COT", "HNCU": "HNCU", "ART": "ART", "GYT": "GYT", "WARST": "WARST", "NZDT": "NZDT", "ACDT": "ACDT", "IST": "IST", "EDT": "EDT", "SRT": "SRT", "MST": "MST", "HNT": "HNT", "HNEG": "HNEG", "LHDT": "LHDT", "ACWDT": "ACWDT", "HEEG": "HEEG", "WEZ": "WEZ", "ACST": "ACST", "AWDT": "AWDT", "PDT": "PDT", "JST": "JST", "CHADT": "CHADT", "HNPM": "HNPM", "CLST": "CLST", "HEOG": "HEOG", "∅∅∅": "∅∅∅", "BT": "BT", "PST": "PST", "WIB": "WIB", "MYT": "MYT", "GMT": "GMT", "HKT": "HKT", "HNNOMX": "HNNOMX"},
	}
}

// Locale returns the current translators string locale
func (vai *vai_Latn) Locale() string {
	return vai.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'vai_Latn'
func (vai *vai_Latn) PluralsCardinal() []locales.PluralRule {
	return vai.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'vai_Latn'
func (vai *vai_Latn) PluralsOrdinal() []locales.PluralRule {
	return vai.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'vai_Latn'
func (vai *vai_Latn) PluralsRange() []locales.PluralRule {
	return vai.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'vai_Latn'
func (vai *vai_Latn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'vai_Latn'
func (vai *vai_Latn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'vai_Latn'
func (vai *vai_Latn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (vai *vai_Latn) MonthAbbreviated(month time.Month) string {
	return vai.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (vai *vai_Latn) MonthsAbbreviated() []string {
	return vai.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (vai *vai_Latn) MonthNarrow(month time.Month) string {
	return vai.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (vai *vai_Latn) MonthsNarrow() []string {
	return vai.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (vai *vai_Latn) MonthWide(month time.Month) string {
	return vai.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (vai *vai_Latn) MonthsWide() []string {
	return vai.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (vai *vai_Latn) WeekdayAbbreviated(weekday time.Weekday) string {
	return vai.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (vai *vai_Latn) WeekdaysAbbreviated() []string {
	return vai.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (vai *vai_Latn) WeekdayNarrow(weekday time.Weekday) string {
	return vai.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (vai *vai_Latn) WeekdaysNarrow() []string {
	return vai.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (vai *vai_Latn) WeekdayShort(weekday time.Weekday) string {
	return vai.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (vai *vai_Latn) WeekdaysShort() []string {
	return vai.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (vai *vai_Latn) WeekdayWide(weekday time.Weekday) string {
	return vai.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (vai *vai_Latn) WeekdaysWide() []string {
	return vai.daysWide
}

// Decimal returns the decimal point of number
func (vai *vai_Latn) Decimal() string {
	return vai.decimal
}

// Group returns the group of number
func (vai *vai_Latn) Group() string {
	return vai.group
}

// Group returns the minus sign of number
func (vai *vai_Latn) Minus() string {
	return vai.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'vai_Latn' and handles both Whole and Real numbers based on 'v'
func (vai *vai_Latn) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vai.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vai.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vai.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'vai_Latn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (vai *vai_Latn) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'vai_Latn'
func (vai *vai_Latn) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vai.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vai.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vai.group[0])
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

	if num < 0 {
		b = append(b, vai.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vai.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'vai_Latn'
// in accounting notation.
func (vai *vai_Latn) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vai.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vai.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vai.group[0])
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

		b = append(b, vai.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vai.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, vai.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vai.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vai.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, vai.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vai.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, vai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, vai.periodsAbbreviated[0]...)
	} else {
		b = append(b, vai.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, vai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vai.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, vai.periodsAbbreviated[0]...)
	} else {
		b = append(b, vai.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, vai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vai.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, vai.periodsAbbreviated[0]...)
	} else {
		b = append(b, vai.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'vai_Latn'
func (vai *vai_Latn) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, vai.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vai.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, vai.periodsAbbreviated[0]...)
	} else {
		b = append(b, vai.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := vai.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
