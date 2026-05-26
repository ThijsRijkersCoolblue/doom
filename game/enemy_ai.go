package game

import (
	"doom/entity"
	"math"
)

func (game *Game) updateEnemies() {
	for i := range game.Enemies {
		enemy := &game.Enemies[i]
		if !enemy.Alive || enemy.Sprite == nil {
			continue
		}

		if enemy.CooldownTicks > 0 {
			enemy.CooldownTicks--
		}

		dx := game.Player.X - enemy.X
		dy := game.Player.Y - enemy.Y
		distance := math.Hypot(dx, dy)
		if distance < minWallDistance {
			continue
		}

		if !game.hasLineOfSight(enemy.X, enemy.Y, game.Player.X, game.Player.Y) {
			continue
		}

		if distance <= enemy.AttackRange {
			game.tryEnemyAttack(enemy)
			continue
		}

		normalizedX := dx / distance
		normalizedY := dy / distance
		stepX := normalizedX * enemy.Speed
		stepY := normalizedY * enemy.Speed
		game.tryMoveEnemy(enemy, stepX, stepY)
	}
}

func (game *Game) tryMoveEnemy(enemy *entity.Enemy, stepX, stepY float64) {
	newX := enemy.X + stepX
	if game.canEnemyMoveTo(newX, enemy.Y) {
		enemy.X = newX
	}

	newY := enemy.Y + stepY
	if game.canEnemyMoveTo(enemy.X, newY) {
		enemy.Y = newY
	}
}

func (game *Game) canEnemyMoveTo(x, y float64) bool {
	points := [][2]float64{
		{x - enemyRadius, y - enemyRadius},
		{x + enemyRadius, y - enemyRadius},
		{x - enemyRadius, y + enemyRadius},
		{x + enemyRadius, y + enemyRadius},
	}

	for _, point := range points {
		if game.isBlocked(point[0], point[1]) {
			return false
		}
	}

	if math.Hypot(game.Player.X-x, game.Player.Y-y) < playerRadius+enemyRadius {
		return false
	}

	return true
}

func (game *Game) hasLineOfSight(fromX, fromY, toX, toY float64) bool {
	dx := toX - fromX
	dy := toY - fromY
	distance := math.Hypot(dx, dy)
	if distance < minWallDistance {
		return true
	}

	stepCount := int(distance/lineStep) + 1
	stepX := dx / float64(stepCount)
	stepY := dy / float64(stepCount)

	x := fromX
	y := fromY
	for i := 0; i < stepCount; i++ {
		x += stepX
		y += stepY
		if game.isBlocked(x, y) {
			return false
		}
	}

	return true
}

func (game *Game) tryEnemyAttack(enemy *entity.Enemy) {
	if enemy.CooldownTicks > 0 {
		return
	}

	if game.UI.GodMode || game.UI.Health <= 0 {
		enemy.CooldownTicks = enemy.AttackCooldown
		return
	}

	damage := enemy.AttackDamage
	if game.UI.Armor > 0 {
		armorAbsorb := damage / 2
		if armorAbsorb > game.UI.Armor {
			armorAbsorb = game.UI.Armor
		}
		game.UI.Armor -= armorAbsorb
		damage -= armorAbsorb
	}

	game.UI.Health -= damage
	if game.UI.Health < 0 {
		game.UI.Health = 0
	}
	game.UI.DamageFlashTicks = 12
	enemy.CooldownTicks = enemy.AttackCooldown
}
