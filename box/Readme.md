# Binary Object X 

Binary Object X is a binary format specification for cross-language communication. 
It's an adjustable & self-describing protocol format with an optional schema.

A BOXed message looks like this:

```
    [VERSION]([SCHEMA])[HEADER][NULL][MESSAGE][MEMORY]
```

### VERSION
Version of the serialisation format.

### SCHEMA
The schema is an optional component of a message and contains
additional information about the message, such as field names
as chosen by the encoder.

### HEADER
The header is a series of bytes that describe the layout of 
the message. Usually, the first four bits is the size and
kind of a field and the last four bits are a field number.
If the field number is greater than 15 than the field number
is left as zero and then the next byte determines the field number. If this byte is also zero, then the next two bytes 
determine the field number and so on.


```
    0000    0000
    ^size   ^field
```

```go
const (
	Command = iota

	BitsN
	Bits8
	Bits16
	Bits32
	Bits64
	Bits96
	Bits128
	Bits192
	Bits256

	Opening
	Pointer
	Exactly
	Mapping
	Dynamic
	Boolean
)

```

### BODY
This contains the data of the message and
should be loaded into the fields as described by the header.
Unknown fields are ignored. If any DATA contains pointers, then the dereferenced data of those pointers are also stored in this section of the message.
