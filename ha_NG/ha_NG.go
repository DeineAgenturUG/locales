package ha_NG

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type ha_NG struct {
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

// New returns a new instance of translator for the 'ha_NG' locale
func New() locales.Translator {
	return &ha_NG{
		locale:                 "ha_NG",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Jan", "Fab", "Mar", "Afi", "May", "Yun", "Yul", "Agu", "Sat", "Okt", "Nuw", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "Y", "Y", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Janairu", "Faburairu", "Maris", "Afirilu", "Mayu", "Yuni", "Yuli", "Agusta", "Satumba", "Oktoba", "Nuwamba", "Disamba"},
		daysAbbreviated:        []string{"Lah", "Lit", "Tal", "Lar", "Alh", "Jum", "Asa"},
		daysNarrow:             []string{"L", "L", "T", "L", "A", "J", "A"},
		daysShort:              []string{"Lh", "Li", "Ta", "Lr", "Al", "Ju", "As"},
		daysWide:               []string{"Lahadi", "Litinin", "Talata", "Laraba", "Alhamis", "Jummaʼa", "Asabar"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"KHAI", "BHAI"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Kafin haihuwar annab", "Bayan haihuwar annab"},
		timezones:              map[string]string{"MYT": "MYT", "ART": "ART", "CHAST": "CHAST", "LHST": "LHST", "HENOMX": "HENOMX", "ECT": "ECT", "TMST": "TMST", "EAT": "EAT", "HKST": "HKST", "HNCU": "HNCU", "ARST": "ARST", "CHADT": "CHADT", "AST": "AST", "WARST": "WARST", "VET": "VET", "WESZ": "WESZ", "WAT": "WAT", "AWDT": "AWDT", "HECU": "HECU", "BOT": "BOT", "WEZ": "WEZ", "OEZ": "OEZ", "GFT": "GFT", "UYT": "UYT", "CLST": "CLST", "JST": "JST", "TMT": "TMT", "AKST": "AKST", "CST": "CST", "EST": "EST", "HNNOMX": "HNNOMX", "AEDT": "AEDT", "OESZ": "OESZ", "WITA": "WITA", "PDT": "PDT", "MEZ": "MEZ", "MST": "MST", "HEPM": "HEPM", "LHDT": "LHDT", "HADT": "HADT", "HEOG": "HEOG", "WART": "WART", "PST": "PST", "MESZ": "MESZ", "ACST": "ACST", "GYT": "GYT", "GMT": "GMT", "HNEG": "HNEG", "CLT": "CLT", "HNPMX": "HNPMX", "HEPMX": "HEPMX", "WAST": "WAST", "NZST": "NZST", "ACDT": "ACDT", "IST": "IST", "COT": "COT", "EDT": "EDT", "ACWDT": "ACWDT", "AWST": "AWST", "JDT": "JDT", "AKDT": "AKDT", "HNPM": "HNPM", "MDT": "MDT", "HAST": "HAST", "HKT": "HKT", "HAT": "HAT", "BT": "BT", "WIB": "WIB", "ChST": "ChST", "AEST": "AEST", "NZDT": "NZDT", "HNT": "HNT", "CAT": "CAT", "CDT": "CDT", "SRT": "SRT", "SGT": "SGT", "SAST": "SAST", "ACWST": "ACWST", "COST": "COST", "HNOG": "HNOG", "ADT": "ADT", "WIT": "WIT", "HEEG": "HEEG", "∅∅∅": "∅∅∅", "UYST": "UYST"},
	}
}

// Locale returns the current translators string locale
func (ha *ha_NG) Locale() string {
	return ha.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ha_NG'
func (ha *ha_NG) PluralsCardinal() []locales.PluralRule {
	return ha.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ha_NG'
func (ha *ha_NG) PluralsOrdinal() []locales.PluralRule {
	return ha.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ha_NG'
func (ha *ha_NG) PluralsRange() []locales.PluralRule {
	return ha.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ha_NG'
func (ha *ha_NG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ha_NG'
func (ha *ha_NG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ha_NG'
func (ha *ha_NG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ha *ha_NG) MonthAbbreviated(month time.Month) string {
	return ha.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ha *ha_NG) MonthsAbbreviated() []string {
	return ha.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ha *ha_NG) MonthNarrow(month time.Month) string {
	return ha.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ha *ha_NG) MonthsNarrow() []string {
	return ha.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ha *ha_NG) MonthWide(month time.Month) string {
	return ha.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ha *ha_NG) MonthsWide() []string {
	return ha.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ha *ha_NG) WeekdayAbbreviated(weekday time.Weekday) string {
	return ha.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ha *ha_NG) WeekdaysAbbreviated() []string {
	return ha.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ha *ha_NG) WeekdayNarrow(weekday time.Weekday) string {
	return ha.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ha *ha_NG) WeekdaysNarrow() []string {
	return ha.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ha *ha_NG) WeekdayShort(weekday time.Weekday) string {
	return ha.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ha *ha_NG) WeekdaysShort() []string {
	return ha.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ha *ha_NG) WeekdayWide(weekday time.Weekday) string {
	return ha.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ha *ha_NG) WeekdaysWide() []string {
	return ha.daysWide
}

// Decimal returns the decimal point of number
func (ha *ha_NG) Decimal() string {
	return ha.decimal
}

// Group returns the group of number
func (ha *ha_NG) Group() string {
	return ha.group
}

// Group returns the minus sign of number
func (ha *ha_NG) Minus() string {
	return ha.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ha_NG' and handles both Whole and Real numbers based on 'v'
func (ha *ha_NG) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 1 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ha.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ha_NG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ha *ha_NG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 1
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ha.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ha_NG'
func (ha *ha_NG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ha.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ha.group[0])
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

	for j := len(ha.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ha.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ha.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ha.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ha_NG'
// in accounting notation.
func (ha *ha_NG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ha.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ha.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ha.group[0])
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

		for j := len(ha.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ha.currencyNegativePrefix[j])
		}

		b = append(b, ha.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ha.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ha.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ha.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ha.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ha.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ha_NG'
func (ha *ha_NG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ha.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ha.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
