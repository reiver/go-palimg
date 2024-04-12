package palimg

// If an 2-dimensional image is stored as 1-dimensional slice, then calculateArrayIndex converts
// 2-dimensional (x,y) coordinates into a 1-dimension array index, for an image that is width×height
// in size, with (0,0) being in the top left corner.
func calculateArrayIndex(width int, height int, x int, y int) int {
	if x < 0 || width <= x {
		return 0
	}
	if y < 0 || height <= y {
		return 0
	}

	return (y-1)*width + x
}

