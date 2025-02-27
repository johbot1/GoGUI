package main

import (
	"crypto/rand"
	"image/color"
	"math/big"
)

// DiceSwitchingButtonColor Simply sets the background color of the rectangles behind each button
var DiceSwitchingButtonColor = color.RGBA{160, 155, 155, 255}

// DiceSwitchingButtonXpos specifies the X positional value of a button
const DiceSwitchingButtonXpos = float32(20)
const (
	DiceSwitchingButtonHeight = 40
	DiceSwitchingButtonWidth  = 120
)
const ScaleFactor = float32(1.5)
const DiceSwitchingButtonTitle = int(DiceSwitchingButtonXpos + 20)
const CircleDrawingXModifier = 1.87
const CircleDrawingYModifier = 2.7

// Buttons Labels for the buttons (dice options)
var Buttons = []string{"1d4", "1d6", "1d8", "1d10", "1d20", "1d100"}

// RollDice generates a cryptographically secure random number for the dice roll
func RollDice(sides int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(sides)))
	if err != nil {
		// Fallback: Return max value in case of error (should rarely happen)
		return sides
	}
	return int(n.Int64()) + 1
}
