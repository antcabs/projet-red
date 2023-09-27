package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	X, Y   float32
	SpeedX float32
	SpeedY float32
	Radius float32
}

func (ball *Ball) Draw() {
	rl.DrawCircle(int32(ball.X), int32(ball.Y), ball.Radius, rl.Yellow)
}

type Paddle struct {
	X, Y    float32
	Speed   float32
	Width   float32
	Height  float32
	Texture rl.Texture2D
}

func (paddle *Paddle) GetRect() rl.Rectangle {
	return rl.NewRectangle(paddle.X-paddle.Width/2, paddle.Y-paddle.Height/2, paddle.Width, paddle.Height)
}

func (paddle *Paddle) Draw() {
	rl.DrawTexture(paddle.Texture, int32(paddle.X-paddle.Width/2), int32(paddle.Y-paddle.Height/2), rl.White)
}

func pong() {
	rl.InitWindow(1060, 650, "tt/TennisGame")
	rl.SetWindowState(rl.FlagVsyncHint)

	backgroundImage := rl.LoadImage("tt/COURTE.jpg")
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
