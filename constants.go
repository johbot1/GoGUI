package main

import (
	"image/color"
)

// DiceSwitchingButtonColor Simply sets the background color of the rectangles behind each button
var DiceSwitchingButtonColor = color.RGBA{R: 160, G: 155, B: 155, A: 255}

var DiceRollingButtonColor = color.RGBA{R: 255, G: 26, B: 26, A: 255}

var multiplierButtonColor = color.RGBA{80, 80, 80, 255} // Example color; adjust as desired

var buttonColors = []color.RGBA{ // Define colors for the color buttons
	{255, 0, 0, 255},     // 1) Red
	{0, 255, 0, 255},     // 2) Green
	{0, 0, 255, 255},     // 3) Blue
	{255, 255, 255, 255}, // 4) White (for the X button)
}

// Constants applied to various parts of the code, grouped by their location
const (
	// Dice Switching Button Parameters
	DiceSwitchingButtonXpos   = float32(20)
	DiceSwitchingButtonHeight = 40
	DiceSwitchingButtonWidth  = 120
	DiceSwitchingButtonTitle  = int(DiceSwitchingButtonXpos + 20)
	// Dice Drawing Parameters
	ScaleFactor             = float32(1.5)
	CircleDrawingXModifier  = 1.82
	CircleDrawingYModifier  = 2.5
	ButtonPlacementModifier = 50
	// Roll Dice ButtonParameters
	RollButtonXpos     = 20
	RollButtonYpos     = 500
	RollButtonTextXpos = 65
	RollButtonTextYpos = 550
	RollButtonWidth    = 120
	RollButtonHeight   = 80
	// Font Parameters
	FontSize = 25
	FontDPI  = FontSize * 3
	//Multiple Dice Button Parameters
	MultiplierButtonX       = float32(20)
	MultiplierButtonYStart  = float32(350) // Start drawing below the dice selection buttons
	MultiplierButtonWidth   = float32(120)
	MultiplierButtonHeight  = float32(40)
	MultiplierButtonSpacing = float32(50) // Vertical spacing between buttons
	//Color Changing Button Parameters
	RightButtonYStart  = float32(100) // Start vertically around 100px from the top
	RightButtonWidth   = float32(100) // Button width
	RightButtonHeight  = float32(100) // Button height
	RightButtonSpacing = float32(20)  // Space between buttons

)

// Buttons Labels for the buttons (dice options)
var Buttons = []string{"1d4", "1d6", "1d8", "1d10", "1d20", "1d100"}
