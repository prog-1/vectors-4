// package main

// import (
// 	"image/color"
// 	_ "image/png"
// 	"log"
// 	"math"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// )

// type (
// 	point struct {
// 		x, y float64
// 	}

// 	Game struct {
// 		sun   point
// 		earth point
// 		moon  point
// 	}
// )

// const (
// 	winTitle            = "planets"
// 	winWidth, winHeight = 500, 500
// )

// var img *ebiten.Image

// func init() {
// 	var err error
// 	img, _, err = ebitenutil.NewImageFromFile("earth.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	ebitenutil.DrawCircle(screen, g.sun.x, g.sun.y, 30, color.RGBA{0xff, 0xff, 0, 0xff})
// 	ebitenutil.DrawCircle(screen, g.earth.x, g.earth.y, 15, color.RGBA{0, 0xff, 0, 0xff})
// 	ebitenutil.DrawCircle(screen, g.moon.x, g.moon.y, 5, color.RGBA{0xff, 0xff, 0xff, 0xff})
// }

// func (g *Game) rotation() {
// 	x, y := g.earth.x, g.earth.y
// 	g.earth.x = x*math.Cos(0.01) - y*math.Sin(0.01) + (g.sun.x - g.sun.x*math.Cos(0.01) + g.sun.y*math.Sin(0.01))
// 	g.earth.y = x*math.Sin(0.01) + y*math.Cos(0.01) + (-g.sun.x*math.Sin(0.01) + g.sun.y - g.sun.y*math.Cos(0.01))
// 	x, y = g.moon.x, g.moon.y
// 	g.moon.x = x*math.Cos(0.01) - y*math.Sin(0.01) + (g.sun.x - g.sun.x*math.Cos(0.01) + g.sun.y*math.Sin(0.01))
// 	g.moon.y = x*math.Sin(0.01) + y*math.Cos(0.01) + (-g.sun.x*math.Sin(0.01) + g.sun.y - g.sun.y*math.Cos(0.01))
// 	x, y = g.moon.x, g.moon.y
// 	g.moon.x = x*math.Cos(-0.03) - y*math.Sin(-0.03) + (g.earth.x - g.earth.x*math.Cos(-0.03) + g.earth.y*math.Sin(-0.03))
// 	g.moon.y = x*math.Sin(-0.03) + y*math.Cos(-0.03) + (-g.earth.x*math.Sin(-0.03) + g.earth.y - g.earth.y*math.Cos(-0.03))
// }
// func (g *Game) Update() error {
// 	g.rotation()
// 	return nil

// }

// func (g *Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

//	func main() {
//		ebiten.SetWindowTitle(winTitle)
//		ebiten.SetWindowSize(winWidth, winHeight)
//		if err := ebiten.RunGame(&Game{sun: point{x: 250, y: 250}, earth: point{x: 100, y: 250}, moon: point{x: 130, y: 250}}); err != nil {
//			log.Fatal(err)
//		}
//	}
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
