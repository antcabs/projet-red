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
	screenWidth = 1800
	screeHeight = 900
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
	playerSpeed float32 = 0.6

	musicPause bool
	music rl.Music

	cam rl.Camera2D
)







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

	grassSprite = rl.LoadTexture("Tilesets/Termap.png")

	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc =  rl.NewRectangle(0, 0, 16, 16)

	workshopSprite = rl.LoadTexture("Characters/workshop.png")
	workshopPosition = rl.NewVector2(1610, 1330)
	

	playerSprite =rl.LoadTexture("Characters/Character.png")

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(420, 1410, 40, 40)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/MUSICGAME.mp3")
	musicPause = false
	rl.PlayMusicStream(music)

	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screeHeight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0.0, 1.0)

	cam.Zoom = 3

}


func jeu() {
	

	for running {
		imput()
		update()
		ren()
		
	}
}