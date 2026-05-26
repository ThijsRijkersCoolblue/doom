package assets

import (
	"fmt"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadWallTextures(textureFiles map[int]string) (map[int]*ebiten.Image, error) {
	textures := make(map[int]*ebiten.Image, len(textureFiles))

	for tileID, path := range textureFiles {
		texture, err := loadTexture(path)
		if err != nil {
			return nil, fmt.Errorf("load texture for tile %d: %w", tileID, err)
		}

		textures[tileID] = texture
	}

	return textures, nil
}

func loadTexture(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decodedImage, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(decodedImage), nil
}
