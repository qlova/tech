package std

import (
	"unsafe"

	"qlova.tech/abi"
	"qlova.tech/ffi"
)

var Char struct {
	ffi.Header `linux:"libc.so.6"`

	IsAlpha        func(rune) bool `cgo:"isalpha"`
	IsAlphaNumeric func(rune) bool `cgo:"isalnum"`
	IsControl      func(rune) bool `cgo:"iscntrl"`
	IsDigit        func(rune) bool `cgo:"isdigit"`
	IsGraph        func(rune) bool `cgo:"isgraph"`
	IsLower        func(rune) bool `cgo:"islower"`
	IsPrint        func(rune) bool `cgo:"isprint"`
	IsPuncuation   func(rune) bool `cgo:"ispunct"`
	IsSpace        func(rune) bool `cgo:"isspace"`
	IsUpper        func(rune) bool `cgo:"isupper"`
	IsHexDigit     func(rune) bool `cgo:"isxdigit"`

	ToLower func(rune) rune `cgo:"tolower"`
	ToUpper func(rune) rune `cgo:"toupper"`
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

type Error int32

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
