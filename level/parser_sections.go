package level

import (
	"fmt"
	"strings"
)

func isSectionHeader(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func processSectionLine(section, line string, level *Level, mapLines *[]string) error {
	switch section {
	case sectionMap:
		*mapLines = append(*mapLines, line)
		return nil
	case sectionWalls:
		return parseWallMapping(line, level.WallTextureFiles)
	case sectionEnemies:
		return parseEnemySpriteMapping(line, level.EnemySpriteFiles)
	case sectionFloors:
		return parseSinglePath(line, "floor", &level.FloorTextureFile)
	case sectionSky:
		return parseSinglePath(line, "sky", &level.SkyTextureFile)
	case sectionSpawns:
		return parseSpawnLine(line, level)
	default:
		return fmt.Errorf("line outside section: %s", line)
	}
}
