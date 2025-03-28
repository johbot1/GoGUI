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
	"math/rand"
	"strconv"
	"time"
)

// Globals
var rollResult int
var counter int

//go:embed assets/Draconis-JRw6B.ttf
var embeddedFont []byte
var gameFont font.Face

// Global timer variable
var lastActionTime time.Time

type Game struct {
	selectedDice       int          // Holds the current dice selection (1d4, 1d6, etc.)
	diceColor          color.RGBA   // Holds the current color for the dice lines
	multiplierClicked  bool         // Keep track of the mulitplier interaction
	selectedMultiplier int          // Holds current multiplier for the number of dice
	rollResults        [maxDice]int // Array for multiple dice
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

// Initialize the rollResults array
func (g *Game) Initialize() {
	// Fill the rollResults array with 0s or reset it
	for i := range g.rollResults {
		g.rollResults[i] = 0
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

// Function to load a font face with a specific size
func LoadFontWithSize(fontData []byte, size float64) font.Face {
	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size, // Size adjustment for roll result text
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
		//counter++
		//log.Println(counter)
		mouseX, mouseY := ebiten.CursorPosition()

		// Convert mouseX, mouseY to float32 for comparison
		fMouseX := float32(mouseX)
		fMouseY := float32(mouseY)

		// Get the current time
		currentTime := time.Now()

		// Check if the "+" button is clicked and if enough time has passed
		if fMouseX >= IncrementButtonX && fMouseX <= IncrementButtonX+IncrementButtonWidth &&
			fMouseY >= DiceCountButtonYStart && fMouseY <= DiceCountButtonYStart+IncrementButtonHeight {
			if currentTime.Sub(lastActionTime) > time.Millisecond*500 { // 500ms threshold
				// Increment only if the current multiplier is less than 5
				if g.selectedMultiplier < 5 {
					g.selectedMultiplier++
				}
				lastActionTime = currentTime
			}
		}

		// Check if the "-" button is clicked and if enough time has passed
		if fMouseX >= DecrementButtonX && fMouseX <= DecrementButtonX+DecrementButtonWidth &&
			fMouseY >= DiceCountButtonYStart && fMouseY <= DiceCountButtonYStart+DecrementButtonHeight {
			if currentTime.Sub(lastActionTime) > time.Millisecond*500 { // 500ms threshold
				// Allow the decrement and reset the timer
				if g.selectedMultiplier > 1 { // Ensure it doesn't go below 1
					g.selectedMultiplier--
				}
				lastActionTime = currentTime
			}
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

	// Draw the result of the dice roll (with larger font)
	// Load the larger font just for this text
	largeFont := LoadFontWithSize(embeddedFont, 100) // Adjust size (72) as needed

	// Convert number to string and measure its width
	rollStr := fmt.Sprintf("%d", g.rollResults[0])
	bounds := text.BoundString(largeFont, rollStr)
	textWidth := bounds.Dx()  // Get width of rendered text
	textHeight := bounds.Dy() // Get height of rendered text

	// Center text on the screen
	textX := screenWidth/2 - textWidth/2
	textY := screenHeight/2 + textHeight/4 // Adjust Y slightly for centering

	// Draw the roll number at the center using the larger font
	text.Draw(screen, rollStr, largeFont, textX, textY, color.White)

	// Draw additional dice numbers (based on multiplier)
	// Calculate the position for the additional dice numbers (below the main die)
	additionalDiceX := diceX
	additionalDiceY := diceY + diceSize + MultipleDiceCountSpacingModifier // 20px spacing between the main die and additional dice numbers

	// Draw the additional dice numbers
	for i := 1; i < g.selectedMultiplier; i++ {
		// Each number will be spaced horizontally by 20px from the previous one
		diceNumberStr := fmt.Sprintf("%d", i) // Use the index as the dice roll number
		bounds := text.BoundString(gameFont, diceNumberStr)
		textWidth := bounds.Dx() // Get width of rendered text

		// Center the additional dice numbers below the main die
		textX := additionalDiceX + float32(i)*(float32(textWidth)+MultipleDiceCountXModifier) // Spacing between numbers
		textY := additionalDiceY + MultipleDiceCountYModifier

		// Draw the number as text
		text.Draw(screen, diceNumberStr, gameFont, int(textX), int(textY), color.White)
	}

	// Draw the selected dice
	DrawDiceShape(screen, diceX, diceY, diceSize, g.selectedDice, lineWidth, g.diceColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Set window size
	return 800, 600
}

// This function handles rolling the dice and storing the results
func (g *Game) RollDiceAndDisplayResult() {
	for i := 0; i < g.selectedMultiplier; i++ {
		// Roll the dice and store the result in the rollResults array
		roll := rand.Intn(20) + 1 // Example: Roll a 20-sided die
		g.rollResults[i] = roll   // Store the result at index i
	}

	// Print statements to debug and confirm the dice roll results
	for i := 0; i < g.selectedMultiplier; i++ {
		fmt.Printf("Dice %d Roll: %d\n", i+1, g.rollResults[i])
	}
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
