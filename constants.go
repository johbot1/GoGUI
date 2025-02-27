package main

import (
	"crypto/rand"
	"image/color"
	"math/big"
)

var ButtonColor = color.RGBA{160, 155, 155, 255}

const ButtonX = float32(20)

// Draw dice type buttons on the left
const ButtonWidth, ButtonHeight = 120, 40
const ScaleFactor = float32(1.5)

// Labels for the buttons (dice options)
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
