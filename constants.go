package main

import (
	"image/color"
)

// DiceSwitchingButtonColor Simply sets the background color of the rectangles behind each button
var DiceSwitchingButtonColor = color.RGBA{R: 160, G: 155, B: 155, A: 255}

var DiceRollingButtonColor = color.RGBA{R: 255, G: 26, B: 26, A: 255}

// DiceSwitchingButtonXpos specifies the X positional value of a button
const (
	DiceSwitchingButtonXpos   = float32(20)
	DiceSwitchingButtonHeight = 40
	DiceSwitchingButtonWidth  = 120
	ScaleFactor               = float32(1.5)
	DiceSwitchingButtonTitle  = int(DiceSwitchingButtonXpos + 20)
	CircleDrawingXModifier    = 1.87
	CircleDrawingYModifier    = 2.8
	ButtonPlacementModifier   = 50
)

// Buttons Labels for the buttons (dice options)
var Buttons = []string{"1d4", "1d6", "1d8", "1d10", "1d20", "1d100"}
