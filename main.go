package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 960
	screenHeight = 640
)

type Point struct{
	x, y float64
}

type Game struct{
	width, height int
	p1, p2 Point
}

func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(x2-x1) <= math.Abs(y2-y1) {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	} else {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(int(x), int(y), c)
		}
	}
}

func (g *Game) Rotate(angle float64){
	g.p1.x = g.p1.x*math.Cos(angle) - g.p1.y*math.Sin(angle)
	g.p1.y = g.p1.x*math.Sin(angle) + g.p1.y*math.Cos(angle)
	g.p2.x = g.p2.x*math.Cos(angle) - g.p2.y*math.Sin(angle)
	g.p2.y = g.p2.x*math.Sin(angle) + g.p2.y*math.Cos(angle)
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update() error {
	g.Rotate(math.Pi /100)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, float64(g.width/2), float64(g.height/2), 100, color.RGBA{255, 0, 0, 50})
	DrawLineDDA(screen, g.p1.x+float64(g.width/2), g.p1.y+float64(g.height/2), g.p2.x+float64(g.width/2), g.p2.y+float64(g.height/2), color.RGBA{255, 255, 255, 255})
	// ebitenutil.DrawLine(screen, g.p1.x+float64(g.width/2), g.p1.y+float64(g.height/2), g.p2.x+float64(g.width/2), g.p2.y+float64(g.height/2), color.RGBA{255, 255, 255, 255})
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		p1: Point{x:0, y:0},
		p2: Point{x:100, y:100},
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