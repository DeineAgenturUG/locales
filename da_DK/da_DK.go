package da_DK

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type da_DK struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
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

// New returns a new instance of translator for the 'da_DK' locale
func New() locales.Translator {
	return &da_DK{
		locale:                 "da_DK",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ".",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mar.", "apr.", "maj", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januar", "februar", "marts", "april", "maj", "juni", "juli", "august", "september", "oktober", "november", "december"},
		daysAbbreviated:        []string{"søn.", "man.", "tir.", "ons.", "tor.", "fre.", "lør."},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sø", "ma", "ti", "on", "to", "fr", "lø"},
		daysWide:               []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"f.Kr.", "e.Kr."},
		erasNarrow:             []string{"fKr", "eKr"},
		erasWide:               []string{"f.Kr.", "e.Kr."},
		timezones:              map[string]string{"HNCU": "Cubansk normaltid", "CDT": "Central-sommertid", "ACWDT": "Vestlig centralaustralsk sommertid", "EAT": "Østafrikansk tid", "ACST": "Centralaustralsk normaltid", "GYT": "Guyana-tid", "AST": "Atlantic-normaltid", "OESZ": "Østeuropæisk sommertid", "CST": "Central-normaltid", "WART": "Vestargentinsk normaltid", "NZDT": "Newzealandsk sommertid", "ACWST": "Vestlig centralaustralsk normaltid", "HEEG": "Østgrønlandsk sommertid", "PST": "Pacific-normaltid", "HNPM": "Saint Pierre- og Miquelon-normaltid", "WESZ": "Vesteuropæisk sommertid", "GFT": "Fransk Guyana-tid", "HKST": "Hongkong-sommertid", "MEZ": "Centraleuropæisk normaltid", "LHDT": "Lord Howe-sommertid", "HNOG": "Vestgrønlandsk normaltid", "HENOMX": "Nordvestmexicansk sommertid", "SRT": "Surinam-tid", "NZST": "Newzealandsk normaltid", "MST": "Macao-normaltid", "WITA": "Centralindonesisk tid", "TMT": "Turkmensk normaltid", "HADT": "Hawaii-Aleutian-sommertid", "COT": "Colombiansk normaltid", "GMT": "GMT", "WAT": "Vestafrikansk normaltid", "TMST": "Turkmensk sommertid", "HECU": "Cubansk sommertid", "HNEG": "Østgrønlandsk normaltid", "BT": "Bhutan-tid", "LHST": "Lord Howe-normaltid", "MYT": "Malaysia-tid", "UYT": "Uruguayansk normaltid", "ART": "Argentinsk normaltid", "WEZ": "Vesteuropæisk normaltid", "ChST": "Chamorro-tid", "VET": "Venezuelansk tid", "WAST": "Vestafrikansk sommertid", "SGT": "Singapore-tid", "WIT": "Østindonesisk tid", "HNNOMX": "Nordvestmexicansk normaltid", "CHADT": "Chatham-sommertid", "EDT": "Eastern-sommertid", "HAST": "Hawaii-Aleutian-normaltid", "UYST": "Uruguayansk sommertid", "HEPM": "Saint Pierre- og Miquelon-sommertid", "AWDT": "Vestaustralsk sommertid", "OEZ": "Østeuropæisk normaltid", "JDT": "Japansk sommertid", "SAST": "Sydafrikansk tid", "AKDT": "Alaska-sommertid", "IST": "Indisk normaltid", "WIB": "Vestindonesisk tid", "COST": "Colombiansk sommertid", "CHAST": "Chatham-normaltid", "HEPMX": "Mexicansk Pacific-sommertid", "AWST": "Vestaustralsk normaltid", "JST": "Japansk normaltid", "CAT": "Centralafrikansk tid", "ACDT": "Centralaustralsk sommertid", "∅∅∅": "Brasiliansk sommertid", "CLT": "Chilensk normaltid", "HNPMX": "Mexicansk Pacific-normaltid", "AEDT": "Østaustralsk sommertid", "HAT": "Newfoundlandsk sommertid", "ARST": "Argentinsk sommertid", "MDT": "Macao-sommertid", "ECT": "Ecuadoriansk tid", "ADT": "Atlantic-sommertid", "AKST": "Alaska-normaltid", "CLST": "Chilensk sommertid", "AEST": "Østaustralsk normaltid", "HEOG": "Vestgrønlandsk sommertid", "PDT": "Pacific-sommertid", "HNT": "Newfoundlandsk normaltid", "BOT": "Boliviansk tid", "MESZ": "Centraleuropæisk sommertid", "EST": "Eastern-normaltid", "HKT": "Hongkong-normaltid", "WARST": "Vestargentinsk sommertid"},
	}
}

// Locale returns the current translators string locale
func (da *da_DK) Locale() string {
	return da.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'da_DK'
func (da *da_DK) PluralsCardinal() []locales.PluralRule {
	return da.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'da_DK'
func (da *da_DK) PluralsOrdinal() []locales.PluralRule {
	return da.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'da_DK'
func (da *da_DK) PluralsRange() []locales.PluralRule {
	return da.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'da_DK'
func (da *da_DK) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	t := locales.T(n, v)

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'da_DK'
func (da *da_DK) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'da_DK'
func (da *da_DK) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := da.CardinalPluralRule(num1, v1)
	end := da.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (da *da_DK) MonthAbbreviated(month time.Month) string {
	return da.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (da *da_DK) MonthsAbbreviated() []string {
	return da.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (da *da_DK) MonthNarrow(month time.Month) string {
	return da.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (da *da_DK) MonthsNarrow() []string {
	return da.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (da *da_DK) MonthWide(month time.Month) string {
	return da.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (da *da_DK) MonthsWide() []string {
	return da.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (da *da_DK) WeekdayAbbreviated(weekday time.Weekday) string {
	return da.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (da *da_DK) WeekdaysAbbreviated() []string {
	return da.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (da *da_DK) WeekdayNarrow(weekday time.Weekday) string {
	return da.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (da *da_DK) WeekdaysNarrow() []string {
	return da.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (da *da_DK) WeekdayShort(weekday time.Weekday) string {
	return da.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (da *da_DK) WeekdaysShort() []string {
	return da.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (da *da_DK) WeekdayWide(weekday time.Weekday) string {
	return da.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (da *da_DK) WeekdaysWide() []string {
	return da.daysWide
}

// Decimal returns the decimal point of number
func (da *da_DK) Decimal() string {
	return da.decimal
}

// Group returns the group of number
func (da *da_DK) Group() string {
	return da.group
}

// Group returns the minus sign of number
func (da *da_DK) Minus() string {
	return da.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'da_DK' and handles both Whole and Real numbers based on 'v'
func (da *da_DK) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, da.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, da.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'da_DK' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (da *da_DK) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, da.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, da.percentSuffix...)

	b = append(b, da.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'da_DK'
func (da *da_DK) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := da.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, da.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, da.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, da.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, da.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'da_DK'
// in accounting notation.
func (da *da_DK) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := da.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, da.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, da.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, da.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, da.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, da.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, da.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'da_DK'
func (da *da_DK) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'da_DK'
func (da *da_DK) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, da.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'da_DK'
func (da *da_DK) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, da.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'da_DK'
func (da *da_DK) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, da.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0x64, 0x65, 0x6e}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, da.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'da_DK'
func (da *da_DK) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'da_DK'
func (da *da_DK) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'da_DK'
func (da *da_DK) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'da_DK'
func (da *da_DK) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := da.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
