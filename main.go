package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	sWidth  = 600
	sHeight = 600
)

type point struct {
	x, y float64
}

type Game struct {
	width, height int
	a, b, c, d, e point
}

var col = color.RGBA{244, 212, 124, 255}

func DrawLine(img *ebiten.Image, x1, x2, y1, y2 int, col color.Color) {
	if math.Abs(float64(x2-x1)) >= math.Abs(float64(y2-y1)) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(x, int(y), col)
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), y, col)
		}
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		a:      point{300, 300},
		b:      point{300, 200},
		c:      point{350, 200},
		d:      point{300, 250},
		e:      point{350, 250},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	ang := math.Pi / 90

	x1, y1 := g.b.x, g.b.y
	x2, y2 := g.c.x, g.c.y
	x3, y3 := g.d.x, g.d.y
	x4, y4 := g.e.x, g.e.y

	g.b.x = ((x1-g.a.x)*math.Cos(ang) - (y1-g.a.y)*math.Sin(ang)) + g.a.x
	g.b.y = ((x1-g.a.x)*math.Sin(ang) + (y1-g.a.y)*math.Cos(ang)) + g.a.y

	g.c.x = ((x2-g.a.x)*math.Cos(ang) - (y2-g.a.y)*math.Sin(ang)) + g.a.x
	g.c.y = ((x2-g.a.x)*math.Sin(ang) + (y2-g.a.y)*math.Cos(ang)) + g.a.y

	g.d.x = ((x3-g.a.x)*math.Cos(ang) - (y3-g.a.y)*math.Sin(ang)) + g.a.x
	g.d.y = ((x3-g.a.x)*math.Sin(ang) + (y3-g.a.y)*math.Cos(ang)) + g.a.y

	g.e.x = ((x4-g.a.x)*math.Cos(ang) - (y4-g.a.y)*math.Sin(ang)) + g.a.x
	g.e.y = ((x4-g.a.x)*math.Sin(ang) + (y4-g.a.y)*math.Cos(ang)) + g.a.y

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLine(screen, int(g.a.x), int(g.b.x), int(g.a.y), int(g.b.y), col)
	DrawLine(screen, int(g.b.x), int(g.c.x), int(g.b.y), int(g.c.y), col)
	DrawLine(screen, int(g.c.x), int(g.e.x), int(g.c.y), int(g.e.y), col)
	DrawLine(screen, int(g.e.x), int(g.d.x), int(g.e.y), int(g.d.y), col)

	DrawLine(screen, int(g.b.x), int(g.e.x), int(g.b.y), int(g.e.y), col)
	DrawLine(screen, int(g.d.x), int(g.c.x), int(g.d.y), int(g.c.y), col)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(sWidth, sHeight)
	if err := ebiten.RunGame(NewGame(sWidth, sHeight)); err != nil {
		log.Fatal(err)
	}
}
