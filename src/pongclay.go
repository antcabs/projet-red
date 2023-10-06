    package main

import (
    "fmt"

    rl "github.com/gen2brain/raylib-go/raylib"
)



func (ball *Ball) Draw() {
	rl.DrawCircle(int32(ball.X), int32(ball.Y), ball.Radius, rl.Yellow)
}


func (paddle *Paddle) GetRect() rl.Rectangle {
	return rl.NewRectangle(paddle.X-paddle.Width/2, paddle.Y-paddle.Height/2, paddle.Width, paddle.Height)
}

func (paddle *Paddle) Draw() {
	rl.DrawTexture(paddle.Texture, int32(paddle.X-paddle.Width/2), int32(paddle.Y-paddle.Height/2), rl.White)
}
func clay() {
    rl.InitWindow(750, 600, "TennisGame")
    rl.SetWindowState(rl.FlagVsyncHint)

    backgroundImage := rl.LoadImage("res/COURT3.png")
    defer rl.UnloadImage(backgroundImage)

    backgroundTexture := rl.LoadTextureFromImage(backgroundImage)

    characterImage1 := rl.LoadImage("res/GOLD.png")
    defer rl.UnloadImage(characterImage1)

	characterImage2 := rl.LoadImage("res/BASIC.png")
	defer rl.UnloadImage(characterImage2)

    characterTexture := rl.LoadTextureFromImage(characterImage1)
	characterTexture2 := rl.LoadTextureFromImage(characterImage2)

    hitSound := rl.LoadSound("res/son.wav")
    defer rl.UnloadSound(hitSound)

    ball := Ball{
        X:      float32(rl.GetScreenWidth()) / 2,
        Y:      float32(rl.GetScreenHeight()) / 2,
        Radius: 8,
        SpeedX: 300,
        SpeedY: 300,
    }

    leftPaddle := Paddle{
        X:       float32(rl.GetScreenWidth()) / 2,
        Y:       50,
        Width:   100,
        Height:  10,
        Speed:   500,
        Texture: characterTexture2,
    }

    rightPaddle := Paddle{
        X:       float32(rl.GetScreenWidth()) / 2,
        Y:       float32(rl.GetScreenHeight()) - 50,
        Width:   100,
        Height:  10,
        Speed:   500,
        Texture: characterTexture,
    }

    var winnerText string
    var scorePlayer1, scorePlayer2 int
    const maxScore = 20

    for !rl.WindowShouldClose() {
        ball.X += ball.SpeedX * rl.GetFrameTime()
        ball.Y += ball.SpeedY * rl.GetFrameTime()

        if ball.X < 0 {
            ball.X = 0
            ball.SpeedX *= -1
        }
        if ball.X > float32(rl.GetScreenWidth()) {
            ball.X = float32(rl.GetScreenWidth())
            ball.SpeedX *= -1
        }

        if leftPaddle.X < ball.X {
            leftPaddle.X += leftPaddle.Speed * rl.GetFrameTime()
        } else {
            leftPaddle.X -= leftPaddle.Speed * rl.GetFrameTime()
        }
        // Player controls for the right paddle
        if rl.IsKeyDown(rl.KeyUp) && rightPaddle.Y > 0 {
            rightPaddle.Y -= rightPaddle.Speed * rl.GetFrameTime()
        }
        if rl.IsKeyDown(rl.KeyDown) && rightPaddle.Y < float32(rl.GetScreenHeight()) {
            rightPaddle.Y += rightPaddle.Speed * rl.GetFrameTime()
        }
        if rl.IsKeyDown(rl.KeyLeft) && rightPaddle.X > 0 {
            rightPaddle.X -= rightPaddle.Speed * rl.GetFrameTime()
        }
        if rl.IsKeyDown(rl.KeyRight) && rightPaddle.X < float32(rl.GetScreenWidth())-rightPaddle.Width {
            rightPaddle.X += rightPaddle.Speed * rl.GetFrameTime()

        }

        if rl.CheckCollisionCircleRec(rl.NewVector2(ball.X, ball.Y), ball.Radius, leftPaddle.GetRect()) {
            if ball.SpeedY < 0 {
                ball.SpeedY *= -1.1
                ball.SpeedX = (ball.X - leftPaddle.X) / (leftPaddle.Width / 2) * ball.SpeedY
                rl.PlaySound(hitSound)
            }
        }

        if rl.CheckCollisionCircleRec(rl.NewVector2(ball.X, ball.Y), ball.Radius, rightPaddle.GetRect()) {
            if ball.SpeedY > 0 {
                ball.SpeedY *= -1.1
                ball.SpeedX = (ball.X - rightPaddle.X) / (rightPaddle.Width / 2) * -ball.SpeedY
                rl.PlaySound(hitSound)
            }
        }

        if ball.Y < 0 {
            scorePlayer2++
            resetBall(&ball)
        }
        if ball.Y > float32(rl.GetScreenHeight()) {
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


