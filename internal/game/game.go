package game

import (
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	frameOX      = 0
	frameOY      = 0
	frameWidth   = 300
	frameHeight  = 300
	frameCount   = 1
)

type Game struct {
	count  int
	mouseX int
	mouseY int
	Assets map[string]*ebiten.Image // Store images as a map in the Game struct
}

func NewGame() *Game {
	return &Game{
		Assets: make(map[string]*ebiten.Image), // Initialize the assets map
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
	screen.DrawImage(g.Assets["firering_png"].SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("frame sx: %d, sy: %d", sx, sy))

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mouse: %d, %d", g.mouseX, g.mouseY), 25, 25)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {

	return screenWidth, screenHeight
}
