package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paused struct {
	isPaused bool
	text     string
}

type Timer struct {
	studyDuration     time.Duration
	breakDuration     time.Duration
	longBreakDuration time.Duration
	Pause             Paused
}

func initTimer() Timer {
	return Timer{
		studyDuration:     25 * time.Minute,
		breakDuration:     5 * time.Minute,
		longBreakDuration: 15 * time.Minute,
		Pause:             initPaused(),
	}
}

func initPaused() Paused {
	return Paused{
		isPaused: true,
		text:     "to unpause the timer",
	}
}

func main() {
	const (
		screenWidth   = int32(400)
		screenHeight  = int32(200)
		studyDuration = 25 * time.Minute
	)

	var (
		timer   Timer
		endTime time.Time
	)

	timer = initTimer()
	endTime = time.Now().Add(timer.studyDuration)

	rl.InitWindow(screenWidth, screenHeight, "Pomodoro")

	for !rl.WindowShouldClose() {
		timer.changePauseStatus()
		remainingTime := endTime.Sub(time.Now())
		statusText := "STUDY"
		timeText := fmt.Sprintf("%.001f", remainingTime.Seconds()/60)
		pauseText := fmt.Sprintf("Press [SPACE]: %s", timer.Pause.text)

		statusTextSize := rl.MeasureText(statusText, 20)
		timeTextSize := rl.MeasureText(timeText, 150)
		pauseTextSize := rl.MeasureText(pauseText, 20)

		if !timer.Pause.isPaused {
			endTime = time.Now().Add(timer.studyDuration)
		} else {
			remainingTime = endTime.Sub(time.Now())
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText(timeText, (screenWidth-timeTextSize)/2, 30, 150, rl.Black)
		rl.DrawText(statusText, (screenWidth-statusTextSize)/2, 10, 20, rl.DarkGreen)
		rl.DrawText(pauseText, (screenWidth-pauseTextSize)/2, 170, 20, rl.DarkGray)

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func (t *Timer) changePauseStatus() {
	if rl.IsKeyPressed(rl.KeySpace) {
		t.Pause.isPaused = !t.Pause.isPaused
	}

	if t.Pause.isPaused {
		t.Pause.text = "to unpause the timer"
	} else {
		t.Pause.text = "to pause the timer"
	}
}
