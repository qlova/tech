//Package in provides common and useful units of measurement.
package qty

import (
	"qlova.tech/qty/si"
)

const pound = 453.59237

//Grams returns a quantity of mass equivalent to the given amount of SI grams.
func Grams(amount float64) Mass { return mass(1, "g", amount) }

//Kilograms returns a quantity of mass equivalent to the given amount of SI Kilograms.
func Kilograms(amount float64) Mass { return mass(si.Kilo, "kg", amount) }

//Tonnes returns a quantity of mass equivalent to the given amount of SI tonnes.
func Tonnes(amount float64) Mass { return mass(1000*si.Kilo, "ft", amount) }

//Daltons returns a quantity of mass equivalent to the given amount of SI daltons.
func Daltons(amount float64) Mass { return mass(1.6605390666050e-27*si.Kilo, "Da", amount) }

//Grains returns a quantity of mass equivalent to the given amount of avoirdupois grains.
func Grains(amount float64) Mass { return mass(pound/7000, "gr", amount) }

//Drams returns a quantity of mass equivalent to the given amount of avoirdupois drams.
func Drams(amount float64) Mass { return mass(pound/256, "dr", amount) }

//Ounces returns a quantity of mass equivalent to the given amount of avoirdupois ounces.
func Ounces(amount float64) Mass { return mass(pound/16, "oz", amount) }

//Pounds returns a quantity of mass equivalent to the given amount of avoirdupois pounds.
func Pounds(amount float64) Mass { return mass(pound, "lb", amount) }

//Quarters returns a quantity of mass equivalent to the given amount of avoirdupois quarters.
func Quarters(amount float64) Mass { return mass(pound*25, "qr", amount) }

//ShortHundredWeights returns a quantity of mass equivalent to the given amount of avoirdupois short hundredweights.
func ShortHundredWeights(amount float64) Mass { return mass(pound*100, "cwt", amount) }

//Tons returns a quantity of mass equivalent to the given amount of avoirdupois tons.
func Tons(amount float64) Mass { return mass(pound*2000, "ton", amount) }

//SolarMasses returns a quantity of mass equivalent to the given amount of solar masses.
func SolarMasses(amount float64) Mass { return mass(1.98847e30, "M☉", amount) }

//EarthMasses returns a quantity of mass equivalent to the given amount of earth masses.
func EarthMasses(amount float64) Mass { return mass(5.9722e24, "M⊕", amount) }

//Mass is a quantity of mass.
//Measured in grams by default.
type Mass struct {
	g float64

	unit   float64
	symbol string
}

func mass(unit float64, format string, quantity float64) Mass {
	return Mass{unit * quantity, unit, format}
}

//Set sets this mass to the given mass, but
//preserves the formatting and units of the
//this mass.
func (mass *Mass) Set(to Mass) { mass.g = to.g }

//Add adds the other mass to this mass.
func (mass *Mass) Add(other Mass) { mass.g += other.g }

//Sub subtracts the other mass from this mass.
func (mass *Mass) Sub(other Mass) { mass.g -= other.g }

//Mul multiplies the mass by the given scalar.
func (mass *Mass) Mul(scalar float64) { mass.g *= scalar }

//Div divides the mass by the given scalar.
func (mass *Mass) Div(scalar float64) { mass.g /= scalar }

//Plus returns a mass that is the sum of both masses.
//The formatting and units of this mass are reset.
func (mass Mass) Plus(other Mass) Mass { return Mass{g: mass.g + other.g} }

//Minus returns a mass that is the difference
//of both masses. The formatting and units of this mass are reset.
func (mass Mass) Minus(other Mass) Mass { return Mass{g: mass.g - other.g} }

//Times returns a copy of this mass multiplied by the given scalar.
func (mass Mass) Times(scalar float64) Mass { mass.Mul(scalar); return mass }

//DividedBy returns a copy of this mass, divided by the given scalar.
func (mass Mass) DividedBy(scalar float64) Mass { mass.Div(scalar); return mass }

//Grams returns the mass as a quantity of grams.
func (mass Mass) Grams() float64 { return mass.g }

//Float64 returns the mass as a quanity of
//units that the mass was measured in.
func (mass Mass) Float64() float64 {
	if mass.unit == 0 {
		return mass.g
	}
	return mass.g / mass.unit
}

//String returns this mass formatted as a string with the
//format and units that the mass was originally measured in.
func (mass Mass) String() string {
	if mass.symbol == "" || mass.unit == 0 {
		return fmtsi("g", mass.g)
	}
	return format(mass.g/mass.unit, mass.symbol)
}

//Display returns a copy of this mass as if it were measured
//with the given format and units.
func (mass Mass) In(in func(float64) Mass) Mass {
	f := in(0)
	mass.unit = f.unit
	mass.symbol = f.symbol
	return mass
}
