package images

import (
	_ "embed"
)

var (
	//go:embed grass.png
	Grass_png []byte

	//go:embed character.png
	Character_png []byte
)
