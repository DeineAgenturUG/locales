package sq_AL

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type sq_AL struct {
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

// New returns a new instance of translator for the 'sq_AL' locale
func New() locales.Translator {
	return &sq_AL{
		locale:                 "sq_AL",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "jan", "shk", "mar", "pri", "maj", "qer", "korr", "gush", "sht", "tet", "nën", "dhj"},
		monthsNarrow:           []string{"", "j", "sh", "m", "p", "m", "q", "k", "g", "sh", "t", "n", "dh"},
		monthsWide:             []string{"", "janar", "shkurt", "mars", "prill", "maj", "qershor", "korrik", "gusht", "shtator", "tetor", "nëntor", "dhjetor"},
		daysAbbreviated:        []string{"Die", "Hën", "Mar", "Mër", "Enj", "Pre", "Sht"},
		daysNarrow:             []string{"d", "h", "m", "m", "e", "p", "sh"},
		daysShort:              []string{"die", "hën", "mar", "mër", "enj", "pre", "sht"},
		daysWide:               []string{"e diel", "e hënë", "e martë", "e mërkurë", "e enjte", "e premte", "e shtunë"},
		periodsAbbreviated:     []string{"p.d.", "m.d."},
		periodsNarrow:          []string{"p.d.", "m.d."},
		periodsWide:            []string{"e paradites", "e pasdites"},
		erasAbbreviated:        []string{"p.K.", "mb.K."},
		erasNarrow:             []string{"p.K.", "mb.K."},
		erasWide:               []string{"para Krishtit", "mbas Krishtit"},
		timezones:              map[string]string{"COST": "Ora verore e Kolumbisë", "GMT": "Ora e Grinuiçit", "OESZ": "Ora verore e Evropës Lindore", "HNCU": "Ora standarde e Kubës", "HENOMX": "Ora verore e Meksikës Veriperëndimore", "CST": "Ora standarde e SHBA-së Qendrore", "∅∅∅": "Ora verore e Brazilisë", "SGT": "Ora e Singaporit", "PST": "Ora standarde e Territoreve Amerikane të Bregut të Paqësorit", "HEEG": "Ora verore e Grenlandës Lindore", "SAST": "Ora standarde e Afrikës Jugore", "ACST": "Ora standarde e Australisë Qendrore", "HNPM": "Ora standarde e Shën-Pier dhe Mikelon", "MYT": "Ora e Malajzisë", "CAT": "Ora e Afrikës Qendrore", "ACWDT": "Ora verore e Australisë Qendroro-Perëndimore", "LHST": "Ora standarde e Lord-Houit", "CLST": "Ora verore e Kilit", "HNPMX": "Ora standarde e Territoreve Meksikane të Bregut të Paqësorit", "SRT": "Ora e Surinamit", "AST": "Ora standarde e Atlantikut", "HKT": "Ora standarde e Hong-Kongut", "HADT": "Ora verore e Ishujve Hauai-Aleutian", "ECT": "Ora e Ekuadorit", "PDT": "Ora verore e Territoreve Amerikane të Bregut të Paqësorit", "UYT": "Ora standarde e Uruguait", "MST": "Ora standarde e Makaos", "MDT": "Ora verore e Makaos", "HKST": "Ora verore e Hong-Kongut", "JST": "Ora standarde e Japonisë", "NZST": "Ora standarde e Zelandës së Re", "HEPMX": "Ora verore e Territoreve Meksikane të Bregut të Paqësorit", "COT": "Ora standarde e Kolumbisë", "WAST": "Ora verore e Afrikës Perëndimore", "ACWST": "Ora standarde e Australisë Qendroro-Perëndimore", "HNOG": "Ora standarde e Grënlandës Perëndimore", "TMST": "Ora verore e Turkmenistanit", "UYST": "Ora verore e Uruguait", "BT": "Ora e Butanit", "WAT": "Ora standarde e Afrikës Perëndimore", "HECU": "Ora verore e Kubës", "AKDT": "Ora verore e Alaskës", "WARST": "Ora verore e Argjentinës Perëndimore", "WESZ": "Ora verore e Evropës Perëndimore", "HAST": "Ora standarde e Ishujve Hauai-Aleutian", "AWST": "Ora standarde e Australisë Perëndimore", "HNT": "Ora standarde e Njufaundlendit [Tokës së Re]", "MESZ": "Ora verore e Evropës Qendrore", "WEZ": "Ora standarde e Evropës Perëndimore", "HNNOMX": "Ora standarde e Meksikës Veriperëndimore", "HNEG": "Ora standarde e Grenlandës Lindore", "HAT": "Ora verore e Njufaundlendit [Tokës së Re]", "ACDT": "Ora verore e Australisë Qendrore", "CHAST": "Ora standarde e Katamit", "HEPM": "Ora verore e Shën-Pier dhe Mikelon", "WART": "Ora standarde e Argjentinës Perëndimore", "ChST": "Ora e Kamorros", "WITA": "Ora e Indonezisë Qendrore", "VET": "Ora e Venezuelës", "BOT": "Ora e Bolivisë", "ARST": "Ora verore e Argjentinës", "EST": "Ora standarde e SHBA-së Lindore", "CDT": "Ora verore e SHBA-së Qendrore", "MEZ": "Ora standarde e Evropës Qendrore", "IST": "Ora standarde e Indisë", "ADT": "Ora verore e Atlantikut", "NZDT": "Ora verore e Zelandës së Re", "CHADT": "Ora verore e Katamit", "GYT": "Ora e Guajanës", "LHDT": "Ora verore e Lord-Houit", "GFT": "Ora e Guajanës Franceze", "CLT": "Ora standarde e Kilit", "HEOG": "Ora verore e Grënlandës Perëndimore", "JDT": "Ora verore e Japonisë", "EAT": "Ora e Afrikës Lindore", "ART": "Ora standarde e Argjentinës", "WIB": "Ora e Indonezisë Perëndimore", "AEST": "Ora standarde e Australisë Lindore", "AEDT": "Ora verore e Australisë Lindore", "WIT": "Ora e Indonezisë Lindore", "AWDT": "Ora verore e Australisë Perëndimore", "TMT": "Ora standarde e Turkmenistanit", "OEZ": "Ora standarde e Evropës Lindore", "AKST": "Ora standarde e Alaskës", "EDT": "Ora verore e SHBA-së Lindore"},
	}
}

