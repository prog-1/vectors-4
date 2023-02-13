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

const (
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height int
	o, a, b, c, d Point
}

func Rotate(a *Point, b Point, angle float64) {
	a.x, a.y = (a.x-b.x)*math.Cos(angle)-(a.y-b.y)*math.Sin(angle)+b.x, (a.x-b.x)*math.Sin(angle)+(a.y-b.y)*math.Cos(angle)+b.y
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		o:      Point{200, 300},
		a:      Point{200, 150},
		b:      Point{300, 150},
		c:      Point{300, 200},
		d:      Point{200, 200},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	Rotate(&g.a, g.o, math.Pi/180)
	Rotate(&g.b, g.o, math.Pi/180)
	Rotate(&g.c, g.o, math.Pi/180)
	Rotate(&g.d, g.o, math.Pi/180)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.o.x, g.o.y, g.a.x, g.a.y, color.White)
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.White)
	ebitenutil.DrawLine(screen, g.b.x, g.b.y, g.c.x, g.c.y, color.White)
	ebitenutil.DrawLine(screen, g.c.x, g.c.y, g.d.x, g.d.y, color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
