package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawBackground(screen *ebiten.Image) {
	game.drawSky(screen)
	game.drawFloor(screen)
}

func (game *Game) drawSky(screen *ebiten.Image) {
	halfHeight := game.ScreenHeight / 2
	if game.SkyTexture == nil {
		return
	}

	skyWidth := game.SkyTexture.Bounds().Dx()
	skyHeight := game.SkyTexture.Bounds().Dy()
	if skyWidth == 0 || skyHeight == 0 {
		return
	}

	horizontalOffset := int((game.Player.Angle / (2 * math.Pi)) * float64(skyWidth))

	for y := 0; y < halfHeight; y++ {
		textureY := y * skyHeight / halfHeight
		for x := 0; x < game.ScreenWidth; x++ {
			textureX := (x + horizontalOffset) % skyWidth
			if textureX < 0 {
				textureX += skyWidth
			}
			screen.Set(x, y, game.SkyTexture.At(textureX, textureY))
		}
	}
}

func (game *Game) drawFloor(screen *ebiten.Image) {
	halfHeight := game.ScreenHeight / 2
	if game.FloorTexture == nil {
		return
	}

	floorWidth := game.FloorTexture.Bounds().Dx()
	floorHeight := game.FloorTexture.Bounds().Dy()
	if floorWidth == 0 || floorHeight == 0 {
		return
	}

	leftRayAngle := game.Player.Angle - fieldOfView/2
	rightRayAngle := game.Player.Angle + fieldOfView/2
	rayDirX0 := math.Cos(leftRayAngle)
	rayDirY0 := math.Sin(leftRayAngle)
	rayDirX1 := math.Cos(rightRayAngle)
	rayDirY1 := math.Sin(rightRayAngle)
	posZ := 0.5 * float64(game.ScreenHeight)

	for y := halfHeight; y < game.ScreenHeight; y++ {
		rowOffset := y - halfHeight
		if rowOffset == 0 {
			continue
		}

		rowDistance := posZ / float64(rowOffset)
		stepX := rowDistance * (rayDirX1 - rayDirX0) / float64(game.ScreenWidth)
		stepY := rowDistance * (rayDirY1 - rayDirY0) / float64(game.ScreenWidth)

		worldX := game.Player.X + rowDistance*rayDirX0
		worldY := game.Player.Y + rowDistance*rayDirY0

		for x := 0; x < game.ScreenWidth; x++ {
			textureX := floorTextureCoordinate(worldX, floorWidth)
			textureY := floorTextureCoordinate(worldY, floorHeight)
			screen.Set(x, y, game.FloorTexture.At(textureX, textureY))

			worldX += stepX
			worldY += stepY
		}
	}
}

func floorTextureCoordinate(worldValue float64, textureSize int) int {
	frac := worldValue - math.Floor(worldValue)
	coord := int(frac * float64(textureSize))
	if coord < 0 {
		return 0
	}
	if coord >= textureSize {
		return textureSize - 1
	}

	return coord
}
