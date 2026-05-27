package game

import (
	"doom/world"
	"math"
)

func distanceToSegment(px, py, ax, ay, bx, by float64) float64 {
	dx := bx - ax
	dy := by - ay
	lengthSquared := dx*dx + dy*dy
	if lengthSquared == 0 {
		return math.Hypot(px-ax, py-ay)
	}

	t := ((px-ax)*dx + (py-ay)*dy) / lengthSquared
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}

	closestX := ax + t*dx
	closestY := ay + t*dy
	return math.Hypot(px-closestX, py-closestY)
}

func pointInPolygon(x, y float64, polygon []world.Vertex) bool {
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
