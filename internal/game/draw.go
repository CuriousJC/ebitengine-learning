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
	horizontalPosition := float64(gs.FirstQuarterX)
	verticalPosition := float64(gs.FirstThirdY)

	// Define colors
	shadowColor := color.RGBA{A: 255, R: 255, G: 255, B: 255} // White shadow
	rectColor := color.RGBA{A: 255, R: 220, G: 20, B: 60}     // Crimson red

	//Shadow Image
	shadow := ebiten.NewImage(rectWidth, rectHeight)
	vector.DrawFilledRect(shadow, 0, 0, float32(rectWidth), float32(rectHeight), shadowColor, false)

	//Red Image
	rect := ebiten.NewImage(rectWidth, rectHeight)
	vector.DrawFilledRect(rect, 0, 0, float32(rectWidth), float32(rectHeight), rectColor, false)

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

// DrawRoundedHealthBar is my first attempt to use SourceIn Blending.
// It creates a mask of pink circles and rectangles that makes the shape
// of the rounded button and lays it over the filled transparency
// I then "apply" this mask by source blending in the redRectangle
func DrawRoundedHealthBar(gs *state.GlobalState, screen *ebiten.Image) {

	rectWidth := 300
	rectHeight := 25
	horizontalPosition := float64(gs.ThirdQuarterX)
	verticalPosition := float64(gs.FirstThirdY)
	rectColor := color.RGBA{A: 255, R: 96, G: 37, B: 37}    // Crimson red
	maskColor := color.RGBA{A: 255, R: 255, G: 192, B: 203} // mask color
	maskFill := color.RGBA{A: 0, R: 255, G: 255, B: 255}    //fill transparent

	//Rounded corners mask Image
	mask := ebiten.NewImage(rectWidth, rectHeight)
	mask.Fill(maskFill)
	DrawRoundedRectMask(mask, 0, 0, float32(rectWidth), float32(rectHeight), 10, maskColor)

	//Red Rectangle
	redRect := ebiten.NewImage(rectWidth, rectHeight)
	vector.DrawFilledRect(redRect, 0, 0, float32(rectWidth), float32(rectHeight), rectColor, false)

	//Blend red rectangle into the mask.  The source is the red rectangle
	//  and the mask will only display that portion of the source that is blended into the non-transparent alpha
	rectMaskOp := &ebiten.DrawImageOptions{}
	rectMaskOp.Blend = ebiten.BlendSourceIn
	mask.DrawImage(redRect, rectMaskOp) //blend source redRect into destination mask
	healthBar := mask                   //mask is now filled with transparent maskFill and the maskColor was overwritten with redRect

	healthBarOp := &ebiten.DrawImageOptions{}
	healthBarOp.GeoM.Translate(-float64(rectWidth)/2, -float64(rectHeight)/2) //center our origin
	healthBarOp.GeoM.Translate(horizontalPosition, verticalPosition)          //position
	screen.DrawImage(healthBar, healthBarOp)
}

func DrawRoundedRectMask(mask *ebiten.Image, x, y, width, height, radius float32, maskColor color.Color) {
	//Round Corners
	vector.DrawFilledCircle(mask, x+radius, y+radius, radius, maskColor, false)
	vector.DrawFilledCircle(mask, x+width-radius, y+radius, radius, maskColor, false)
	vector.DrawFilledCircle(mask, x+radius, y+height-radius, radius, maskColor, false)
	vector.DrawFilledCircle(mask, x+width-radius, y+height-radius, radius, maskColor, false)

	//Rectangle Edges
	vector.DrawFilledRect(mask, x+radius, y, width-2*radius, radius, maskColor, false)               //top edge
	vector.DrawFilledRect(mask, x+radius, y+height-radius, width-2*radius, radius, maskColor, false) //bottom edge
	vector.DrawFilledRect(mask, x, y+radius, radius+width, height-2*radius, maskColor, false)        //left edge
	vector.DrawFilledRect(mask, x+width-radius, y+radius, radius, height-2*radius, maskColor, false) //right edge
}

