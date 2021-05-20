//Package box provides a marshalling and unmarshaling of binary object explanations.
package box

const (
	Message = iota

	Comment
	Closing

	Padding

	Decrypt
	Deflate
	Integer

	Version
	Timeout
	Passkey
	Testing
	Tracker

	Creator
	Created
	ModTime

	Failure
)

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
