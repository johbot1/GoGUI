package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type Game struct {
	selectedDice int // Holds the current dice selection (1d4, 1d6, etc.)
}

func (g *Game) Update() error {
	// Game logic will go here later
	// Handling input (button clicks)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		g.DiceSwitchingMouseLogic(mouseX, mouseY)
	}
	return nil
}

// DrawDiceShape dynamically creates different dice outlines
func DrawDiceShape(screen *ebiten.Image, x, y, size float32, sides int, lineWidth float32, color color.Color) {
	scaledSize := size * ScaleFactor
	switch sides {
	case 4: // d4 - Triangle
		DrawTriangle(screen, x, y, scaledSize, lineWidth, color)
	case 6: // d6 - Square
		DrawSquare(screen, x, y, scaledSize, lineWidth, color)
	case 8: // d8 - Diamond
		DrawDiamond(screen, x, y, scaledSize, lineWidth, color)
	case 10: // d10 - Pentagon
		DrawPentagon(screen, x, y, scaledSize, lineWidth, color)
	case 20: // d20 - Hexagon
		DrawHexagon(screen, x, y, scaledSize, lineWidth, color)
	case 100: // d100 - Circle
		DrawCircle(screen, x, y, scaledSize, lineWidth, color)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Screen dimensions
	screenWidth, screenHeight := screen.Size()
	//Dice sizing parameters
	diceX := float32(screenWidth)/2 - 180
	diceY := float32(screenHeight)/2 - 200
	diceSize := float32(screenWidth) * 0.3
	lineWidth := float32(screenWidth) * 0.01

	// Draw the Dice Switching buttons
	for i, buttonText := range Buttons {
		buttonY := float32(50 + i*50)
		vector.DrawFilledRect(screen, DiceSwitchingButtonXpos, buttonY,
			float32(DiceSwitchingButtonWidth), float32(DiceSwitchingButtonHeight), DiceSwitchingButtonColor, true)
		ebitenutil.DebugPrintAt(screen, buttonText, DiceSwitchingButtonTitle, int(buttonY+10))
	}

	//Draw the Roll Button
	vector.DrawFilledRect(screen, 330, 235, 150, 100, DiceRollingButtonColor, true)
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

	fmt.Println("Rolled dice!", RollDice(20))

	// Start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
