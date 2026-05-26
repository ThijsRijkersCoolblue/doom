package level

import (
	"fmt"
	"strconv"
	"strings"
)

func parseWallMapping(line string, mapping map[int]string) error {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid wall mapping: %s", line)
	}

	id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("invalid wall id: %s", line)
	}

	mapping[id] = strings.TrimSpace(parts[1])
	return nil
}

func parseEnemySpriteMapping(line string, mapping map[string]string) error {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid enemy sprite mapping: %s", line)
	}

	key := strings.TrimSpace(parts[0])
	if key == "" {
		return fmt.Errorf("enemy sprite key is empty: %s", line)
	}

	mapping[key] = strings.TrimSpace(parts[1])
	return nil
}

func parseSinglePath(line, expectedKey string, output *string) error {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid %s mapping: %s", expectedKey, line)
	}

	key := strings.TrimSpace(parts[0])
	if key != expectedKey {
		return fmt.Errorf("expected key %s, got: %s", expectedKey, key)
	}

	*output = strings.TrimSpace(parts[1])
	return nil
}

func parseSpawnLine(line string, level *Level) error {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return fmt.Errorf("invalid spawn line: %s", line)
	}

	kind := strings.ToLower(parts[0])
	x, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return fmt.Errorf("invalid spawn x: %s", line)
	}

	y, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return fmt.Errorf("invalid spawn y: %s", line)
	}

	if kind == "player" {
		level.PlayerSpawnX = x
		level.PlayerSpawnY = y
		level.PlayerSpawnLoaded = true
		return nil
	}

	if kind == "enemy" {
		if len(parts) < 4 {
			return fmt.Errorf("enemy spawn missing sprite key: %s", line)
		}

		level.EnemySpawns = append(level.EnemySpawns, EnemySpawn{
			X:         x,
			Y:         y,
			SpriteKey: parts[3],
		})
		return nil
	}

	return fmt.Errorf("unknown spawn kind: %s", line)
}
