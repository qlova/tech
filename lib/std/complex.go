package std

import (
	"qlova.tech/abi"
	"qlova.tech/ffi"
)


var Complex struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Real func(abi.ComplexDouble) abi.Double        `cgo:"creal"`
	Imag func(abi.ComplexDouble) abi.Double        `cgo:"cimag"`
	Abs  func(abi.ComplexDouble) abi.Double        `cgo:"cabs"`
	Arg  func(abi.ComplexDouble) abi.Double        `cgo:"carg"`
	Conj func(abi.ComplexDouble) abi.ComplexDouble `cgo:"conj"`
	Proj func(abi.ComplexDouble) abi.ComplexDouble `cgo:"cproj"`

	Exp func(abi.ComplexDouble) abi.ComplexDouble                    `cgo:"cexp"`
	Log func(abi.ComplexDouble) abi.ComplexDouble                    `cgo:"clog"`
	Pow func(abi.ComplexDouble, abi.ComplexDouble) abi.ComplexDouble `cgo:"cpow"`

	Sqrt func(abi.ComplexDouble) abi.ComplexDouble `cgo:"csqrt"`
	Sin  func(abi.ComplexDouble) abi.ComplexDouble `cgo:"csin"`
	Cos  func(abi.ComplexDouble) abi.ComplexDouble `cgo:"ccos"`
	Tan  func(abi.ComplexDouble) abi.ComplexDouble `cgo:"ctan"`
	Asin func(abi.ComplexDouble) abi.ComplexDouble `cgo:"casin"`
	Acos func(abi.ComplexDouble) abi.ComplexDouble `cgo:"cacos"`
	Atan func(abi.ComplexDouble) abi.ComplexDouble `cgo:"catan"`

	Sinh  func(abi.ComplexDouble) abi.ComplexDouble `cgo:"csinh"`
	Cosh  func(abi.ComplexDouble) abi.ComplexDouble `cgo:"ccosh"`
	Tanh  func(abi.ComplexDouble) abi.ComplexDouble `cgo:"ctanh"`
	Asinh func(abi.ComplexDouble) abi.ComplexDouble `cgo:"casinh"`
	Acosh func(abi.ComplexDouble) abi.ComplexDouble `cgo:"cacosh"`
	Atanh func(abi.ComplexDouble) abi.ComplexDouble `cgo:"catanh"`
}

var ComplexFloat struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Real func(abi.ComplexFloat) abi.Float        `cgo:"crealf"`
	Imag func(abi.ComplexFloat) abi.Float        `cgo:"cimagf"`
	Abs  func(abi.ComplexFloat) abi.Float        `cgo:"cabsf"`
	Arg  func(abi.ComplexFloat) abi.Float        `cgo:"cargf"`
	Conj func(abi.ComplexFloat) abi.ComplexFloat `cgo:"conjf"`
	Proj func(abi.ComplexFloat) abi.ComplexFloat `cgo:"cprojf"`

	Exp func(abi.ComplexFloat) abi.ComplexFloat                   `cgo:"cexpf"`
	Log func(abi.ComplexFloat) abi.ComplexFloat                   `cgo:"clogf"`
	Pow func(abi.ComplexFloat, abi.ComplexFloat) abi.ComplexFloat `cgo:"cpowf"`

	Sqrt func(abi.ComplexFloat) abi.ComplexFloat `cgo:"csqrtf"`
	Sin  func(abi.ComplexFloat) abi.ComplexFloat `cgo:"csinf"`
	Cos  func(abi.ComplexFloat) abi.ComplexFloat `cgo:"ccosf"`
	Tan  func(abi.ComplexFloat) abi.ComplexFloat `cgo:"ctanf"`
	Asin func(abi.ComplexFloat) abi.ComplexFloat `cgo:"casinf"`
	Acos func(abi.ComplexFloat) abi.ComplexFloat `cgo:"cacosf"`
	Atan func(abi.ComplexFloat) abi.ComplexFloat `cgo:"catanf"`

	Sinh  func(abi.ComplexFloat) abi.ComplexFloat `cgo:"csinhf"`
	Cosh  func(abi.ComplexFloat) abi.ComplexFloat `cgo:"ccoshf"`
	Tanh  func(abi.ComplexFloat) abi.ComplexFloat `cgo:"ctanhf"`
	Asinh func(abi.ComplexFloat) abi.ComplexFloat `cgo:"casinhf"`
	Acosh func(abi.ComplexFloat) abi.ComplexFloat `cgo:"cacoshf"`
	Atanh func(abi.ComplexFloat) abi.ComplexFloat `cgo:"catanhf"`
}

var ComplexDoubleLong struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Real func(abi.ComplexDoubleLong) abi.DoubleLong        `cgo:"creall"`
	Imag func(abi.ComplexDoubleLong) abi.DoubleLong        `cgo:"cimagl"`
	Abs  func(abi.ComplexDoubleLong) abi.DoubleLong        `cgo:"cabsl"`
	Arg  func(abi.ComplexDoubleLong) abi.DoubleLong        `cgo:"cargl"`
	Conj func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"conjl"`
	Proj func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"cprojl"`

	Exp func(abi.ComplexDoubleLong) abi.ComplexDoubleLong                        `cgo:"cexpl"`
	Log func(abi.ComplexDoubleLong) abi.ComplexDoubleLong                        `cgo:"clogl"`
	Pow func(abi.ComplexDoubleLong, abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"cpowl"`

	Sqrt func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"csqrtl"`
	Sin  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"csinl"`
	Cos  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"ccosl"`
	Tan  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"ctanl"`
	Asin func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"casinl"`
	Acos func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"cacosl"`
	Atan func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"catanl"`

	Sinh  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"csinhl"`
	Cosh  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"ccoshl"`
	Tanh  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"ctanhl"`
	Asinh func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"casinhl"`
	Acosh func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"cacoshl"`
	Atanh func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `cgo:"catanhl"`
}