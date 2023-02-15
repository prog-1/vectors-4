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
	width, height    int
	angle float64
	l1 *line
}

type point struct {
	x, y float64
}

type line struct {
	a point//starting point
	b point//ending point
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	g.l1.b.x, g.l1.b.y = rotate(g.l1, g.angle)//rotating first line ending point
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, g.l1.a.x, g.l1.a.y, g.l1.b.x, g.l1.b.y, color.RGBA{255, 255, 255, 255})
}

//-------------------------Functions----------------------------------

func rotate(l *line, angle float64) (newx,newy float64){
	
	//moving ending point to top left corner
	l.b.x, l.b.y = l.b.x-l.a.x, l.b.y-l.a.y

	//rotating ending point
	newx = l.b.x*math.Cos(angle) - l.b.y*math.Sin(angle)
	newy = l.b.x*math.Sin(angle) + l.b.y*math.Cos(angle)

	//returning ending point that is moved on it's place
	return newx+l.a.x,newy+l.a.y
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

	//All pre-declared stuff is stored here

	angle := math.Pi/360
	
	var l1 line
	l1.a.x, l1.a.y = (sW/2)-200, (sH/2)
	l1.b.x, l1.b.y = (sW/2)-200, (sH/2)-100

	//creating and returning game instance
	return &Game{sW, sH, angle, &l1}

}