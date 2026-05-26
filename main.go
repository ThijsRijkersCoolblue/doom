package main

import (
	"doom/assets"
	"doom/drawer"
	"doom/game"
	"doom/player"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var worldMap = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 2, 2, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
}

func main() {
	textures, err := assets.LoadWallTextures(map[int]string{
		1: "assets/player/PLAYA1.png",
		2: "assets/monsters/SPOS/SPOSA1.png",
	})
	if err != nil {
		log.Fatal(err)
	}

	game := &game.Game{
		Player: player.Player{
			X:        3.5,
			Y:        3.5,
			Angle:    0,
			Speed:    0.2,
			Rotation: 0.1,
		},
		Drawer: drawer.Drawer{
			Step:    0.05,
			MaxDist: 20,
		},
		WorldMap:     worldMap,
		MapTextures:  textures,
		ScreenWidth:  640,
		ScreenHeight: 480,
	}

	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)

	if err = ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
