package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type vector struct {
	x, y float64
}

type line struct {
	a, b vector
}

func (l *line) rotate(rad float64) {
	x1, y1 := rotate(l.b.x-l.a.x, l.b.y-l.a.y, rad)
	l.b = vector{x1 + l.a.x, y1 + l.a.y}
}

func rotate(x, y, rad float64) (x1, y1 float64) {
	x1 = x*math.Cos(rad) - y*math.Sin(rad)
	y1 = y*math.Cos(rad) + x*math.Sin(rad)
	return x1, y1
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type game struct {
	l line
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.rotate(math.Pi / 180)
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.l.a.x, g.l.a.y, g.l.b.x, g.l.b.y, color.RGBA{200, 250, 200, 250})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &game{}
	g.l = line{vector{screenWidth / 2, screenHeight / 2}, vector{screenWidth / 4 * 2, screenHeight / 4}}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
