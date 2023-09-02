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

	IsAlphaNumeric func(abi.Int) abi.Int `ffi:"isalnum"`
	IsAlpha        func(abi.Int) abi.Int `ffi:"isalpha"`
	IsUpper        func(abi.Int) abi.Int `ffi:"isupper"`
	IsLower        func(abi.Int) abi.Int `ffi:"islower"`
	IsDigit        func(abi.Int) abi.Int `ffi:"isdigit"`
	IsHexDigit     func(abi.Int) abi.Int `ffi:"isxdigit"`
	IsControl      func(abi.Int) abi.Int `ffi:"iscntrl"`
	IsGraph        func(abi.Int) abi.Int `ffi:"isgraph"`
	IsSpace        func(abi.Int) abi.Int `ffi:"isspace"`
	IsBlank        func(abi.Int) abi.Int `ffi:"isblank"`
	IsPrint        func(abi.Int) abi.Int `ffi:"isprint"`
	IsPuncuation   func(abi.Int) abi.Int `ffi:"ispunct"`

	ToLower func(abi.Int) abi.Int `ffi:"tolower"`
	ToUpper func(abi.Int) abi.Int `ffi:"toupper"`
}

var FloatingPoint struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	ClearExceptions   func(abi.FloatException) abi.Error                                `ffi:"feclearexcept"`
	Exceptions        func(abi.FloatException) abi.FloatException                       `ffi:"fetestexcept"`
	RaiseExceptions   func(abi.FloatException) abi.Error                                `ffi:"feraiseexcept"`
	GetExceptionFlag  func(*abi.FloatingPointEnvironment, abi.FloatException) abi.Error `ffi:"fgetexceptflag"`
	SetExceptionFlag  func(*abi.FloatingPointEnvironment, abi.FloatException) abi.Error `ffi:"fsetexceptflag"`
	SetRoundingMode   func(abi.FloatRoundingMode) abi.Error                             `ffi:"fesetround"`
	GetRoundingMode   func() abi.FloatRoundingMode                                      `ffi:"fegetround"`
	GetEnvironment    func(*abi.FloatingPointEnvironment) abi.Error                     `ffi:"fegetenv"`
	SetEnvironment    func(*abi.FloatingPointEnvironment) abi.Error                     `ffi:"fesetenv"`
	UpdateEnvironment func(*abi.FloatingPointEnvironment) abi.Error                     `ffi:"feupdateenv"`
	HoldExceptions    func(*abi.FloatingPointEnvironment) abi.Error                     `ffi:"feholdexcept"`
}

var Locale struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Set func(abi.LocaleCategory, *abi.Locale) abi.String `ffi:"setlocale"`
	Get func() *abi.Locale                               `ffi:"localeconv"`
}

var Program struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Abort              func()                             `ffi:"abort"`
	Exit               func(abi.Int)                      `ffi:"exit"`
	ExitFast           func(abi.Int)                      `ffi:"quick_exit"`
	ExitWithoutCleanup func(abi.Int)                      `ffi:"_Exit"`
	OnExit             func(func())                       `ffi:"atexit"`
	OnExitFast         func(func())                       `ffi:"at_quick_exit"`
	LongJump           func(abi.JumpBuffer, abi.Int)      `ffi:"longjmp"`
	OnSignal           func(abi.Signal, func(abi.Signal)) `ffi:"signal"`
	Raise              func(abi.Signal)                   `ffi:"raise"`
}

