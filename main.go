package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

var (
	text = "Click the button!"
)

type Game struct{}

func (g *Game) Update() error {
	// Check if the mouse was clicked
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		// Check if the click is within the button area
		if mouseX > 100 && mouseX < 300 && mouseY > 200 && mouseY < 250 {
			text = "Button clicked!"
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Set background color (light gray in this case)
	screen.Fill(color.RGBA{R: 200, G: 200, B: 200, A: 255})

	// Draw the button (a simple rectangle)
	buttonColor := color.RGBA{R: 51, G: 153, B: 229, A: 255}
	ebitenutil.DrawRect(screen, 100, 200, 200, 50, buttonColor)

	// Draw button text
	ebitenutil.DebugPrintAt(screen, "Click me!", 150, 215)

	// Draw the text at the top
	ebitenutil.DebugPrintAt(screen, text, 150, 100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the screen dimensions (800x600)
	return 800, 600
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
