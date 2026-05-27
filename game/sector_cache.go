package game

import "doom/world"

func (game *Game) ensureSectorCache() {
	if len(game.SectorPolys) == len(game.Sectors) && len(game.SectorAreas) == len(game.Sectors) {
		return
	}

	game.SectorPolys = buildSectorPolygons(game.Vertices, game.Linedefs, game.Sectors)
	game.SectorAreas = make([]float64, len(game.Sectors))
	for i := range game.SectorPolys {
		game.SectorAreas[i] = polygonArea(game.SectorPolys[i])
	}
}

func buildSectorPolygons(vertices []world.Vertex, linedefs []world.Linedef, sectors []world.Sector) [][]world.Vertex {
	polygons := make([][]world.Vertex, len(sectors))
	for sectorIndex := range sectors {
		poly := make([]world.Vertex, 0)
		for _, line := range linedefs {
			if line.FrontSector != sectorIndex {
				continue
			}
			if line.StartVertex < 0 || line.StartVertex >= len(vertices) {
				continue
			}
			poly = append(poly, vertices[line.StartVertex])
		}
		polygons[sectorIndex] = poly
	}

	return polygons
}

func polygonArea(polygon []world.Vertex) float64 {
	if len(polygon) < 3 {
		return 0
	}

	area := 0.0
	for i := 0; i < len(polygon); i++ {
		next := (i + 1) % len(polygon)
		area += polygon[i].X*polygon[next].Y - polygon[next].X*polygon[i].Y
	}

	if area < 0 {
		area = -area
	}
	return area * 0.5
}
