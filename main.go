package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	// Game logic will go here later
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Get the screen width and height for centering the text
	screenWidth, screenHeight := screen.Size()

	// Title text
	title := "{TITLE}"

	// Calculate the width and height of the title text for centering
	// Using a smaller multiplier to simulate larger text
	charWidth := 16  // Default width of the character
	charHeight := 16 // Default height of the character

	// Simulate a larger font size by scaling the width and height of characters
	scaleFactor := 5 // Increase this value to make the text larger
	textWidth := len(title) * charWidth * scaleFactor
	textHeight := charHeight * scaleFactor

	// Center the title on the screen
	x := (screenWidth - textWidth) / 2
	y := (screenHeight - textHeight) / 2

	// Draw the title text in the center with the simulated larger font
	for i := 0; i < len(title); i++ {
		// Multiply each character's position to simulate enlargement
		ebitenutil.DebugPrintAt(screen, string(title[i]), x+(i*charWidth*scaleFactor), y)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set the window size, 800x600 as a default
	return 800, 600
}

func main() {
	// Create the game object
	game := &Game{}

	// Initialize the game and set the window title to "{TITLE}"
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
