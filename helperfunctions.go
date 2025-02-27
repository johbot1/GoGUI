package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func DrawTriangle(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x, y+size, x+size/2, y, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/2, y, x+size, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size, x, y+size, lineWidth, color, true)
}

func drawSquare(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x, y, x+size, y, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size, x+size, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x, y, x, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y, x+size, y+size, lineWidth, color, true)
}

func drawDiamond(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x+size/2, y, x, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size/2, x+size/2, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/2, y+size, x+size, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size/2, x+size/2, y, lineWidth, color, true)
}

func drawPentagon(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x+size/2, y, x, y+size/3, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size/3, x+size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/4, y+size, x+3*size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+3*size/4, y+size, x+size, y+size/3, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size/3, x+size/2, y, lineWidth, color, true)
}

func drawHexagon(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x+size/4, y, x+3*size/4, y, lineWidth, color, true)
	vector.StrokeLine(screen, x+3*size/4, y, x+size, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size/2, x+3*size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+3*size/4, y+size, x+size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/4, y+size, x, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size/2, x+size/4, y, lineWidth, color, true)
}
