package palimg

func Core(width int, height int, data ...uint8) PalettedImageCore {
	return internalCore{
		data:data,
		width:width,
		height:height,
	}
}

type internalCore struct {
	data []uint8
	width int
	height int
}

func (receiver internalCore) ColorIndexAt(x int, y int) uint8 {
	var length int = len(receiver.data)
	if length <= 0 {
		return 0
	}

	var arrayIndex int = calculateArrayIndex(receiver.width, receiver.height, x, y)
	if arrayIndex < 0 || length <= arrayIndex {
		arrayIndex = 0
	}

	return receiver.data[arrayIndex]
}

