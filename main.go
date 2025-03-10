package main

import (
	"kimond/gladiatext/internal"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &internal.Game{}

	ebiten.SetWindowSize(internal.ScreenWidth, internal.ScreenHeight)
	ebiten.SetWindowTitle("Turn-Based Combat")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
