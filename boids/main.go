package main

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green   = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeight + 1]int
	rwlock  = sync.RWMutex{}
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
