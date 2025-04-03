package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

// DrawD4Dice Draws a triangle given 3 distinct points, and a specified color; d4
func DrawD4Dice(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	// Calculate triangle vertices for an equilateral triangle
	x1, y1 := x+size/2, y    // Top vertex
	x2, y2 := x, y+size      // Bottom-left
	x3, y3 := x+size, y+size // Bottom-right

	// Draw triangle edges
	vector.StrokeLine(screen, x1, y1, x2, y2, lineWidth, color, true)
	vector.StrokeLine(screen, x2, y2, x3, y3, lineWidth, color, true)
	vector.StrokeLine(screen, x3, y3, x1, y1, lineWidth, color, true)
}

// DrawD6Dice Draws a square given 4 distinct points, and a specified color; d6
func DrawD6Dice(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	// Calculate square corners
	x1, y1 := x, y           // Top-left
	x2, y2 := x+size, y      // Top-right
	x3, y3 := x, y+size      // Bottom-left
	x4, y4 := x+size, y+size // Bottom-right

	// Draw square edges
	vector.StrokeLine(screen, x1, y1, x2, y2, lineWidth, color, true)
	vector.StrokeLine(screen, x2, y2, x4, y4, lineWidth, color, true)
	vector.StrokeLine(screen, x4, y4, x3, y3, lineWidth, color, true)
	vector.StrokeLine(screen, x3, y3, x1, y1, lineWidth, color, true)
}

// DrawD8Dice Draws a diamond given 4 distinct points, and a specified color; d8
func DrawD8Dice(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	// Calculate diamond points
	x1, y1 := x+size/2, y      // Top
	x2, y2 := x, y+size/2      // Left
	x3, y3 := x+size/2, y+size // Bottom
	x4, y4 := x+size, y+size/2 // Right

	// Draw diamond edges
	vector.StrokeLine(screen, x1, y1, x2, y2, lineWidth, color, true)
	vector.StrokeLine(screen, x2, y2, x3, y3, lineWidth, color, true)
	vector.StrokeLine(screen, x3, y3, x4, y4, lineWidth, color, true)
	vector.StrokeLine(screen, x4, y4, x1, y1, lineWidth, color, true)
}

// DrawD10Dice Draws a pentagon given 5 distinct points, and a specified color; d10
func DrawD10Dice(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	// Calculate pentagon points
	x1, y1 := x+size/2, y        // Top
	x2, y2 := x, y+size*0.4      // Upper Left
	x3, y3 := x+size*0.2, y+size // Lower Left
	x4, y4 := x+size*0.8, y+size // Lower Right
	x5, y5 := x+size, y+size*0.4 // Upper Right

	// Draw pentagon edges
	vector.StrokeLine(screen, x1, y1, x2, y2, lineWidth, color, true)
	vector.StrokeLine(screen, x2, y2, x3, y3, lineWidth, color, true)
	vector.StrokeLine(screen, x3, y3, x4, y4, lineWidth, color, true)
	vector.StrokeLine(screen, x4, y4, x5, y5, lineWidth, color, true)
	vector.StrokeLine(screen, x5, y5, x1, y1, lineWidth, color, true)
}

// DrawD20Dice Draws a hexagon given 6 distinct points, and a specified color; d20
func DrawD20Dice(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	// Calculate hexagon points
	x1, y1 := x+size*0.25, y      // Top Left
	x2, y2 := x+size*0.75, y      // Top Right
	x3, y3 := x+size, y+size*0.5  // Right
	x4, y4 := x+size*0.75, y+size // Bottom Right
	x5, y5 := x+size*0.25, y+size // Bottom Left
	x6, y6 := x, y+size*0.5       // Left

	// Draw hexagon edges
	vector.StrokeLine(screen, x1, y1, x2, y2, lineWidth, color, true)
	vector.StrokeLine(screen, x2, y2, x3, y3, lineWidth, color, true)
	vector.StrokeLine(screen, x3, y3, x4, y4, lineWidth, color, true)
	vector.StrokeLine(screen, x4, y4, x5, y5, lineWidth, color, true)
	vector.StrokeLine(screen, x5, y5, x6, y6, lineWidth, color, true)
	vector.StrokeLine(screen, x6, y6, x1, y1, lineWidth, color, true)
}

