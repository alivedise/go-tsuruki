package main

import (
	_ "image/png"
	"log"

	"github.com/alivedise/tsuruki/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(400, 600)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
