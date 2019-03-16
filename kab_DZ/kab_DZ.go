package kab_DZ

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type kab_DZ struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'kab_DZ' locale
func New() locales.Translator {
	return &kab_DZ{
		locale:             "kab_DZ",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		decimal:            ",",
		group:              " ",
		minus:              "-",
		percent:            "%",
		perMille:           "‰",
		timeSeparator:      ":",
		inifinity:          "∞",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "Yen", "Fur", "Meɣ", "Yeb", "May", "Yun", "Yul", "Ɣuc", "Cte", "Tub", "Nun", "Duǧ"},
		monthsNarrow:       []string{"", "Y", "F", "Ɣ", "B", "M", "N", "L", "C", "T", "R", "W", "D"},
		monthsWide:         []string{"", "Yennayer", "Fuṛar", "Meɣres", "Yebrir", "Mayyu", "Yunyu", "Yulyu", "Ɣuct", "Ctembeṛ", "Tubeṛ", "Nunembeṛ", "Duǧembeṛ"},
		daysAbbreviated:    []string{"Yan", "San", "Kraḍ", "Kuẓ", "Sam", "Sḍis", "Say"},
		daysNarrow:         []string{"C", "R", "A", "H", "M", "S", "D"},
		daysShort:          []string{"Cr", "Ri", "Ra", "Hd", "Mh", "Sm", "Sd"},
		daysWide:           []string{"Yanass", "Sanass", "Kraḍass", "Kuẓass", "Samass", "Sḍisass", "Sayass"},
		periodsAbbreviated: []string{"n tufat", "n tmeddit"},
		periodsNarrow:      []string{"f", "m"},
		periodsWide:        []string{"n tufat", "n tmeddit"},
		erasAbbreviated:    []string{"snd. T.Ɛ", "sld. T.Ɛ"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"send talalit n Ɛisa", "seld talalit n Ɛisa"},
		timezones:          map[string]string{"MEZ": "Akud amagnu n Turuft Talemmast", "AKDT": "Akud n Unebdu n Alaska", "BT": "Akud n Butan", "WIB": "Akud n Umalu n Indunisya", "EDT": "Akud n Unebdu n Usammar Agafa n Marikan", "ACWST": "Akud Amagnu n Tlemmast n Umalu n Ustṛalya", "HEPMX": "Akud Amelwi n Unebdu n Miksik", "SRT": "Akud n Surinam", "AEST": "Akud Amagnu n Ustṛalya n Usammar", "WIT": "Akud n Usammar n Indunisya", "ACDT": "Akud n Unebdu n Ustralya Talemmast", "WEZ": "Akud amagnu n turuft n umalu", "TMT": "Akud Amagnu n Ṭurkmanistan", "NZDT": "Akud n Unebdu Ziland Tamaynut", "CHAST": "Akud Amagnu n Catham", "CHADT": "Akud n Unebdu Catham", "HNPM": "Akud Amagnu n San Pyir & Miklun", "COST": "Akud n Unebdu n Kulumbya", "HNCU": "Akud Amagnu n Kuba", "HNPMX": "Akud amagnu Amelwi n Miksik", "GYT": "Akud n Gwiyan", "HNOG": "Akud Amagnu n Grinland n Umalu", "AWDT": "Akud n Unebdu Ustṛalya n Umalu", "OEZ": "Akud amagnu n Turuft n Usammar", "WESZ": "Akud n unebdu turuft n umalu", "LHDT": "Akud n Unebdu n Lord Howe", "HADT": "Akud n Unebu n Haway-Aliwsyan", "CLST": "Akud n Unebdu n Cili", "WAT": "Akud amagnu n tefriqt n umalu", "ART": "Akud Amagnu n Arjuntin", "AKST": "Akud Amagnu n Alaska", "ChST": "Akud Amagnu n Camuṛṛu", "MST": "MST", "AEDT": "Akud n Unebdu n Ustṛalya n Usammar", "JDT": "Akud n Unebdu n Japun", "CAT": "Akud n tefriqt talemmast", "MDT": "MDT", "HAT": "Akud n Unebdu n Wakal Amaynut", "BOT": "Akud n Bulivi", "MESZ": "Akud n unebdu n Turuft Talemmast", "CDT": "Akud n Unebdu n Tlemmast n Marikan", "∅∅∅": "Akud n Unebdu n Bṛazilya", "HKT": "Akud Amagnu n Hung Kung", "PST": "Akud Amelwi Amagnu n Marikan n Ugafa", "HNT": "Akud Amagnu n Wakal Amaynut", "ACST": "Akud Amagnu n Ustralya Talemmast", "OESZ": "Akud n unebdu n Turuft n Usammar", "EAT": "Akud n tefriqt n usammar", "ACWDT": "Akud n Unebdu n Tlemmast n Umalu n Ustṛalya", "WITA": "Akud n Tlemmast n Indunisya", "TMST": "Akud n Unebdu n Ṭurkmanistan", "VET": "Akud n Vinizwila", "SAST": "Akud amagnu n tefriqt n unẓul", "HKST": "Akud n Unebdu n Hung Kung", "JST": "Akud Amagnu n Japun", "HEEG": "Akud n Unebdu n Grinland n Usammar", "HEPM": "Akud n Unebdu n San Pyir & Miklun", "WAST": "Akud n unebdu n tefriqt n umalu", "HNEG": "Akud Amagnu n Grinland n Usammar", "IST": "Akud Amagnu n Hend", "UYT": "Akud amagnu n Urugway", "UYST": "Akud n Unebdu n Urugway", "HEOG": "Akud n Unebdu n Grinland n Umalu", "SGT": "Akud Amagnu n Sangapur", "AWST": "Akud Amagnu n Ustṛalya n Umalu", "HECU": "Akud n Unebdu n Kuba", "PDT": "Akud Amelwi n Unebdu n Marikan n Ugafa", "LHST": "Akud Amagnu n Lord Howe", "EST": "Akud Amagnu n Usammar Agafa n Marikan", "GFT": "Akud n Gwiyan Tafṛansist", "HAST": "Akud Amagnu n Haway-Aliwsyan", "COT": "Akud Amagnu n Kulumbya", "GMT": "Akud alemmas n Greenwich", "WARST": "Akud n Unebdu n Arjuntin n Usammar", "HENOMX": "Akud n Unebdu n Ugafa Amalu n Miksik", "ARST": "Akud n Unebdu n Arjuntin", "AST": "Akud Amagnu Aṭlasan", "ADT": "Akud Aṭlasan n Unebdu", "WART": "Akud Amagnu n Arjuntin n Usammar", "CST": "Akud Amagnu n Tlemmast n Marikan", "MYT": "Akud n Malizya", "CLT": "Akud Amagnu n Cili", "ECT": "Akud n Ikwaṭur", "HNNOMX": "Akud Amagnu n Ugafa Amalu n Miksik", "NZST": "Akud Amagnu n Ziland Tamaynut"},
	}
}

