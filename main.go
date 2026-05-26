package main

import (
	"doom/assets"
	"doom/drawer"
	"doom/entity"
	"doom/game"
	"doom/level"
	"doom/player"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	loadedLevel, err := level.LoadByName("level01")
	if err != nil {
		log.Fatal(err)
	}

	wallTextures, err := assets.LoadWallTextures(loadedLevel.WallTextureFiles)
	if err != nil {
		log.Fatal(err)
	}

	enemyTextures, err := assets.LoadSpriteTextures(loadedLevel.EnemySpriteFiles)
	if err != nil {
		log.Fatal(err)
	}

	pistolFrames, err := assets.LoadPistolFrames()
	if err != nil {
		log.Fatal(err)
	}

	floorTexture, err := assets.LoadTexture(loadedLevel.FloorTextureFile)
	if err != nil {
		log.Fatal(err)
	}

	skyTexture, err := assets.LoadTexture(loadedLevel.SkyTextureFile)
	if err != nil {
		log.Fatal(err)
	}

	uiAssets, err := assets.LoadUIAssets()
	if err != nil {
		log.Fatal(err)
	}

	gameEnemies := make([]entity.Enemy, 0, len(loadedLevel.EnemySpawns))
	for _, spawn := range loadedLevel.EnemySpawns {
		sprite := enemyTextures[spawn.SpriteKey]
		gameEnemies = append(gameEnemies, entity.Enemy{
			X:      spawn.X,
			Y:      spawn.Y,
			Sprite: sprite,
			Alive:  true,
		})
	}

	gameInstance := &game.Game{
		Player: player.Player{
			X:        loadedLevel.PlayerSpawnX,
			Y:        loadedLevel.PlayerSpawnY,
			Angle:    0,
			Speed:    0.2,
			Rotation: 0.06,
		},
		Drawer: drawer.Drawer{
			Step:    0.02,
			MaxDist: 30,
		},
		WorldMap:     loadedLevel.WorldMap,
		WallTextures: wallTextures,
		FloorTexture: floorTexture,
		SkyTexture:   skyTexture,
		UIAssets:     uiAssets,
		UI: game.UIState{
			Health: 100,
			Armor:  25,
			Ammo:   50,
		},
		Weapon: game.WeaponState{
			Frames:        pistolFrames,
			AnimationStep: 3,
		},
		Enemies:      gameEnemies,
		ScreenWidth:  960,
		ScreenHeight: 540,
	}

	ebiten.SetWindowSize(gameInstance.ScreenWidth, gameInstance.ScreenHeight)
	ebiten.SetWindowTitle("doom")

	if err = ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}
