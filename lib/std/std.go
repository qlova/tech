package std

import (
	"runtime/debug"

	"qlova.tech/abi"
	"qlova.tech/ffi"
)

// Assert aborts the program if val is zero.
func Assert[T comparable](val T) {
	var zero T
	if val == zero {
		debug.PrintStack()
		Program.Abort()
	}
}

// FloatRoundingMode returns the current rounding mode.
// -1	the default rounding direction is not known
// 0	toward zero; same meaning as FE_TOWARDZERO
// 1	to nearest; same meaning as FE_TONEAREST
// 2	towards positive infinity; same meaning as FE_UPWARD
// 3	towards negative infinity; same meaning as FE_DOWNWARD
// other values	implementation-defined behavior
func FloatRoundingMode() abi.Int {
	switch FloatingPoint.GetRoundingMode() {
	case abi.FloatRoundTowardZero:
		return 0
	case abi.FloatRoundToNearest:
		return 1
	case abi.FloatRoundUpward:
		return 2
	case abi.FloatRoundDownward:
		return 3
	default:
		return -1
	}
}

var Char struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

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
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	ClearExceptions   func(abi.FloatException) abi.Error                                `cgo:"feclearexcept"`
	Exceptions        func(abi.FloatException) abi.FloatException                       `cgo:"fetestexcept"`
	RaiseExceptions   func(abi.FloatException) abi.Error                                `cgo:"feraiseexcept"`
	GetExceptionFlag  func(*abi.FloatingPointEnvironment, abi.FloatException) abi.Error `cgo:"fgetexceptflag"`
	SetExceptionFlag  func(*abi.FloatingPointEnvironment, abi.FloatException) abi.Error `cgo:"fsetexceptflag"`
	SetRoundingMode   func(abi.FloatRoundingMode) abi.Error                             `cgo:"fesetround"`
	GetRoundingMode   func() abi.FloatRoundingMode                                      `cgo:"fegetround"`
	GetEnvironment    func(*abi.FloatingPointEnvironment) abi.Error                     `cgo:"fegetenv"`
	SetEnvironment    func(*abi.FloatingPointEnvironment) abi.Error                     `cgo:"fesetenv"`
	UpdateEnvironment func(*abi.FloatingPointEnvironment) abi.Error                     `cgo:"feupdateenv"`
	HoldExceptions    func(*abi.FloatingPointEnvironment) abi.Error                     `cgo:"feholdexcept"`
}

var Locale struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Set func(abi.LocaleCategory, *abi.Locale) abi.String `cgo:"setlocale"`
	Get func() *abi.Locale                               `cgo:"localeconv"`
}

var Program struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Abort              func()                             `cgo:"abort"`
	Exit               func(abi.Int)                      `cgo:"exit"`
	ExitFast           func(abi.Int)                      `cgo:"quick_exit"`
	ExitWithoutCleanup func(abi.Int)                      `cgo:"_Exit"`
	OnExit             func(func())                       `cgo:"atexit"`
	OnExitFast         func(func())                       `cgo:"at_quick_exit"`
	LongJump           func(abi.JumpBuffer, abi.Int)      `cgo:"longjmp"`
	OnSignal           func(abi.Signal, func(abi.Signal)) `cgo:"signal"`
	Raise              func(abi.Signal)                   `cgo:"raise"`
}

