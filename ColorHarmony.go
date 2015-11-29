package main

import (
	"fmt"
	"image/color"
)

// Return a RandomRGB
func RandomRGB() color.RGBA {
	return color.RGBA{
		uint8(randInt(0, 255)),
		uint8(randInt(0, 255)),
		uint8(randInt(0, 255)),
		255}
}

func Hex(c color.RGBA) string {
	r := pad(intToHex(c.R), 2)
	g := pad(intToHex(c.G), 2)
	b := pad(intToHex(c.B), 2)
	return fmt.Sprintf("#%s%s%s", r, g, b)

}

// Return the complementary color
func Complementary(v color.RGBA) [2]color.RGBA {
	var result color.RGBA
	result.R = 255 - v.R
	result.G = 255 - v.G
	result.B = 255 - v.B
	result.A = v.A
	return [2]color.RGBA{v, result}
}

// Return a monochromatic color set
func Monochromatic(v color.RGBA) color.RGBA {
	var result color.RGBA
	return result
}

// Return an analogous color set
func Analogous(v color.RGBA) color.RGBA {
	var result color.RGBA
	return result
}

// Return a split complementary color set
func SplitComplementary(v color.RGBA) color.RGBA {
	var result color.RGBA
	return result
}

// Return a triadic color set
func Triadic(v color.RGBA) color.RGBA {
	var result color.RGBA
	return result
}

// Return a triadic color set
func Tetradic(v color.RGBA) color.RGBA {
	var result color.RGBA
	return result
}
