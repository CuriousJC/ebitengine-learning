package main

import (
	"github.com/curiousjc/ebitengine-learning/assets"
	"github.com/curiousjc/ebitengine-learning/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 1280
	screenHeight = 960
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ebitengine Learning")

	//Create our Game instance
	g := game.NewGame()
	g.GlobalState.ActiveDebug = true

	//Load assets into memory one time at startup
	g.GlobalState.Assets = assets.LoadAssets()
	g.GlobalState.Fonts = assets.LoadFonts()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
