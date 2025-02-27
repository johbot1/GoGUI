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

type Game struct {
	selectedDice int // Holds the current dice selection (1d4, 1d6, etc.)
}

func (g *Game) Update() error {
	// Game logic will go here later
	// Handling input (button clicks)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		if mouseX < 150 && mouseY > 50 && mouseY < 90 { // 1d4 button
			g.selectedDice = 4
		} else if mouseX < 150 && mouseY > 100 && mouseY < 140 { // 1d6 button
			g.selectedDice = 6
		} else if mouseX < 150 && mouseY > 150 && mouseY < 190 { // 1d8 button
			g.selectedDice = 8
		} else if mouseX < 150 && mouseY > 200 && mouseY < 240 { // 1d10 button
			g.selectedDice = 10
		} else if mouseX < 150 && mouseY > 250 && mouseY < 290 { // 1d20 button
			g.selectedDice = 20
		} else if mouseX < 150 && mouseY > 300 && mouseY < 340 { // 1d100 button
			g.selectedDice = 100
		}
	}
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
	scaleFactor := float32(1.5)
	size *= scaleFactor

	switch sides {
	case 4: // d4 - Triangle
		vector.StrokeLine(screen, x, y+size, x+size/2, y, lineWidth, color, true)
		vector.StrokeLine(screen, x+size/2, y, x+size, y+size, lineWidth, color, true)
		vector.StrokeLine(screen, x+size, y+size, x, y+size, lineWidth, color, true)

	case 6: // d6 - Square
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
	// Screen dimensions
	screenWidth, screenHeight := screen.Size()
	// Draw dice type buttons on the left
	buttonWidth, buttonHeight := 120, 40
	buttonX := float32(20)
	//Dice sizing parameters
	diceX := float32(screenWidth)/2 - 180
	diceY := float32(screenHeight)/2 - 200
	diceSize := float32(screenWidth) * 0.3
	lineWidth := float32(screenWidth) * 0.01
	// Labels for the buttons (dice options)
	buttons := []string{"1d4", "1d6", "1d8", "1d10", "1d20", "1d100"}

	// Draw the buttons
	for i, buttonText := range buttons {
		buttonY := float32(50 + i*50)
		ebitenutil.DrawRect(screen, float64(buttonX), float64(buttonY), float64(buttonWidth), float64(buttonHeight), color.RGBA{0, 0, 0, 255})
		ebitenutil.DebugPrintAt(screen, buttonText, int(buttonX+20), int(buttonY+10))
	}
	//Title shown at the top of the screen
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Dice Roll Dice"), int(screenWidth/2-50), int(screenHeight-580))
	// Show the current dice selection (below buttons)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Selected Dice: 1d%d", g.selectedDice), int(screenWidth/2-50), int(screenHeight-50))

	DrawDiceShape(screen, diceX, diceY, diceSize, g.selectedDice, lineWidth, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set window size
	return 800, 600
}

func main() {
	// Initialize the game object
	game := &Game{}

	// Set the initial dice type (default to d20)
	game.selectedDice = 20

	// Start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
