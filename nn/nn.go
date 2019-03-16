package nn

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type nn struct {
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

// New returns a new instance of translator for the 'nn' locale
func New() locales.Translator {
	return &nn{
		locale:                 "nn",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "kr", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "Db", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mars", "apr.", "mai", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
		daysAbbreviated:        []string{"sø.", "må.", "ty.", "on.", "to.", "fr.", "la."},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sø.", "må.", "ty.", "on.", "to.", "fr.", "la."},
		daysWide:               []string{"søndag", "måndag", "tysdag", "onsdag", "torsdag", "fredag", "laurdag"},
		periodsAbbreviated:     []string{"f.m.", "e.m."},
		periodsNarrow:          []string{"f.m.", "e.m."},
		periodsWide:            []string{"formiddag", "ettermiddag"},
		erasAbbreviated:        []string{"f.Kr.", "e.Kr."},
		erasNarrow:             []string{"f.Kr.", "e.Kr."},
		erasWide:               []string{"f.Kr.", "e.Kr."},
		timezones:              map[string]string{"COT": "kolombiansk normaltid", "COST": "kolombiansk sumartid", "AKST": "alaskisk normaltid", "LHDT": "sumartid for Lord Howe-øya", "CLST": "chilensk sumartid", "JDT": "japansk sumartid", "TMT": "turkmensk normaltid", "BT": "bhutansk tid", "IST": "indisk tid", "CST": "normaltid for sentrale Nord-Amerika", "ADT": "sumartid for den nordamerikanske atlanterhavskysten", "NZST": "nyzealandsk normaltid", "VET": "venezuelansk tid", "SRT": "surinamsk tid", "HKT": "hongkongkinesisk normaltid", "JST": "japansk normaltid", "ARST": "argentinsk sumartid", "ACDT": "sentralaustralsk sommartid", "EDT": "sumartid for den nordamerikansk austkysten", "HAST": "normaltid for Hawaii og Aleutene", "HEPMX": "sumartid for den meksikanske stillehavskysten", "GYT": "guyansk tid", "TMST": "turkmensk sumartid", "MEZ": "sentraleuropeisk standardtid", "HNPM": "normaltid for Saint-Pierre-et-Miquelon", "MDT": "MDT", "AEST": "austaustralsk standardtid", "AWST": "vestaustralsk standardtid", "EAT": "austafrikansk tid", "PST": "normaltid for den nordamerikanske stillehavskysten", "PDT": "sumartid for den nordamerikanske stillehavskysten", "HEPM": "sumartid for Saint-Pierre-et-Miquelon", "ECT": "ecuadoriansk tid", "ACST": "sentralaustralsk standardtid", "WIB": "vestindonesisk tid", "HEOG": "vestgrønlandsk sumartid", "OEZ": "austeuropeisk standardtid", "CHADT": "sumartid for Chatham", "MYT": "malaysisk tid", "WAT": "vestafrikansk standardtid", "HENOMX": "sumartid for nordvestlege Mexico", "HNT": "normaltid for Newfoundland", "CHAST": "normaltid for Chatham", "CDT": "sumartid for sentrale Nord-Amerika", "WESZ": "vesteuropeisk sommartid", "HADT": "sumartid for Hawaii og Aleutene", "UYT": "uruguayansk normaltid", "AST": "normaltid for den nordamerikanske atlanterhavskysten", "UYST": "uruguayansk sumartid", "ACWDT": "vest-sentralaustralsk sommartid", "HNPMX": "normaltid for den meksikanske stillehavskysten", "OESZ": "austeuropeisk sommartid", "WARST": "vestargentinsk sumartid", "HNCU": "kubansk normaltid", "HNNOMX": "normaltid for nordvestlege Mexico", "HAT": "sumartid for Newfoundland", "EST": "normaltid for den nordamerikansk austkysten", "HNOG": "vestgrønlandsk normaltid", "AWDT": "vestaustralsk sommartid", "HECU": "kubansk sumartid", "CAT": "sentralafrikansk tid", "ChST": "tidssone for Chamorro", "ACWST": "vest-sentralaustralsk standardtid", "SGT": "singaporsk tid", "WIT": "austindonesisk tid", "WITA": "sentralindonesisk tid", "BOT": "boliviansk tid", "AKDT": "alaskisk sumartid", "WEZ": "vesteuropeisk standardtid", "WAST": "vestafrikansk sommartid", "CLT": "chilensk normaltid", "GMT": "Greenwich middeltid", "MST": "MST", "HKST": "hongkongkinesisk sumartid", "NZDT": "nyzealandsk sumartid", "SAST": "sørafrikansk tid", "HNEG": "austgrønlandsk normaltid", "GFT": "tidssone for Fransk Guyana", "∅∅∅": "∅∅∅", "AEDT": "austaustralsk sommartid", "WART": "vestargentinsk normaltid", "HEEG": "austgrønlandsk sumartid", "MESZ": "sentraleuropeisk sommartid", "ART": "argentinsk normaltid", "LHST": "normaltid for Lord Howe-øya"},
	}
}

// Locale returns the current translators string locale
func (nn *nn) Locale() string {
	return nn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nn'
func (nn *nn) PluralsCardinal() []locales.PluralRule {
	return nn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nn'
func (nn *nn) PluralsOrdinal() []locales.PluralRule {
	return nn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nn'
func (nn *nn) PluralsRange() []locales.PluralRule {
	return nn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nn'
func (nn *nn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nn'
func (nn *nn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nn'
func (nn *nn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nn *nn) MonthAbbreviated(month time.Month) string {
	return nn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nn *nn) MonthsAbbreviated() []string {
	return nn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nn *nn) MonthNarrow(month time.Month) string {
	return nn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nn *nn) MonthsNarrow() []string {
	return nn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nn *nn) MonthWide(month time.Month) string {
	return nn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nn *nn) MonthsWide() []string {
	return nn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nn *nn) WeekdayAbbreviated(weekday time.Weekday) string {
	return nn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nn *nn) WeekdaysAbbreviated() []string {
	return nn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nn *nn) WeekdayNarrow(weekday time.Weekday) string {
	return nn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nn *nn) WeekdaysNarrow() []string {
	return nn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nn *nn) WeekdayShort(weekday time.Weekday) string {
	return nn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nn *nn) WeekdaysShort() []string {
	return nn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nn *nn) WeekdayWide(weekday time.Weekday) string {
	return nn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nn *nn) WeekdaysWide() []string {
	return nn.daysWide
}

// Decimal returns the decimal point of number
func (nn *nn) Decimal() string {
	return nn.decimal
}

// Group returns the group of number
func (nn *nn) Group() string {
	return nn.group
}

// Group returns the minus sign of number
func (nn *nn) Minus() string {
	return nn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nn' and handles both Whole and Real numbers based on 'v'
func (nn *nn) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nn.group) - 1; j >= 0; j-- {
					b = append(b, nn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nn *nn) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nn.percentSuffix...)

	b = append(b, nn.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nn'
func (nn *nn) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nn.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nn.group) - 1; j >= 0; j-- {
					b = append(b, nn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, nn.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nn'
// in accounting notation.
func (nn *nn) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nn.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nn.group) - 1; j >= 0; j-- {
					b = append(b, nn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, nn.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, nn.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'nn'
func (nn *nn) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'nn'
func (nn *nn) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nn'
func (nn *nn) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nn'
func (nn *nn) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nn'
func (nn *nn) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nn'
func (nn *nn) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nn'
func (nn *nn) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nn'
func (nn *nn) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, []byte{0x6b, 0x6c}...)
	b = append(b, []byte{0x2e, 0x20}...)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
