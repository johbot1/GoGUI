package main

import (
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"strconv"
)

// Globals
var rollResult int
var counter int

//go:embed assets/Draconis-JRw6B.ttf
var embeddedFont []byte
var gameFont font.Face

type Game struct {
	selectedDice       int        // Holds the current dice selection (1d4, 1d6, etc.)
	diceColor          color.RGBA // Holds the current color for the dice lines
	selectedMultiplier int        // Holds current multiplier for the number of dice
	multiplierClicked  bool       // Keep track of the mulitplier interaction
}

// NewGame initializes a new game with a default dice selection, color, and multiplier
func NewGame() *Game {
	// Set default dice color to white (or any color you prefer)
	return &Game{
		selectedDice:       20,              // Default to a 6-sided die
		diceColor:          buttonColors[3], // Default color is white
		selectedMultiplier: 1,               // Default amount of dice to roll
		multiplierClicked:  false,
	}
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

// Update updates the program depending on the actions of the user
func (g *Game) Update() error {
	// Handling mouse input for dice switching and color buttons
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		counter++
		log.Println(counter)
		mouseX, mouseY := ebiten.CursorPosition()
		// Check if the mouse click is within the Roll Dice button bounds
		if mouseX >= RollDiceXStart && mouseX <= RollDiceXEnd && mouseY >= RollDiceYStart && mouseY <= RollDiceYEnd {
			// Roll the dice when the button is clicked
			g.RollDiceAndDisplayResult()
		}
		//Handles Dice Switching
		g.DiceSwitchingMouseLogic(mouseX, mouseY)

		//Handles color Switching
		g.ColorSwitchingMouseLogic(mouseX, mouseY)

		if g.multiplierClicked {
			g.multiplierClicked = false //Ensure its click behavior reset every time an update happens
		} else {
			g.MultiplierSwitchingMouseLogic(mouseX, mouseY)

		}
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
	diceX := float32(screenWidth)/2 - DiceXYAdjustment
	diceY := float32(screenHeight)/2 - DiceXYAdjustment
	diceSize := float32(screenWidth) * 0.3
	lineWidth := float32(screenWidth) * 0.01

	// Draw the Dice Switching buttons
	for i, buttonText := range Buttons {
		buttonY := float32(ButtonPlacementModifier + i*ButtonPlacementModifier)
		vector.DrawFilledRect(screen, DiceSwitchingButtonXpos, buttonY,
			float32(DiceSwitchingButtonWidth), float32(DiceSwitchingButtonHeight), DiceSwitchingButtonColor, true)
		text.Draw(screen, buttonText, gameFont, DiceSwitchingButtonTitle, int(buttonY+25), color.White)

	}

	// Draw the decrement button (-)
	vector.DrawFilledRect(screen, DecrementButtonX, DiceCountButtonYStart, DecrementButtonWidth, DecrementButtonHeight, multiplierButtonColor, true)
	text.Draw(screen, "-", gameFont, int(DecrementButtonX+20), int(DiceCountButtonYStart+30), color.White)

	// Draw the increment button (+)
	vector.DrawFilledRect(screen, IncrementButtonX, DiceCountButtonYStart, IncrementButtonWidth, IncrementButtonHeight, multiplierButtonColor, true)
	text.Draw(screen, "+", gameFont, int(IncrementButtonX+20), int(DiceCountButtonYStart+30), color.White)

	// Draw the dice count display button
	vector.DrawFilledRect(screen, DiceCountDisplayButtonX, DiceCountDisplayButtonY, DiceCountDisplayWidth, DiceCountDisplayHeight, multiplierButtonColor, true)
	text.Draw(screen, strconv.Itoa(g.selectedMultiplier), gameFont, int(DiceCountDisplayButtonX+40), int(DiceCountDisplayButtonY+30), color.White)

	// Draw color buttons
	rightButtonX := float32(screenWidth - (DiceXYAdjustment - 40)) // Position on the right side
	for i, btnColor := range buttonColors {
		btnY := RightButtonYStart + float32(i)*(ColorButtonHeight+ColorButtonSpacing)
		vector.DrawFilledRect(screen, rightButtonX, btnY, ColorButtonWidth, ColorButtonHeight, btnColor, true)
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
	DrawDiceShape(screen, diceX, diceY, diceSize, g.selectedDice, lineWidth, g.diceColor)
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
	counter = 0
	//Load in font before running the game
	gameFont = LoadEmbeddedFont()

	// Initialize the game object
	game := NewGame()

	// Start the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
