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

## Section 1: The Options Screen
    Display Particle Type Options:
        Create a menu with options for circle, square, and star particles.
        Allow user to select one particle type.

    Color Picker/Selection:
        Offer a selection of predefined colors (e.g., red, green, blue, yellow).
        Allow the user to select one color for the particles.

    Speed Control:
        Provide options for slow, medium, and fast speeds.
        Implement slider or radio buttons for speed selection.

    Start Animation Button:
        Create a button that, when clicked, transitions to the visual screen (animation).

## Section 2: Visual Screen

    Particle Emitter:
        Emit particles from the center of the screen or on mouse click.
        Create a loop to continuously emit particles based on the selected speed.

    Particle Movement:
        Make particles move in random directions (with slight speed variations).
        Implement smooth animation for particles’ movement.

    Particle Properties:
        Particles should have random sizes and fade out as they move.
        Set a lifetime for each particle, after which it disappears.

    Interactive Particle Generation (optional):
        Allow particles to be generated on mouse click or mouse movement.
        Adjust emission speed based on the speed option selected on the first screen.

    Looping:
        Particles should keep emitting until the user presses a key to restart or exit.


## Section 3: Feedback
(To be filled in after receiving feedback.)