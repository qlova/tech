# Binary Object X 

Binary Object X is a binary format specification for cross-language communication. 
It's an adjustable & self-describing protocol format with an optional schema.

A BOXed message looks like this:

```
    [ARCHITECTURE]([SCHEMA])[HEADER][DATA][MEMORY]
```

### ARCHITECTURE
Is a single byte that describes the architecture of the message.

1. Is 64 bit? (else 32 bit)
2. Is Big Endian? (else little endian)
3. Fat strings? (else C strings)
4. Reserved
5. Reserved
6. Reserved
7. Reserved
8. Reserved

### SCHEMA
The schema is an optional component of a message and contains
additional information about the message, such as field names
as chosen by the encoder.

### HEADER
The header is a series of bytes that describe the layout of 
the message. Usually, the first three bits is the size and
kind of a field and the last five bits are a field number.

```
    000     00000
    ^size   ^field
```

The header ends with a null byte of zeros.

### DATA
This contains the data of the message and
should be loaded into the fields as described by the header.
Unknown fields are ignored.

### Memory
If DATA contains pointers, then the dereferenced data of
those pointers are stored in this section of the message.
MEMORY is structured like a sub message, with it's own 
header and each subsequent pointer as an incrementing and
implicit field identifier. The pointer in the DATA section
should point to the field value it has been written to in
the memory section.
