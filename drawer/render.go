package drawer

import "github.com/hajimehoshi/ebiten/v2"

func DrawTexturedColumn(
	screen *ebiten.Image,
	screenX int,
	topY int,
	bottomY int,
	clipTop int,
	clipBottom int,
	screenHeight int,
	texture *ebiten.Image,
	wallX float64,
) {
	if texture == nil {
		return
	}

	textureWidth, textureHeight := texture.Bounds().Dx(), texture.Bounds().Dy()
	if textureWidth == 0 || textureHeight == 0 {
		return
	}

	startY, endY := visibleWallRange(topY, bottomY, clipTop, clipBottom, screenHeight)
	if startY >= endY {
		return
	}

	textureX := textureXCoordinate(wallX, textureWidth)
	wallHeight := bottomY - topY
	if wallHeight <= 0 {
		return
	}

	for screenY := startY; screenY < endY; screenY++ {
		textureY := textureYCoordinate(screenY, topY, wallHeight, textureHeight)
		screen.Set(screenX, screenY, texture.At(textureX, textureY))
	}
}

func visibleWallRange(topY, bottomY, clipTop, clipBottom, screenHeight int) (int, int) {
	startY := maxInt(topY, clipTop)
	endY := minInt(bottomY, clipBottom)

	if startY < 0 {
		startY = 0
	}
	if endY > screenHeight {
		endY = screenHeight
	}

	return startY, endY
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
