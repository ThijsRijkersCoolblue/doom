package game

import (
	"doom/drawer"
	"doom/player"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const minWallDistance = 0.001

type Game struct {
	Player       player.Player
	Drawer       drawer.Drawer
	WorldMap     [][]int
	MapTextures  map[int]*ebiten.Image
	ScreenWidth  int
	ScreenHeight int
}

func (game *Game) Update() error {
	game.movePlayer()
	game.rotatePlayer()

	return nil
}

func (game *Game) movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		game.Player.X += math.Cos(game.Player.Angle) * game.Player.Speed
		game.Player.Y += math.Sin(game.Player.Angle) * game.Player.Speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		game.Player.X -= math.Cos(game.Player.Angle) * game.Player.Speed
		game.Player.Y -= math.Sin(game.Player.Angle) * game.Player.Speed
	}
}

func (game *Game) rotatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.Player.Angle -= game.Player.Rotation
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		game.Player.Angle += game.Player.Rotation
	}
}

func (game *Game) Draw(screen *ebiten.Image) {
	fov := math.Pi / 3
	numberOfRays := game.ScreenWidth

	for screenX := 0; screenX < numberOfRays; screenX++ {
		rayAngle := game.rayAngleForColumn(screenX, numberOfRays, fov)
		hit := game.Drawer.CastRay(&game.Player, rayAngle, game.WorldMap)

		if hit.TileID == 0 {
			continue
		}

		wallDistance := game.correctedDistance(hit.Distance, rayAngle)
		wallHeight := int(float64(game.ScreenHeight) / wallDistance)

		texture := game.MapTextures[hit.TileID]
		drawer.DrawTexturedColumn(screen, screenX, wallHeight, game.ScreenHeight, texture, hit.WallX)
	}
}

func (game *Game) rayAngleForColumn(screenX, numberOfRays int, fov float64) float64 {
	return game.Player.Angle - fov/2 + fov*float64(screenX)/float64(numberOfRays)
}

func (game *Game) correctedDistance(distance, rayAngle float64) float64 {
	correctedDistance := distance * math.Cos(game.Player.Angle-rayAngle)
	if correctedDistance < minWallDistance {
		return minWallDistance
	}

	return correctedDistance
}

func (game *Game) Layout(_, _ int) (int, int) {
	return game.ScreenWidth, game.ScreenHeight
}
