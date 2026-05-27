package world

import "math"

func PointInSector(x, y float64, sectorIndex int, vertices []Vertex, linedefs []Linedef) bool {
	polygon := sectorPolygon(sectorIndex, vertices, linedefs)
	if len(polygon) < 3 {
		return false
	}

	inside := false
	for i, j := 0, len(polygon)-1; i < len(polygon); j, i = i, i+1 {
		xi := polygon[i].X
		yi := polygon[i].Y
		xj := polygon[j].X
		yj := polygon[j].Y

		intersects := ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi+1e-9)+xi)
		if intersects {
			inside = !inside
		}
	}

	return inside
}

func IsPointInAnySector(x, y float64, vertices []Vertex, linedefs []Linedef, sectors []Sector) bool {
	return FindSector(x, y, vertices, linedefs, sectors) >= 0
}

func FindSector(x, y float64, vertices []Vertex, linedefs []Linedef, sectors []Sector) int {
	closestSector := -1
	closestArea := math.MaxFloat64

	for i := range sectors {
		if !PointInSector(x, y, i, vertices, linedefs) {
			continue
		}

		area := sectorArea(i, vertices, linedefs)
		if area > 0 && area < closestArea {
			closestArea = area
			closestSector = i
		}
	}

	return closestSector
}

func FindSectorFromPolygons(x, y float64, polygons [][]Vertex, areas []float64) int {
	closestSector := -1
	closestArea := math.MaxFloat64

	for i, polygon := range polygons {
		if len(polygon) < 3 {
			continue
		}
		if !pointInPolygon(x, y, polygon) {
			continue
		}

		area := 0.0
		if i < len(areas) {
			area = areas[i]
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

func DistanceToSegment(x, y float64, start Vertex, end Vertex) float64 {
	dx := end.X - start.X
	dy := end.Y - start.Y
	lengthSquared := dx*dx + dy*dy
	if lengthSquared == 0 {
		return math.Hypot(x-start.X, y-start.Y)
	}

	t := ((x-start.X)*dx + (y-start.Y)*dy) / lengthSquared
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}

	closestX := start.X + t*dx
	closestY := start.Y + t*dy
	return math.Hypot(x-closestX, y-closestY)
}

func sectorPolygon(sectorIndex int, vertices []Vertex, linedefs []Linedef) []Vertex {
	polygon := make([]Vertex, 0)
	for _, line := range linedefs {
		if line.FrontSector != sectorIndex {
			continue
		}
		if line.StartVertex < 0 || line.StartVertex >= len(vertices) {
			continue
		}
		polygon = append(polygon, vertices[line.StartVertex])
	}

	return polygon
}

func pointInPolygon(x, y float64, polygon []Vertex) bool {
	inside := false
	for i, j := 0, len(polygon)-1; i < len(polygon); j, i = i, i+1 {
		xi := polygon[i].X
		yi := polygon[i].Y
		xj := polygon[j].X
		yj := polygon[j].Y

		intersects := ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi+1e-9)+xi)
		if intersects {
			inside = !inside
		}
	}

	return inside
}

func sectorArea(sectorIndex int, vertices []Vertex, linedefs []Linedef) float64 {
	polygon := sectorPolygon(sectorIndex, vertices, linedefs)
	if len(polygon) < 3 {
		return 0
	}

	area := 0.0
	for i := 0; i < len(polygon); i++ {
		next := (i + 1) % len(polygon)
		area += polygon[i].X*polygon[next].Y - polygon[next].X*polygon[i].Y
	}

	return math.Abs(area) * 0.5
}
