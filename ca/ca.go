package ca

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type ca struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'ca' locale
func New() locales.Translator {
	return &ca{
		locale:                 "ca",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "FJD", "FKP", "FRF", "£", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JP¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "₩", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "de gen.", "de febr.", "de març", "d’abr.", "de maig", "de juny", "de jul.", "d’ag.", "de set.", "d’oct.", "de nov.", "de des."},
		monthsNarrow:           []string{"", "GN", "FB", "MÇ", "AB", "MG", "JN", "JL", "AG", "ST", "OC", "NV", "DS"},
		monthsWide:             []string{"", "de gener", "de febrer", "de març", "d’abril", "de maig", "de juny", "de juliol", "d’agost", "de setembre", "d’octubre", "de novembre", "de desembre"},
		daysAbbreviated:        []string{"dg.", "dl.", "dt.", "dc.", "dj.", "dv.", "ds."},
		daysNarrow:             []string{"dg", "dl", "dt", "dc", "dj", "dv", "ds"},
		daysShort:              []string{"dg.", "dl.", "dt.", "dc.", "dj.", "dv.", "ds."},
		daysWide:               []string{"diumenge", "dilluns", "dimarts", "dimecres", "dijous", "divendres", "dissabte"},
		periodsAbbreviated:     []string{"a. m.", "p. m."},
		periodsNarrow:          []string{"a. m.", "p. m."},
		periodsWide:            []string{"a. m.", "p. m."},
		erasAbbreviated:        []string{"aC", "dC"},
		erasNarrow:             []string{"aC", "dC"},
		erasWide:               []string{"abans de Crist", "després de Crist"},
		timezones:              map[string]string{"BOT": "Hora de Bolívia", "ARST": "Hora d’estiu de l’Argentina", "AKST": "Hora estàndard d’Alaska", "ACST": "Hora estàndard d’Austràlia Central", "WESZ": "Hora d’estiu de l’Oest d’Europa", "HAST": "Hora estàndard de Hawaii-Aleutianes", "HNNOMX": "Hora estàndard del nord-oest de Mèxic", "PST": "Hora estàndard del Pacífic d’Amèrica del Nord", "CST": "Hora estàndard central d’Amèrica del Nord", "EST": "Hora estàndard oriental d’Amèrica del Nord", "WAT": "Hora estàndard de l’Àfrica Occidental", "ECT": "Hora de l’Equador", "JDT": "Hora d’estiu del Japó", "AKDT": "Hora d’estiu d’Alaska", "PDT": "Hora d’estiu del Pacífic d’Amèrica del Nord", "∅∅∅": "Hora d’estiu de les Açores", "WART": "Hora estàndard de l’oest de l’Argentina", "GYT": "Hora de Guyana", "HECU": "Hora d’estiu de Cuba", "CHADT": "Hora d’estiu de Chatham", "MESZ": "Hora d’estiu del Centre d’Europa", "UYT": "Hora estàndard de l’Uruguai", "CLST": "Hora d’estiu de Xile", "WARST": "Hora d’estiu de l’oest de l’Argentina", "NZDT": "Hora d’estiu de Nova Zelanda", "COST": "Hora d’estiu de Colòmbia", "GMT": "Hora del Meridià de Greenwich", "HEOG": "Hora d’estiu de l’Oest de Grenlàndia", "SGT": "Hora de Singapur", "ART": "Hora estàndard de l’Argentina", "HNPMX": "Hora estàndard del Pacífic de Mèxic", "TMT": "Hora estàndard del Turkmenistan", "HNEG": "Hora estàndard de l’Est de Grenlàndia", "WITA": "Hora central d’Indonèsia", "OEZ": "Hora estàndard de l’Est d’Europa", "HEEG": "Hora d’estiu de l’Est de Grenlàndia", "MYT": "Hora de Malàisia", "LHST": "Hora estàndard de Lord Howe", "LHDT": "Horari d’estiu de Lord Howe", "MST": "Hora estàndard de muntanya d’Amèrica del Nord", "AWDT": "Hora d’estiu d’Austràlia Occidental", "CLT": "Hora estàndard de Xile", "HNPM": "Hora estàndard de Saint-Pierre-et-Miquelon", "CAT": "Hora de l’Àfrica Central", "BT": "Hora de Bhutan", "CDT": "Hora d’estiu central d’Amèrica del Nord", "SRT": "Hora de Surinam", "MDT": "Hora d’estiu de muntanya d’Amèrica del Nord", "HNCU": "Hora estàndard de Cuba", "AEDT": "Hora d’estiu d’Austràlia Oriental", "SAST": "Hora estàndard del sud de l’Àfrica", "WIB": "Hora de l’oest d’Indonèsia", "EDT": "Hora d’estiu oriental d’Amèrica del Nord", "HADT": "Hora d’estiu de Hawaii-Aleutianes", "UYST": "Hora d’estiu de l’Uruguai", "HNOG": "Hora estàndard de l’Oest de Grenlàndia", "AEST": "Hora estàndard d’Austràlia Oriental", "IST": "Hora de l’Índia", "ACWST": "Hora estàndard d’Austràlia centre-occidental", "HAT": "Hora d’estiu de Terranova", "ACDT": "Hora d’estiu d’Austràlia Central", "GFT": "Hora de la Guaiana Francesa", "EAT": "Hora de l’Àfrica Oriental", "HENOMX": "Hora d’estiu del nord-oest de Mèxic", "ACWDT": "Hora d’estiu d’Austràlia centre-occidental", "HNT": "Hora estàndard de Terranova", "MEZ": "Hora estàndard del Centre d’Europa", "CHAST": "Hora estàndard de Chatham", "WIT": "Hora de l’est d’Indonèsia", "TMST": "Hora d’estiu del Turkmenistan", "HKST": "Hora d’estiu de Hong Kong", "OESZ": "Hora d’estiu de l’Est d’Europa", "VET": "Hora de Veneçuela", "HEPM": "Hora d’estiu de Saint-Pierre-et-Miquelon", "COT": "Hora estàndard de Colòmbia", "WAST": "Hora d’estiu de l’Àfrica Occidental", "ADT": "Hora d’estiu de l’Atlàntic", "HKT": "Hora estàndard de Hong Kong", "JST": "Hora estàndard del Japó", "NZST": "Hora estàndard de Nova Zelanda", "WEZ": "Hora estàndard de l’Oest d’Europa", "ChST": "Hora de Chamorro", "HEPMX": "Hora d’estiu del Pacífic de Mèxic", "AST": "Hora estàndard de l’Atlàntic", "AWST": "Hora estàndard d’Austràlia Occidental"},
	}
}

