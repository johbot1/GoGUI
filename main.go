package main

import (
	"crypto/rand"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/big"
)

type Game struct{}

func (g *Game) Update() error {
	// Game logic will go here later
	return nil
}

// RollDice generates a cryptographically secure random number for the dice roll
func RollDice(sides int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(sides)))
	if err != nil {
		// Fallback: Return max value in case of error (should rarely happen)
		return sides
	}
	return int(n.Int64()) + 1
}

// DrawDiceShape dynamically creates different dice outlines
func DrawDiceShape(screen *ebiten.Image, x, y, size float32, sides int, lineWidth float32, color color.Color) {
	scaleFactor := float32(1.5) // Adjust this for larger dice (Try values like 2.0 or 3.0)
	size *= scaleFactor         // Scale up the dice size

	switch sides {
	case 3: // d3 - Triangle
		vector.StrokeLine(screen, x, y+size, x+size/2, y, lineWidth, color, true)
		vector.StrokeLine(screen, x+size/2, y, x+size, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+size, y+size, x, y+size, lineWidth, color, true)

	case 4: // d4 - Square
		vector.StrokeLine(screen, x, y, x+size, y, lineWidth, color, true)
		vector.StrokeLine(screen, x, y+size, x+size, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x, y, x, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+size, y, x+size, y+size, lineWidth, color, true)

	case 8: // d8 - Diamond
		vector.StrokeLine(screen, x+size/2, y, x, y+size/2, lineWidth, color, true)
		vector.StrokeLine(screen, x, y+size/2, x+size/2, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+size/2, y+size, x+size, y+size/2, lineWidth, color, true)
		vector.StrokeLine(screen, x+size, y+size/2, x+size/2, y, lineWidth, color, true)

	case 10: // d10 - Pentagon
		vector.StrokeLine(screen, x+size/2, y, x, y+size/3, lineWidth, color, true)
		vector.StrokeLine(screen, x, y+size/3, x+size/4, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+size/4, y+size, x+3*size/4, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+3*size/4, y+size, x+size, y+size/3, lineWidth, color, true)
		vector.StrokeLine(screen, x+size, y+size/3, x+size/2, y, lineWidth, color, true)

	case 20: // d20 - Hexagon
		vector.StrokeLine(screen, x+size/4, y, x+3*size/4, y, lineWidth, color, true)
		vector.StrokeLine(screen, x+3*size/4, y, x+size, y+size/2, lineWidth, color, true)
		vector.StrokeLine(screen, x+size, y+size/2, x+3*size/4, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+3*size/4, y+size, x+size/4, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+size/4, y+size, x, y+size/2, lineWidth, color, true)
		vector.StrokeLine(screen, x, y+size/2, x+size/4, y, lineWidth, color, true)

	case 100: // d100 - Circle
		//vector.DrawCircle(screen, x+size/2, y+size/2, size/2, color, true)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Get the screen width and height for centering the text
	screenWidth, screenHeight := screen.Size()

	// Title text
	title := "Dice Roll Dice"

	// Center the title
	ebitenutil.DebugPrintAt(screen, title, screenWidth/2-len(title)*4, 20)
	// Draw the dice (square outline in center)
	diceSize := 100 // Size of the dice square
	diceX := float32(screenWidth/2 - diceSize/2)
	diceY := float32(screenHeight/2 - diceSize/2)
	selectedDice := 20 // Change this value dynamically based on selection
	DrawDiceShape(screen, diceX, diceY, float32(diceSize), selectedDice, 3, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set the window size, 800x600 as a default
	return 800, 600
}

func main() {
	// Create the game object
	game := &Game{}
	diceType := 8 // Rolling a d20
	result := RollDice(diceType)
	fmt.Println("Rolled a d20:", result)

	// Initialize the game and set the window title to "{TITLE}"
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
