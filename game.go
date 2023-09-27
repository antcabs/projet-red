package main

import (
	
	"github.com/gen2brain/raylib-go/raylib"
	
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

const (
	screenWidth = 1000
	screeHeight = 600
	screen2Width = 800
	screen2Height = 400
)

var (
	running = true
	bkgColor = rl.NewColor(147, 211, 196, 255)

	workshopSprite rl.Texture2D
	workshopSrc rl.Rectangle
	workshopDest rl.Rectangle
	workshopPosition rl.Vector2


	grassSprite rl.Texture2D
	tex rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc rl.Rectangle
	playerDest rl.Rectangle
	playerMoving bool
	playerDir int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame int 
	inventory = []string{"Racket", "Balls", "Water", "Powerade"}
	
	frameCount int
	tileDest rl.Rectangle
	tileSrc rl.Rectangle
	tileMap []int
	srcMap []string
	mapW, mapH int

	ballsTexture = rl.LoadTexture("balls.png")
	blueTexture = rl.LoadTexture("blue_powerade.png")
	redTexture = rl.LoadTexture("red_powerade.png")
	magicTexture = rl.LoadTexture("magic_balls.png")
	barsTexture = rl.LoadTexture("bars.png")
	waterTexture = rl.LoadTexture("WATER.png")
	BagTexture = rl.LoadTexture("BAG.png")
	WOODTexture = rl.LoadTexture("WOOD.png")
	GOLDTexture = rl.LoadTexture("GOLD.png")
	legendTexture = rl.LoadTexture("LEGEND.png")
	jetTexture = rl.LoadTexture("jetpack.png")
	
	showMenu = false
	playerSpeed float32 = 2

	musicPause bool
	music rl.Music

	cam rl.Camera2D
)

type Paddle struct {
	X, Y float32
	Speed float32
	Width float32
	Height float32
	Texture rl.Texture2D
}

type Ball struct {
	X, Y float32
	SpeedX float32
	SpeedY float32
	Radius float32
	
}

func (ball *Ball) Draw() {
	rl.DrawCircle(int32(ball.X), int32(ball.Y), ball.Radius, rl.Yellow)
}

func (paddle *Paddle) GetRect() rl.Rectangle {
	return rl.NewRectangle(paddle.X-paddle.Width/2, paddle.Y-paddle.Height/2, paddle.Width, paddle.Height)
}




func(paddle *Paddle) Draw (){
	rl.DrawTexture(paddle.Texture, int32(paddle.X-paddle.Width/2), int32(paddle.Y-paddle.Height/2), rl.White)

}




func drawScene() {
		rl.DrawTexture(grassSprite, 100, 50, rl.White)
		

		
	for i := 0; i<len(tileMap); i++ {
		if tileMap[i]!=0 {
			tileDest.X = tileDest.Width * float32(i % mapW)
			tileDest.Y = tileDest.Height * float32(i / mapW)

			if srcMap[i] == "g" { tex = grassSprite }

			tileSrc.X =tileSrc.Width * float32((tileMap[i]-1) % int(tex.Width/ int32(tileSrc.Width)))
			tileSrc.Y =tileSrc.Height * float32((tileMap[i]-1) / int(tex.Width/ int32(tileSrc.Width)))
			rl.DrawTexturePro(grassSprite, tileSrc, tileDest,rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
		}
		 
	}

	rl.DrawTexturePro(playerSprite, playerSrc, playerDest,rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)
	rl.DrawTexturePro(workshopSprite, rl.NewRectangle(0, 0, float32(workshopSprite.Width), float32(workshopSprite.Height)), rl.NewRectangle(workshopPosition.X, workshopPosition.Y, float32(workshopSprite.Width), float32(workshopSprite.Height)), rl.Vector2{}, 0, rl.White)

}


func imput() {

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePosition := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(rl.NewVector2(float32(mousePosition.X), float32(mousePosition.Y)), rl.NewRectangle(float32(workshopPosition.X-float32(workshopSprite.Width)/2), float32(workshopPosition.Y-float32(workshopSprite.Height)/2), float32(workshopSprite.Width), float32(workshopSprite.Height))) {
		showMenu = !showMenu
		
		}
	}	

	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving =true
		playerDir = 1
		playerUp=true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving =true
		playerDir = 0
		playerDown=true

	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving =true
		playerDir = 2
		playerLeft=true

	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving =true
		playerDir =3
		playerRight=true

	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPause = !musicPause
	}


}
func update() {
	running = !rl.WindowShouldClose()

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
		}






		playerSrc.X = 0


	if playerMoving {
		if playerUp {playerDest.Y-=playerSpeed}
		if playerDown {playerDest.Y+=playerSpeed}
		if playerLeft {playerDest.X-=playerSpeed}
		if playerRight {playerDest.X+=playerSpeed}
		if frameCount % 8 == 1 {playerFrame++}
		playerSrc.X = playerSrc.Width * float32(playerFrame)
	}
		frameCount++
	
	if playerFrame > 3 {playerFrame = 0}
		playerSrc.Y = playerSrc.Height * float32(playerDir) 

	

	rl.UpdateMusicStream(music)
	if musicPause {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
}

