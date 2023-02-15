package main

import (
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type (
	point struct {
		x, y float64
	}
	Game struct {
		p1, p2, p3, p4, p5 point
	}
)

const (
	winTitle            = "vectors-4"
	winWidth, winHeight = 500, 500
)

var cos = math.Cos(0.05)
var sin = math.Sin(0.05)
var c = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.p1.x, g.p1.y, g.p2.x, g.p2.y, c)
	ebitenutil.DrawLine(screen, g.p2.x, g.p2.y, g.p3.x, g.p3.y, c)
	ebitenutil.DrawLine(screen, g.p3.x, g.p3.y, g.p4.x, g.p4.y, c)
	ebitenutil.DrawLine(screen, g.p4.x, g.p4.y, g.p5.x, g.p5.y, c)
}

func (g *Game) rotation() {
	x, y := g.p1.x, g.p1.y
	g.p1.x = x*cos - y*sin + (winWidth/2 - winWidth/2*cos + winHeight/2*sin)
	g.p1.y = x*sin + y*cos + (-winWidth/2*sin + winHeight/2 - winHeight/2*cos)
	x, y = g.p2.x, g.p2.y
	g.p2.x = x*cos - y*sin + (winWidth/2 - winWidth/2*cos + winHeight/2*sin)
	g.p2.y = x*sin + y*cos + (-winWidth/2*sin + winHeight/2 - winHeight/2*cos)
	x, y = g.p3.x, g.p3.y
	g.p3.x = x*cos - y*sin + (winWidth/2 - winWidth/2*cos + winHeight/2*sin)
	g.p3.y = x*sin + y*cos + (-winWidth/2*sin + winHeight/2 - winHeight/2*cos)
	x, y = g.p4.x, g.p4.y
	g.p4.x = x*cos - y*sin + (winWidth/2 - winWidth/2*cos + winHeight/2*sin)
	g.p4.y = x*sin + y*cos + (-winWidth/2*sin + winHeight/2 - winHeight/2*cos)
	x, y = g.p5.x, g.p5.y
	g.p5.x = x*cos - y*sin + (winWidth/2 - winWidth/2*cos + winHeight/2*sin)
	g.p5.y = x*sin + y*cos + (-winWidth/2*sin + winHeight/2 - winHeight/2*cos)

}
func (g *Game) Update() error {
	g.rotation()
	return nil

}

func (g *Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	if err := ebiten.RunGame(&Game{p5: point{x: 250, y: 350}, p1: point{x: 250, y: 250}, p2: point{x: 250, y: 450}, p3: point{x: 320, y: 450}, p4: point{x: 320, y: 350}}); err != nil {
		log.Fatal(err)
	}
}
