package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/phk13/multithreading/deadlocks_train/common"
	"github.com/phk13/multithreading/deadlocks_train/deadlock"
)

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		DrawTracks(screen)
		DrawIntersections(screen)
		DrawTrains(screen)
	}
	return nil
}

func main() {
	if err := ebiten.Run(update, 320, 320, 3, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
