package nl_CW

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type nl_CW struct {
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

// New returns a new instance of translator for the 'nl_CW' locale
func New() locales.Translator {
	return &nl_CW{
		locale:                 "nl_CW",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "NAf.", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mrt.", "apr.", "mei", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
		daysAbbreviated:        []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		daysNarrow:             []string{"Z", "M", "D", "W", "D", "V", "Z"},
		daysShort:              []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		daysWide:               []string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
		periodsAbbreviated:     []string{"a.m.", "p.m."},
		periodsNarrow:          []string{"a.m.", "p.m."},
		periodsWide:            []string{"a.m.", "p.m."},
		erasAbbreviated:        []string{"v.Chr.", "n.Chr."},
		erasNarrow:             []string{"v.C.", "n.C."},
		erasWide:               []string{"voor Christus", "na Christus"},
		timezones:              map[string]string{"UYST": "Uruguayaanse zomertijd", "CHADT": "Chatham-zomertijd", "CDT": "Central-zomertijd", "EST": "Eastern-standaardtijd", "HAST": "Hawaii-Aleoetische standaardtijd", "WART": "West-Argentijnse standaardtijd", "HNPM": "Saint Pierre en Miquelon-standaardtijd", "WIB": "West-Indonesische tijd", "WESZ": "West-Europese zomertijd", "ACWST": "Midden-Australische westelijke standaardtijd", "ADT": "Atlantic-zomertijd", "HKST": "Hongkongse zomertijd", "HNNOMX": "Noordwest-Mexicaanse standaardtijd", "NZST": "Nieuw-Zeelandse standaardtijd", "GMT": "Greenwich Mean Time", "AST": "Atlantic-standaardtijd", "EDT": "Eastern-zomertijd", "MDT": "Macause zomertijd", "OESZ": "Oost-Europese zomertijd", "HNEG": "Oost-Groenlandse standaardtijd", "IST": "Indiase tijd", "LHDT": "Lord Howe-eilandse zomertijd", "JDT": "Japanse zomertijd", "HNOG": "West-Groenlandse standaardtijd", "TMST": "Turkmeense zomertijd", "VET": "Venezolaanse tijd", "PDT": "Pacific-zomertijd", "ACWDT": "Midden-Australische westelijke zomertijd", "ChST": "Chamorro-tijd", "UYT": "Uruguayaanse standaardtijd", "HNPMX": "Mexicaanse Pacific-standaardtijd", "WAST": "West-Afrikaanse zomertijd", "GYT": "Guyaanse tijd", "HECU": "Cubaanse zomertijd", "SAST": "Zuid-Afrikaanse tijd", "HADT": "Hawaii-Aleoetische zomertijd", "ACDT": "Midden-Australische zomertijd", "WIT": "Oost-Indonesische tijd", "HEPM": "Saint Pierre en Miquelon-zomertijd", "BT": "Bhutaanse tijd", "MST": "Macause standaardtijd", "HEPMX": "Mexicaanse Pacific-zomertijd", "AWDT": "West-Australische zomertijd", "TMT": "Turkmeense standaardtijd", "PST": "Pacific-standaardtijd", "HEEG": "Oost-Groenlandse zomertijd", "CHAST": "Chatham-standaardtijd", "ACST": "Midden-Australische standaardtijd", "COST": "Colombiaanse zomertijd", "JST": "Japanse standaardtijd", "WITA": "Centraal-Indonesische tijd", "HNT": "Newfoundland-standaardtijd", "AKST": "Alaska-standaardtijd", "∅∅∅": "Acre-zomertijd", "CLST": "Chileense zomertijd", "SGT": "Singaporese standaardtijd", "EAT": "Oost-Afrikaanse tijd", "OEZ": "Oost-Europese standaardtijd", "HNCU": "Cubaanse standaardtijd", "ARST": "Argentijnse zomertijd", "CAT": "Centraal-Afrikaanse tijd", "WAT": "West-Afrikaanse standaardtijd", "WEZ": "West-Europese standaardtijd", "CST": "Central-standaardtijd", "BOT": "Boliviaanse tijd", "ART": "Argentijnse standaardtijd", "MYT": "Maleisische tijd", "AWST": "West-Australische standaardtijd", "COT": "Colombiaanse standaardtijd", "HAT": "Newfoundland-zomertijd", "GFT": "Frans-Guyaanse tijd", "HENOMX": "Noordwest-Mexicaanse zomertijd", "LHST": "Lord Howe-eilandse standaardtijd", "WARST": "West-Argentijnse zomertijd", "AEST": "Oost-Australische standaardtijd", "AEDT": "Oost-Australische zomertijd", "AKDT": "Alaska-zomertijd", "MEZ": "Midden-Europese standaardtijd", "MESZ": "Midden-Europese zomertijd", "CLT": "Chileense standaardtijd", "ECT": "Ecuadoraanse tijd", "HEOG": "West-Groenlandse zomertijd", "HKT": "Hongkongse standaardtijd", "NZDT": "Nieuw-Zeelandse zomertijd", "SRT": "Surinaamse tijd"},
	}
}

// Locale returns the current translators string locale
func (nl *nl_CW) Locale() string {
	return nl.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nl_CW'
func (nl *nl_CW) PluralsCardinal() []locales.PluralRule {
	return nl.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nl_CW'
func (nl *nl_CW) PluralsOrdinal() []locales.PluralRule {
	return nl.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nl_CW'
func (nl *nl_CW) PluralsRange() []locales.PluralRule {
	return nl.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nl_CW'
func (nl *nl_CW) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nl_CW'
func (nl *nl_CW) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nl_CW'
func (nl *nl_CW) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := nl.CardinalPluralRule(num1, v1)
	end := nl.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nl *nl_CW) MonthAbbreviated(month time.Month) string {
	return nl.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nl *nl_CW) MonthsAbbreviated() []string {
	return nl.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nl *nl_CW) MonthNarrow(month time.Month) string {
	return nl.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nl *nl_CW) MonthsNarrow() []string {
	return nl.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nl *nl_CW) MonthWide(month time.Month) string {
	return nl.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nl *nl_CW) MonthsWide() []string {
	return nl.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nl *nl_CW) WeekdayAbbreviated(weekday time.Weekday) string {
	return nl.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nl *nl_CW) WeekdaysAbbreviated() []string {
	return nl.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nl *nl_CW) WeekdayNarrow(weekday time.Weekday) string {
	return nl.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nl *nl_CW) WeekdaysNarrow() []string {
	return nl.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nl *nl_CW) WeekdayShort(weekday time.Weekday) string {
	return nl.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nl *nl_CW) WeekdaysShort() []string {
	return nl.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nl *nl_CW) WeekdayWide(weekday time.Weekday) string {
	return nl.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nl *nl_CW) WeekdaysWide() []string {
	return nl.daysWide
}

// Decimal returns the decimal point of number
func (nl *nl_CW) Decimal() string {
	return nl.decimal
}

// Group returns the group of number
func (nl *nl_CW) Group() string {
	return nl.group
}

// Group returns the minus sign of number
func (nl *nl_CW) Minus() string {
	return nl.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nl_CW' and handles both Whole and Real numbers based on 'v'
func (nl *nl_CW) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nl_CW' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nl *nl_CW) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nl.percentSuffix...)

	b = append(b, nl.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nl_CW'
func (nl *nl_CW) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, nl.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, nl.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nl_CW'
// in accounting notation.
func (nl *nl_CW) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nl.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nl.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, nl.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, nl.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nl.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, nl.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, nl.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nl.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, nl.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nl_CW'
func (nl *nl_CW) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nl.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nl.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
