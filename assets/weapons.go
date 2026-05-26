package assets

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadWeaponFrames(framePaths []string) ([]*ebiten.Image, error) {
	frames := make([]*ebiten.Image, 0, len(framePaths))

	for _, path := range framePaths {
		frame, err := LoadTexture(path)
		if err != nil {
			return nil, fmt.Errorf("load weapon frame %s: %w", path, err)
		}

		frames = append(frames, frame)
	}

	return frames, nil
}

func LoadPistolFrames() ([]*ebiten.Image, error) {
	return LoadWeaponFrames([]string{
		"assets/weapons/PISGA0.png",
		"assets/weapons/PISGB0.png",
		"assets/weapons/PISGC0.png",
		"assets/weapons/PISGD0.png",
		"assets/weapons/PISGE0.png",
	})
}
