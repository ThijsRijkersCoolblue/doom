package game

import "math"

func (game *Game) tryHitEnemy() {
	closestIndex := -1
	closestDistance := game.Drawer.MaxDist

	for i, enemy := range game.Enemies {
		if !enemy.Alive || enemy.Sprite == nil {
			continue
		}

		dx := enemy.X - game.Player.X
		dy := enemy.Y - game.Player.Y
		distance := math.Hypot(dx, dy)
		if distance > game.Drawer.MaxDist {
			continue
		}

		angleToEnemy := normalizeAngle(math.Atan2(dy, dx) - game.Player.Angle)
		if math.Abs(angleToEnemy) > 0.1 {
			continue
		}

		screenX := game.ScreenWidth / 2
		if screenX < 0 || screenX >= len(game.DepthBuffer) {
			continue
		}

		if distance >= game.DepthBuffer[screenX] {
			continue
		}

		if distance < closestDistance {
			closestDistance = distance
			closestIndex = i
		}
	}

	if closestIndex >= 0 {
		game.Enemies[closestIndex].Alive = false
	}
}
