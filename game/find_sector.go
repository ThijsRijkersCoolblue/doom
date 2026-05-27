package game

import "math"

func (game *Game) findSectorID(x, y float64) int {
	game.ensureSectorCache()

	closestSector := -1
	closestArea := math.MaxFloat64

	for i, polygon := range game.SectorPolys {
		if len(polygon) < 3 {
			continue
		}
		if !pointInPolygon(x, y, polygon) {
			continue
		}

		area := 0.0
		if i < len(game.SectorAreas) {
			area = game.SectorAreas[i]
		}
		if area <= 0 {
			area = math.MaxFloat64 - 1
		}

		if area < closestArea {
			closestArea = area
			closestSector = i
		}
	}

	return closestSector
}
