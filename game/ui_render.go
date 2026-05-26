package game

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (game *Game) drawUI(screen *ebiten.Image) {
	if game.UIAssets == nil || game.UIAssets.Bar == nil {
		return
	}

	barRect := game.statusBarRect()
	game.drawStatusBarBase(screen, barRect)
	game.drawArmsPanel(screen, barRect)
	game.drawFace(screen, barRect)
	game.drawKeys(screen, barRect)
	game.drawStatsText(screen, barRect)
}

func (game *Game) statusBarRect() uiRect {
	height := int(math.Max(64, float64(game.ScreenHeight)*0.18))
	return uiRect{
		x: 0,
		y: game.ScreenHeight - height,
		w: game.ScreenWidth,
		h: height,
	}
}

func (game *Game) drawStatusBarBase(screen *ebiten.Image, rect uiRect) {
	drawScaled(screen, game.UIAssets.Bar, rect)
}

func (game *Game) drawArmsPanel(screen *ebiten.Image, rect uiRect) {
	if game.UIAssets.Arms == nil {
		return
	}

	panelW := int(float64(rect.h) * 1.2)
	panelH := int(float64(rect.h) * 0.85)
	panel := uiRect{
		x: rect.x + int(float64(rect.w)*0.13) - panelW/2,
		y: rect.y + (rect.h-panelH)/2,
		w: panelW,
		h: panelH,
	}
	drawScaled(screen, game.UIAssets.Arms, panel)
}

func (game *Game) drawFace(screen *ebiten.Image, rect uiRect) {
	face := game.currentFace()
	if face == nil {
		return
	}

	faceSize := int(float64(rect.h) * 0.78)
	faceRect := uiRect{
		x: rect.x + rect.w/2 - faceSize/2,
		y: rect.y + (rect.h-faceSize)/2,
		w: faceSize,
		h: faceSize,
	}
	drawScaled(screen, face, faceRect)
}

func (game *Game) drawKeys(screen *ebiten.Image, rect uiRect) {
	iconSize := int(float64(rect.h) * 0.22)
	startX := rect.x + int(float64(rect.w)*0.82)
	startY := rect.y + int(float64(rect.h)*0.22)
	gap := int(float64(iconSize) * 1.25)

	if game.UI.HasBlueKey {
		drawScaled(screen, game.UIAssets.KeyBlue, uiRect{x: startX, y: startY, w: iconSize, h: iconSize})
	}
	if game.UI.HasYelKey {
		drawScaled(screen, game.UIAssets.KeyYel, uiRect{x: startX, y: startY + gap, w: iconSize, h: iconSize})
	}
	if game.UI.HasRedKey {
		drawScaled(screen, game.UIAssets.KeyRed, uiRect{x: startX, y: startY + 2*gap, w: iconSize, h: iconSize})
	}
}

func (game *Game) drawStatsText(screen *ebiten.Image, rect uiRect) {
	leftX := rect.x + int(float64(rect.w)*0.04)
	topY := rect.y + int(float64(rect.h)*0.22)
	lineGap := int(float64(rect.h) * 0.24)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("HEALTH %3d", clampUIStat(game.UI.Health)), leftX, topY)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("ARMOR  %3d", clampUIStat(game.UI.Armor)), leftX, topY+lineGap)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("AMMO   %3d", clampUIStat(game.UI.Ammo)), leftX, topY+2*lineGap)
	ebitenutil.DebugPrintAt(screen, "FIRE: SPACE", leftX, topY-lineGap/2)
}

func (game *Game) currentFace() *ebiten.Image {
	if game.UI.Health <= 0 {
		return game.UIAssets.FaceDead
	}
	if game.UI.GodMode {
		return game.UIAssets.FaceGod
	}
	if game.UI.DamageFlashTicks > 0 {
		return game.UIAssets.FacePain
	}

	return game.UIAssets.FaceIdle
}

func clampUIStat(value int) int {
	if value < 0 {
		return 0
	}
	if value > 999 {
		return 999
	}
	return value
}

type uiRect struct {
	x int
	y int
	w int
	h int
}

func drawScaled(screen, image *ebiten.Image, rect uiRect) {
	if image == nil || rect.w <= 0 || rect.h <= 0 {
		return
	}

	imageW := image.Bounds().Dx()
	imageH := image.Bounds().Dy()
	if imageW == 0 || imageH == 0 {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(rect.w)/float64(imageW), float64(rect.h)/float64(imageH))
	op.GeoM.Translate(float64(rect.x), float64(rect.y))
	screen.DrawImage(image, op)
}
