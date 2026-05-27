package game

import (
	"doom/assets"
	"doom/drawer"
	"doom/entity"
	"doom/player"
	"doom/world"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player        player.Player
	Drawer        drawer.Drawer
	Vertices      []world.Vertex
	Linedefs      []world.Linedef
	Sectors       []world.Sector
	SectorPolys   [][]world.Vertex
	SectorAreas   []float64
	RayDirXCache  []float64
	RayDirYCache  []float64
	RayCacheAngle float64
	WallTextures  map[int]*ebiten.Image
	FloorTexture  *ebiten.Image
	SkyTexture    *ebiten.Image
	UIAssets      *assets.UIAssets
	UI            UIState
	Weapon        WeaponState
	EnemyHeight   float64
	Enemies       []entity.Enemy
	DepthBuffer   []float64
	ClipTop       []int
	ClipBottom    []int
	ClipDistance  []float64
	ScreenWidth   int
	ScreenHeight  int
}

func (game *Game) Update() error {
	game.syncPlayerSector()
	game.rotatePlayer()
	game.movePlayer()
	game.syncPlayerSector()
	game.updateEnemies()
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
		game.ensureClipBuffers()
		game.ensureSectorCache()
		return
	}

	game.DepthBuffer = make([]float64, game.ScreenWidth)
	game.ensureClipBuffers()
	game.ensureSectorCache()
}

func (game *Game) ensureClipBuffers() {
	if len(game.ClipTop) == game.ScreenWidth && len(game.ClipBottom) == game.ScreenWidth {
		if len(game.ClipDistance) == game.ScreenWidth {
			return
		}
	}

	game.ClipTop = make([]int, game.ScreenWidth)
	game.ClipBottom = make([]int, game.ScreenWidth)
	game.ClipDistance = make([]float64, game.ScreenWidth)
	for i := range game.ClipDistance {
		game.ClipDistance[i] = game.Drawer.MaxDist
	}
}
