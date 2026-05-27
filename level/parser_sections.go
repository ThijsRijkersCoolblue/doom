package level

import (
	"fmt"
	"strings"
)

func isSectionHeader(line string) bool {
	return strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]")
}

func processSectionLine(section, line string, level *Level) error {
	switch section {
	case sectionVertices:
		return parseVertexLine(line, level)
	case sectionSectors:
		return parseSectorLine(line, level)
	case sectionLinedefs:
		return parseLinedefLine(line, level)
	case sectionWalls:
		return parseWallMapping(line, level.WallTextureFiles)
	case sectionEnemies:
		return parseEnemySpriteMapping(line, level.EnemySpriteFiles)
	case sectionEnemySize:
		return parseEnemyHeight(line, level)
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
