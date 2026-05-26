package game

import "github.com/hajimehoshi/ebiten/v2"

type WeaponState struct {
	Frames          []*ebiten.Image
	CurrentFrame    int
	Shooting        bool
	AnimationTimer  int
	AnimationStep   int
	FireButtonLatch bool
}

func (game *Game) updateWeapon() {
	if len(game.Weapon.Frames) == 0 {
		return
	}

	if !game.Weapon.Shooting {
		game.Weapon.CurrentFrame = 0
	}

	firePressed := ebiten.IsKeyPressed(ebiten.KeySpace)
	if firePressed && !game.Weapon.FireButtonLatch && !game.Weapon.Shooting && game.UI.Ammo > 0 {
		game.startWeaponShot()
	}
	game.Weapon.FireButtonLatch = firePressed

	if game.Weapon.Shooting {
		game.Weapon.AnimationTimer++
		if game.Weapon.AnimationTimer >= game.Weapon.AnimationStep {
			game.Weapon.AnimationTimer = 0
			game.Weapon.CurrentFrame++

			if game.Weapon.CurrentFrame >= len(game.Weapon.Frames) {
				game.Weapon.CurrentFrame = 0
				game.Weapon.Shooting = false
			}
		}
	}
}

func (game *Game) startWeaponShot() {
	game.Weapon.Shooting = true
	game.Weapon.AnimationTimer = 0
	game.Weapon.CurrentFrame = 1
	game.UI.Ammo--
	game.tryHitEnemy()
}
