package drawer

import (
	"doom/player"
	"doom/world"
	"math"
)

type Hit struct {
	Distance float64
	Texture  int
	WallX    float64
	Front    int
	Back     int
}

func (drawer *Drawer) CastRay(
	player *player.Player,
	rayAngle float64,
	vertices []world.Vertex,
	linedefs []world.Linedef,
) Hit {
	return drawer.CastRayFrom(player, rayAngle, vertices, linedefs, drawer.Step)
}

func (drawer *Drawer) CastRayFrom(
	player *player.Player,
	rayAngle float64,
	vertices []world.Vertex,
	linedefs []world.Linedef,
	startDistance float64,
) Hit {
	closestDistance := drawer.MaxDist
	hit := Hit{Distance: drawer.MaxDist}

	rayDirX := math.Cos(rayAngle)
	rayDirY := math.Sin(rayAngle)

	for _, line := range linedefs {
		if line.StartVertex < 0 || line.StartVertex >= len(vertices) {
			continue
		}
		if line.EndVertex < 0 || line.EndVertex >= len(vertices) {
			continue
		}

		start := vertices[line.StartVertex]
		end := vertices[line.EndVertex]
		distance, wallX, ok := raySegmentIntersection(player.X, player.Y, rayDirX, rayDirY, start, end)
		if !ok || distance < startDistance || distance >= closestDistance {
			continue
		}

		closestDistance = distance
		hit = Hit{
			Distance: distance,
			Texture:  line.TextureID,
			WallX:    wallX,
			Front:    line.FrontSector,
			Back:     line.BackSector,
		}
	}

	return hit
}

func raySegmentIntersection(
	rayOriginX, rayOriginY, rayDirX, rayDirY float64,
	start world.Vertex,
	end world.Vertex,
) (float64, float64, bool) {
	segmentX := end.X - start.X
	segmentY := end.Y - start.Y

	determinant := cross2D(rayDirX, rayDirY, segmentX, segmentY)
	if math.Abs(determinant) < 1e-9 {
		return 0, 0, false
	}

	deltaX := start.X - rayOriginX
	deltaY := start.Y - rayOriginY

	t := cross2D(deltaX, deltaY, segmentX, segmentY) / determinant
	u := cross2D(deltaX, deltaY, rayDirX, rayDirY) / determinant

	if t <= 0 || u < 0 || u > 1 {
		return 0, 0, false
	}

	return t, u, true
}

func cross2D(ax, ay, bx, by float64) float64 {
	return ax*by - ay*bx
}
