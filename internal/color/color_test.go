package color

import (
	"strconv"
	"testing"
)

func TestFromHexSizeError(t *testing.T) {

	tests := []struct {
		Hex string
		err error
	}{

		{"", ErrInvalidHexSize},
		{"f", ErrInvalidHexSize},
		{"f0", ErrInvalidHexSize},
		{"f00", nil}, // 3 chars
		{"ff00", ErrInvalidHexSize},
		{"ff000", ErrInvalidHexSize},
		{"ff0000", nil}, // 6 chars
		{"ff00000", ErrInvalidHexSize},
	}

	for i, test := range tests {

		_, err := FromHex(test.Hex)
		if err != nil {
			if test.err == nil {
				t.Errorf("[%d] unexpected error <%s>", i, err)
			} else if err != test.err {
				t.Errorf("[%d] got error <%s> expected <%s>", i, err, test.err)
			}
			break
		}

		if test.err != nil {
			t.Errorf("[%d] expected error <%s>", i, test.err)
		}

	}

}
func TestFromHexCharsetError(t *testing.T) {

	tests := []struct {
		Hex         string
		validSyntax bool
	}{

		{"012345", true},
		{"789abc", true},
		{"abcdef", true},
		{"aBcDeF", true},
		{"bcdefg", false},
	}

	for i, test := range tests {

		_, err := FromHex(test.Hex)
		if err != nil {
			if test.validSyntax {
				t.Errorf("[%d] unexpected error <%s>", i, err)
				break
			}

			if nerr, ok := err.(*strconv.NumError); ok {

				if nerr.Err != strconv.ErrSyntax {
					t.Errorf("[%d] got unexpected error <%s> expected <%s>", i, nerr.Err, strconv.ErrSyntax)
				}
				break

			}

			t.Errorf("[%d] got an unknown error <%s> expected <%s>", i, err, "strconv.NumError")
			break
		}

		if !test.validSyntax {
			t.Errorf("[%d] expected error", i)
		}

	}

}
func TestFromHexRGB(t *testing.T) {

	tests := []struct {
		Hex string
		R   uint8
		G   uint8
		B   uint8
	}{
		{"000", 0, 0, 0},
		{"000000", 0, 0, 0},

		{"f00", 255, 0, 0},
		{"ff0000", 255, 0, 0},

		{"0f0", 0, 255, 0},
		{"00ff00", 0, 255, 0},

		{"00f", 0, 0, 255},
		{"0000ff", 0, 0, 255},

		{"4cb5ae", 76, 181, 174},
	}

	for i, test := range tests {

		color, err := FromHex(test.Hex)
		if err != nil {
			t.Errorf("[%d] unexpected error <%s>", i, err)
			break
		}

		if color.Red() != test.R || color.Green() != test.G || color.Blue() != test.B {
			t.Errorf("[%d] invalid color rgb(%d,%d,%d) expected rgb(%d,%d,%d)", i, color.Red(), color.Green(), color.Blue(), test.R, test.G, test.B)
		}

	}

}

func getTheme() Theme {
	t := make(Theme)
	t["white"] = 0xffffff
	t["black"] = 0x0
	t["red"] = 0xff0000
	t["green"] = 0x00ff00
	t["blue"] = 0x0000ff
	t["CaSeSeNsItIvE"] = 0x4cb5ae

	return t
}

func TestFromName(t *testing.T) {
	// fetch custom theme
	theme := getTheme()

	tests := []struct {
		Name  string
		Fail  bool
		Color T
	}{
		{"white", false, 0xffffff},
		{"black", false, 0x0},
		{"red", false, 0xff0000},
		{"green", false, 0x00ff00},
		{"blue", false, 0x0000ff},

		{"CaSeSeNsItIvE", false, 0x4cb5ae},
		{"casesensitive", true, 0x4cb5ae},
	}

	for i, test := range tests {

		color, err := FromName(theme, test.Name)
		if err != nil {
			if !test.Fail {
				t.Errorf("[%d] unexpected error <%s>", i, err)
			}
			break
		}

		if color != test.Color {
			t.Errorf("[%d] invalid color %x expected %x", i, color, test.Color)
		}

	}
}
