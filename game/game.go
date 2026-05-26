package game

import (
	"doom/assets"
	"doom/drawer"
	"doom/entity"
	"doom/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player       player.Player
	Drawer       drawer.Drawer
	WorldMap     [][]int
	WallTextures map[int]*ebiten.Image
	FloorTexture *ebiten.Image
	SkyTexture   *ebiten.Image
	UIAssets     *assets.UIAssets
	UI           UIState
	Weapon       WeaponState
	Enemies      []entity.Enemy
	DepthBuffer  []float64
	ScreenWidth  int
	ScreenHeight int
}

func (game *Game) Update() error {
	game.rotatePlayer()
	game.movePlayer()
	game.updateUIState()
	game.updateWeapon()

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.ensureDepthBuffer()
	game.drawBackground(screen)
	game.drawWalls(screen)
	game.drawEnemies(screen)
	game.drawWeapon(screen)
	game.drawUI(screen)
}

func (game *Game) Layout(_, _ int) (int, int) {
	return game.ScreenWidth, game.ScreenHeight
}

func (game *Game) ensureDepthBuffer() {
	if len(game.DepthBuffer) == game.ScreenWidth {
		return
	}

	game.DepthBuffer = make([]float64, game.ScreenWidth)
}
