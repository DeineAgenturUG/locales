package br

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type br struct {
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

// New returns a new instance of translator for the 'br' locale
func New() locales.Translator {
	return &br{
		locale:                 "br",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "$A", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "$CA", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£ RU", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "$ HK", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "$ ZN", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "Db", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$ SU", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "Gen.", "Cʼhwe.", "Meur.", "Ebr.", "Mae", "Mezh.", "Goue.", "Eost", "Gwen.", "Here", "Du", "Kzu."},
		monthsNarrow:           []string{"", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12"},
		monthsWide:             []string{"", "Genver", "Cʼhwevrer", "Meurzh", "Ebrel", "Mae", "Mezheven", "Gouere", "Eost", "Gwengolo", "Here", "Du", "Kerzu"},
		daysAbbreviated:        []string{"Sul", "Lun", "Meu.", "Mer.", "Yaou", "Gwe.", "Sad."},
		daysNarrow:             []string{"Su", "L", "Mz", "Mc", "Y", "G", "Sa"},
		daysShort:              []string{"Sul", "Lun", "Meu.", "Mer.", "Yaou", "Gwe.", "Sad."},
		daysWide:               []string{"Sul", "Lun", "Meurzh", "Mercʼher", "Yaou", "Gwener", "Sadorn"},
		periodsAbbreviated:     []string{"A.M.", "G.M."},
		periodsNarrow:          []string{"am", "gm"},
		periodsWide:            []string{"A.M.", "G.M."},
		erasAbbreviated:        []string{"a-raok J.K.", "goude J.K."},
		erasNarrow:             []string{"a-raok J.K.", "goude J.K."},
		erasWide:               []string{"a-raok Jezuz-Krist", "goude Jezuz-Krist"},
		timezones:              map[string]string{"HNPMX": "HNPMX", "CLST": "eur hañv Chile", "CDT": "CDT", "∅∅∅": "eur hañv Brasília", "EST": "eur cʼhoañv ar Reter", "HADT": "HADT", "AKST": "eur cʼhoañv Alaska", "MESZ": "eur hañv Kreizeuropa", "CHADT": "eur hañv Chatham", "CST": "CST", "GYT": "eur Guyana", "AEST": "eur cʼhoañv Aostralia ar Reter", "WIT": "eur Indonezia ar Reter", "JDT": "eur hañv Japan", "SAST": "eur cʼhoañv Suafrika", "HNT": "eur cʼhoañv Newfoundland", "ACST": "eur cʼhoañv Kreizaostralia", "HAST": "HAST", "HNOG": "eur cʼhoañv Greunland ar Cʼhornôg", "WARST": "eur hañv Arcʼhantina ar Cʼhornôg", "SRT": "eur Surinam", "HECU": "eur hañv Kuba", "PDT": "PDT", "ARST": "eur hañv Arcʼhantina", "CAT": "eur Kreizafrika", "CHAST": "eur cʼhoañv Chatham", "GMT": "Amzer keitat Greenwich (AKG)", "NZST": "eur cʼhoañv Zeland-Nevez", "NZDT": "eur hañv Zeland-Nevez", "HAT": "eur hañv Newfoundland", "EDT": "eur hañv ar Reter", "WAST": "eur hañv Afrika ar Cʼhornôg", "COT": "eur cʼhoañv Kolombia", "MDT": "eur hañv Macau", "MST": "eur cʼhoañv Macau", "HKST": "eur hañv Hong Kong", "EAT": "eur Afrika ar Reter", "WITA": "WITA", "WESZ": "eur hañv Europa ar Cʼhornôg", "LHDT": "LHDT", "ChST": "ChST", "UYT": "eur cʼhoañv Uruguay", "VET": "eur Venezuela", "OESZ": "eur hañv Europa ar Reter", "ART": "eur cʼhoañv Arcʼhantina", "IST": "eur cʼhoañv India", "ADT": "ADT", "TMT": "eur cʼhoañv Turkmenistan", "HNEG": "eur cʼhoañv Greunland ar Reter", "HNPM": "eur cʼhoañv Sant-Pêr-ha-Mikelon", "HEPM": "eur hañv Sant-Pêr-ha-Mikelon", "WEZ": "eur cʼhoañv Europa ar Cʼhornôg", "HENOMX": "eur hañv Gwalarn Mecʼhiko", "HKT": "eur cʼhoañv Hong Kong", "TMST": "eur hañv Turkmenistan", "OEZ": "eur cʼhoañv Europa ar Reter", "COST": "eur hañv Kolombia", "AST": "AST", "AWST": "eur cʼhoañv Aostralia ar Cʼhornôg", "AWDT": "eur hañv Aostralia ar Cʼhornôg", "CLT": "eur cʼhoañv Chile", "HEEG": "eur hañv Greunland ar Reter", "BT": "eur Bhoutan", "ACWDT": "eur hañv Kreizaostralia ar Cʼhornôg", "AEDT": "eur hañv Aostralia ar Reter", "PST": "PST", "JST": "eur cʼhoañv Japan", "WART": "eur cʼhoañv Arcʼhantina ar Cʼhornôg", "WIB": "eur Indonezia ar Cʼhornôg", "GFT": "eur Gwiana cʼhall", "HEPMX": "HEPMX", "ACWST": "eur cʼhoañv Kreizaostralia ar Cʼhornôg", "MEZ": "eur cʼhoañv Kreizeuropa", "HEOG": "eur hañv Greunland ar Cʼhornôg", "SGT": "eur cʼhoañv Singapour", "HNNOMX": "eur cʼhoañv Gwalarn Mecʼhiko", "BOT": "eur Bolivia", "AKDT": "eur hañv Alaska", "ACDT": "eur hañv Kreizaostralia", "ECT": "eur Ecuador", "HNCU": "eur cʼhoañv Kuba", "LHST": "LHST", "MYT": "eur Malaysia", "UYST": "eur hañv Uruguay", "WAT": "eur cʼhoañv Afrika ar Cʼhornôg"},
	}
}

// Locale returns the current translators string locale
func (br *br) Locale() string {
	return br.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'br'
func (br *br) PluralsCardinal() []locales.PluralRule {
	return br.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'br'
func (br *br) PluralsOrdinal() []locales.PluralRule {
	return br.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'br'
func (br *br) PluralsRange() []locales.PluralRule {
	return br.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'br'
func (br *br) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)
	nMod1000000 := math.Mod(n, 1000000)

	if nMod10 == 1 && (nMod100 != 11 && nMod100 != 71 && nMod100 != 91) {
		return locales.PluralRuleOne
	} else if nMod10 == 2 && (nMod100 != 12 && nMod100 != 72 && nMod100 != 92) {
		return locales.PluralRuleTwo
	} else if nMod10 >= 3 && nMod10 <= 4 && (nMod10 == 9) && (nMod100 < 10 || nMod100 > 19) || (nMod100 < 70 || nMod100 > 79) || (nMod100 < 90 || nMod100 > 99) {
		return locales.PluralRuleFew
	} else if n != 0 && nMod1000000 == 0 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'br'
func (br *br) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'br'
func (br *br) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (br *br) MonthAbbreviated(month time.Month) string {
	return br.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (br *br) MonthsAbbreviated() []string {
	return br.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (br *br) MonthNarrow(month time.Month) string {
	return br.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (br *br) MonthsNarrow() []string {
	return br.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (br *br) MonthWide(month time.Month) string {
	return br.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (br *br) MonthsWide() []string {
	return br.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (br *br) WeekdayAbbreviated(weekday time.Weekday) string {
	return br.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (br *br) WeekdaysAbbreviated() []string {
	return br.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (br *br) WeekdayNarrow(weekday time.Weekday) string {
	return br.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (br *br) WeekdaysNarrow() []string {
	return br.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (br *br) WeekdayShort(weekday time.Weekday) string {
	return br.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (br *br) WeekdaysShort() []string {
	return br.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (br *br) WeekdayWide(weekday time.Weekday) string {
	return br.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (br *br) WeekdaysWide() []string {
	return br.daysWide
}

// Decimal returns the decimal point of number
func (br *br) Decimal() string {
	return br.decimal
}

// Group returns the group of number
func (br *br) Group() string {
	return br.group
}

// Group returns the minus sign of number
func (br *br) Minus() string {
	return br.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'br' and handles both Whole and Real numbers based on 'v'
func (br *br) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, br.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'br' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (br *br) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, br.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, br.percentSuffix...)

	b = append(b, br.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'br'
func (br *br) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := br.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, br.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, br.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, br.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'br'
// in accounting notation.
func (br *br) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := br.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, br.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(br.group) - 1; j >= 0; j-- {
					b = append(b, br.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, br.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, br.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, br.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, br.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'br'
func (br *br) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'br'
func (br *br) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, br.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'br'
func (br *br) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, br.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'br'
func (br *br) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, br.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, br.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'br'
func (br *br) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'br'
func (br *br) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, br.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'br'
func (br *br) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, br.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'br'
func (br *br) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, br.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, br.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := br.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
