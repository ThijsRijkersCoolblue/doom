package drawer

import "github.com/hajimehoshi/ebiten/v2"

func DrawTexturedColumn(
	screen *ebiten.Image,
	screenX int,
	wallHeight int,
	screenHeight int,
	texture *ebiten.Image,
	wallX float64,
) {
	if wallHeight <= 0 || texture == nil {
		return
	}

	textureWidth, textureHeight := texture.Bounds().Dx(), texture.Bounds().Dy()
	if textureWidth == 0 || textureHeight == 0 {
		return
	}

	startY, endY := visibleWallRange(wallHeight, screenHeight)
	if startY >= endY {
		return
	}

	textureX := textureXCoordinate(wallX, textureWidth)
	visibleHeight := endY - startY

	for screenY := startY; screenY < endY; screenY++ {
		textureY := textureYCoordinate(screenY, startY, visibleHeight, textureHeight)
		screen.Set(screenX, screenY, texture.At(textureX, textureY))
	}
}

func visibleWallRange(wallHeight, screenHeight int) (int, int) {
	startY := (screenHeight - wallHeight) / 2
	endY := startY + wallHeight

	if startY < 0 {
		startY = 0
	}
	if endY > screenHeight {
		endY = screenHeight
	}

	return startY, endY
}

func textureXCoordinate(wallX float64, textureWidth int) int {
	textureX := int(wallX * float64(textureWidth-1))

	if textureX < 0 {
		return 0
	}
	if textureX >= textureWidth {
		return textureWidth - 1
	}

	return textureX
}

func textureYCoordinate(screenY, startY, visibleHeight, textureHeight int) int {
	v := float64(screenY-startY) / float64(visibleHeight)
	textureY := int(v * float64(textureHeight-1))

	if textureY < 0 {
		return 0
	}
	if textureY >= textureHeight {
		return textureHeight - 1
	}

	return textureY
}
