package main

import "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 400
	screenHeight = 800
	
)

func workshop() {
	// Initialize the window
	rl.InitWindow(screenWidth, screenHeight, "Boutique d'objets")

	// Load the character texture from the PNG image
	characterTexture := rl.LoadTexture("workshop.png")

	ballsTexture := rl.LoadTexture("balls.png")
	blueTexture := rl.LoadTexture("blue_powerade.png")
	redTexture := rl.LoadTexture("red_powerade.png")
	magicTexture := rl.LoadTexture("magic_balls.png")
	barsTexture := rl.LoadTexture("bars.png")
	waterTexture := rl.LoadTexture("WATER.png")
	BagTexture := rl.LoadTexture("BAG.png")
	WOODTexture := rl.LoadTexture("WOOD.png")
	GOLDTexture := rl.LoadTexture("GOLD.png")
	legendTexture := rl.LoadTexture("LEGEND.png")
	jetTexture := rl.LoadTexture("jetpack.png")

	// Position of the character
	characterPosition := rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))

	// Variable to track if the purchase menu should be displayed
	showMenu := false

	// Loop
	for !rl.WindowShouldClose() {
	// Logic

	// Check if the left mouse button is pressed
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
	mousePosition := rl.GetMousePosition()

	// Check if the cursor is on the character
	if rl.CheckCollisionPointRec(rl.NewVector2(float32(mousePosition.X), float32(mousePosition.Y)), rl.NewRectangle(float32(characterPosition.X-float32(characterTexture.Width)/2), float32(characterPosition.Y-float32(characterTexture.Height)/2), float32(characterTexture.Width), float32(characterTexture.Height))) {
	// Open or close the purchase menu based on its current state
	showMenu = !showMenu
	} else if showMenu {
	// If the menu is open and the mouse button is pressed outside the menu, close the menu
	showMenu = false
	}
	}

	// Render

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	// Draw the character
	rl.DrawTexture(characterTexture, int32(characterPosition.X)-int32(float32(characterTexture.Width)/2), int32(characterPosition.Y)-int32(float32(characterTexture.Height)/2), rl.White)

	// Draw the purchase menu if it's open
	if showMenu {
	// Draw the menu background
	rl.DrawRectangle(200, 100, 400, 320, rl.LightGray)

	// Draw the menu items
	rl.DrawText("1. Powerade bleu - Price: 3 coins", 220, 120, 20, rl.Black)
    rl.DrawTexture(blueTexture, 200, 120, rl.White)
	rl.DrawText("2. Powerade rouge - Price: 3 coins", 220, 140, 20, rl.Black)
    rl.DrawTexture(redTexture, 200, 140, rl.White)
	rl.DrawText("3. Balls - Price: 6 coins", 220, 160, 20, rl.Black)
    rl.DrawTexture(ballsTexture, 200, 160, rl.White)
	rl.DrawText("4. Bars - Price: 2 coins", 220, 180, 20, rl.Black)
    rl.DrawTexture(barsTexture, 200, 180, rl.White)
	rl.DrawText("5. Water - Price: 6 coins", 220, 200, 20, rl.Black)
    rl.DrawTexture(waterTexture, 200, 200, rl.White)
	rl.DrawText("6. Wooden Racket - Price: 15 coins", 220, 220, 20, rl.Black)
    rl.DrawTexture(WOODTexture, 200, 220, rl.White)
	rl.DrawText("7. Golden Racket - Price: 50 coins", 220, 240, 20, rl.Black)
    rl.DrawTexture(GOLDTexture, 200, 240, rl.White)
	rl.DrawText("8. Legend Racket - Price: 200 coins", 220, 260, 20, rl.Black)
    rl.DrawTexture(legendTexture, 200, 260, rl.White)
	rl.DrawText("9. Magic Balls - Price: 50 coins", 220, 300, 20, rl.Black)
    rl.DrawTexture(magicTexture, 200, 300, rl.White)
	rl.DrawText("10. Jetpack  - Price: 400 coins", 220, 330, 20, rl.Black)
    rl.DrawTexture(jetTexture, 200, 330, rl.White)
	rl.DrawText("11. BagTexture - Price: 30 coins", 220, 350, 20, rl.Black)
    rl.DrawTexture(BagTexture, 200, 350, rl.White)


	// Add more menu items as needed
	}

	rl.EndDrawing()
	}

	// Unload the character texture
	rl.UnloadTexture(characterTexture)

	// Close the window
	rl.CloseWindow()
}