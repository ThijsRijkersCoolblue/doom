package entity

import "github.com/hajimehoshi/ebiten/v2"

type Enemy struct {
	X              float64
	Y              float64
	Sprite         *ebiten.Image
	Alive          bool
	Speed          float64
	AttackRange    float64
	AttackCooldown int
	AttackDamage   int
	CooldownTicks  int
}