func DrawMaskColorized(gs *state.GlobalState, screen *ebiten.Image) {

	rectWidth := 300
	rectHeight := 25
	horizontalPosition := float64(gs.ThirdQuarterX)
	verticalPosition := float64(gs.FirstThirdY) + 50
	maskColor := color.RGBA{A: 255, R: 0, G: 0, B: 0} // mask color

	//Rounded corners mask Image
	mask := ebiten.NewImage(rectWidth, rectHeight)
	fill := color.RGBA{A: 255, R: 255, G: 255, B: 255}
	mask.Fill(fill)
	DrawRoundedRectMaskColorized(mask, 0, 0, float32(rectWidth), float32(rectHeight), 10, maskColor)

	maskOp := &ebiten.DrawImageOptions{}
	maskOp.GeoM.Translate(horizontalPosition-float64(rectWidth)/2, verticalPosition-float64(rectHeight)/2) // Position the mask
	screen.DrawImage(mask, maskOp)
}

func DrawRoundedRectMaskColorized(mask *ebiten.Image, x, y, width, height, radius float32, maskColor color.Color) {
	maskColor1 := maskColor
	maskColor2 := color.RGBA{A: 255, R: 0, G: 255, B: 0}     // Green
	maskColor3 := color.RGBA{A: 255, R: 0, G: 0, B: 255}     // Blue
	maskColor4 := color.RGBA{A: 255, R: 255, G: 0, B: 0}     // Red
	maskColor5 := color.RGBA{A: 255, R: 0, G: 255, B: 255}   // Cyan
	maskColor6 := color.RGBA{A: 255, R: 255, G: 0, B: 255}   // Magenta
	maskColor7 := color.RGBA{A: 255, R: 128, G: 128, B: 128} // Gray
	maskColor8 := color.RGBA{A: 255, R: 255, G: 165, B: 0}   // Orange

	//rounded circles
	vector.DrawFilledCircle(mask, x+radius, y+radius, radius, maskColor1, false)              //top left
	vector.DrawFilledCircle(mask, x+width-radius, y+radius, radius, maskColor2, false)        //top right
	vector.DrawFilledCircle(mask, x+radius, y+height-radius, radius, maskColor3, false)       //bottom left
	vector.DrawFilledCircle(mask, x+width-radius, y+height-radius, radius, maskColor4, false) //bottom right

	//edges                                        //right edge
	vector.DrawFilledRect(mask, x+radius, y, width-2*radius, radius, maskColor5, false)               //top edge
	vector.DrawFilledRect(mask, x+radius, y+height-radius, width-2*radius, radius, maskColor6, false) //bottom edge
	vector.DrawFilledRect(mask, x, y+radius, radius+width, height-2*radius, maskColor7, false)        //left edge
	vector.DrawFilledRect(mask, x+width-radius, y+radius, radius, height-2*radius, maskColor8, false) //right edge
}

func DrawBorderedBox(gs *state.GlobalState, screen *ebiten.Image) {

	rectWidth := 300
	rectHeight := 25
	borderWidth := 6
	horizontalPosition := float64(gs.ThirdQuarterX)
	verticalPosition := float64(gs.ThirdQuarterY)
	borderColor := color.RGBA{A: 255, R: 255, G: 255, B: 255} // White border
	rectColor := color.RGBA{A: 255, R: 96, G: 37, B: 37}      // Crimson red

	//Border Image
	border := ebiten.NewImage(rectWidth+borderWidth, rectHeight+borderWidth)
	vector.DrawFilledRect(border, 0, 0, float32(rectWidth+borderWidth), float32(rectHeight+borderWidth), borderColor, false)

	//Red Image
	rect := ebiten.NewImage(rectWidth, rectHeight)
	vector.DrawFilledRect(rect, 0, 0, float32(rectWidth), float32(rectHeight), rectColor, false)

	// Draw the border first
	borderOp := &ebiten.DrawImageOptions{}
	borderOp.GeoM.Translate(-float64(borderWidth)/2, -float64(borderWidth)/2) //move the width of the border
	borderOp.GeoM.Translate(-float64(rectWidth)/2, -float64(rectHeight)/2)    //center our origin
	borderOp.GeoM.Translate(horizontalPosition, verticalPosition)             //position
	screen.DrawImage(border, borderOp)

	// Draw the main rectangle
	rectOp := &ebiten.DrawImageOptions{}
	rectOp.GeoM.Translate(-float64(rectWidth)/2, -float64(rectHeight)/2) //center our origin
	rectOp.GeoM.Translate(horizontalPosition, verticalPosition)          //position
	screen.DrawImage(rect, rectOp)

}