// Locale returns the current translators string locale
func (sq *sq_AL) Locale() string {
	return sq.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sq_AL'
func (sq *sq_AL) PluralsCardinal() []locales.PluralRule {
	return sq.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sq_AL'
func (sq *sq_AL) PluralsOrdinal() []locales.PluralRule {
	return sq.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sq_AL'
func (sq *sq_AL) PluralsRange() []locales.PluralRule {
	return sq.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sq_AL'
func (sq *sq_AL) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sq_AL'
func (sq *sq_AL) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod100 := math.Mod(n, 100)
	nMod10 := math.Mod(n, 10)

	if n == 1 {
		return locales.PluralRuleOne
	} else if nMod10 == 4 && nMod100 != 14 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sq_AL'
func (sq *sq_AL) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sq.CardinalPluralRule(num1, v1)
	end := sq.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sq *sq_AL) MonthAbbreviated(month time.Month) string {
	return sq.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sq *sq_AL) MonthsAbbreviated() []string {
	return sq.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sq *sq_AL) MonthNarrow(month time.Month) string {
	return sq.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sq *sq_AL) MonthsNarrow() []string {
	return sq.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sq *sq_AL) MonthWide(month time.Month) string {
	return sq.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sq *sq_AL) MonthsWide() []string {
	return sq.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayAbbreviated(weekday time.Weekday) string {
	return sq.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sq *sq_AL) WeekdaysAbbreviated() []string {
	return sq.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayNarrow(weekday time.Weekday) string {
	return sq.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sq *sq_AL) WeekdaysNarrow() []string {
	return sq.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayShort(weekday time.Weekday) string {
	return sq.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sq *sq_AL) WeekdaysShort() []string {
	return sq.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sq *sq_AL) WeekdayWide(weekday time.Weekday) string {
	return sq.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sq *sq_AL) WeekdaysWide() []string {
	return sq.daysWide
}

// Decimal returns the decimal point of number
func (sq *sq_AL) Decimal() string {
	return sq.decimal
}

// Group returns the group of number
func (sq *sq_AL) Group() string {
	return sq.group
}

// Group returns the minus sign of number
func (sq *sq_AL) Minus() string {
	return sq.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sq_AL' and handles both Whole and Real numbers based on 'v'
func (sq *sq_AL) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sq_AL' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sq *sq_AL) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sq.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sq_AL'
func (sq *sq_AL) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sq.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sq.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sq.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sq_AL'
// in accounting notation.
func (sq *sq_AL) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sq.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sq.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sq.group) - 1; j >= 0; j-- {
					b = append(b, sq.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, sq.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sq.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sq.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sq.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sq.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sq.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sq_AL'
func (sq *sq_AL) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sq.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sq.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sq.periodsAbbreviated[0]...)
	} else {
		b = append(b, sq.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x2c, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sq.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
