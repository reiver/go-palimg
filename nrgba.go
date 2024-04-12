package palimg

import (
	"image"
	"image/color"
)

// NRGBA turns a palimg.PalImg into a image.PalttedImage by applying the provided palette.
func NRGBA(palimg PalImg, rectangle image.Rectangle, palette ...color.NRGBA) image.PalettedImage {
	var colors []color.Color
	for _, c := range palette {
		colors = append(colors, c)
	}

	return internalPalettedImage{
		palimg:palimg,
		rectangle:rectangle.Canon(),
		palette:colors,
	}
}