// Locale returns the current translators string locale
func (ca *ca) Locale() string {
	return ca.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ca'
func (ca *ca) PluralsCardinal() []locales.PluralRule {
	return ca.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ca'
func (ca *ca) PluralsOrdinal() []locales.PluralRule {
	return ca.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ca'
func (ca *ca) PluralsRange() []locales.PluralRule {
	return ca.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ca'
func (ca *ca) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ca'
func (ca *ca) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 || n == 3 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ca'
func (ca *ca) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ca *ca) MonthAbbreviated(month time.Month) string {
	return ca.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ca *ca) MonthsAbbreviated() []string {
	return ca.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ca *ca) MonthNarrow(month time.Month) string {
	return ca.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ca *ca) MonthsNarrow() []string {
	return ca.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ca *ca) MonthWide(month time.Month) string {
	return ca.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ca *ca) MonthsWide() []string {
	return ca.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ca *ca) WeekdayAbbreviated(weekday time.Weekday) string {
	return ca.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ca *ca) WeekdaysAbbreviated() []string {
	return ca.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ca *ca) WeekdayNarrow(weekday time.Weekday) string {
	return ca.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ca *ca) WeekdaysNarrow() []string {
	return ca.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ca *ca) WeekdayShort(weekday time.Weekday) string {
	return ca.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ca *ca) WeekdaysShort() []string {
	return ca.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ca *ca) WeekdayWide(weekday time.Weekday) string {
	return ca.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ca *ca) WeekdaysWide() []string {
	return ca.daysWide
}

// Decimal returns the decimal point of number
func (ca *ca) Decimal() string {
	return ca.decimal
}

// Group returns the group of number
func (ca *ca) Group() string {
	return ca.group
}

// Group returns the minus sign of number
func (ca *ca) Minus() string {
	return ca.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ca' and handles both Whole and Real numbers based on 'v'
func (ca *ca) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ca' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ca *ca) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ca.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ca'
func (ca *ca) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ca.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ca'
// in accounting notation.
func (ca *ca) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, ca.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ca.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ca.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ca'
func (ca *ca) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ca'
func (ca *ca) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ca'
func (ca *ca) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ca'
func (ca *ca) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ca.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ca'
func (ca *ca) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ca'
func (ca *ca) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ca'
func (ca *ca) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ca'
func (ca *ca) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ca.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
