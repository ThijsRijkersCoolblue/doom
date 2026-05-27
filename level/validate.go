package level

import "fmt"

func validateLevel(level Level) error {
	if len(level.Vertices) == 0 {
		return fmt.Errorf("level has no vertices")
	}

	if len(level.Sectors) == 0 {
		return fmt.Errorf("level has no sectors")
	}

	if len(level.Linedefs) == 0 {
		return fmt.Errorf("level has no linedefs")
	}

	if !level.PlayerSpawnLoaded {
		return fmt.Errorf("player spawn not defined")
	}

	if level.PlayerSector < 0 || level.PlayerSector >= len(level.Sectors) {
		return fmt.Errorf("player spawn sector is invalid")
	}

	if level.FloorTextureFile == "" {
		return fmt.Errorf("floor texture is missing")
	}

	if level.SkyTextureFile == "" {
		return fmt.Errorf("sky texture is missing")
	}

	if level.EnemyVisualHeight <= 0 {
		return fmt.Errorf("enemy visual height must be positive")
	}

	for i, line := range level.Linedefs {
		if line.StartVertex < 0 || line.StartVertex >= len(level.Vertices) {
			return fmt.Errorf("linedef %d start vertex is invalid", i)
		}
		if line.EndVertex < 0 || line.EndVertex >= len(level.Vertices) {
			return fmt.Errorf("linedef %d end vertex is invalid", i)
		}
		if line.FrontSector < 0 || line.FrontSector >= len(level.Sectors) {
			return fmt.Errorf("linedef %d sector is invalid", i)
		}
		if line.BackSector >= len(level.Sectors) {
			return fmt.Errorf("linedef %d back sector is invalid", i)
		}
		if line.BackSector < -1 {
			return fmt.Errorf("linedef %d back sector is invalid", i)
		}
		if _, ok := level.WallTextureFiles[line.TextureID]; !ok {
			return fmt.Errorf("missing wall texture mapping for tile %d", line.TextureID)
		}
	}

	for i, sector := range level.Sectors {
		if sector.CeilingHeight <= sector.FloorHeight {
			return fmt.Errorf("sector %d has invalid floor and ceiling heights", i)
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
