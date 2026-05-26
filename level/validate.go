package level

import "fmt"

func validateLevel(level Level) error {
	if len(level.WorldMap) == 0 {
		return fmt.Errorf("level has no map")
	}

	if !level.PlayerSpawnLoaded {
		return fmt.Errorf("player spawn not defined")
	}

	if level.FloorTextureFile == "" {
		return fmt.Errorf("floor texture is missing")
	}

	if level.SkyTextureFile == "" {
		return fmt.Errorf("sky texture is missing")
	}

	for _, row := range level.WorldMap {
		for _, tileID := range row {
			if tileID > 0 {
				if _, ok := level.WallTextureFiles[tileID]; !ok {
					return fmt.Errorf("missing wall texture mapping for tile %d", tileID)
				}
			}
		}
	}

	for _, spawn := range level.EnemySpawns {
		if spawn.SpriteKey == "" {
			return fmt.Errorf("enemy spawn missing sprite key")
		}
		if _, ok := level.EnemySpriteFiles[spawn.SpriteKey]; !ok {
			return fmt.Errorf("missing enemy sprite mapping for key %s", spawn.SpriteKey)
		}
	}

	return nil
}
