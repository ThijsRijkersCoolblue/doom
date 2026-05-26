package level

import "strings"

func ParseASCIILevel(input string) (Level, error) {
	level := Level{
		WallTextureFiles: make(map[int]string),
		EnemySpriteFiles: make(map[string]string),
	}

	lines := strings.Split(input, "\n")
	section := sectionNone
	var mapLines []string

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if isSectionHeader(line) {
			section = strings.ToLower(line)
			continue
		}

		if err := processSectionLine(section, line, &level, &mapLines); err != nil {
			return Level{}, err
		}
	}

	worldMap, playerX, playerY, enemySpawns, err := parseMapGrid(mapLines)
	if err != nil {
		return Level{}, err
	}

	level.WorldMap = worldMap
	level.PlayerSpawnX = playerX
	level.PlayerSpawnY = playerY
	level.EnemySpawns = append(level.EnemySpawns, enemySpawns...)
	level.PlayerSpawnLoaded = playerX >= 0 && playerY >= 0

	if err := validateLevel(level); err != nil {
		return Level{}, err
	}

	return level, nil
}
