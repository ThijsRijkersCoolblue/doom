package level

import "fmt"

var catalog = map[string]string{
	"level01": "levels/level01.txt",
}

func LoadByName(name string) (Level, error) {
	path, ok := catalog[name]
	if !ok {
		return Level{}, fmt.Errorf("unknown level name: %s", name)
	}

	return Load(path)
}
