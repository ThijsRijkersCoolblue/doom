package game

const playerEyeHeight = 0.5

func (game *Game) syncPlayerSector() {
	sectorID := game.findSectorID(game.Player.X, game.Player.Y)
	if sectorID < 0 || sectorID >= len(game.Sectors) {
		return
	}

	game.Player.SectorID = sectorID
	game.Player.Z = game.Sectors[sectorID].FloorHeight + playerEyeHeight
}

func (game *Game) currentSectorFloor() float64 {
	if game.Player.SectorID < 0 || game.Player.SectorID >= len(game.Sectors) {
		return 0
	}

	return game.Sectors[game.Player.SectorID].FloorHeight
}
