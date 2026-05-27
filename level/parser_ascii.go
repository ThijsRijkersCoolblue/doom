package level

import "strings"

func ParseASCIILevel(input string) (Level, error) {
	level := Level{
		WallTextureFiles:  make(map[int]string),
		EnemySpriteFiles:  make(map[string]string),
		EnemyVisualHeight: 0.7,
	}

	lines := strings.Split(input, "\n")
	section := sectionNone

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if isSectionHeader(line) {
			section = strings.ToLower(line)
			continue
		}

		if err := processSectionLine(section, line, &level); err != nil {
			return Level{}, err
		}
	}

	if err := validateLevel(level); err != nil {
		return Level{}, err
	}

	return level, nil
}
