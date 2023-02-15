package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//---------------------------Declaration--------------------------------

const (
	sW = 640 //screen width
	sH = 480 //screen height
)

type Game struct {
	width, height     int
	angle             float64
	a, b, c, d, e, rp point
	//a-e - flag points, rp - rotation point
}

type point struct {
	x, y float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	//flag point rotation
	g.a.x, g.a.y = rotate(g.rp, g.a, g.angle)
	g.b.x, g.b.y = rotate(g.rp, g.b, g.angle)
	g.c.x, g.c.y = rotate(g.rp, g.c, g.angle)
	g.e.x, g.e.y = rotate(g.rp, g.e, g.angle)
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.b.x, g.b.y, color.RGBA{255, 255, 255, 255})
	//ebitenutil.DrawLine(screen, g.b.x, g.b.y, g.c.x, g.c.y, color.RGBA{255, 255, 255, 255})//not bug, but feature
	ebitenutil.DrawLine(screen, g.c.x, g.c.y, g.d.x, g.d.y, color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawLine(screen, g.a.x, g.a.y, g.e.x, g.e.y, color.RGBA{255, 255, 255, 255})
}

//-------------------------Functions----------------------------------

func rotate(rp, p point, angle float64) (newx, newy float64) {
	//p - point, rp - rotation point

	//moving point to top left corner
	p.x, p.y = p.x-rp.x, p.y-rp.y

	//rotating point
	newx = p.x*math.Cos(angle) - p.y*math.Sin(angle)
	newy = p.x*math.Sin(angle) + p.y*math.Cos(angle)

	//returning point that is moved on it's place
	return newx + rp.x, newy + rp.y
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {

	ebiten.SetWindowSize(sW, sH)
	ebiten.SetWindowTitle("Spinning Stuff!")
	ebiten.SetWindowResizable(true) //enablening window resizes

	//running game
	g := NewGame(sW, sH)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func NewGame(width, height int) *Game {

	angle := math.Pi / 2 //here should be somewhere 45-90. me sorry :]

	var a, b, c, d, e, rp point

	a.x, a.y = (sW / 2), (sH/2)-100
	b.x, b.y = (sW/2)+100, (sH/2)-100
	c.x, c.y = (sW/2)+100, (sH / 2)
	d.x, d.y = (sW / 2), (sH / 2)
	e.x, e.y = (sW / 2), (sH/2)+100

	rp.x, rp.y = (sW / 2), (sH / 2)

	//creating and returning game instance
	return &Game{sW, sH, angle, a, b, c, d, e, rp}

}