var Files struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Open          func(abi.String, abi.String) *abi.File                         `ffi:"fopen"`
	Reopen        func(abi.String, abi.String, *abi.File) *abi.File              `ffi:"freopen"`
	Flush         func(*abi.File) abi.Int                                        `ffi:"fflush"`
	SetBuffer     func(*abi.File, abi.Pointer) abi.Int                           `ffi:"setbuf"`
	SetBufferMode func(*abi.File, abi.Pointer, abi.BufferMode, abi.Size) abi.Int `ffi:"setvbuf"`
	SetCharWide   func(*abi.File, abi.Int) abi.Int                               `ffi:"fwide"`

	Read  func(abi.Pointer, abi.Size, abi.Size, *abi.File) abi.Int `ffi:"fread"`
	Write func(abi.Pointer, abi.Size, abi.Size, *abi.File) abi.Int `ffi:"fwrite"`

	GetChar   func(*abi.File) abi.Int                           `ffi:"fgetc"`
	GetString func(abi.Pointer, abi.Int, *abi.File) abi.Pointer `ffi:"fgets"`
	PutChar   func(abi.Int, *abi.File) abi.Int                  `ffi:"fputc"`
	PutString func(abi.String, *abi.File) abi.Int               `ffi:"fputs"`
	UngetChar func(abi.Int, *abi.File) abi.Int                  `ffi:"ungetc"`

	GetCharWide   func(*abi.File) abi.CharWide                            `ffi:"fgetwc"`
	GetStringWide func(abi.StringWide, abi.Int, *abi.File) abi.StringWide `ffi:"fgetws"`
	PutCharWide   func(abi.CharWide, *abi.File) abi.CharWide              `ffi:"fputwc"`
	PutStringWide func(abi.StringWide, *abi.File) abi.Int                 `ffi:"fputws"`
	UngetCharWide func(abi.CharWide, *abi.File) abi.CharWide              `ffi:"ungetwc"`

	Scanf      func(*abi.File, abi.String, ...abi.Pointer) abi.Int     `ffi:"fscanf_s"`
	Printf     func(*abi.File, abi.String, ...abi.Pointer) abi.Int     `ffi:"fprintf_s"`
	ScanWidef  func(*abi.File, abi.StringWide, ...abi.Pointer) abi.Int `ffi:"fwscanf_s"`
	PrintWidef func(*abi.File, abi.StringWide, ...abi.Pointer) abi.Int `ffi:"fwprintf_s"`

	Tell   func(*abi.File) abi.Long                        `ffi:"ftell"`
	GetPos func(*abi.File, *abi.FilePosition) abi.Int      `ffi:"fgetpos"`
	Seek   func(*abi.File, abi.Long, abi.SeekMode) abi.Int `ffi:"fseek"`
	SetPos func(*abi.File, *abi.FilePosition) abi.Int      `ffi:"fsetpos"`
	Rewind func(*abi.File)                                 `ffi:"rewind"`

	ClearErr func(*abi.File)         `ffi:"clearerr"`
	IsEOF    func(*abi.File) abi.Int `ffi:"feof"`
	IsErr    func(*abi.File) abi.Int `ffi:"ferror"`
	Error    func(*abi.String)       `ffi:"perror"`

	Remove   func(abi.String) abi.Int             `ffi:"remove"`
	Rename   func(abi.String, abi.String) abi.Int `ffi:"rename"`
	Temp     func() *abi.File                     `ffi:"tmpfile"`
	TempName func(abi.String) abi.String          `ffi:"tmpnam"`
}

var IO struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	GetChar   func() abi.Int                          `ffi:"getchar"`
	GetString func(abi.Pointer, abi.Size) abi.Pointer `ffi:"gets_s"`
	PutChar   func(abi.Int) abi.Int                   `ffi:"putchar"`
	PutString func(abi.String) abi.Int                `ffi:"puts"`

	GetCharWide func() abi.CharWide             `ffi:"getwchar"`
	PutCharWide func(abi.CharWide) abi.CharWide `ffi:"putwchar"`

	Scanf      func(abi.String, ...abi.Pointer) abi.Int     `ffi:"scanf_s"`
	Printf     func(abi.String, ...abi.Pointer) abi.Int     `ffi:"printf_s"`
	ScanWidef  func(abi.StringWide, ...abi.Pointer) abi.Int `ffi:"wscanf_s"`
	PrintWidef func(abi.StringWide, ...abi.Pointer) abi.Int `ffi:"wprintf_s"`
}

