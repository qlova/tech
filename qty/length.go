package qty

const yard = 0.9144
const foot = yard / 3
const inch = foot / 12
const km = 1000

//Metres returns a quantity of length equivalent to the given amount of SI metres.
func Metres(amount float64) Length { return length(1, "m", amount) }

//Meters returns a quantity of length equivalent to the given amount of SI meters.
func Meters(amount float64) Length { return length(1, "m", amount) }

//Inches returns a quantity of length equivalent to the given amount of imperial/customary inches.
func Inches(amount float64) Length { return length(inch, "in", amount) }

//Feet returns a quantity of length equivalent to the given amount of imperial/customary feet.
func Feet(amount float64) Length { return length(foot, "ft", amount) }

//Yards returns a quantity of length equivalent to the given amount of imperial/customary yards.
func Yards(amount float64) Length { return length(yard, "yd", amount) }

//Lines returns a quantity of length equivalent to the given amount of imperial/customary lines.
func Lines(amount float64) Length { return length(inch/12, "in", amount) }

//Thous returns a quantity of length equivalent to the given amount of imperial/customary thous/mils.
func Thous(amount float64) Length { return length(inch/1000, "thou", amount) }

//Miles returns a quantity of length equivalent to the given amount of imperial/customary miles.
func Miles(amount float64) Length { return length(yard*1760, "mi", amount) }

//NauticalMiles returns a quantity of length equivalent to the given amount of nautical miles.
func NauticalMiles(amount float64) Length { return length(1852, "NM", amount) }

//LightYears returns a quantity of length equivalent to the given amount of light years.
func LightYears(amount float64) Length { return length(9460730472580.8*km, "ly", amount) }

//Parsecs returns a quantity of length equivalent to the given amount of parsecs.
func Parsecs(amount float64) Length { return length(30856775814671.9*km, "pc", amount) }

//AstronomicalUnits returns a quantity of length equivalent to the given amount of astronomical units.
func AstronomicalUnits(amount float64) Length { return length(149597870700, "au", amount) }

//Length is a quantity of length.
//Measured in metres by default.
type Length struct {
	m float64

	unit   float64
	symbol string
}

//Distance is an alias of Length.
type Distance = Length

func length(unit float64, format string, quantity float64) Length {
	return Length{unit * quantity, unit, format}
}

//Set sets this length to the given length, but
//preserves the formatting and units of the
//this length.
func (length *Length) Set(to Length) { length.m = to.m }

//Add adds the other length to this length.
func (length *Length) Add(other Length) { length.m += other.m }

//Sub subtracts the other length from this length.
func (length *Length) Sub(other Length) { length.m -= other.m }

//Mul multiplies the length by the given scalar.
func (length *Length) Mul(scalar float64) { length.m *= scalar }

//Div divides the length by the given scalar.
func (length *Length) Div(scalar float64) { length.m /= scalar }

//Plus returns a length that is the sum of both lengthes.
//The formatting and units of this length are reset.
func (length Length) Plus(other Length) Length { return Length{m: length.m + other.m} }

//Minus returns a length that is the difference
//of both lengthes. The formatting and units of this length are reset.
func (length Length) Minus(other Length) Length { return Length{m: length.m - other.m} }

//Times returns a copy of this length multiplied by the given scalar.
func (length Length) Times(scalar float64) Length { length.Mul(scalar); return length }

//DividedBy returns a copy of this length, divided by the given scalar.
func (length Length) DividedBy(scalar float64) Length { length.Div(scalar); return length }

//Metres returns the length as a quantity of metres.
func (length Length) Metres() float64 { return length.m }

//Float64 returns the length as a quanity of
//units that the length was measured in.
func (length Length) Float64() float64 {
	if length.unit == 0 {
		return length.m
	}
	return length.m / length.unit
}

//String returns this length formatted as a string with the
//format and units that the length was originally measured in.
func (length Length) String() string {
	if length.symbol == "" || length.unit == 0 {
		return fmtsi("m", length.m)
	}
	return format(length.m/length.unit, length.symbol)
}

//Display returns a copy of this length as if it were measured
//with the given format and units.
func (length Length) In(in func(float64) Length) Length {
	f := in(0)
	length.unit = f.unit
	length.symbol = f.symbol
	return length
}
