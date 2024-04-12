package palimg

// The PalettedImageCoreFunc type is an adapter that turns a func of type func(int,int)uint8 into a palimg.PalettedImageCore.
type PalettedImageCoreFunc func(x int, y int)uint8

var _ PalettedImageCore = PalettedImageCoreFunc(nil)

func (receiver PalettedImageCoreFunc) ColorIndexAt(x, y int) uint8 {
	return receiver(x,y)
}

