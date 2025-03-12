package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Globals
var rollResult int

//go:embed assets/Draconis-JRw6B.ttf
var embeddedFont []byte
var gameFont font.Face

type Game struct {
	selectedDice int        // Holds the current dice selection (1d4, 1d6, etc.)
	diceColor    color.RGBA // Holds the current color for the dice lines
}

// LoadEmbeddedFont loads the font from the embedded binary
func LoadEmbeddedFont() font.Face {
	tt, err := opentype.Parse(embeddedFont)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    FontSize, // Adjust font size
		DPI:     FontDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return face
}

func (g *Game) Update() error {
	// Handling mouse input for dice switching and color buttons
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		// Check if the mouse click is within the Roll Dice button bounds
		if mouseX >= 330 && mouseX <= 480 && mouseY >= 235 && mouseY <= 335 {
			// Roll the dice when the button is clicked
			g.RollDiceAndDisplayResult()
		}
		//Handles Dice Switching
		g.DiceSwitchingMouseLogic(mouseX, mouseY)

		//Handles color Switching
		g.ColorSwitchingMouseLogic(mouseX, mouseY)
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
	diceY := float32(screenHeight)/2 - 180
	diceSize := float32(screenWidth) * 0.3
	lineWidth := float32(screenWidth) * 0.01

	// Draw the Dice Switching buttons
	for i, buttonText := range Buttons {
		buttonY := float32(ButtonPlacementModifier + i*ButtonPlacementModifier)
		vector.DrawFilledRect(screen, DiceSwitchingButtonXpos, buttonY,
			float32(DiceSwitchingButtonWidth), float32(DiceSwitchingButtonHeight), DiceSwitchingButtonColor, true)
		text.Draw(screen, buttonText, gameFont, DiceSwitchingButtonTitle, int(buttonY+25), color.White)

	}

	// Draw multiplier buttons (x1, x2, x3, x4)
	newButtons := []string{"x1", "x2", "x3", "x4"}
	for i, btnLabel := range newButtons {
		btnY := MultiplierButtonYStart + float32(i)*MultiplierButtonSpacing
		// Draw the button background
		vector.DrawFilledRect(screen, MultiplierButtonX, btnY, MultiplierButtonWidth, MultiplierButtonHeight, multiplierButtonColor, true)
		// Draw the button label; adjust the x/y offsets as needed for centering
		text.Draw(screen, btnLabel, gameFont, int(MultiplierButtonX+40), int(btnY+28), color.White)
	}

	// Draw color buttons
	rightButtonX := float32(screenWidth - 140) // Position on the right side

	// Draw the right-side buttons
	for i, btnColor := range buttonColors {
		btnY := RightButtonYStart + float32(i)*(RightButtonHeight+RightButtonSpacing)
		vector.DrawFilledRect(screen, rightButtonX, btnY, RightButtonWidth, RightButtonHeight, btnColor, true)

		// Draw an 'X' on the last button (white button)
		if i == 3 {
			text.Draw(screen, "X", gameFont, int(rightButtonX+40), int(btnY+60), color.RGBA{255, 0, 0, 255}) // Red 'X'
		}
	}

	// Draw the "Roll" button with larger text
	vector.DrawFilledRect(screen, RollButtonXpos, RollButtonYpos, RollButtonWidth, RollButtonHeight, DiceRollingButtonColor, true)
	text.Draw(screen, "Roll", gameFont, RollButtonTextXpos, RollButtonTextYpos, color.White)

	//Title shown at the top of the screen
	text.Draw(screen, "Dice Roll Dice", gameFont, screenWidth/2-75, screenHeight-570, color.White)

	// Show the current dice selection (below buttons)
	text.Draw(screen, fmt.Sprintf("Selected Dice: 1d%d", g.selectedDice), gameFont, int(screenWidth/2-100), int(screenHeight-530), color.White)

	// Show the result of the dice roll (if available)
	if rollResult > 0 {
		// Display the roll result
		text.Draw(screen, fmt.Sprintf("Roll Result: %d", rollResult), gameFont, screenWidth/2-75, screenHeight/2+250, color.White)

	}
	// Draw the selected dice
	DrawDiceShape(screen, diceX, diceY, diceSize, g.selectedDice, lineWidth, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set window size
	return 800, 600
}

func (g *Game) RollDiceAndDisplayResult() {
	// Roll the selected dice
	rollResult = RollDice(g.selectedDice)
}

func main() {
	//Load in font before running the game
	gameFont = LoadEmbeddedFont()

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
