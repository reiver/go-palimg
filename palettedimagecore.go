package palimg

// PalettedImageCore is a paletted-image without the palette.
type PalettedImageCore interface {
	// ColorIndexAt returns the palette index of the pixel at (x,y), where (0,0) is in the top-left corner of the image.
	ColorIndexAt(x int, y int) uint8
}
