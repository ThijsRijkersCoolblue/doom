package game

import (
	"image/color"
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

	game.ensureRayDirectionCache()
	game.ensureSectorCache()

	sampleStep := 2
	projectionScale := float64(game.ScreenHeight) * 0.5
	screenCenter := float64(game.ScreenHeight) * 0.5

	for y := halfHeight + 1; y < game.ScreenHeight; y += sampleStep {
		denominator := float64(y) - screenCenter
		if denominator <= 0 {
			continue
		}

		for x := 0; x < game.ScreenWidth; x += sampleStep {
			worldX, worldY, ok := game.floorWorldPointForPixel(
				game.RayDirXCache[x],
				game.RayDirYCache[x],
				denominator,
				projectionScale,
			)
			if !ok {
				continue
			}

			textureX := floorTextureCoordinate(worldX, floorWidth)
			textureY := floorTextureCoordinate(worldY, floorHeight)
			pixel := game.FloorTexture.At(textureX, textureY)

			drawFloorBlock(screen, x, y, sampleStep, game.ScreenWidth, game.ScreenHeight, pixel)
		}
	}
}

func (game *Game) floorWorldPointForPixel(rayDirX, rayDirY, denominator, projectionScale float64) (float64, float64, bool) {
	floorHeight := game.currentSectorFloor()

	for i := 0; i < 2; i++ {
		heightDelta := game.Player.Z - floorHeight
		if heightDelta <= minWallDistance {
			return 0, 0, false
		}

		distance := (heightDelta * projectionScale) / denominator
		if distance <= minWallDistance {
			return 0, 0, false
		}

		worldX := game.Player.X + rayDirX*distance
		worldY := game.Player.Y + rayDirY*distance

		sectorID := game.findSectorID(worldX, worldY)
		if sectorID < 0 || sectorID >= len(game.Sectors) {
			return worldX, worldY, true
		}

		nextFloor := game.Sectors[sectorID].FloorHeight
		if math.Abs(nextFloor-floorHeight) < minWallDistance {
			return worldX, worldY, true
		}

		floorHeight = nextFloor
	}

	heightDelta := game.Player.Z - floorHeight
	if heightDelta <= minWallDistance {
		return 0, 0, false
	}

	distance := (heightDelta * projectionScale) / denominator
	worldX := game.Player.X + rayDirX*distance
	worldY := game.Player.Y + rayDirY*distance
	return worldX, worldY, true
}

func drawFloorBlock(screen *ebiten.Image, x, y, step, width, height int, pixel color.Color) {
	for dy := 0; dy < step; dy++ {
		sy := y + dy
		if sy >= height {
			break
		}
		for dx := 0; dx < step; dx++ {
			sx := x + dx
			if sx >= width {
				break
			}
			screen.Set(sx, sy, pixel)
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
