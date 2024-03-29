package bg_BG

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type bg_BG struct {
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

// New returns a new instance of translator for the 'bg_BG' locale
func New() locales.Translator {
	return &bg_BG{
		locale:                 "bg_BG",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
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
		monthsAbbreviated:      []string{"", "яну", "фев", "март", "апр", "май", "юни", "юли", "авг", "сеп", "окт", "ное", "дек"},
		monthsNarrow:           []string{"", "я", "ф", "м", "а", "м", "ю", "ю", "а", "с", "о", "н", "д"},
		monthsWide:             []string{"", "януари", "февруари", "март", "април", "май", "юни", "юли", "август", "септември", "октомври", "ноември", "декември"},
		daysAbbreviated:        []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysNarrow:             []string{"н", "п", "в", "с", "ч", "п", "с"},
		daysShort:              []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysWide:               []string{"неделя", "понеделник", "вторник", "сряда", "четвъртък", "петък", "събота"},
		periodsAbbreviated:     []string{"am", "pm"},
		periodsNarrow:          []string{"am", "pm"},
		periodsWide:            []string{"пр.об.", "сл.об."},
		erasAbbreviated:        []string{"пр.Хр.", "сл.Хр."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"преди Христа", "след Христа"},
		timezones:              map[string]string{"HKT": "Хонконгско стандартно време", "MYT": "Малайзийско време", "HADT": "Хавайско-алеутско лятно часово време", "GMT": "Средно гринуичко време", "EAT": "Източноафриканско време", "HNEG": "Източногренландско стандартно време", "CHADT": "Чатъмско лятно часово време", "∅∅∅": "∅∅∅", "ACWST": "Австралия – западно централно стандартно време", "HNPMX": "Мексиканско тихоокеанско стандартно време", "HNOG": "Западногренландско стандартно време", "WAST": "Западноафриканско лятно часово време", "AWST": "Западноавстралийско стандартно време", "SGT": "Сингапурско време", "EST": "Северноамериканско източно стандартно време", "ACWDT": "Австралия – западно централно лятно часово време", "HNT": "Нюфаундлендско стандартно време", "CHAST": "Чатъмско стандартно време", "IST": "Индийско време", "NZDT": "Новозеландско лятно часово време", "JDT": "Японско лятно часово време", "SAST": "Южноафриканско време", "ART": "Аржентинско стандартно време", "WEZ": "Западноевропейско стандартно време", "WESZ": "Западноевропейско лятно време", "WIT": "Източноиндонезийско време", "HAT": "Нюфаундлендско лятно часово време", "AKDT": "Аляска – лятно часово време", "ACST": "Централноавстралийско стандартно време", "PDT": "Северноамериканско тихоокеанско лятно часово време", "UYST": "Уругвайско лятно часово време", "WAT": "Западноафриканско стандартно време", "CLST": "Чилийско лятно часово време", "COT": "Колумбийско стандартно време", "HEOG": "Западногренландско лятно часово време", "WARST": "Западноаржентинско лятно часово време", "NZST": "Новозеландско стандартно време", "ChST": "Чаморско време", "AKST": "Аляска – стандартно време", "ARST": "Аржентинско лятно часово време", "HNNOMX": "Северозападно стандартно мексиканско време", "ADT": "Северноамериканско атлантическо лятно часово време", "HKST": "Хонконгско лятно часово време", "TMST": "Туркменистанско лятно часово време", "BOT": "Боливийско време", "GYT": "Гаяна", "WITA": "Централноиндонезийско време", "HENOMX": "Северозападно лятно часово мексиканско време", "LHST": "Лорд Хау – стандартно време", "MST": "MST", "HEPMX": "Мексиканско тихоокеанско лятно часово време", "TMT": "Туркменистанско стандартно време", "OEZ": "Източноевропейско стандартно време", "HEPM": "Сен Пиер и Микелон – лятно часово време", "WIB": "Западноиндонезийско време", "GFT": "Френска Гвиана", "AEST": "Източноавстралийско стандартно време", "AEDT": "Източноавстралийско лятно часово време", "PST": "Северноамериканско тихоокеанско стандартно време", "HEEG": "Източногренландско лятно часово време", "BT": "Бутанско време", "ACDT": "Централноавстралийско лятно часово време", "HNPM": "Сен Пиер и Микелон – стандартно време", "UYT": "Уругвайско стандартно време", "CDT": "Северноамериканско централно лятно часово време", "CST": "Северноамериканско централно стандартно време", "OESZ": "Източноевропейско лятно часово време", "HECU": "Кубинско лятно часово време", "EDT": "Северноамериканско източно лятно часово време", "LHDT": "Лорд Хау – лятно часово време", "AWDT": "Западноавстралийско лятно часово време", "WART": "Западноаржентинско стандартно време", "HNCU": "Кубинско стандартно време", "MEZ": "Централноевропейско стандартно време", "MESZ": "Централноевропейско лятно часово време", "CLT": "Чилийско стандартно време", "SRT": "Суринамско време", "ECT": "Еквадорско време", "JST": "Японско стандартно време", "VET": "Венецуелско време", "HAST": "Хавайско-алеутско стандартно време", "MDT": "MDT", "AST": "Северноамериканско атлантическо стандартно време", "CAT": "Централноафриканско време", "COST": "Колумбийско лятно часово време"},
	}
}

// Locale returns the current translators string locale
func (bg *bg_BG) Locale() string {
	return bg.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'bg_BG'
func (bg *bg_BG) PluralsCardinal() []locales.PluralRule {
	return bg.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'bg_BG'
func (bg *bg_BG) PluralsOrdinal() []locales.PluralRule {
	return bg.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'bg_BG'
func (bg *bg_BG) PluralsRange() []locales.PluralRule {
	return bg.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'bg_BG'
func (bg *bg_BG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'bg_BG'
func (bg *bg_BG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'bg_BG'
func (bg *bg_BG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (bg *bg_BG) MonthAbbreviated(month time.Month) string {
	return bg.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (bg *bg_BG) MonthsAbbreviated() []string {
	return bg.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (bg *bg_BG) MonthNarrow(month time.Month) string {
	return bg.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (bg *bg_BG) MonthsNarrow() []string {
	return bg.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (bg *bg_BG) MonthWide(month time.Month) string {
	return bg.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (bg *bg_BG) MonthsWide() []string {
	return bg.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (bg *bg_BG) WeekdayAbbreviated(weekday time.Weekday) string {
	return bg.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (bg *bg_BG) WeekdaysAbbreviated() []string {
	return bg.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (bg *bg_BG) WeekdayNarrow(weekday time.Weekday) string {
	return bg.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (bg *bg_BG) WeekdaysNarrow() []string {
	return bg.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (bg *bg_BG) WeekdayShort(weekday time.Weekday) string {
	return bg.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (bg *bg_BG) WeekdaysShort() []string {
	return bg.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (bg *bg_BG) WeekdayWide(weekday time.Weekday) string {
	return bg.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (bg *bg_BG) WeekdaysWide() []string {
	return bg.daysWide
}

// Decimal returns the decimal point of number
func (bg *bg_BG) Decimal() string {
	return bg.decimal
}

// Group returns the group of number
func (bg *bg_BG) Group() string {
	return bg.group
}

// Group returns the minus sign of number
func (bg *bg_BG) Minus() string {
	return bg.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'bg_BG' and handles both Whole and Real numbers based on 'v'
func (bg *bg_BG) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(bg.group) - 1; j >= 0; j-- {
					b = append(b, bg.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'bg_BG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (bg *bg_BG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, bg.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'bg_BG'
func (bg *bg_BG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bg.currencies[currency]
	l := len(s) + len(symbol) + 4

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, bg.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bg.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, bg.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'bg_BG'
// in accounting notation.
func (bg *bg_BG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := bg.currencies[currency]
	l := len(s) + len(symbol) + 6

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, bg.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, bg.currencyNegativePrefix[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, bg.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, bg.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, bg.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x20, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

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

	b = append(b, []byte{0x20, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bg.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, bg.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, bg.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd0, 0xb3}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20, 0xd1, 0x87}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xd1, 0x87}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xd1, 0x87}...)
	b = append(b, []byte{0x2e, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'bg_BG'
func (bg *bg_BG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, bg.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xd1, 0x87}...)
	b = append(b, []byte{0x2e, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := bg.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
