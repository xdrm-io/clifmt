package color

import (
	"fmt"
	"io"
	"strconv"
)

// NameError raised when a color name or hexa is invalid
type NameError struct {
	Err  error
	Name string
}

func (err *NameError) Error() string {
	return fmt.Sprintf("%s '%s'", err.Err, err.Name)
}

// NameError errors
var (
	ErrInvalidHexSize   = fmt.Errorf("expect a size of 3 or 6 (without the '#' prefix)")
	ErrUnknownColorName = fmt.Errorf("unknown color name")
)

// T represents a color
type T uint32

// FromName returns the integer value of a color name
// from the built-in color map ; it is case insensitive
func FromName(t Theme, s string) (T, error) {
	value, ok := t[s]
	if !ok {
		return 0, &NameError{ErrUnknownColorName, s}
	}
	return value, nil
}

// FromHex returns the integer value associated with
// an hexadecimal string (full-sized or short version)
// the format is 'abc' or 'abcdef'
func FromHex(s string) (T, error) {
	if len(s) != 3 && len(s) != 6 {
		return 0, ErrInvalidHexSize
	}

	// short version
	input := s
	if len(s) == 3 {
		input = fmt.Sprintf("%c%c%c%c%c%c", s[0], s[0], s[1], s[1], s[2], s[2])
	}

	n, err := strconv.ParseUint(input, 16, 32)
	if err != nil {
		return 0, err
	}

	return T(n), nil
}

// Parse tries to parse a color string (can be a name or an hexa value)
func Parse(t Theme, s string) (T, error) {
	if len(s) < 1 {
		return 0, io.ErrUnexpectedEOF
	}
	if s[0] == '#' {
		return FromHex(s[1:])
	}
	return FromName(t, s)
}

// Red component of the color
func (c T) Red() uint8 {
	return uint8((c >> 16) & 0xff)
}

// Green component of the color
func (c T) Green() uint8 {
	return uint8((c >> 8) & 0xff)
}

// Blue component of the color
func (c T) Blue() uint8 {
	return uint8(c & 0xff)
}
