# GoGUI
One of three Go! related projects for my final semester, this one to build a GUI.
I'm really starting to not like coding.

## Requirements:
Golang v.1.23.6 or higher
- If you do not have Go installed, visit the following link for setup instructions  
  `https://go.dev/dl/`
- To check if you have Go installed properly, open a command line and enter:  
  `go version`

## How To Run:
- Download the files from this github and unzip them
- Navigate to the unzipped folder in a command terminal
- Use the following command to run the project:  
  `go run main.go`
- If you wish to build an executable for repeated tests, run this command:  
  `go build main.go`



# Overview:

## Section 1: The Options Screen ✅
### 🔹 Dice Type Selection: ✅

    Provide six dice options on the left side of the screen:
        🎲 1d4
        🎲 1d6
        🎲 1d8
        🎲 1d10
        🎲 1d20
        🎲 1d100
    Only one dice type can be selected at a time.

### 🔹 Color Scheme Selection:

    On the right side, allow the user to select one of three line color options:
        🟢 Green Lines
        🔵 Blue Lines
        🔴 Red Lines
    The background remains black at all times.

### 🔹 Roll Button:

    A large rectangular button below the dice labeled "ROLL".
    Clicking this button will trigger the dice roll animation and display the result.

## Section 2: The Visual Screen

### 🎲 Dice Representation:

    The dice is drawn in a simple wireframe style (blocky, geometric).
    The selected dice type determines the shape.

### 🎲 Rolling Animation:

    When the roll button is clicked:
        The dice wobbles or rotates slightly before stopping.
        A random result is selected and displayed above the dice.
        The dice momentarily fades out and back in to indicate the roll is finished.

### 🎲 Result Display:

    The rolled number appears above the dice in large, clear text.

## Section 3: Fixes / Feedback

(This section will be filled in after receiving feedback.)