package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const playerRadius = 0.2

func (game *Game) rotatePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.Player.Angle -= game.Player.Rotation
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		game.Player.Angle += game.Player.Rotation
	}
}

func (game *Game) movePlayer() {
	direction := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		direction += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		direction -= 1
	}

	if direction == 0 {
		return
	}

	deltaX := math.Cos(game.Player.Angle) * game.Player.Speed * direction
	deltaY := math.Sin(game.Player.Angle) * game.Player.Speed * direction
	game.tryMove(deltaX, deltaY)
}

func (game *Game) tryMove(deltaX, deltaY float64) {
	newX := game.Player.X + deltaX
	if game.canMoveTo(newX, game.Player.Y) {
		game.Player.X = newX
	}

	newY := game.Player.Y + deltaY
	if game.canMoveTo(game.Player.X, newY) {
		game.Player.Y = newY
	}
}

func (game *Game) canMoveTo(x, y float64) bool {
	points := [][2]float64{
		{x - playerRadius, y - playerRadius},
		{x + playerRadius, y - playerRadius},
		{x - playerRadius, y + playerRadius},
		{x + playerRadius, y + playerRadius},
	}

	for _, point := range points {
		if game.isBlocked(point[0], point[1]) {
			return false
		}
	}

	return true
}

func (game *Game) isBlocked(x, y float64) bool {
	mapX := int(x)
	mapY := int(y)

	if mapY < 0 || mapY >= len(game.WorldMap) || mapX < 0 || mapX >= len(game.WorldMap[0]) {
		return true
	}

	return game.WorldMap[mapY][mapX] > 0
}
