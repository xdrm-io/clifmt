package color

// Theme is used to associate color names to integer color values ;
// the color names in the theme are available in the colorizing format
type Theme map[string]T

// DefaultTheme sets the default theme associations
func DefaultTheme() Theme {
	theme := make(Theme)
	theme["black"] = 0x000000
	theme["white"] = 0xffffff
	theme["red"] = 0xff0000
	theme["green"] = 0x00ff00
	theme["blue"] = 0x0000ff
	theme["yellow"] = 0xffff00
	theme["orange"] = 0xff8c00
	theme["purple"] = 0x800080
	theme["navy"] = 0x000080
	theme["aqua"] = 0x00ffff
	theme["gray"] = 0x808080
	theme["silver"] = 0xc0c0c0
	theme["fuchsia"] = 0xff00ff
	theme["olive"] = 0x808000
	theme["teal"] = 0x008080
	theme["brown"] = 0x800000

	return theme
}
