package es_BR

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type es_BR struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	percentSuffix      string
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

// New returns a new instance of translator for the 'es_BR' locale
func New() locales.Translator {
	return &es_BR{
		locale:             "es_BR",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     []locales.PluralRule{6},
		pluralsRange:       []locales.PluralRule{6},
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		perMille:           "‰",
		timeSeparator:      ":",
		inifinity:          "∞",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		percentSuffix:      " ",
		monthsAbbreviated:  []string{"", "ene.", "feb.", "mar.", "abr.", "may.", "jun.", "jul.", "ago.", "sep.", "oct.", "nov.", "dic."},
		monthsNarrow:       []string{"", "E", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:         []string{"", "enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
		daysAbbreviated:    []string{"dom.", "lun.", "mar.", "mié.", "jue.", "vie.", "sáb."},
		daysNarrow:         []string{"d", "l", "m", "m", "j", "v", "s"},
		daysShort:          []string{"DO", "LU", "MA", "MI", "JU", "VI", "SA"},
		daysWide:           []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		periodsAbbreviated: []string{"a.m.", "p.m."},
		periodsNarrow:      []string{"a.m.", "p.m."},
		periodsWide:        []string{"a.m.", "p.m."},
		erasAbbreviated:    []string{"a. C.", "d. C."},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"antes de Cristo", "después de Cristo"},
		timezones:          map[string]string{"SRT": "hora de Surinam", "ECT": "hora de Ecuador", "HAT": "hora de verano de Terranova", "LHST": "hora estándar de Lord Howe", "LHDT": "hora de verano de Lord Howe", "MYT": "hora de Malasia", "WAST": "hora de verano de África occidental", "HEEG": "hora de verano de Groenlandia oriental", "EST": "hora estándar oriental", "HADT": "hora de verano de Hawái-Aleutianas", "CST": "hora estándar central", "UYT": "hora estándar de Uruguay", "CLST": "hora de verano de Chile", "WITA": "hora de Indonesia central", "PDT": "hora de verano del Pacífico", "GMT": "hora del meridiano de Greenwich", "HNOG": "hora estándar de Groenlandia occidental", "OEZ": "hora estándar de Europa del Este", "WART": "hora estándar de Argentina occidental", "WARST": "hora de verano de Argentina occidental", "EDT": "hora de verano oriental", "VET": "hora de Venezuela", "HNT": "hora estándar de Terranova", "CDT": "hora de verano central", "HAST": "hora estándar de Hawái-Aleutianas", "HEPMX": "hora de verano del Pacífico de México", "AST": "hora estándar del Atlántico", "AWDT": "hora de verano de Australia occidental", "GYT": "hora de Guyana", "SGT": "hora de Singapur", "JST": "hora estándar de Japón", "HNPM": "hora estándar de San Pedro y Miquelón", "CHAST": "hora estándar de Chatham", "WESZ": "hora de verano de Europa del Oeste", "PST": "hora estándar del Pacífico", "ADT": "hora de verano del Atlántico", "AEDT": "hora de verano de Australia oriental", "HKST": "hora de verano de Hong Kong", "TMT": "hora estándar de Turkmenistán", "NZST": "hora estándar de Nueva Zelanda", "NZDT": "hora de verano de Nueva Zelanda", "HECU": "hora de verano de Cuba", "CHADT": "hora de verano de Chatham", "COST": "hora de verano de Colombia", "MST": "Hora estándar de Macao", "HEOG": "hora de verano de Groenlandia occidental", "HKT": "hora estándar de Hong Kong", "SAST": "hora de Sudáfrica", "HNEG": "hora estándar de Groenlandia oriental", "CAT": "hora de África central", "ChST": "hora estándar de Chamorro", "∅∅∅": "hora de verano de Perú", "COT": "hora estándar de Colombia", "WIT": "hora de Indonesia oriental", "AKST": "hora estándar de Alaska", "AKDT": "hora de verano de Alaska", "GFT": "hora de la Guayana Francesa", "UYST": "hora de verano de Uruguay", "CLT": "hora estándar de Chile", "OESZ": "hora de verano de Europa del Este", "HNCU": "hora estándar de Cuba", "ACST": "hora estándar de Australia central", "ACWST": "hora estándar de Australia centroccidental", "MDT": "Hora de verano de Macao", "AWST": "hora estándar de Australia occidental", "BOT": "hora de Bolivia", "MEZ": "hora estándar de Europa central", "ARST": "hora de verano de Argentina", "WEZ": "hora estándar de Europa del Oeste", "MESZ": "hora de verano de Europa central", "ACWDT": "hora de verano de Australia centroccidental", "AEST": "hora estándar de Australia oriental", "TMST": "hora de verano de Turkmenistán", "EAT": "hora de África oriental", "JDT": "hora de verano de Japón", "HNNOMX": "hora estándar del noroeste de México", "HENOMX": "hora de verano del noroeste de México", "HEPM": "hora de verano de San Pedro y Miquelón", "ACDT": "hora de verano de Australia central", "IST": "hora de India", "HNPMX": "hora estándar del Pacífico de México", "WAT": "hora estándar de África occidental", "BT": "hora de Bután", "WIB": "hora de Indonesia occidental", "ART": "hora estándar de Argentina"},
	}
}

// Locale returns the current translators string locale
func (es *es_BR) Locale() string {
	return es.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'es_BR'
func (es *es_BR) PluralsCardinal() []locales.PluralRule {
	return es.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'es_BR'
func (es *es_BR) PluralsOrdinal() []locales.PluralRule {
	return es.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'es_BR'
func (es *es_BR) PluralsRange() []locales.PluralRule {
	return es.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'es_BR'
func (es *es_BR) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'es_BR'
func (es *es_BR) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'es_BR'
func (es *es_BR) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (es *es_BR) MonthAbbreviated(month time.Month) string {
	return es.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (es *es_BR) MonthsAbbreviated() []string {
	return es.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (es *es_BR) MonthNarrow(month time.Month) string {
	return es.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (es *es_BR) MonthsNarrow() []string {
	return es.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (es *es_BR) MonthWide(month time.Month) string {
	return es.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (es *es_BR) MonthsWide() []string {
	return es.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (es *es_BR) WeekdayAbbreviated(weekday time.Weekday) string {
	return es.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (es *es_BR) WeekdaysAbbreviated() []string {
	return es.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (es *es_BR) WeekdayNarrow(weekday time.Weekday) string {
	return es.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (es *es_BR) WeekdaysNarrow() []string {
	return es.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (es *es_BR) WeekdayShort(weekday time.Weekday) string {
	return es.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (es *es_BR) WeekdaysShort() []string {
	return es.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (es *es_BR) WeekdayWide(weekday time.Weekday) string {
	return es.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (es *es_BR) WeekdaysWide() []string {
	return es.daysWide
}

// Decimal returns the decimal point of number
func (es *es_BR) Decimal() string {
	return es.decimal
}

// Group returns the group of number
func (es *es_BR) Group() string {
	return es.group
}

// Group returns the minus sign of number
func (es *es_BR) Minus() string {
	return es.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'es_BR' and handles both Whole and Real numbers based on 'v'
func (es *es_BR) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'es_BR' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (es *es_BR) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, es.percentSuffix...)

	b = append(b, es.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'es_BR'
func (es *es_BR) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
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
		b = append(b, es.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'es_BR'
// in accounting notation.
func (es *es_BR) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := es.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, es.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, es.group[0])
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

		b = append(b, es.minus[0])

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
			b = append(b, es.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'es_BR'
func (es *es_BR) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'es_BR'
func (es *es_BR) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'es_BR'
func (es *es_BR) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'es_BR'
func (es *es_BR) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, es.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, es.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'es_BR'
func (es *es_BR) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'es_BR'
func (es *es_BR) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'es_BR'
func (es *es_BR) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'es_BR'
func (es *es_BR) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, es.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, es.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := es.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
