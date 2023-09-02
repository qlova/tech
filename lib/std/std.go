package std

import (
	"unsafe"

	"qlova.tech/abi"
	"qlova.tech/ffi"
)

type (
	Error            = abi.Int
	RangeCheckedSize = abi.Size
)

var Complex struct {
	ffi.Header `linux:"libm.so.6"`

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
	ffi.Header `linux:"libm.so.6"`

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
	ffi.Header `linux:"libm.so.6"`

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

var Char struct {
	ffi.Header `linux:"libc.so.6"`

	IsAlphaNumeric func(abi.Int) abi.Int `cgo:"isalnum"`
	IsAlpha        func(abi.Int) abi.Int `cgo:"isalpha"`
	IsUpper        func(abi.Int) abi.Int `cgo:"isupper"`
	IsLower        func(abi.Int) abi.Int `cgo:"islower"`
	IsDigit        func(abi.Int) abi.Int `cgo:"isdigit"`
	IsHexDigit     func(abi.Int) abi.Int `cgo:"isxdigit"`
	IsControl      func(abi.Int) abi.Int `cgo:"iscntrl"`
	IsGraph        func(abi.Int) abi.Int `cgo:"isgraph"`
	IsSpace        func(abi.Int) abi.Int `cgo:"isspace"`
	IsBlank        func(abi.Int) abi.Int `cgo:"isblank"`
	IsPrint        func(abi.Int) abi.Int `cgo:"isprint"`
	IsPuncuation   func(abi.Int) abi.Int `cgo:"ispunct"`

	ToLower func(abi.Int) abi.Int `cgo:"tolower"`
	ToUpper func(abi.Int) abi.Int `cgo:"toupper"`
}

var FloatingPoint struct {
	ffi.Header `linux:"libc.so.6"`

	ClearExceptions   func(abi.Int) abi.Int                            `cgo:"feclearexcept"`
	Exceptions        func(abi.Int) abi.Int                            `cgo:"fetestexcept"`
	RaiseExceptions   func(abi.Int) abi.Int                            `cgo:"feraiseexcept"`
	GetExceptionFlag  func(flagp abi.Pointer, excepts abi.Int) abi.Int `cgo:"fgetexceptflag"`
	SetExceptionFlag  func(flagp abi.Pointer, excepts abi.Int) abi.Int `cgo:"fsetexceptflag"`
	SetRoundingMode   func(abi.Int) abi.Int                            `cgo:"fesetround"`
	GetEnvironment    func(fenv abi.Pointer) abi.Int                   `cgo:"fegetenv"`
	SetEnvironment    func(fenv abi.Pointer) abi.Int                   `cgo:"fesetenv"`
	UpdateEnvironment func(fenv abi.Pointer) abi.Int                   `cgo:"feupdateenv"`
	HoldExceptions    func(fenv abi.Pointer) abi.Int                   `cgo:"feholdexcept"`
}

var Locale struct {
	ffi.Header `linux:"libc.so.6"`

	Set func(abi.Int, abi.Pointer) string `cgo:"setlocale"`
	Get func() abi.Pointer                `cgo:"localeconv"`
}

var Program struct {
	ffi.Header `linux:"libc.so.6"`

	Abort              func()        `cgo:"abort"`
	Exit               func(abi.Int) `cgo:"exit"`
	ExitFast           func(abi.Int) `cgo:"quick_exit"`
	ExitWithoutCleanup func(abi.Int) `cgo:"_Exit"`
	OnExit             func(func())  `cgo:"atexit"`
	OnExitFast         func(func())  `cgo:"at_quick_exit"`
}

var Math struct {
	ffi.Header `linux:"libm.so.6"`

	Acos  func(float64) float64            `cgo:"acos"`
	Asin  func(float64) float64            `cgo:"asin"`
	Atan  func(float64) float64            `cgo:"atan"`
	Atan2 func(float64, float64) float64   `cgo:"atan2"`
	Cos   func(float64) float64            `cgo:"cos"`
	Cosh  func(float64) float64            `cgo:"cosh"`
	Sin   func(float64) float64            `cgo:"sin"`
	Sinh  func(float64) float64            `cgo:"sinh"`
	Tanh  func(float64) float64            `cgo:"tanh"`
	Exp   func(float64) float64            `cgo:"exp"`
	Frexp func(float64) (float64, int32)   `cgo:"frexp"`
	Ldexp func(float64, int32) float64     `cgo:"ldexp"`
	Log   func(float64) float64            `cgo:"log"`
	Log10 func(float64) float64            `cgo:"log10"`
	Modf  func(float64) (float64, float64) `cgo:"modf"`
	Pow   func(float64, float64) float64   `cgo:"pow"`
	Sqrt  func(float64) float64            `cgo:"sqrt"`
	Ceil  func(float64) float64            `cgo:"ceil"`
	Fabs  func(float64) float64            `cgo:"fabs"`
	Floor func(float64) float64            `cgo:"floor"`
	Fmod  func(float64, float64) float64   `cgo:"fmod"`
}

type File abi.Pointer
type FilePos int

var IO struct {
	ffi.Header `linux:"libc.so.6"`

	Close         func(File) Error                                      `cgo:"fclose"`
	ClearError    func(File)                                            `cgo:"clearerr"`
	IsEOF         func(File) bool                                       `cgo:"feof"`
	IsError       func(File) bool                                       `cgo:"ferror"`
	Flush         func(File) Error                                      `cgo:"fflush"`
	GetPos        func(File) (Error, FilePos)                           `cgo:"ftell"`
	Open          func(string, string) File                             `cgo:"fopen"`
	Read          func(unsafe.Pointer, int, int, File) Error            `cgo:"fread"`
	Reopen        func(string, string, File) File                       `cgo:"freopen"`
	Seek          func(File, int, int32) Error                          `cgo:"fseek"`
	SetPos        func(File, FilePos)                                   `cgo:"fseek"`
	Tell          func(File) FilePos                                    `cgo:"ftell"`
	Write         func(unsafe.Pointer, int, int, File) Error            `cgo:"fwrite"`
	Remove        func(string) Error                                    `cgo:"remove"`
	Rename        func(string, string) Error                            `cgo:"rename"`
	Rewind        func(File)                                            `cgo:"rewind"`
	SetBuffer     func(File, unsafe.Pointer)                            `cgo:"setbuf"`
	SetBufferSize func(File, unsafe.Pointer, int32, int)                `cgo:"setvbuf"`
	TempFile      func() File                                           `cgo:"tmpfile"`
	TempName      func(string) string                                   `cgo:"tmpnam"`
	Fprintf       func(File, string, ...unsafe.Pointer) Error           `cgo:"fprintf"`
	Printf        func(string, ...unsafe.Pointer) int32                 `cgo:"printf"`
	Sprintf       func(unsafe.Pointer, string, ...unsafe.Pointer) Error `cgo:"sprintf"`
	Fscanf        func(File, string, ...unsafe.Pointer) Error           `cgo:"fscanf"`
	Scanf         func(string, ...unsafe.Pointer) Error                 `cgo:"scanf"`
	Sscanf        func(unsafe.Pointer, string, ...unsafe.Pointer) Error `cgo:"sscanf"`
	Fgetc         func(File) byte                                       `cgo:"fgetc"`
	Fgets         func(unsafe.Pointer, int, File) string                `cgo:"fgets"`
	Fputc         func(byte, File) byte                                 `cgo:"fputc"`
	Fputs         func(string, File) *byte                              `cgo:"fputs"`
	Getc          func(File) byte                                       `cgo:"getc"`
	GetChar       func() byte                                           `cgo:"getchar"`
	GetString     func(unsafe.Pointer) string                           `cgo:"gets"`
	Putc          func(byte, File) Error                                `cgo:"putc"`
	PutChar       func(byte) Error                                      `cgo:"putchar"`
	PutString     func(string) Error                                    `cgo:"puts"`
	Ungetc        func(byte, File) Error                                `cgo:"ungetc"`
	Error         func(string)                                          `cgo:"perror"`
}

var String struct {
	ffi.Header `linux:"libc.so.6"`

	ToFloat64    func(string) float64          `cgo:"atof"`
	ToInt32      func(string) int32            `cgo:"atoi"`
	ToInt        func(string) int              `cgo:"atol"`
	ParseFloat64 func([]byte) (float64, *byte) `cgo:"strtod"`
	ParseInt32   func([]byte) (int32, *byte)   `cgo:"strtol"`
	ParseUint    func([]byte) (uint, *byte)    `cgo:"strtoul"`

	Append     func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer          `cgo:"strcat"`
	AppendUpto func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer `cgo:"strncat"`

	Find func(string, byte) *byte `cgo:"strchr"`

	Compare        func(unsafe.Pointer, unsafe.Pointer) int32                   `cgo:"strcmp"`
	CompareCollate func(unsafe.Pointer, unsafe.Pointer) int32                   `cgo:"strcoll"`
	Copy           func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer          `cgo:"strcpy"`
	CopyUpto       func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer `cgo:"strncpy"`
	Match          func(string, string) int32                                   `cgo:"strcspn"`
	Error          func(Error) string                                           `cgo:"strerror"`
	Length         func(unsafe.Pointer) uintptr                                 `cgo:"strlen"`
	MatchFirst     func([]byte, string) *byte                                   `cgo:"strpbrk"`
	MatchLast      func([]byte, string) *byte                                   `cgo:"strrchr"`
	MatchLength    func(string, string) uintptr                                 `cgo:"strspn"`
	Contains       func([]byte, string) *byte                                   `cgo:"strstr"`
	Tokenize       func(unsafe.Pointer, string) unsafe.Pointer                  `cgo:"strtok"`
	Localize       func(unsafe.Pointer, unsafe.Pointer, uintptr) uintptr        `cgo:"strxfrm"`
}

var Lib struct {
	ffi.Header `linux:"libc.so.6"`

	Calloc  func(uintptr, uintptr) unsafe.Pointer        `cgo:"calloc"`
	Free    func(unsafe.Pointer)                         `cgo:"free"`
	Malloc  func(uintptr) unsafe.Pointer                 `cgo:"malloc"`
	Realloc func(unsafe.Pointer, uintptr) unsafe.Pointer `cgo:"realloc"`
	Abort   func()                                       `cgo:"abort"`
	AtExit  func(func()) Error                           `cgo:"atexit"`
	Exit    func(int)                                    `cgo:"exit"`
	Getenv  func(string) string                          `cgo:"getenv"`
	System  func(string) Error                           `cgo:"system"`

	BinarySearch func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr, func(unsafe.Pointer, unsafe.Pointer) int) unsafe.Pointer `cgo:"bsearch"`

	Sort func(unsafe.Pointer, uintptr, uintptr, func(unsafe.Pointer, unsafe.Pointer) int32) unsafe.Pointer `cgo:"qsort"`

	MemoryCompare func(unsafe.Pointer, unsafe.Pointer, uintptr) int32 `cgo:"memcmp"`
	MemoryCopy    func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	MemoryMove    func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	MemorySet     func(unsafe.Pointer, byte, uintptr) unsafe.Pointer
	MemoryFind    func(string, byte, uintptr) *byte `cgo:"memchar"`
}

var I32 struct {
	ffi.Header `linux:"libc.so.6"`

	Abs         func(int32) int32             `cgo:"abs"`
	Div         func(int32, int32) Div[int32] `cgo:"div"`
	Rand        func() int32                  `cgo:"rand"`
	SetRandSeed func(int32)                   `cgo:"srand"`
}

var Int struct {
	ffi.Header `linux:"libc.so.6"`

	Abs func(int) int           `cgo:"labs"`
	Div func(int, int) Div[int] `cgo:"ldiv"`
}

type Time int32
type Ticks int32

var Clock struct {
	ffi.Header `linux:"libc.so.6"`

	String func(Time) string        `cgo:"asctime"`
	Local  func(Time) string        `cgo:"ctime"`
	Diff   func(Time, Time) float32 `cgo:"difftime"`
	Ticks  func() Ticks             `cgo:"clock"`
	Time   func(*Time) Time         `cgo:"time"`
}

type Div[T int32 | int] struct {
	Quo int
	Rem int
}
