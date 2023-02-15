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

func (l *line) draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, l.a.x, l.a.y, l.b.x, l.b.y, color.RGBA{200, 250, 200, 250})
}

func rotate(l *line, rad float64) *vector {
	x, y := l.b.x-l.a.x, l.b.y-l.a.y
	x, y = x*math.Cos(rad)-y*math.Sin(rad), y*math.Cos(rad)+x*math.Sin(rad)
	return &vector{x + l.a.x, y + l.a.y}
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type game struct {
	p []*vector
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	for i := range g.p {
		g.p[i] = rotate(&line{*g.p[0], *g.p[i]}, math.Pi/180)
	}
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	for i := 1; i < len(g.p); i++ {
		ebitenutil.DrawLine(screen, g.p[i-1].x, g.p[i-1].y, g.p[i].x, g.p[i].y, color.RGBA{200, 250, 200, 250})
	}
}

func NewGame() *game {
	return &game{
		[]*vector{
			{screenWidth / 2, screenHeight / 2},
			{screenWidth / 2, screenHeight / 4},
			{screenWidth / 4 * 3, screenHeight / 4},
			{screenWidth / 4 * 3, screenHeight / 8 * 3},
			{screenWidth / 2, screenHeight / 8 * 3},
		},
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
