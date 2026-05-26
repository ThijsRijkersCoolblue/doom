package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawWeapon(screen *ebiten.Image) {
	if len(game.Weapon.Frames) == 0 {
		return
	}

	frame := game.Weapon.Frames[game.Weapon.CurrentFrame]
	if frame == nil {
		return
	}

	weaponHeight := int(math.Max(72, float64(game.ScreenHeight)*0.28))
	weaponWidth := int(float64(weaponHeight) * float64(frame.Bounds().Dx()) / float64(frame.Bounds().Dy()))

	x := game.ScreenWidth/2 - weaponWidth/2
	hudRect := game.statusBarRect()
	y := hudRect.y - weaponHeight + int(float64(hudRect.h)*0.08)

	drawScaled(screen, frame, uiRect{x: x, y: y, w: weaponWidth, h: weaponHeight})
}
