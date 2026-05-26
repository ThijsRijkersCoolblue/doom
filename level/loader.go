package level

import "os"

func Load(path string) (Level, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return Level{}, err
	}

	return ParseASCIILevel(string(content))
}
