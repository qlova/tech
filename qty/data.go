package qty

import (
	"math"
	"strings"
)

//Bytes2 returns a quantity of data equivalent to the given amount of bytes.
//This Data will be formatted in powers of 2, ie MiB, Gib, etc
func Bytes2(amount float64) Data { return data(8, "iB", amount) }

//Bytes returns a quantity of data equivalent to the given amount of bytes.
//This Data will be formatted in powers of 10, ie MB, Gb, etc
func Bytes(amount float64) Data { return data(8, "B", amount) }

//Nibbles returns a quantity of data equivalent to the given amount of nibbles.
func Nibbles(amount float64) Data { return data(4, "nib(s)", amount) }

//Bits returns a quantity of data equivalent to the given amount of bits.
//This Data will be formatted in powers of 10, ie MB, Gb, etc
func Bits(amount float64) Data { return data(8, "b", amount) }

//Bits2 returns a quantity of data equivalent to the given amount of bits.
//This Data will be formatted in powers of 2, ie MiB, Gib, etc
func Bits2(amount float64) Data { return data(8, "ib", amount) }

//Trits returns a quantity of data equivalent to the given amount of trits.
func Trits(amount float64) Data { return data(math.Log2(3), "trit(s)", amount) }

//Nats returns a quantity of data equivalent to the given amount of natural units of information.
func Nats(amount float64) Data { return data(math.Log2(math.E), "nat", amount) }

//Dits returns a quantity of data equivalent to the given amount of bans/hartleys/dits.
func Dits(amount float64) Data { return data(math.Log2(10), "dit(s)", amount) }

//Data is a quantity of data.
//Measured in bits by default.
type Data struct {
	b float64

	unit   float64
	symbol string
}

//Information is an alias of Data.
type Information = Data

func data(unit float64, format string, quantity float64) Data {
	return Data{unit * quantity, unit, format}
}

//Set sets this data to the given data, but
//preserves the formatting and units of the
//this data.
func (data *Data) Set(to Data) { data.b = to.b }

//Add adds the other data to this data.
func (data *Data) Add(other Data) { data.b += other.b }

//Sub subtracts the other data from this data.
func (data *Data) Sub(other Data) { data.b -= other.b }

//Mul multiplies the data by the given scalar.
func (data *Data) Mul(scalar float64) { data.b *= scalar }

//Div divides the data by the given scalar.
func (data *Data) Div(scalar float64) { data.b /= scalar }

//Plus returns a data that is the sum of both dataes.
//The formatting and units of this data are reset.
func (data Data) Plus(other Data) Data { return Data{b: data.b + other.b} }

//Minus returns a data that is the difference
//of both dataes. The formatting and units of this data are reset.
func (data Data) Minus(other Data) Data { return Data{b: data.b - other.b} }

//Times returns a copy of this data multiplied by the given scalar.
func (data Data) Times(scalar float64) Data { data.Mul(scalar); return data }

//DividedBy returns a copy of this data, divided by the given scalar.
func (data Data) DividedBy(scalar float64) Data { data.Div(scalar); return data }

//Bits returns the data as a quantity of bits.
func (data Data) Bits() float64 { return data.b }

//Float64 returns the data as a quantity of
//units that the data was measured in.
func (data Data) Float64() float64 {
	if data.unit == 0 {
		return data.b
	}
	return data.b / data.unit
}

//String returns this data formatted as a string with the
//format and units that the data was originally measured in.
func (data Data) String() string {
	if data.symbol == "" || data.unit == 0 {
		return fmtsi("b", data.b)
	}
	switch data.symbol {
	case "b":
		return fmtsi("b", data.b)
	case "B":
		return strings.Replace(fmtsi("b", data.b/8), "b", "B", 1)
	case "iB":
		return fmtbites(data.b / 8)
	case "ib":
		return strings.Replace(fmtbites(data.b), "B", "b", 1)
	default:
		return format(data.b/data.unit, data.symbol)
	}
}

//Display returns a copy of this data as if it were measured
//with the given format and units.
func (data Data) In(in func(float64) Data) Data {
	f := in(0)
	data.unit = f.unit
	data.symbol = f.symbol
	return data
}