func ping() {
	rl.InitWindow(2000, 900, "tt/TennisGame")
	rl.SetWindowState(rl.FlagVsyncHint)

	backgroundImage := rl.LoadImage("res/Tilesets/Termap.png")
	defer rl.UnloadImage(backgroundImage)

	backgroundTexture := rl.LoadTextureFromImage(backgroundImage)

	characterImage := rl.LoadImage("tt/BASIC.png")
	defer rl.UnloadImage(characterImage)

	characterTexture := rl.LoadTextureFromImage(characterImage)

	hitSound := rl.LoadSound("tt/son.wav")
	defer rl.UnloadSound(hitSound)

	ball := Ball{
		X:      float32(rl.GetScreenWidth()) / 2,
		Y:      float32(rl.GetScreenHeight()) / 2,
		Radius: 8,
		SpeedX: 300,
		SpeedY: 300,
	}

	leftPaddle := Paddle{
		X:       50,
		Y:       float32(rl.GetScreenHeight()) / 2,
		Width:   10,
		Height:  100,
		Speed:   500,
		Texture: characterTexture,
	}

	rightPaddle := Paddle{
		X:       float32(rl.GetScreenWidth()) - 50,
		Y:       float32(rl.GetScreenHeight()) / 2,
		Width:   10,
		Height:  100,
		Speed:   500,
		Texture: characterTexture,
	}

	var winnerText string
	var scorePlayer1, scorePlayer2 int
	const maxScore = 10

	for !rl.WindowShouldClose() {
		ball.X += ball.SpeedX * rl.GetFrameTime()
		ball.Y += ball.SpeedY * rl.GetFrameTime()

		if ball.Y < 0 {
			ball.Y = 0
			ball.SpeedY *= -1
		}
		if ball.Y > float32(rl.GetScreenHeight()) {
			ball.Y = float32(rl.GetScreenHeight())
			ball.SpeedY *= -1
		}

		if leftPaddle.Y < ball.Y {
			leftPaddle.Y += leftPaddle.Speed * rl.GetFrameTime()
		} else {
			leftPaddle.Y -= leftPaddle.Speed * rl.GetFrameTime()
		}

		if rl.IsKeyDown(rl.KeyUp) {
			rightPaddle.Y -= rightPaddle.Speed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyDown) {
			rightPaddle.Y += rightPaddle.Speed * rl.GetFrameTime()
		}

		if rl.CheckCollisionCircleRec(rl.NewVector2(ball.X, ball.Y), ball.Radius, leftPaddle.GetRect()) {
			if ball.SpeedX < 0 {
				ball.SpeedX *= -1.1
				ball.SpeedY = (ball.Y - leftPaddle.Y) / (leftPaddle.Height / 2) * ball.SpeedX
				rl.PlaySound(hitSound)
			}
		}

		if rl.CheckCollisionCircleRec(rl.NewVector2(ball.X, ball.Y), ball.Radius, rightPaddle.GetRect()) {
			if ball.SpeedX > 0 {
				ball.SpeedX *= -1.1
				ball.SpeedY = (ball.Y - rightPaddle.Y) / (rightPaddle.Height / 2) * -ball.SpeedX
				rl.PlaySound(hitSound)
			}
		}

		if ball.X < 0 {
			scorePlayer2++
			resetBall(&ball)
		}
		if ball.X > float32(rl.GetScreenWidth()) {
			scorePlayer1++
			resetBall(&ball)
		}

		if scorePlayer1 == maxScore {
			winnerText = "Jeu, set et match ! Player wins!"
			break
		} else if scorePlayer2 == maxScore {
			winnerText = "Jeu, set et match ! Opponent wins!"
			break
		}

		if winnerText != "" && rl.IsKeyPressed(rl.KeySpace) {
			resetBall(&ball)
			winnerText = ""
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawTexture(backgroundTexture, 0, 0, rl.White)

		ball.Draw()
		leftPaddle.Draw()
		rightPaddle.Draw()

		if winnerText != "" {
			textWidth := rl.MeasureText(winnerText, 60)
			rl.DrawText(winnerText, int32(rl.GetScreenWidth())/2-int32(textWidth)/2, int32(rl.GetScreenHeight())/2-30, 60, rl.Yellow)
		}

		rl.DrawText(fmt.Sprintf("Opponent: %d", scorePlayer1), 10, 20, 20, rl.Black)
		rl.DrawText(fmt.Sprintf("Player: %d", scorePlayer2), int32(rl.GetScreenWidth())-140, 20, 20, rl.Black)

		rl.DrawFPS(1, 1)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func resetBall(ball *Ball) {
	ball.X = float32(rl.GetScreenWidth()) / 2
	ball.Y = float32(rl.GetScreenHeight()) / 2
	ball.SpeedX = 300
	ball.SpeedY = 300
}



func ren() {
	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()


}

func loadMap(mapFile string) { 
	file, err := ioutil.ReadFile(mapFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	remNewLines := strings.Replace(string(file), "\n", " ", -1)
	sliced := strings.Split(remNewLines, " ")
	mapW = -1
	mapH = -1
	for i:=0; i<len(sliced); i++{
		s, _ := strconv.ParseInt(sliced[i], 10, 64)
		m := int(s)
		if mapW == -1 {
			mapW = m
		} else if mapH == -1{
			mapH = m 
		} else if i < mapW*mapH+2 {
			tileMap = append(tileMap, m)
		}else {
			srcMap = append(srcMap, sliced[i])
		}
	}
	if len(tileMap) > mapW*mapH {tileMap = tileMap[:len(tileMap)-1]
	}

}
func init() {
	rl.InitWindow(screenWidth, screeHeight, "Ynovten")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Tilesets/Termap.png")

	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc =  rl.NewRectangle(0, 0, 16, 16)

	workshopSprite = rl.LoadTexture("res/Characters/workshop.png")
	workshopPosition = rl.NewVector2(1610, 1330)
	

	playerSprite =rl.LoadTexture("res/Characters/Character.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 60, 60)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/MUSICGAME.mp3")
	musicPause = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screeHeight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0.0, 1.0)

	cam.Zoom = 0.7

}


func main() {
	

	for running {
		imput()
		update()
		ren()
		pong()
	}
}