// Locale returns the current translators string locale
func (kab *kab_DZ) Locale() string {
	return kab.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kab_DZ'
func (kab *kab_DZ) PluralsCardinal() []locales.PluralRule {
	return kab.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kab_DZ'
func (kab *kab_DZ) PluralsOrdinal() []locales.PluralRule {
	return kab.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kab_DZ'
func (kab *kab_DZ) PluralsRange() []locales.PluralRule {
	return kab.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kab_DZ'
func (kab *kab_DZ) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 0 || i == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kab_DZ'
func (kab *kab_DZ) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kab_DZ'
func (kab *kab_DZ) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kab *kab_DZ) MonthAbbreviated(month time.Month) string {
	return kab.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kab *kab_DZ) MonthsAbbreviated() []string {
	return kab.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kab *kab_DZ) MonthNarrow(month time.Month) string {
	return kab.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kab *kab_DZ) MonthsNarrow() []string {
	return kab.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kab *kab_DZ) MonthWide(month time.Month) string {
	return kab.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kab *kab_DZ) MonthsWide() []string {
	return kab.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kab *kab_DZ) WeekdayAbbreviated(weekday time.Weekday) string {
	return kab.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kab *kab_DZ) WeekdaysAbbreviated() []string {
	return kab.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kab *kab_DZ) WeekdayNarrow(weekday time.Weekday) string {
	return kab.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kab *kab_DZ) WeekdaysNarrow() []string {
	return kab.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kab *kab_DZ) WeekdayShort(weekday time.Weekday) string {
	return kab.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kab *kab_DZ) WeekdaysShort() []string {
	return kab.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kab *kab_DZ) WeekdayWide(weekday time.Weekday) string {
	return kab.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kab *kab_DZ) WeekdaysWide() []string {
	return kab.daysWide
}

// Decimal returns the decimal point of number
func (kab *kab_DZ) Decimal() string {
	return kab.decimal
}

// Group returns the group of number
func (kab *kab_DZ) Group() string {
	return kab.group
}

// Group returns the minus sign of number
func (kab *kab_DZ) Minus() string {
	return kab.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kab_DZ' and handles both Whole and Real numbers based on 'v'
func (kab *kab_DZ) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kab.group) - 1; j >= 0; j-- {
					b = append(b, kab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kab_DZ' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kab *kab_DZ) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, kab.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kab_DZ'
func (kab *kab_DZ) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kab.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kab.group) - 1; j >= 0; j-- {
					b = append(b, kab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kab.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kab.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kab_DZ'
// in accounting notation.
func (kab *kab_DZ) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kab.currencies[currency]
	l := len(s) + len(symbol) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kab.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(kab.group) - 1; j >= 0; j-- {
					b = append(b, kab.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, kab.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kab.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kab.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kab.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, kab.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kab.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kab_DZ'
func (kab *kab_DZ) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kab.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kab.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, kab.periodsAbbreviated[0]...)
	} else {
		b = append(b, kab.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kab.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
