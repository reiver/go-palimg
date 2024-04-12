package palimg

import (
	"image"
	"image/color"
)

// NRGBA turns a palimg.PalettedImageCore into a image.PalettedImage by applying the provided palette.
func NRGBA(palimgcore PalettedImageCore, rectangle image.Rectangle, palette ...color.NRGBA) image.PalettedImage {
	var colors []color.Color
	for _, c := range palette {
		colors = append(colors, c)
	}

	return internalPalettedImage{
		palimgcore:palimgcore,
		rectangle:rectangle.Canon(),
		palette:colors,
	}
}
