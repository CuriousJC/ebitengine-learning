// embed.go
package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"image"
	_ "image/png"
	"log"
)

//go:embed spritesheet.png
var spritesheet_png []byte

//go:embed fire-ring.png
var firering_png []byte

// FONTS
//
//go:embed FiraSans-Regular.ttf
var firaSansRegular []byte

// LoadAssets returns a mapped set of images for the game
func LoadAssets() map[string]*ebiten.Image {
	assets := make(map[string]*ebiten.Image)

	assets["spritesheet_png"] = loadImage(spritesheet_png)
	assets["firering_png"] = loadImage(firering_png)

	return assets
}

// LoadFonts returns a mapped set of fonts for the game
func LoadFonts() map[string]*text.GoTextFaceSource {
	fonts := make(map[string]*text.GoTextFaceSource)

	fonts["firaSansRegular"] = loadFont(firaSansRegular)

	return fonts

}

// loadFont Function flip embedded font into GoTextFaceSource
func loadFont(data []byte) *text.GoTextFaceSource {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	return s
}

// loadImage Function flip embedded image into ebiten Image
func loadImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed to load image:", err)
	}
	return ebiten.NewImageFromImage(img)
}
