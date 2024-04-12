package palimg

// The PalImgFunc type is an adapter that turns a func of type func(int,int)uint8 into a palimg.PalImg.
type PalImgFunc func(x int, y int)uint8

var _ PalImg = PalImgFunc(nil)

func (receiver PalImgFunc) ColorIndexAt(x, y int) uint8 {
	return receiver(x,y)
}

