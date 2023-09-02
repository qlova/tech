package std

import (
	"qlova.tech/abi"
	"qlova.tech/ffi"
)

var Complex struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Real func(abi.ComplexDouble) abi.Double        `ffi:"creal"`
	Imag func(abi.ComplexDouble) abi.Double        `ffi:"cimag"`
	Abs  func(abi.ComplexDouble) abi.Double        `ffi:"cabs"`
	Arg  func(abi.ComplexDouble) abi.Double        `ffi:"carg"`
	Conj func(abi.ComplexDouble) abi.ComplexDouble `ffi:"conj"`
	Proj func(abi.ComplexDouble) abi.ComplexDouble `ffi:"cproj"`

	Exp func(abi.ComplexDouble) abi.ComplexDouble                    `ffi:"cexp"`
	Log func(abi.ComplexDouble) abi.ComplexDouble                    `ffi:"clog"`
	Pow func(abi.ComplexDouble, abi.ComplexDouble) abi.ComplexDouble `ffi:"cpow"`

	Sqrt func(abi.ComplexDouble) abi.ComplexDouble `ffi:"csqrt"`
	Sin  func(abi.ComplexDouble) abi.ComplexDouble `ffi:"csin"`
	Cos  func(abi.ComplexDouble) abi.ComplexDouble `ffi:"ccos"`
	Tan  func(abi.ComplexDouble) abi.ComplexDouble `ffi:"ctan"`
	Asin func(abi.ComplexDouble) abi.ComplexDouble `ffi:"casin"`
	Acos func(abi.ComplexDouble) abi.ComplexDouble `ffi:"cacos"`
	Atan func(abi.ComplexDouble) abi.ComplexDouble `ffi:"catan"`

	Sinh  func(abi.ComplexDouble) abi.ComplexDouble `ffi:"csinh"`
	Cosh  func(abi.ComplexDouble) abi.ComplexDouble `ffi:"ccosh"`
	Tanh  func(abi.ComplexDouble) abi.ComplexDouble `ffi:"ctanh"`
	Asinh func(abi.ComplexDouble) abi.ComplexDouble `ffi:"casinh"`
	Acosh func(abi.ComplexDouble) abi.ComplexDouble `ffi:"cacosh"`
	Atanh func(abi.ComplexDouble) abi.ComplexDouble `ffi:"catanh"`
}

var ComplexFloat struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Real func(abi.ComplexFloat) abi.Float        `ffi:"crealf"`
	Imag func(abi.ComplexFloat) abi.Float        `ffi:"cimagf"`
	Abs  func(abi.ComplexFloat) abi.Float        `ffi:"cabsf"`
	Arg  func(abi.ComplexFloat) abi.Float        `ffi:"cargf"`
	Conj func(abi.ComplexFloat) abi.ComplexFloat `ffi:"conjf"`
	Proj func(abi.ComplexFloat) abi.ComplexFloat `ffi:"cprojf"`

	Exp func(abi.ComplexFloat) abi.ComplexFloat                   `ffi:"cexpf"`
	Log func(abi.ComplexFloat) abi.ComplexFloat                   `ffi:"clogf"`
	Pow func(abi.ComplexFloat, abi.ComplexFloat) abi.ComplexFloat `ffi:"cpowf"`

	Sqrt func(abi.ComplexFloat) abi.ComplexFloat `ffi:"csqrtf"`
	Sin  func(abi.ComplexFloat) abi.ComplexFloat `ffi:"csinf"`
	Cos  func(abi.ComplexFloat) abi.ComplexFloat `ffi:"ccosf"`
	Tan  func(abi.ComplexFloat) abi.ComplexFloat `ffi:"ctanf"`
	Asin func(abi.ComplexFloat) abi.ComplexFloat `ffi:"casinf"`
	Acos func(abi.ComplexFloat) abi.ComplexFloat `ffi:"cacosf"`
	Atan func(abi.ComplexFloat) abi.ComplexFloat `ffi:"catanf"`

	Sinh  func(abi.ComplexFloat) abi.ComplexFloat `ffi:"csinhf"`
	Cosh  func(abi.ComplexFloat) abi.ComplexFloat `ffi:"ccoshf"`
	Tanh  func(abi.ComplexFloat) abi.ComplexFloat `ffi:"ctanhf"`
	Asinh func(abi.ComplexFloat) abi.ComplexFloat `ffi:"casinhf"`
	Acosh func(abi.ComplexFloat) abi.ComplexFloat `ffi:"cacoshf"`
	Atanh func(abi.ComplexFloat) abi.ComplexFloat `ffi:"catanhf"`
}

var ComplexDoubleLong struct {
	ffi.Header `linux:"libm.so.6" darwin:"libSystem.dylib"`

	Real func(abi.ComplexDoubleLong) abi.DoubleLong        `ffi:"creall"`
	Imag func(abi.ComplexDoubleLong) abi.DoubleLong        `ffi:"cimagl"`
	Abs  func(abi.ComplexDoubleLong) abi.DoubleLong        `ffi:"cabsl"`
	Arg  func(abi.ComplexDoubleLong) abi.DoubleLong        `ffi:"cargl"`
	Conj func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"conjl"`
	Proj func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"cprojl"`

	Exp func(abi.ComplexDoubleLong) abi.ComplexDoubleLong                        `ffi:"cexpl"`
	Log func(abi.ComplexDoubleLong) abi.ComplexDoubleLong                        `ffi:"clogl"`
	Pow func(abi.ComplexDoubleLong, abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"cpowl"`

	Sqrt func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"csqrtl"`
	Sin  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"csinl"`
	Cos  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"ccosl"`
	Tan  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"ctanl"`
	Asin func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"casinl"`
	Acos func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"cacosl"`
	Atan func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"catanl"`

	Sinh  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"csinhl"`
	Cosh  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"ccoshl"`
	Tanh  func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"ctanhl"`
	Asinh func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"casinhl"`
	Acosh func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"cacoshl"`
	Atanh func(abi.ComplexDoubleLong) abi.ComplexDoubleLong `ffi:"catanhl"`
}
