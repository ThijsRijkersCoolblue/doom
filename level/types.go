package level

import "doom/world"

type EnemySpawn struct {
	X         float64
	Y         float64
	SpriteKey string
}

type Level struct {
	Vertices          []world.Vertex
	Linedefs          []world.Linedef
	Sectors           []world.Sector
	EnemyVisualHeight float64
	PlayerSpawnX      float64
	PlayerSpawnY      float64
	PlayerSector      int
	WallTextureFiles  map[int]string
	EnemySpriteFiles  map[string]string
	FloorTextureFile  string
	SkyTextureFile    string
	EnemySpawns       []EnemySpawn
	PlayerSpawnLoaded bool
}
