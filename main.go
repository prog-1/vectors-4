package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Point is a struct for representing 2D vectors.
type Point struct {
	x, y float64
}

// Game is a game instance.
type Game struct {
	width, height      int
	p1, p2, p3, p4, p5 Point
}

// NewGame returns a new Game instance.
func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		p1:     Point{float64(width) / 2, float64(height) / 2},
		p2:     Point{430, 130},
		p3:     Point{490, 180},
		p4:     Point{460, 210},
		p5:     Point{400, 160},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func Rotate(p1 Point, p2 *Point) {
	var newpos Point
	newpos.x = (p2.x-p1.x)*math.Cos(math.Pi/180) - (p2.y-p1.y)*math.Sin(math.Pi/180)
	newpos.y = (p2.x-p1.x)*math.Sin(math.Pi/180) + (p2.y-p1.y)*math.Cos(math.Pi/180)
	p2.x, p2.y = newpos.x+p1.x, newpos.y+p1.y
}

// Update updates a game state.
func (g *Game) Update() error {
	Rotate(g.p1, &g.p2)
	Rotate(g.p1, &g.p3)
	Rotate(g.p1, &g.p4)
	Rotate(g.p1, &g.p5)
	return nil
}

// Draw renders a game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.p1.x, g.p1.y, g.p2.x, g.p2.y, color.White)
	ebitenutil.DrawLine(screen, g.p2.x, g.p2.y, g.p3.x, g.p3.y, color.White)
	ebitenutil.DrawLine(screen, g.p3.x, g.p3.y, g.p4.x, g.p4.y, color.White)
	ebitenutil.DrawLine(screen, g.p4.x, g.p4.y, g.p5.x, g.p5.y, color.White)
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
