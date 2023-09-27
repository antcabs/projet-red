package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth = 1000
	screeHeight = 480
)

var (
	running = true
	bkgColor = rl.NewColor(147, 211, 196, 255)

	grassSprite rl.Texture2D
)


func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
}


func imput() {}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)

	drawScene()

	rl.EndDrawing()


}

func init() {
	rl.InitWindow(screenWidth, screeHeight, "Ynovten")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Tilesets/Grass.png")
}

func quit() {
	rl.CloseWindow()
}

func main() {
	

	for running {
		imput()
		update()
		render()
	}
	quit ()
}
