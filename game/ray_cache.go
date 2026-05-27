package game

import "math"

func (game *Game) ensureRayDirectionCache() {
	if len(game.RayDirXCache) == game.ScreenWidth &&
		len(game.RayDirYCache) == game.ScreenWidth &&
		math.Abs(game.RayCacheAngle-game.Player.Angle) < 1e-9 {
		return
	}

	game.RayDirXCache = make([]float64, game.ScreenWidth)
	game.RayDirYCache = make([]float64, game.ScreenWidth)

	for x := 0; x < game.ScreenWidth; x++ {
		rayAngle := game.rayAngleForColumn(x, game.ScreenWidth)
		game.RayDirXCache[x] = math.Cos(rayAngle)
		game.RayDirYCache[x] = math.Sin(rayAngle)
	}

	game.RayCacheAngle = game.Player.Angle
}