var String struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Error func(abi.Error) abi.String `ffi:"strerror"`

	Scanf      func(abi.String, abi.String, ...abi.Pointer) abi.Int         `ffi:"sscanf_s"`
	Printf     func(abi.String, abi.String, ...abi.Pointer) abi.Int         `ffi:"sprintf_s"`
	ScanWidef  func(abi.StringWide, abi.StringWide, ...abi.Pointer) abi.Int `ffi:"swscanf_s"`
	PrintWidef func(abi.StringWide, abi.StringWide, ...abi.Pointer) abi.Int `ffi:"swprintf_s"`

	ToFloat               func(abi.String) abi.Float                                `ffi:"atof"`
	ToInt                 func(abi.String) abi.Int                                  `ffi:"atoi"`
	ToLong                func(abi.String) abi.Long                                 `ffi:"atol"`
	ToLongLong            func(abi.String) abi.LongLong                             `ffi:"atoll"`
	ParseLong             func(abi.String, *abi.Char, abi.Int) abi.Long             `ffi:"strtol"`
	ParseLongLong         func(abi.String, *abi.Char, abi.Int) abi.LongLong         `ffi:"strtoll"`
	ParseUnsignedLong     func(abi.String, *abi.Char, abi.Int) abi.LongUnsigned     `ffi:"strtoul"`
	ParseUnsignedLongLong func(abi.String, *abi.Char, abi.Int) abi.LongLongUnsigned `ffi:"strtoull"`
	ParseFloat            func(abi.String, *abi.Char) abi.Float                     `ffi:"strtof"`
	ParseDouble           func(abi.String, *abi.Char) abi.Double                    `ffi:"strtod"`
	ParseDoubleLong       func(abi.String, *abi.Char) abi.DoubleLong                `ffi:"strtold"`
	ParseIntmax           func(abi.String, *abi.Char, abi.Int) abi.IntMax           `ffi:"strtoimax"`
	ParseUintmax          func(abi.String, *abi.Char, abi.Int) abi.UIntMax          `ffi:"strtoumax"`

	Copy           func(abi.String, abi.Size, abi.String) abi.Error           `ffi:"strcpy_s"`
	CopyRange      func(abi.String, abi.Size, abi.String, abi.Size) abi.Error `ffi:"strncpy_s"`
	Append         func(abi.String, abi.Size, abi.String) abi.Error           `ffi:"strcat_s"`
	AppendRange    func(abi.String, abi.Size, abi.String, abi.Size) abi.Error `ffi:"strncat_s"`
	Localize       func(abi.String, abi.String, abi.Size) abi.Size            `ffi:"strxfrm"`
	Duplicate      func(abi.String) abi.String                                `ffi:"strdup"`
	DuplicateRange func(abi.String, abi.Size) abi.String                      `ffi:"strndup"`

	Length          func(abi.String, abi.Size) abi.Size                        `ffi:"strlen_s"`
	Compare         func(abi.String, abi.String) abi.Int                       `ffi:"strcmp"`
	CompareInLocale func(abi.String, abi.String) abi.Int                       `ffi:"strcoll"`
	FindFirst       func(abi.String, abi.Int) *abi.Char                        `ffi:"strchr"`
	FindLast        func(abi.String, abi.Int) *abi.Char                        `ffi:"strrchr"`
	MatchLength     func(abi.String, abi.String) abi.Size                      `ffi:"strspn"`
	Match           func(abi.String, abi.String) abi.Size                      `ffi:"strcspn"`
	MatchFirst      func(abi.String, abi.String) *abi.Char                     `ffi:"strpbrk"`
	Contains        func(abi.String, abi.String) *abi.Char                     `ffi:"strstr"`
	ScanToken       func(abi.String, abi.Size, abi.String, abi.Size) *abi.Char `ffi:"strtok_s"`
}

var Lib struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Calloc  func(abi.Size, abi.Size) abi.Pointer    `ffi:"calloc"`
	Free    func(abi.Pointer)                       `ffi:"free"`
	Malloc  func(abi.Size) abi.Pointer              `ffi:"malloc"`
	Realloc func(abi.Pointer, abi.Size) abi.Pointer `ffi:"realloc"`
	Abort   func()                                  `ffi:"abort"`
	AtExit  func(func()) abi.Error                  `ffi:"atexit"`
	Exit    func(abi.Int)                           `ffi:"exit"`
	Getenv  func(abi.String) abi.String             `ffi:"getenv"`
	System  func(abi.String) abi.Error              `ffi:"system"`

	BinarySearch func(abi.Pointer, abi.Pointer, abi.Size, abi.Size, func(abi.Pointer, abi.Pointer) abi.Int) abi.Pointer `ffi:"bsearch"`

	Sort func(abi.Pointer, abi.Size, abi.Size, func(abi.Pointer, abi.Pointer) abi.Int) abi.Pointer `ffi:"qsort"`

	MemoryCompare func(abi.Pointer, abi.Pointer, abi.Size) abi.Int               `ffi:"memcmp"`
	MemoryCopy    func(abi.Pointer, abi.Size, abi.Pointer, abi.Size) abi.Pointer `ffi:"memcpy_s"`
	MemoryMove    func(abi.Pointer, abi.Size, abi.Pointer, abi.Size) abi.Pointer `ffi:"memmove_s"`
	MemorySet     func(abi.Pointer, abi.Size, abi.Int, abi.Size) abi.Pointer     `ffi:"memset_s"`
	MemoryFind    func(abi.Pointer, abi.Int, abi.Size) abi.Pointer               `ffi:"memchr"`
}

var Time struct {
	ffi.Header `linux:"libc.so.6" darwin:"libSystem.dylib"`

	Diff          func(abi.Time, abi.Time) abi.Double       `ffi:"difftime"`
	Now           func(*abi.Time) abi.Time                  `ffi:"time"`
	Clock         func() abi.Clock                          `ffi:"clock"`
	Nanos         func(*abi.NanoTime, abi.TimeType)         `ffi:"timespec_get"`
	GetResolution func(*abi.NanoTime, abi.TimeType) abi.Int `ffi:"clock_getres"`

	DateString     func(abi.String, abi.Size, abi.String, *abi.Date) abi.Size         `ffi:"strftime"`
	DateStringWide func(abi.StringWide, abi.Size, abi.StringWide, *abi.Date) abi.Size `ffi:"wcsftime"`

	UTC   func(abi.Time) *abi.Date `ffi:"gmtime"`
	Local func(abi.Time) *abi.Date `ffi:"localtime"`
	Value func(*abi.Date) abi.Time `ffi:"mktime"`
}

type Div[T abi.Int | abi.Long | abi.LongLong | abi.IntMax] struct {
	Quo T
	Rem T
}
