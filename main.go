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
	screenWidth  = 960
	screenHeight = 640
)

type Point struct {
	x, y float64
}

type Game struct {
	width, height      int
	p1, p2, p3, p4, p5 Point
}

func (g *Game) Rotate(angle float64) {
	var c Point
	c.x = g.p1.x*math.Cos(angle) - g.p1.y*math.Sin(angle)
	c.y = g.p1.x*math.Sin(angle) + g.p1.y*math.Cos(angle)
	g.p1.x, g.p1.y = c.x, c.y
	c.x = g.p2.x*math.Cos(angle) - g.p2.y*math.Sin(angle)
	c.y = g.p2.x*math.Sin(angle) + g.p2.y*math.Cos(angle)
	g.p2.x, g.p2.y = c.x, c.y
	c.x = g.p3.x*math.Cos(angle) - g.p3.y*math.Sin(angle)
	c.y = g.p3.x*math.Sin(angle) + g.p3.y*math.Cos(angle)
	g.p3.x, g.p3.y = c.x, c.y
	c.x = g.p4.x*math.Cos(angle) - g.p4.y*math.Sin(angle)
	c.y = g.p4.x*math.Sin(angle) + g.p4.y*math.Cos(angle)
	g.p4.x, g.p4.y = c.x, c.y
	c.x = g.p5.x*math.Cos(angle) - g.p5.y*math.Sin(angle)
	c.y = g.p5.x*math.Sin(angle) + g.p5.y*math.Cos(angle)
	g.p5.x, g.p5.y = c.x, c.y
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	g.Rotate(math.Pi / 180)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.p1.x+float64(g.width/2), g.p1.y+float64(g.height/2), g.p2.x+float64(g.width/2), g.p2.y+float64(g.height/2), color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.p2.x+float64(g.width/2), g.p2.y+float64(g.height/2), g.p3.x+float64(g.width/2), g.p3.y+float64(g.height/2), color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.p3.x+float64(g.width/2), g.p3.y+float64(g.height/2), g.p4.x+float64(g.width/2), g.p4.y+float64(g.height/2), color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.p4.x+float64(g.width/2), g.p4.y+float64(g.height/2), g.p5.x+float64(g.width/2), g.p5.y+float64(g.height/2), color.RGBA{255, 255, 255, 255})
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		p1:     Point{x: 0, y: 0},
		p2:     Point{x: 0, y: 200},
		p3:     Point{x: -75, y: 200},
		p4:     Point{x: -75, y: 160},
		p5:     Point{x: 0, y: 160},
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
