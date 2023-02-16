package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Point struct {
	x, y float64
}

type Figure struct {
	lines []Line
}

func (f *Figure) RotateWithCenterPoint(angle, x, y float64) {
	for i := range f.lines {
		f.lines[i].RotateWithCenterPoint(angle, x, y)
	}
}

func (f *Figure) Draw(screen *ebiten.Image) {
	for i := range f.lines {
		f.lines[i].Draw(screen)
	}
}

type Line struct {
	start, end Point
}

func (l *Line) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, l.start.x, l.start.y, l.end.x, l.end.y, color.RGBA{255, 0, 0, 255})
	// ebitenutil.Dra
}

type Game struct {
	f    Figure
	last time.Time
	// BackBuffer *ebiten.Image
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

func (l *Line) Rotate(angle float64) {
	l.start.Rotate(angle)
	l.end.Rotate(angle)
}

func (l *Line) RotateWithCenterPoint(angle, x, y float64) {
	l.start.RotateWithCenterPoint(angle, x, y)
	l.end.RotateWithCenterPoint(angle, x, y)
}

func (p *Point) RotateWithCenterPoint(angle, x, y float64) {
	tmp := p.x
	p.x = ((p.x-x)*math.Cos(angle) - (p.y-y)*math.Sin(angle)) + x
	p.y = ((tmp-x)*math.Sin(angle) + (p.y-y)*math.Cos(angle)) + y
}

func (p *Point) Rotate(angle float64) {
	tmp := p.x
	p.x = p.x*math.Cos(angle) - p.y*math.Sin(angle)
	p.y = tmp*math.Sin(angle) + p.y*math.Cos(angle)
}

func (g *Game) Update() error {
	// if g.last.Sub(time.Now()).Milliseconds() > -50 {
	// 	return nil
	// }
	g.f.RotateWithCenterPoint(math.Pi/360, 100, 100)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.f.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	g := &Game{Figure{[]Line{Line{Point{100, 10}, Point{100, 100}}, Line{Point{100, 10}, Point{200, 10}}, Line{Point{200, 10}, Point{200, 55}}, Line{Point{200, 55}, Point{100, 55}}}}, time.Now()}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world")
}
