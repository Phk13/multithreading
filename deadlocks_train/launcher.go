package main

import (
	"log"
	"sync"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	. "github.com/phk13/multithreading/deadlocks_train/arbitrator"
	. "github.com/phk13/multithreading/deadlocks_train/common"
)

var (
	trains        [4]*Train
	intersections [4]*Intersection
)

const trainLength = 70

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTracks(screen)
	DrawIntersections(screen)
	DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return 320, 320
}

func main() {
	for i := 0; i < 4; i++ {
		trains[i] = &Train{Id: i, TrainLength: trainLength, Front: 0}
	}

	for i := 0; i < 4; i++ {
		intersections[i] = &Intersection{Id: i, Mutex: sync.Mutex{}, LockedBy: -1}
	}

	go MoveTrain(trains[0], 300, []*Crossing{{Position: 125, Intersection: intersections[0]},
		{Position: 175, Intersection: intersections[1]}})

	go MoveTrain(trains[1], 300, []*Crossing{{Position: 125, Intersection: intersections[1]},
		{Position: 175, Intersection: intersections[2]}})

	go MoveTrain(trains[2], 300, []*Crossing{{Position: 125, Intersection: intersections[2]},
		{Position: 175, Intersection: intersections[3]}})

	go MoveTrain(trains[3], 300, []*Crossing{{Position: 125, Intersection: intersections[3]},
		{Position: 175, Intersection: intersections[0]}})

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