// DrawD100Dice Draws a circle to represent a d100
func DrawD100Dice(screen *ebiten.Image, x, y, size, lineWidth float32, color color.Color) {
	vector.StrokeCircle(screen, x*CircleDrawingXModifier, y*CircleDrawingYModifier, size/2, lineWidth, color, true)
}

// DiceSwitchingMouseLogic Processes the mouse input logic for whenever the user clicks one of the dice buttons
func (g *Game) DiceSwitchingMouseLogic(mouseX, mouseY int) {
	//Dice Switching Logic
	if mouseX < 150 {
		switch {
		case mouseY > 50 && mouseY < 90:
			g.selectedDice = 4
			rollResult = 0
		case mouseY > 100 && mouseY < 140:
			g.selectedDice = 6
			rollResult = 0
		case mouseY > 150 && mouseY < 190:
			g.selectedDice = 8
			rollResult = 0
		case mouseY > 200 && mouseY < 240:
			g.selectedDice = 10
			rollResult = 0
		case mouseY > 250 && mouseY < 290:
			g.selectedDice = 20
			rollResult = 0
		case mouseY > 300 && mouseY < 340:
			g.selectedDice = 100
			rollResult = 0
		}
	}
	// Roll Dice Logic
	if mouseX >= 20 && mouseX <= 140 && mouseY >= 490 && mouseY <= 600 {
		// Roll the dice when the "Roll Dice" button is clicked
		g.RollDiceAndDisplayResult()
	}
}

// ColorSwitchingMouseLogic processes the mouse input logic for whenever the user clicks on a colored button
func (g *Game) ColorSwitchingMouseLogic(mouseX, mouseY int) {
	// Check mouse click for each color button
	switch {
	case mouseX > 650 && mouseX < 800 && mouseY > 100 && mouseY < 200:
		// Red button
		//fmt.Println("Clicked on Red button")
		g.diceColor = buttonColors[0] // Red color
	case mouseX > 650 && mouseX < 800 && mouseY > 220 && mouseY < 320:
		// Green button
		//fmt.Println("Clicked on Green button")
		g.diceColor = buttonColors[1] // Green color
	case mouseX > 650 && mouseX < 800 && mouseY > 340 && mouseY < 440:
		// Blue button
		//fmt.Println("Clicked on Blue button")
		g.diceColor = buttonColors[2] // Blue color
	case mouseX > 650 && mouseX < 800 && mouseY > 460 && mouseY < 560:
		// Reset button (White with Red X)
		//fmt.Println("Clicked on Reset button")
		g.diceColor = buttonColors[3] // Reset to white
	}
}

// RollDice generates a random number for the dice roll
func RollDice(sides int) int {
	// Generate a random number between 1 and the amount of sides for the dice selected
	return rand.Intn(sides) + 1
}

// MultiplierSwitchingMouseLogic processes mouse input for the multiplier buttons
func (g *Game) MultiplierSwitchingMouseLogic(mouseX, mouseY int) {
	// Check if the decrement button is clicked
	if mouseX > int(DecrementButtonX) && mouseX < int(IncrementButtonX+IncrementButtonWidth) &&
		mouseY > int(DiceCountButtonYStart) && mouseY < int(DiceCountButtonYStart+DiceCountDisplayHeight) {
		if mouseX < int(DecrementButtonX+DecrementButtonWidth) {
			// Decrease only if above 1
			if g.selectedMultiplier > 1 {
				g.selectedMultiplier--
				g.multiplierClicked = true
			}
		} else if mouseX > int(IncrementButtonX) {
			g.selectedMultiplier++ // Increment by 1
		}
	}
	return
}
