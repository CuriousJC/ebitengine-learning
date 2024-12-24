package game

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 640
	screenHeight = 480
	frameOX      = 510
	frameOY      = 640
	frameWidth   = 65
	frameHeight  = 65
	frameCount   = 8
)

var (
	elementImage *ebiten.Image
)

type Game struct {
	count  int
	mouseX int
	mouseY int
}

func NewGame() *Game {
	return &Game{
		count:  0,
		mouseX: 0,
		mouseY: 0,
	}
}

func (g *Game) Update() error {
	g.count++
	g.mouseX, g.mouseY = ebiten.CursorPosition()
	return nil
}

// random comment
func (g *Game) Draw(screen *ebiten.Image) {
	// Step 1: Compute the frame
	//i := (g.count / 50) % frameCount
	//sx, sy := frameOX+i*frameWidth, frameOY

	sx, sy := frameOX, frameOY

	// Step 2: Set up transformations
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2) // Center the origin
	op.GeoM.Translate(float64(g.mouseX), float64(g.mouseY))            // Follow the mouse
	//op.GeoM.Scale(0.5, 0.5)                                 // Scale down

	// Step 3: Render the image
	screen.DrawImage(elementImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("frame sx: %d, sy: %d", sx, sy))

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mouse: %d, %d", g.mouseX, g.mouseY), 25, 25)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	// Decode an image from the image file's byte slice.
	//TODO: switch this to something embedded in our own source code
	img, _, err := image.Decode(bytes.NewReader(images.Spritesheet_png))
	if err != nil {
		log.Fatal(err)
	}
	elementImage = ebiten.NewImageFromImage(img)

	return screenWidth, screenHeight
}