var Files struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Open          func(abi.String, abi.String) *abi.File                         `cgo:"fopen"`
	Reopen        func(abi.String, abi.String, *abi.File) *abi.File              `cgo:"freopen"`
	Flush         func(*abi.File) abi.Int                                        `cgo:"fflush"`
	SetBuffer     func(*abi.File, abi.Pointer) abi.Int                           `cgo:"setbuf"`
	SetBufferMode func(*abi.File, abi.Pointer, abi.BufferMode, abi.Size) abi.Int `cgo:"setvbuf"`
	SetCharWide   func(*abi.File, abi.Int) abi.Int                               `cgo:"fwide"`

	Read  func(abi.Pointer, abi.Size, abi.Size, *abi.File) abi.Int `cgo:"fread"`
	Write func(abi.Pointer, abi.Size, abi.Size, *abi.File) abi.Int `cgo:"fwrite"`

	GetChar   func(*abi.File) abi.Int                           `cgo:"fgetc"`
	GetString func(abi.Pointer, abi.Int, *abi.File) abi.Pointer `cgo:"fgets"`
	PutChar   func(abi.Int, *abi.File) abi.Int                  `cgo:"fputc"`
	PutString func(abi.String, *abi.File) abi.Int               `cgo:"fputs"`
	UngetChar func(abi.Int, *abi.File) abi.Int                  `cgo:"ungetc"`

	GetCharWide   func(*abi.File) abi.CharWide                            `cgo:"fgetwc"`
	GetStringWide func(abi.StringWide, abi.Int, *abi.File) abi.StringWide `cgo:"fgetws"`
	PutCharWide   func(abi.CharWide, *abi.File) abi.CharWide              `cgo:"fputwc"`
	PutStringWide func(abi.StringWide, *abi.File) abi.Int                 `cgo:"fputws"`
	UngetCharWide func(abi.CharWide, *abi.File) abi.CharWide              `cgo:"ungetwc"`

	Scanf      func(*abi.File, abi.String, ...abi.Pointer) abi.Int     `cgo:"fscanf_s"`
	Printf     func(*abi.File, abi.String, ...abi.Pointer) abi.Int     `cgo:"fprintf_s"`
	ScanWidef  func(*abi.File, abi.StringWide, ...abi.Pointer) abi.Int `cgo:"fwscanf_s"`
	PrintWidef func(*abi.File, abi.StringWide, ...abi.Pointer) abi.Int `cgo:"fwprintf_s"`

	Tell   func(*abi.File) abi.Long                        `cgo:"ftell"`
	GetPos func(*abi.File, *abi.FilePosition) abi.Int      `cgo:"fgetpos"`
	Seek   func(*abi.File, abi.Long, abi.SeekMode) abi.Int `cgo:"fseek"`
	SetPos func(*abi.File, *abi.FilePosition) abi.Int      `cgo:"fsetpos"`
	Rewind func(*abi.File)                                 `cgo:"rewind"`

	ClearErr func(*abi.File)         `cgo:"clearerr"`
	IsEOF    func(*abi.File) abi.Int `cgo:"feof"`
	IsErr    func(*abi.File) abi.Int `cgo:"ferror"`
	Error    func(*abi.String)       `cgo:"perror"`

	Remove   func(abi.String) abi.Int             `cgo:"remove"`
	Rename   func(abi.String, abi.String) abi.Int `cgo:"rename"`
	Temp     func() *abi.File                     `cgo:"tmpfile"`
	TempName func(abi.String) abi.String          `cgo:"tmpnam"`
}

var IO struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	GetChar   func() abi.Int                          `cgo:"getchar"`
	GetString func(abi.Pointer, abi.Size) abi.Pointer `cgo:"gets_s"`
	PutChar   func(abi.Int) abi.Int                   `cgo:"putchar"`
	PutString func(abi.String) abi.Int                `cgo:"puts"`

	GetCharWide func() abi.CharWide             `cgo:"getwchar"`
	PutCharWide func(abi.CharWide) abi.CharWide `cgo:"putwchar"`

	Scanf      func(abi.String, ...abi.Pointer) abi.Int     `cgo:"scanf_s"`
	Printf     func(abi.String, ...abi.Pointer) abi.Int     `cgo:"printf_s"`
	ScanWidef  func(abi.StringWide, ...abi.Pointer) abi.Int `cgo:"wscanf_s"`
	PrintWidef func(abi.StringWide, ...abi.Pointer) abi.Int `cgo:"wprintf_s"`
}

