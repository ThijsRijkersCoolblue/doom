package level

type EnemySpawn struct {
	X         float64
	Y         float64
	SpriteKey string
}

type Level struct {
	WorldMap          [][]int
	PlayerSpawnX      float64
	PlayerSpawnY      float64
	WallTextureFiles  map[int]string
	EnemySpriteFiles  map[string]string
	FloorTextureFile  string
	SkyTextureFile    string
	EnemySpawns       []EnemySpawn
	PlayerSpawnLoaded bool
}
