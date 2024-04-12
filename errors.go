package palimg

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilReceiver     = erorr.Error("palimg: nil receiver")
	errPaletteOverflow = erorr.Error("palimg: palette overflow — color palette can only contain a maximum of 256 colors")
)
