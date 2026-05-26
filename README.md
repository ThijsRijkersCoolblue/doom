# Doom

## Current progress

![Gameplay progress](etc/Screen%20Recording%202026-05-26%20at%2023.01.41.gif)

## Features

- Raycasted 3D-style world rendering
- Textured walls, floor, and sky
- Enemy sprites with hit detection
- Player HUD (health, armor, ammo)
- Animated pistol shooting
- Custom ASCII level format loader

## Requirements

- Go 1.25+

## Run

```bash
go run .
```

## Controls

- `W` / `S`: Move forward/backward
- `A` / `D`: Turn left/right
- `Space`: Shoot

## Development notes

- The game starts by loading `level01` from `levels/level01.txt`.
- Level metadata (walls, sprites, floor, sky) is parsed from section based ASCII files.
