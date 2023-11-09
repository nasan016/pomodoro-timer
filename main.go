package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paused struct {
	isPaused       bool
	text           string
	pauseStartTime time.Time
	pauseDuration  time.Duration
}

func initPaused() Paused {
	return Paused{
		isPaused: false,
		text:     "to unpause the timer",
	}
}

type Timer struct {
	studyDuration     time.Duration
	breakDuration     time.Duration
	longBreakDuration time.Duration
	Pause             Paused
}

func initTimer() Timer {
	return Timer{
		studyDuration:     5 * time.Second,
		breakDuration:     4 * time.Second,
		longBreakDuration: 5 * time.Second,
		Pause:             initPaused(),
	}
}

func main() {
	const (
		screenWidth  = int32(400)
		screenHeight = int32(200)
	)

	var (
		timer         Timer
		endTime       time.Time
		remainingTime time.Duration
	)

	timer = initTimer()
	endTime = time.Now().Add(timer.studyDuration)
	remainingTime = endTime.Sub(time.Now())

	rl.InitWindow(screenWidth, screenHeight, "Pomodoro")

	for !rl.WindowShouldClose() {
		//timer methods
		timer.changePauseStatus()
		timer.changeText()
		timer.startPauseTimer()

		minutes := int(remainingTime.Minutes())
		seconds := int(remainingTime.Seconds()) % 60

		statusText := "STUDY"
		timeText := fmt.Sprintf("%02d:%02d", minutes, seconds)
		pauseText := fmt.Sprintf("Press [SPACE]: %s", timer.Pause.text)

		statusTextSize := rl.MeasureText(statusText, 20)
		timeTextSize := rl.MeasureText(timeText, 150)
		pauseTextSize := rl.MeasureText(pauseText, 20)

		if timer.Pause.isPaused {
			remainingTime = endTime.Sub(time.Now())
			timer.Pause.pauseDuration = time.Since(timer.Pause.pauseStartTime)
		} else {
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
}

func (t *Timer) changeText() {
	if t.Pause.isPaused {
		t.Pause.text = "to unpause the timer"
	} else {
		t.Pause.text = "to pause the timer"
	}
}

func (t *Timer) startPauseTimer() {
	if t.Pause.isPaused {

	} else {

	}
}
