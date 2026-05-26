package drawer

import (
	"doom/player"
	"math"
)

type Hit struct {
	Distance float64
	TileID   int
	WallX    float64
}

func (drawer *Drawer) CastRay(
	player *player.Player,
	rayAngle float64,
	worldMap [][]int,
) Hit {
	for distance := drawer.Step; distance < drawer.MaxDist; distance += drawer.Step {
		x := player.X + math.Cos(rayAngle)*distance
		y := player.Y + math.Sin(rayAngle)*distance

		mapX := int(x)
		mapY := int(y)

		if mapY < 0 || mapY >= len(worldMap) || mapX < 0 || mapX >= len(worldMap[0]) {
			return Hit{Distance: drawer.MaxDist}
		}

		tile := worldMap[mapY][mapX]
		if tile > 0 {
			wallX := wallCoordinate(x, y, rayAngle, drawer.Step, mapX)
			return Hit{Distance: distance, TileID: tile, WallX: wallX}
		}
	}

	return Hit{Distance: drawer.MaxDist}
}

func wallCoordinate(x, y, rayAngle, step float64, mapX int) float64 {
	prevX := x - math.Cos(rayAngle)*step

	var wallX float64
	if int(prevX) != mapX {
		wallX = y - math.Floor(y)
	} else {
		wallX = x - math.Floor(x)
	}

	if wallX < 0 {
		return 0
	}
	if wallX > 1 {
		return 1
	}

	return wallX
}
