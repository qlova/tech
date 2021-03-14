//Package si provides the prefix constants for the International System of Units.
package si

/*
	Prefixes are added to unit names to produce multiples and submultiples
	of the original unit. All of these are integer powers of ten, and above
	a hundred or below a hundredth all are integer powers of a thousand.
	For example, kilo- denotes a multiple of a thousand and milli- denotes a
	multiple of a thousandth, so there are one thousand millimetres to the metre
	and one thousand metres to the kilometre. The prefixes are never combined,
	so for example a millionth of a metre is a micrometre, not a millimillimetre.
	Multiples of the kilogram are named as if the gram were the base unit, so a
	millionth of a kilogram is a milligram, not a microkilogram. When prefixes
	are used to form multiples and submultiples of SI base and derived units,
	the resulting units are no longer coherent.
*/
const (
	Yocto = 1e-24
	Zepto = 1e-21
	Atto  = 1e-18
	Femto = 1e-15
	Pico  = 1e-12
	Nano  = 1e-9
	Micro = 1e-6
	Milli = 1e-3
	Centi = 1e-2
	Deci  = 1e-1
	Deca  = 1e1
	Hecto = 1e2
	Kilo  = 1e3
	Mega  = 1e6
	Giga  = 1e9
	Tera  = 1e12
	Peta  = 1e15
	Exa   = 1e18
	Zetta = 1e21
	Yotta = 1e24
)
