package level

import "fmt"

func parseMapGrid(lines []string) ([][]int, float64, float64, []EnemySpawn, error) {
	if len(lines) == 0 {
		return nil, 0, 0, nil, fmt.Errorf("map section is empty")
	}

	width := len(lines[0])
	worldMap := make([][]int, len(lines))
	playerX, playerY := -1.0, -1.0
	var enemySpawns []EnemySpawn

	for y, line := range lines {
		if len(line) != width {
			return nil, 0, 0, nil, fmt.Errorf("map row width mismatch at row %d", y)
		}

		row, rowPlayerX, rowPlayerY, rowEnemies, err := parseMapRow(line, y)
		if err != nil {
			return nil, 0, 0, nil, err
		}

		if rowPlayerX >= 0 && rowPlayerY >= 0 {
			playerX = rowPlayerX
			playerY = rowPlayerY
		}

		enemySpawns = append(enemySpawns, rowEnemies...)
		worldMap[y] = row
	}

	return worldMap, playerX, playerY, enemySpawns, nil
}

func parseMapRow(line string, rowIndex int) ([]int, float64, float64, []EnemySpawn, error) {
	row := make([]int, len(line))
	playerX, playerY := -1.0, -1.0
	var enemySpawns []EnemySpawn

	for x, ch := range line {
		tileID, rowSpawnX, rowSpawnY, enemySpawn, err := parseMapCharacter(ch, x, rowIndex)
		if err != nil {
			return nil, 0, 0, nil, err
		}

		row[x] = tileID

		if rowSpawnX >= 0 && rowSpawnY >= 0 {
			playerX = rowSpawnX
			playerY = rowSpawnY
		}

		if enemySpawn != nil {
			enemySpawns = append(enemySpawns, *enemySpawn)
		}
	}

	return row, playerX, playerY, enemySpawns, nil
}

func parseMapCharacter(ch rune, x, y int) (int, float64, float64, *EnemySpawn, error) {
	spawnX := -1.0
	spawnY := -1.0

	switch ch {
	case charSolidTile:
		return 1, spawnX, spawnY, nil, nil
	case charEmptyTile:
		return 0, spawnX, spawnY, nil, nil
	case charPlayerSpawn:
		return 0, float64(x) + 0.5, float64(y) + 0.5, nil, nil
	case charEnemySpawn:
		enemy := EnemySpawn{
			X:         float64(x) + 0.5,
			Y:         float64(y) + 0.5,
			SpriteKey: defaultEnemyKey,
		}
		return 0, spawnX, spawnY, &enemy, nil
	default:
		if ch >= '2' && ch <= '9' {
			return int(ch - '0'), spawnX, spawnY, nil, nil
		}
		return 0, spawnX, spawnY, nil, fmt.Errorf("invalid map character '%c' at (%d,%d)", ch, x, y)
	}
}
