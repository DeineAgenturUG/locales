package uk

import (
	"math"
	"strconv"
	"time"

	"github.com/DeineAgenturUG/locales"
	"github.com/DeineAgenturUG/locales/currency"
)

type uk struct {
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

// New returns a new instance of translator for the 'uk' locale
func New() locales.Translator {
	return &uk{
		locale:                 "uk",
		pluralsCardinal:        []locales.PluralRule{2, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{4, 6},
		pluralsRange:           []locales.PluralRule{2, 4, 5, 6},
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "₴", "крб.", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VEF", "VES", "VND", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "CFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "січ.", "лют.", "бер.", "квіт.", "трав.", "черв.", "лип.", "серп.", "вер.", "жовт.", "лист.", "груд."},
		monthsNarrow:           []string{"", "с", "л", "б", "к", "т", "ч", "л", "с", "в", "ж", "л", "г"},
		monthsWide:             []string{"", "січня", "лютого", "березня", "квітня", "травня", "червня", "липня", "серпня", "вересня", "жовтня", "листопада", "грудня"},
		daysAbbreviated:        []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysNarrow:             []string{"Н", "П", "В", "С", "Ч", "П", "С"},
		daysShort:              []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
		daysWide:               []string{"неділя", "понеділок", "вівторок", "середа", "четвер", "пʼятниця", "субота"},
		periodsAbbreviated:     []string{"дп", "пп"},
		periodsNarrow:          []string{"дп", "пп"},
		periodsWide:            []string{"дп", "пп"},
		erasAbbreviated:        []string{"до н. е.", "н. е."},
		erasNarrow:             []string{"до н.е.", "н.е."},
		erasWide:               []string{"до нашої ери", "нашої ери"},
		timezones:              map[string]string{"HNT": "за стандартним часом на острові Ньюфаундленд", "HEPM": "за літнім часом на островах Сен-П’єр і Мікелон", "AWDT": "за літнім західноавстралійським часом", "WARST": "за літнім за західноаргентинським часом", "HECU": "за літнім часом на Кубі", "PDT": "за північноамериканським тихоокеанським літнім часом", "WIB": "за західноіндонезійським часом", "ADT": "за атлантичним літнім часом", "HENOMX": "за літнім північнозахідним часом у Мексиці", "CAT": "за центральноафриканським часом", "NZST": "за стандартним часом у Новій Зеландії", "∅∅∅": "час: Акрі, літній", "OEZ": "за східноєвропейським стандартним часом", "MESZ": "за центральноєвропейським літнім часом", "CHAST": "за стандартним часом на архіпелазі Чатем", "ChST": "за часом на Північних Маріанських островах", "BOT": "за болівійським часом", "CHADT": "за літнім часом на архіпелазі Чатем", "HNPM": "за стандартним часом на островах Сен-П’єр і Мікелон", "MYT": "за часом у Малайзії", "VET": "за часом у Венесуелі", "ACWDT": "за літнім центральнозахідним австралійським часом", "HKT": "за стандартним часом у Гонконзі", "WAST": "за західноафриканським літнім часом", "CLT": "за стандартним чилійським часом", "AEST": "за стандартним східноавстралійським часом", "WIT": "за східноіндонезійським часом", "ART": "за стандартним аргентинським часом", "HADT": "за літнім гавайсько-алеутським часом", "HEOG": "за літнім західним часом у Ґренландії", "HKST": "за літнім часом у Гонконзі", "EAT": "за східноафриканським часом", "NZDT": "за літнім часом у Новій Зеландії", "HNNOMX": "за стандартним північнозахідним часом у Мексиці", "IST": "за індійським стандартним часом", "MDT": "MDT", "MST": "MST", "AEDT": "за літнім східноавстралійським часом", "HNCU": "за стандартним часом на Кубі", "ARST": "за літнім аргентинським часом", "CDT": "за північноамериканським центральним літнім часом", "LHDT": "за літнім часом на острові Лорд-Хау", "SRT": "за часом у Суринамі", "GFT": "за часом Французької Гвіани", "WITA": "за центральноіндонезійським часом", "TMT": "за стандартним часом у Туркменістані", "PST": "за північноамериканським тихоокеанським стандартним часом", "SAST": "за південноафриканським часом", "HAST": "за стандартним гавайсько-алеутським часом", "SGT": "за часом у Сінґапурі", "HNEG": "за стандартним східним часом у Ґренландії", "EDT": "за північноамериканським східним літнім часом", "AWST": "за стандартним західноавстралійським часом", "AST": "за атлантичним стандартним часом", "HNOG": "за стандартним західним часом у Ґренландії", "JST": "за японським стандартним часом", "JDT": "за японським літнім часом", "MEZ": "за центральноєвропейським стандартним часом", "WEZ": "за західноєвропейським стандартним часом", "ECT": "за часом в Еквадорі", "HNPMX": "за стандартним тихоокеанським часом у Мексиці", "GYT": "за часом у Ґаяні", "TMST": "за літнім часом у Туркменістані", "HAT": "за літнім часом у Ньюфаундленд", "WAT": "за західноафриканським стандартним часом", "CLST": "за літнім чилійським часом", "COT": "за стандартним колумбійським часом", "COST": "за літнім колумбійським часом", "OESZ": "за східноєвропейським літнім часом", "HEEG": "за літнім східним часом у Ґренландії", "ACST": "за стандартним центральноавстралійським часом", "ACDT": "за літнім центральноавстралійським часом", "UYST": "за літнім часом в Уруґваї", "EST": "за північноамериканським східним стандартним часом", "AKDT": "за літнім часом на Алясці", "BT": "за часом у Бутані", "WESZ": "за західноєвропейським літнім часом", "AKST": "за стандартним часом на Алясці", "ACWST": "за стандартним центральнозахідним австралійським часом", "HEPMX": "за літнім тихоокеанським часом у Мексиці", "GMT": "за Ґрінвічем", "WART": "за стандартним західноаргентинським часом", "CST": "за північноамериканським центральним стандартним часом", "LHST": "за стандартним часом на острові Лорд-Хау", "UYT": "за стандартним часом в Уруґваї"},
	}
}

// Locale returns the current translators string locale
func (uk *uk) Locale() string {
	return uk.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uk'
func (uk *uk) PluralsCardinal() []locales.PluralRule {
	return uk.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uk'
func (uk *uk) PluralsOrdinal() []locales.PluralRule {
	return uk.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uk'
func (uk *uk) PluralsRange() []locales.PluralRule {
	return uk.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uk'
func (uk *uk) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	iMod10 := i % 10
	iMod100 := i % 100

	if v == 0 && iMod10 == 1 && iMod100 != 11 {
		return locales.PluralRuleOne
	} else if v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14) {
		return locales.PluralRuleFew
	} else if (v == 0 && iMod10 == 0) || (v == 0 && iMod10 >= 5 && iMod10 <= 9) || (v == 0 && iMod100 >= 11 && iMod100 <= 14) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uk'
func (uk *uk) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if nMod10 == 3 && nMod100 != 13 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uk'
func (uk *uk) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uk.CardinalPluralRule(num1, v1)
	end := uk.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uk *uk) MonthAbbreviated(month time.Month) string {
	return uk.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uk *uk) MonthsAbbreviated() []string {
	return uk.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uk *uk) MonthNarrow(month time.Month) string {
	return uk.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uk *uk) MonthsNarrow() []string {
	return uk.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uk *uk) MonthWide(month time.Month) string {
	return uk.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uk *uk) MonthsWide() []string {
	return uk.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uk *uk) WeekdayAbbreviated(weekday time.Weekday) string {
	return uk.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uk *uk) WeekdaysAbbreviated() []string {
	return uk.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uk *uk) WeekdayNarrow(weekday time.Weekday) string {
	return uk.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uk *uk) WeekdaysNarrow() []string {
	return uk.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uk *uk) WeekdayShort(weekday time.Weekday) string {
	return uk.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uk *uk) WeekdaysShort() []string {
	return uk.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uk *uk) WeekdayWide(weekday time.Weekday) string {
	return uk.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uk *uk) WeekdaysWide() []string {
	return uk.daysWide
}

// Decimal returns the decimal point of number
func (uk *uk) Decimal() string {
	return uk.decimal
}

// Group returns the group of number
func (uk *uk) Group() string {
	return uk.group
}

// Group returns the minus sign of number
func (uk *uk) Minus() string {
	return uk.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uk' and handles both Whole and Real numbers based on 'v'
func (uk *uk) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uk' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uk *uk) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uk.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uk'
func (uk *uk) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uk.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, uk.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uk'
// in accounting notation.
func (uk *uk) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uk.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uk.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(uk.group) - 1; j >= 0; j-- {
					b = append(b, uk.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, uk.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uk.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, uk.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, uk.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uk'
func (uk *uk) FmtDateShort(t time.Time) string {

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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'uk'
func (uk *uk) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd1, 0x80}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uk'
func (uk *uk) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd1, 0x80}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uk'
func (uk *uk) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, uk.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uk.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd1, 0x80}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uk'
func (uk *uk) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uk'
func (uk *uk) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uk'
func (uk *uk) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uk'
func (uk *uk) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uk.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := uk.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
