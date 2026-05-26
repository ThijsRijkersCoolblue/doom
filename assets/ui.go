package assets

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type UIAssets struct {
	Bar      *ebiten.Image
	Arms     *ebiten.Image
	KeyBlue  *ebiten.Image
	KeyYel   *ebiten.Image
	KeyRed   *ebiten.Image
	FaceIdle *ebiten.Image
	FacePain *ebiten.Image
	FaceGod  *ebiten.Image
	FaceDead *ebiten.Image
}

func LoadUIAssets() (*UIAssets, error) {
	bar, err := LoadTexture("assets/ui/STBAR.png")
	if err != nil {
		return nil, fmt.Errorf("load status bar: %w", err)
	}

	arms, err := LoadTexture("assets/ui/STARMS.png")
	if err != nil {
		return nil, fmt.Errorf("load arms panel: %w", err)
	}

	keyBlue, err := LoadTexture("assets/ui/STKEYS0.png")
	if err != nil {
		return nil, fmt.Errorf("load blue key icon: %w", err)
	}

	keyYel, err := LoadTexture("assets/ui/STKEYS1.png")
	if err != nil {
		return nil, fmt.Errorf("load yellow key icon: %w", err)
	}

	keyRed, err := LoadTexture("assets/ui/STKEYS2.png")
	if err != nil {
		return nil, fmt.Errorf("load red key icon: %w", err)
	}

	faceIdle, err := LoadTexture("assets/ui/STFST00.png")
	if err != nil {
		return nil, fmt.Errorf("load idle face: %w", err)
	}

	facePain, err := LoadTexture("assets/ui/STFOUCH0.png")
	if err != nil {
		return nil, fmt.Errorf("load pain face: %w", err)
	}

	faceGod, err := LoadTexture("assets/ui/STFGOD0.png")
	if err != nil {
		return nil, fmt.Errorf("load god face: %w", err)
	}

	faceDead, err := LoadTexture("assets/ui/STFDEAD0.png")
	if err != nil {
		return nil, fmt.Errorf("load dead face: %w", err)
	}

	return &UIAssets{
		Bar:      bar,
		Arms:     arms,
		KeyBlue:  keyBlue,
		KeyYel:   keyYel,
		KeyRed:   keyRed,
		FaceIdle: faceIdle,
		FacePain: facePain,
		FaceGod:  faceGod,
		FaceDead: faceDead,
	}, nil
}
