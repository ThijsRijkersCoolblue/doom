package game

import (
	"doom/drawer"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (game *Game) drawWalls(screen *ebiten.Image) {
	numberOfRays := game.ScreenWidth

	for screenX := 0; screenX < numberOfRays; screenX++ {
		rayAngle := game.rayAngleForColumn(screenX, numberOfRays)
		game.DepthBuffer[screenX] = game.Drawer.MaxDist
		startDistance := game.Drawer.Step
		clipTop := 0
		clipBottom := game.ScreenHeight
		game.ClipTop[screenX] = -1
		game.ClipBottom[screenX] = -1
		game.ClipDistance[screenX] = game.Drawer.MaxDist

		for pass := 0; pass < 16; pass++ {
			if clipTop >= clipBottom {
				break
			}

			hit := game.Drawer.CastRayFrom(&game.Player, rayAngle, game.Vertices, game.Linedefs, startDistance)
			if hit.Texture == 0 || hit.Front < 0 || hit.Front >= len(game.Sectors) {
				break
			}

			frontSectorID, backSectorID := game.resolveHitSectors(hit.Front, hit.Back)
			if frontSectorID < 0 || frontSectorID >= len(game.Sectors) {
				break
			}

			wallDistance := game.correctedDistance(hit.Distance, rayAngle)
			frontSector := game.Sectors[frontSectorID]
			texture := game.WallTextures[hit.Texture]

			if backSectorID >= 0 && backSectorID < len(game.Sectors) {
				backSector := game.Sectors[backSectorID]
				drewPortalWall := false

				upperTop := maxFloat(frontSector.CeilingHeight, backSector.CeilingHeight)
				upperBottom := minFloat(frontSector.CeilingHeight, backSector.CeilingHeight)
				if upperTop-upperBottom > minWallDistance {
					topY, bottomY := game.projectWallSpan(upperTop, upperBottom, wallDistance)
					drawer.DrawTexturedColumn(screen, screenX, topY, bottomY, clipTop, clipBottom, game.ScreenHeight, texture, hit.WallX)
					drewPortalWall = true
				}

				lowerTop := maxFloat(frontSector.FloorHeight, backSector.FloorHeight)
				lowerBottom := minFloat(frontSector.FloorHeight, backSector.FloorHeight)
				hideFrontPitWall := pass == 0 && frontSector.FloorHeight > backSector.FloorHeight
				if hideFrontPitWall {
					ledgeY, _ := game.projectWallSpan(lowerTop, lowerBottom, wallDistance)
					game.ClipTop[screenX] = ledgeY
					game.ClipDistance[screenX] = wallDistance
				}
				if lowerTop-lowerBottom > minWallDistance && !hideFrontPitWall {
					topY, bottomY := game.projectWallSpan(lowerTop, lowerBottom, wallDistance)
					drawer.DrawTexturedColumn(screen, screenX, topY, bottomY, clipTop, clipBottom, game.ScreenHeight, texture, hit.WallX)
					drewPortalWall = true
				}

				if drewPortalWall {
					game.DepthBuffer[screenX] = minFloat(game.DepthBuffer[screenX], wallDistance)
				}

				openFloor := maxFloat(frontSector.FloorHeight, backSector.FloorHeight)
				openCeiling := minFloat(frontSector.CeilingHeight, backSector.CeilingHeight)
				if openCeiling-openFloor <= minWallDistance {
					game.DepthBuffer[screenX] = wallDistance
					break
				}

				openTopY, openBottomY := game.projectWallSpan(openCeiling, openFloor, wallDistance)
				clipTop = maxInt(clipTop, openTopY)
				clipBottom = minInt(clipBottom, openBottomY)

				startDistance = hit.Distance + game.Drawer.Step
				continue
			}

			topY, bottomY := game.projectWallSpan(frontSector.CeilingHeight, frontSector.FloorHeight, wallDistance)
			drawer.DrawTexturedColumn(screen, screenX, topY, bottomY, clipTop, clipBottom, game.ScreenHeight, texture, hit.WallX)
			game.DepthBuffer[screenX] = wallDistance
			break
		}
	}
}

func (game *Game) resolveHitSectors(frontSectorID, backSectorID int) (int, int) {
	if backSectorID < 0 || backSectorID >= len(game.Sectors) {
		return frontSectorID, -1
	}

	if game.Player.SectorID == backSectorID && game.Player.SectorID != frontSectorID {
		return backSectorID, frontSectorID
	}

	return frontSectorID, backSectorID
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func (game *Game) rayAngleForColumn(screenX, numberOfRays int) float64 {
	return game.Player.Angle - fieldOfView/2 + fieldOfView*float64(screenX)/float64(numberOfRays)
}

func (game *Game) correctedDistance(distance, rayAngle float64) float64 {
	corrected := distance * math.Cos(game.Player.Angle-rayAngle)
	if corrected < minWallDistance {
		return minWallDistance
	}

	return corrected
}

func (game *Game) projectWallSpan(ceilingHeight, floorHeight, distance float64) (int, int) {
	projectionScale := float64(game.ScreenHeight) * 0.5
	screenCenter := float64(game.ScreenHeight) * 0.5

	top := screenCenter - ((ceilingHeight-game.Player.Z)*projectionScale)/distance
	bottom := screenCenter - ((floorHeight-game.Player.Z)*projectionScale)/distance
	return int(top), int(bottom)
}