var String struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Error func(abi.Error) abi.String `cgo:"strerror"`

	Scanf      func(abi.String, abi.String, ...abi.Pointer) abi.Int         `cgo:"sscanf_s"`
	Printf     func(abi.String, abi.String, ...abi.Pointer) abi.Int         `cgo:"sprintf_s"`
	ScanWidef  func(abi.StringWide, abi.StringWide, ...abi.Pointer) abi.Int `cgo:"swscanf_s"`
	PrintWidef func(abi.StringWide, abi.StringWide, ...abi.Pointer) abi.Int `cgo:"swprintf_s"`

	ToFloat               func(abi.String) abi.Float                                `cgo:"atof"`
	ToInt                 func(abi.String) abi.Int                                  `cgo:"atoi"`
	ToLong                func(abi.String) abi.Long                                 `cgo:"atol"`
	ToLongLong            func(abi.String) abi.LongLong                             `cgo:"atoll"`
	ParseLong             func(abi.String, *abi.Char, abi.Int) abi.Long             `cgo:"strtol"`
	ParseLongLong         func(abi.String, *abi.Char, abi.Int) abi.LongLong         `cgo:"strtoll"`
	ParseUnsignedLong     func(abi.String, *abi.Char, abi.Int) abi.LongUnsigned     `cgo:"strtoul"`
	ParseUnsignedLongLong func(abi.String, *abi.Char, abi.Int) abi.LongLongUnsigned `cgo:"strtoull"`
	ParseFloat            func(abi.String, *abi.Char) abi.Float                     `cgo:"strtof"`
	ParseDouble           func(abi.String, *abi.Char) abi.Double                    `cgo:"strtod"`
	ParseDoubleLong       func(abi.String, *abi.Char) abi.DoubleLong                `cgo:"strtold"`
	ParseIntmax           func(abi.String, *abi.Char, abi.Int) abi.IntMax           `cgo:"strtoimax"`
	ParseUintmax          func(abi.String, *abi.Char, abi.Int) abi.UIntMax          `cgo:"strtoumax"`

	Copy           func(abi.String, abi.Size, abi.String) abi.Error           `cgo:"strcpy_s"`
	CopyRange      func(abi.String, abi.Size, abi.String, abi.Size) abi.Error `cgo:"strncpy_s"`
	Append         func(abi.String, abi.Size, abi.String) abi.Error           `cgo:"strcat_s"`
	AppendRange    func(abi.String, abi.Size, abi.String, abi.Size) abi.Error `cgo:"strncat_s"`
	Localize       func(abi.String, abi.String, abi.Size) abi.Size            `cgo:"strxfrm"`
	Duplicate      func(abi.String) abi.String                                `cgo:"strdup"`
	DuplicateRange func(abi.String, abi.Size) abi.String                      `cgo:"strndup"`

	Length          func(abi.String, abi.Size) abi.Size                        `cgo:"strlen_s"`
	Compare         func(abi.String, abi.String) abi.Int                       `cgo:"strcmp"`
	CompareInLocale func(abi.String, abi.String) abi.Int                       `cgo:"strcoll"`
	FindFirst       func(abi.String, abi.Int) *abi.Char                        `cgo:"strchr"`
	FindLast        func(abi.String, abi.Int) *abi.Char                        `cgo:"strrchr"`
	MatchLength     func(abi.String, abi.String) abi.Size                      `cgo:"strspn"`
	Match           func(abi.String, abi.String) abi.Size                      `cgo:"strcspn"`
	MatchFirst      func(abi.String, abi.String) *abi.Char                     `cgo:"strpbrk"`
	Contains        func(abi.String, abi.String) *abi.Char                     `cgo:"strstr"`
	ScanToken       func(abi.String, abi.Size, abi.String, abi.Size) *abi.Char `cgo:"strtok_s"`
}

var Lib struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Calloc  func(abi.Size, abi.Size) abi.Pointer    `cgo:"calloc"`
	Free    func(abi.Pointer)                       `cgo:"free"`
	Malloc  func(abi.Size) abi.Pointer              `cgo:"malloc"`
	Realloc func(abi.Pointer, abi.Size) abi.Pointer `cgo:"realloc"`
	Abort   func()                                  `cgo:"abort"`
	AtExit  func(func()) abi.Error                  `cgo:"atexit"`
	Exit    func(abi.Int)                           `cgo:"exit"`
	Getenv  func(abi.String) abi.String             `cgo:"getenv"`
	System  func(abi.String) abi.Error              `cgo:"system"`

	BinarySearch func(abi.Pointer, abi.Pointer, abi.Size, abi.Size, func(abi.Pointer, abi.Pointer) abi.Int) abi.Pointer `cgo:"bsearch"`

	Sort func(abi.Pointer, abi.Size, abi.Size, func(abi.Pointer, abi.Pointer) abi.Int) abi.Pointer `cgo:"qsort"`

	MemoryCompare func(abi.Pointer, abi.Pointer, abi.Size) abi.Int               `cgo:"memcmp"`
	MemoryCopy    func(abi.Pointer, abi.Size, abi.Pointer, abi.Size) abi.Pointer `cgo:"memcpy_s"`
	MemoryMove    func(abi.Pointer, abi.Size, abi.Pointer, abi.Size) abi.Pointer `cgo:"memmove_s"`
	MemorySet     func(abi.Pointer, abi.Size, abi.Int, abi.Size) abi.Pointer     `cgo:"memset_s"`
	MemoryFind    func(abi.Pointer, abi.Int, abi.Size) abi.Pointer               `cgo:"memchr"`
}

var Time struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Diff          func(abi.Time, abi.Time) abi.Double       `cgo:"difftime"`
	Now           func(*abi.Time) abi.Time                  `cgo:"time"`
	Clock         func() abi.Clock                          `cgo:"clock"`
	Nanos         func(*abi.NanoTime, abi.TimeType)         `cgo:"timespec_get"`
	GetResolution func(*abi.NanoTime, abi.TimeType) abi.Int `cgo:"clock_getres"`

	DateString     func(abi.String, abi.Size, abi.String, *abi.Date) abi.Size         `cgo:"strftime"`
	DateStringWide func(abi.StringWide, abi.Size, abi.StringWide, *abi.Date) abi.Size `cgo:"wcsftime"`

	UTC   func(abi.Time) *abi.Date `cgo:"gmtime"`
	Local func(abi.Time) *abi.Date `cgo:"localtime"`
	Value func(*abi.Date) abi.Time `cgo:"mktime"`
}

type Div[T abi.Int | abi.Long | abi.LongLong | abi.IntMax] struct {
	Quo T
	Rem T
}
