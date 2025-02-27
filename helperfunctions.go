package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

// DrawTriangle Draws a triangle given 3 distinct points, and a specified color
func DrawTriangle(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x, y+size, x+size/2, y, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/2, y, x+size, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size, x, y+size, lineWidth, color, true)
}

// DrawSquare Draws a square given 4 distinct points, and a specified color
func DrawSquare(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x, y, x+size, y, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size, x+size, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x, y, x, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y, x+size, y+size, lineWidth, color, true)
}

// DrawDiamond Draws a diamond given 4 distinct points, and a specified color
func DrawDiamond(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x+size/2, y, x, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size/2, x+size/2, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/2, y+size, x+size, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size/2, x+size/2, y, lineWidth, color, true)
}

// DrawPentagon Draws a pentagon given 5 distinct points, and a specified color
func DrawPentagon(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x+size/2, y, x, y+size/3, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size/3, x+size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/4, y+size, x+3*size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+3*size/4, y+size, x+size, y+size/3, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size/3, x+size/2, y, lineWidth, color, true)
}

// DrawHexagon Draws a hexagon given 6 distinct points, and a specified color
func DrawHexagon(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeLine(screen, x+size/4, y, x+3*size/4, y, lineWidth, color, true)
	vector.StrokeLine(screen, x+3*size/4, y, x+size, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x+size, y+size/2, x+3*size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+3*size/4, y+size, x+size/4, y+size, lineWidth, color, true)
	vector.StrokeLine(screen, x+size/4, y+size, x, y+size/2, lineWidth, color, true)
	vector.StrokeLine(screen, x, y+size/2, x+size/4, y, lineWidth, color, true)
}

func DrawCircle(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeCircle(screen, x*CircleDrawingXModifier, y*CircleDrawingYModifier, size/2, lineWidth, color, true)
}

// DiceSwitchingMouseLogic Processes the mouse input logic for whenever the user clicks one of the buttons
func (g *Game) DiceSwitchingMouseLogic(mouseX, mouseY int) {
	if mouseX < 150 {
		switch {
		case mouseY > 50 && mouseY < 90:
			g.selectedDice = 4
		case mouseY > 100 && mouseY < 140:
			g.selectedDice = 6
		case mouseY > 150 && mouseY < 190:
			g.selectedDice = 8
		case mouseY > 200 && mouseY < 240:
			g.selectedDice = 10
		case mouseY > 250 && mouseY < 290:
			g.selectedDice = 20
		case mouseY > 300 && mouseY < 340:
			g.selectedDice = 100
		}
	}
}
