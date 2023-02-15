package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type vector struct {
	x, y int
}

type line struct {
	a, b vector
}

func (l *line) draw(screen *ebiten.Image) {
	DrawLine(screen, l.a.x, l.a.y, l.b.x, l.b.y, color.RGBA{1, 100, 100, 255})
}

func (l *line) rotate(rad float64) {
	x1, y1 := l.b.x-l.a.x, l.b.y-l.a.y
	x1, y1 = rotate(float64(x1), float64(y1), rad)
	l.b = vector{x1 + l.a.x, y1 + l.a.y}
}

func rotate(x, y, rad float64) (x1, y1 int) {
	x1 = int(x*math.Cos(rad) - y*math.Sin(rad))
	y1 = int(y*math.Cos(rad) + x*math.Sin(rad))
	return x1, y1
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

type game struct {
	l            line
	screenBuffer *ebiten.Image
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.rotate(math.Pi / 180)
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	g.l.draw(screen)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &game{}
	g.screenBuffer = ebiten.NewImage(screenWidth, screenHeight)
	g.l = line{vector{screenWidth / 2, screenHeight / 2}, vector{screenWidth / 4 * 3, screenHeight / 4}}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
func DrawLine(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	// abs(Dy) < abs(dx) | / abs(dx) => abs(Dy)/abs(Dx) < 1 == abs(k) < 1
	if abs(y2-y1) < abs(x2-x1) {
		if x1 > x2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		dx, dy := x2-x1, y2-y1
		dirY := 1
		// Dy < 0 => y2 - y1 < 0 => y1 > y2 => Growing downwards
		if dy < 0 {
			dirY = -1
			dy = -dy // For us to pretend that line is growing upwards
		}
		d := 2*dy - dx
		for x, y := x1, y1; x < x2; x++ {
			img.Set(x, y, c)
			if d >= 0 { // NE
				y += dirY
				d += dy - dx
			} else { // E
				d += dy
			}
		}
	} else {
		if y1 > y2 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		dx, dy := x2-x1, y2-y1
		dirX := 1
		if dx < 0 {
			dirX = -1
			dx = -dx
		}
		d := 2*dx - dy
		for x, y := x1, y1; y < y2; y++ {
			img.Set(x, y, c)
			if d >= 0 { // NE
				x += dirX
				d += dx - dy
			} else { // E
				d += dx
			}
		}
	}
}
