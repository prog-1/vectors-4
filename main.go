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

type Point struct {
	x, y float64
}

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	pos1, pos2, pos3, pos4, pos5 Point
}

func (g *Game) Update() error {
	ang := math.Pi / 120

	g.pos2.x = ((g.pos2.x-g.pos1.x)*math.Cos(ang) - (g.pos2.y-g.pos1.y)*math.Sin(ang)) + g.pos1.x
	g.pos2.y = ((g.pos2.x-g.pos1.x)*math.Sin(ang) + (g.pos2.y-g.pos1.y)*math.Cos(ang)) + g.pos1.y

	g.pos3.x = ((g.pos3.x-g.pos1.x)*math.Cos(ang) - (g.pos3.y-g.pos1.y)*math.Sin(ang)) + g.pos1.x
	g.pos3.y = ((g.pos3.x-g.pos1.x)*math.Sin(ang) + (g.pos3.y-g.pos1.y)*math.Cos(ang)) + g.pos1.y

	g.pos4.x = ((g.pos4.x-g.pos1.x)*math.Cos(ang) - (g.pos4.y-g.pos1.y)*math.Sin(ang)) + g.pos1.x
	g.pos4.y = ((g.pos4.x-g.pos1.x)*math.Sin(ang) + (g.pos4.y-g.pos1.y)*math.Cos(ang)) + g.pos1.y

	g.pos5.x = ((g.pos5.x-g.pos1.x)*math.Cos(ang) - (g.pos5.y-g.pos1.y)*math.Sin(ang)) + g.pos1.x
	g.pos5.y = ((g.pos5.x-g.pos1.x)*math.Sin(ang) + (g.pos5.y-g.pos1.y)*math.Cos(ang)) + g.pos1.y

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.pos1.x, g.pos1.y, g.pos2.x, g.pos2.y, color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen, g.pos2.x, g.pos2.y, g.pos3.x, g.pos3.y, color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen, g.pos3.x, g.pos3.y, g.pos4.x, g.pos4.y, color.RGBA{R: 227, G: 76, B: 235, A: 255})
	ebitenutil.DrawLine(screen, g.pos4.x, g.pos4.y, g.pos5.x, g.pos5.y, color.RGBA{R: 227, G: 76, B: 235, A: 255})
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{pos1: Point{x: 320, y: 240},
		pos2: Point{x: 320, y: 160},
		pos3: Point{x: 380, y: 160},
		pos4: Point{x: 380, y: 200},
		pos5: Point{x: 320, y: 200}}); err != nil {
		log.Fatal(err)
	}
}
