package game

import (
	"doom/drawer"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawWalls(screen *ebiten.Image) {
	numberOfRays := game.ScreenWidth

	for screenX := 0; screenX < numberOfRays; screenX++ {
		rayAngle := game.rayAngleForColumn(screenX, numberOfRays)
		hit := game.Drawer.CastRay(&game.Player, rayAngle, game.WorldMap)

		if hit.TileID == 0 {
			game.DepthBuffer[screenX] = game.Drawer.MaxDist
			continue
		}

		wallDistance := game.correctedDistance(hit.Distance, rayAngle)
		game.DepthBuffer[screenX] = wallDistance

		wallHeight := int(float64(game.ScreenHeight) / wallDistance)
		texture := game.WallTextures[hit.TileID]
		drawer.DrawTexturedColumn(screen, screenX, wallHeight, game.ScreenHeight, texture, hit.WallX)
	}
}

func (game *Game) rayAngleForColumn(screenX, numberOfRays int) float64 {
	return game.Player.Angle - fieldOfView/2 + fieldOfView*float64(screenX)/float64(numberOfRays)
}

func (game *Game) correctedDistance(distance, rayAngle float64) float64 {
	corrected := distance * math.Cos(game.Player.Angle-rayAngle)
	if corrected < minWallDistance {
		return minWallDistance
	}

	return corrected
}
