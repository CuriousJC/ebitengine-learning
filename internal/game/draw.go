package game

import (
	"image"
	"image/color"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/curiousjc/ebitengine-learning/internal/state"
)

func DrawMouseCursorImage(gs *state.GlobalState, screen *ebiten.Image) {

	cursorImage := gs.Assets["firering_png"]
	imageScale := 0.25
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScale, imageScale)
	op.GeoM.Translate(-float64(cursorImage.Bounds().Dx())*imageScale/2, -float64(cursorImage.Bounds().Dy())*imageScale/2) // Center the origin
	//op.GeoM.Rotate(math.Pi / 4)                                                                                           // Rotate 45 degrees
	op.GeoM.Translate(float64(gs.MouseX), float64(gs.MouseY)) // Follow the mouse

	screen.DrawImage(cursorImage.SubImage(image.Rect(0, 0, cursorImage.Bounds().Dx(), cursorImage.Bounds().Dy())).(*ebiten.Image), op)

}

func DrawHealthBar(gs *state.GlobalState, screen *ebiten.Image) {
	// Define the rectangle dimensions
	rectWidth := 300
	rectHeight := 50
	shadowOffsetX := 3 // Horizontal shadow offset
	shadowOffsetY := 3 // Vertical shadow offset

	// Define colors
	shadowColor := color.RGBA{A: 255, R: 255, G: 255, B: 255} // White shadow
	rectColor := color.RGBA{A: 255, R: 220, G: 20, B: 60}     // Crimson red

	//Shadow Image
	shadow := ebiten.NewImage(rectWidth, rectHeight)
	vector.DrawFilledRect(shadow, 0, 0, float32(rectWidth), float32(rectHeight), shadowColor, false)

	//Red Image
	rect := ebiten.NewImage(rectWidth, rectHeight)
	vector.DrawFilledRect(rect, 0, 0, float32(rectWidth), float32(rectHeight), rectColor, false)

	//Positioning
	horizontalPosition := float64(gs.FirstQuarterX)
	verticalPosition := float64(gs.FirstThirdY)

	// Draw the shadow first
	shadowOp := &ebiten.DrawImageOptions{}
	shadowOp.GeoM.Translate(-float64(rectWidth)/2, -float64(rectHeight)/2)  //center
	shadowOp.GeoM.Translate(horizontalPosition, verticalPosition)           //position
	shadowOp.GeoM.Translate(float64(shadowOffsetX), float64(shadowOffsetY)) //offset shadow
	screen.DrawImage(shadow, shadowOp)

	// Draw the main rectangle (no transformation)
	rectOp := &ebiten.DrawImageOptions{}
	rectOp.GeoM.Translate(-float64(rectWidth)/2, -float64(rectHeight)/2) //center
	rectOp.GeoM.Translate(horizontalPosition, verticalPosition)          //position
	screen.DrawImage(rect, rectOp)
}
