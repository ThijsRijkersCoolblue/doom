package game

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawEnemies(screen *ebiten.Image) {
	if len(game.Enemies) == 0 {
		return
	}

	for _, enemy := range game.Enemies {
		if enemy.Sprite == nil {
			continue
		}

		dx := enemy.X - game.Player.X
		dy := enemy.Y - game.Player.Y
		distance := math.Hypot(dx, dy)
		if distance < minWallDistance {
			continue
		}

		angle := normalizeAngle(math.Atan2(dy, dx) - game.Player.Angle)
		if math.Abs(angle) > fieldOfView/2 {
			continue
		}

		projectedSize := int(float64(game.ScreenHeight) / distance)
		if projectedSize <= 0 {
			continue
		}

		screenCenter := game.ScreenWidth / 2
		screenX := int((angle/(fieldOfView/2))*float64(screenCenter) + float64(screenCenter))
		left := screenX - projectedSize/2
		top := (game.ScreenHeight - projectedSize) / 2

		game.drawEnemySprite(screen, enemy.Sprite, left, top, projectedSize, distance)
	}
}

func (game *Game) drawEnemySprite(screen, sprite *ebiten.Image, left, top, size int, distance float64) {
	textureWidth := sprite.Bounds().Dx()
	textureHeight := sprite.Bounds().Dy()
	if textureWidth == 0 || textureHeight == 0 {
		return
	}

	for spriteX := 0; spriteX < size; spriteX++ {
		screenX := left + spriteX
		if screenX < 0 || screenX >= game.ScreenWidth {
			continue
		}
		if distance >= game.DepthBuffer[screenX] {
			continue
		}

		textureX := spriteX * textureWidth / size
		for spriteY := 0; spriteY < size; spriteY++ {
			screenY := top + spriteY
			if screenY < 0 || screenY >= game.ScreenHeight {
				continue
			}

			textureY := spriteY * textureHeight / size
			pixel := sprite.At(textureX, textureY)
			if isTransparent(pixel) {
				continue
			}
			screen.Set(screenX, screenY, pixel)
		}
	}
}

func isTransparent(pixel color.Color) bool {
	_, _, _, alpha := pixel.RGBA()
	return alpha == 0
}

func normalizeAngle(angle float64) float64 {
	for angle > math.Pi {
		angle -= 2 * math.Pi
	}
	for angle < -math.Pi {
		angle += 2 * math.Pi
	}

	return angle
}
