//Package qty provides types that represent standard unit quantities.
package qty

import (
	"strconv"
	"strings"

	"qlova.tech/qty/si"
)

func format(amount float64, symbol string) string {
	return strings.TrimRight(strings.TrimRight(strconv.FormatFloat(amount, 'f', 2, 64), "0"), ".") + " " + symbol
}

func fmtsi(symbol string, amount float64) string {

	switch {
	case amount <= si.Yocto:
		return format(amount/si.Yocto, "y"+symbol)

	case amount <= si.Zepto:
		return format(amount/si.Zepto, "z"+symbol)

	case amount <= si.Atto:
		return format(amount/si.Atto, "a"+symbol)

	case amount <= si.Femto:
		return format(amount/si.Femto, "f"+symbol)

	case amount <= si.Pico:
		return format(amount/si.Pico, "p"+symbol)

	case amount <= si.Nano:
		return format(amount/si.Nano, "n"+symbol)

	case amount <= si.Micro:
		return format(amount/si.Micro, "μ"+symbol)

	case amount <= si.Milli:
		return format(amount/si.Milli, "m"+symbol)

	case amount <= si.Centi:
		return format(amount/si.Centi, "c"+symbol)

	case amount <= si.Deci:
		return format(amount/si.Deci, "d"+symbol)

	case amount >= si.Yotta:
		return format(amount/si.Yotta, "Y"+symbol)

	case amount >= si.Zetta:
		return format(amount/si.Zetta, "Z"+symbol)

	case amount >= si.Exa:
		return format(amount/si.Exa, "E"+symbol)

	case amount >= si.Peta:
		return format(amount/si.Peta, "P"+symbol)

	case amount >= si.Tera:
		return format(amount/si.Tera, "T"+symbol)

	case amount > si.Giga:
		return format(amount/si.Giga, "G"+symbol)

	case amount >= si.Mega:
		return format(amount/si.Mega, "M"+symbol)

	case amount >= si.Kilo:
		return format(amount/si.Kilo, "k"+symbol)

	case amount >= si.Hecto:
		return format(amount/si.Hecto, "h"+symbol)

	case amount >= si.Deca:
		return format(amount/si.Deca, "da"+symbol)

	default:
		return format(amount, symbol)
	}
}

func fmtbites(amount float64) string {

	/*
		Prefixes that are not SI, but serve the same purpose, instead in base 2.
	*/
	const (
		Kibi = 1 << 10
		Mebi = 1 << 20
		Gibi = 1 << 30
		Tebi = 1 << 40
		Pebi = 1 << 50
		Exbi = 1 << 60
		Zebi = 1 << 70
		Yobi = 1 << 80
	)

	const symbol = "B"

	switch {
	case amount <= si.Yocto:
		return format(amount/si.Yocto, "y"+symbol)

	case amount <= si.Zepto:
		return format(amount/si.Zepto, "z"+symbol)

	case amount <= si.Atto:
		return format(amount/si.Atto, "a"+symbol)

	case amount <= si.Femto:
		return format(amount/si.Femto, "f"+symbol)

	case amount <= si.Pico:
		return format(amount/si.Pico, "p"+symbol)

	case amount <= si.Nano:
		return format(amount/si.Nano, "n"+symbol)

	case amount <= si.Micro:
		return format(amount/si.Micro, "μ"+symbol)

	case amount <= si.Milli:
		return format(amount/si.Milli, "m"+symbol)

	case amount <= si.Centi:
		return format(amount/si.Centi, "c"+symbol)

	case amount <= si.Deci:
		return format(amount/si.Deci, "d"+symbol)

	case amount >= Yobi:
		return format(amount/Yobi, "Yi"+symbol)

	case amount >= Zebi:
		return format(amount/Zebi, "Zi"+symbol)

	case amount >= Exbi:
		return format(amount/Exbi, "Ei"+symbol)

	case amount >= Pebi:
		return format(amount/Pebi, "Pi"+symbol)

	case amount >= Tebi:
		return format(amount/Tebi, "Ti"+symbol)

	case amount > Gibi:
		return format(amount/Gibi, "Gi"+symbol)

	case amount >= Mebi:
		return format(amount/Mebi, "Mi"+symbol)

	case amount >= Kibi:
		return format(amount/Kibi, "ki"+symbol)

	default:
		return format(amount, symbol)
	}
}
