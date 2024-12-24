package main

import (
	"github.com/curiousjc/ebitengine-learning/assets"
	"github.com/curiousjc/ebitengine-learning/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Ebitengine Learning")

	g := game.NewGame()
	g.Assets = assets.LoadAssets()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
