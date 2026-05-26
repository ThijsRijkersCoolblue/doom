package game

import "github.com/hajimehoshi/ebiten/v2"

type UIState struct {
	Health int
	Armor  int
	Ammo   int

	HasBlueKey bool
	HasYelKey  bool
	HasRedKey  bool

	DamageFlashTicks int
	GodMode          bool
}

func (game *Game) updateUIState() {
	if game.UI.Health > 0 && ebiten.IsKeyPressed(ebiten.KeyH) {
		game.UI.Health--
		game.UI.DamageFlashTicks = 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		if game.UI.Health < 100 {
			game.UI.Health++
		}
	}

	if game.UI.DamageFlashTicks > 0 {
		game.UI.DamageFlashTicks--
	}
}
