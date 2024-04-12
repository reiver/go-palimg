package palimg

import (
	"image"
	"image/color"
)

type internalPalettedImage struct {
	palimgcore PalettedImageCore
	rectangle image.Rectangle
	palette []color.Color
}

var _ image.PalettedImage = internalPalettedImage{}

func (receiver internalPalettedImage) At(x, y int) color.Color {
	if len(receiver.palette) <= 0 {
		return nil
	}

	var colorindex uint8 = receiver.ColorIndexAt(x, y)
	var ci int = int(colorindex)

	if ci < 0 || len(receiver.palette) <= ci {
		return nil
	}

	return receiver.palette[colorindex]
}


func (receiver internalPalettedImage) Bounds() image.Rectangle {
	return receiver.rectangle
}

func (receiver internalPalettedImage) ColorIndexAt(x, y int) uint8 {
	var point image.Point = image.Pt(x,y)

	if !point.In(receiver.rectangle) {
		return 0
	}

	var palimgcore PalettedImageCore = receiver.palimgcore
	if nil == palimgcore {
		return 0
	}

	return palimgcore.ColorIndexAt(x,y)
}

func (receiver internalPalettedImage) ColorModel() color.Model {
	return color.Palette(receiver.palette)
}
