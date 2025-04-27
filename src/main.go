package main

import (
	"embed"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

var (
	//go:embed media/*
	mediaList embed.FS
)

const (
	upCode     = 1
	rightCode  = 2
	downCode   = 3
	leftCode   = 4
	xCode      = 5
	aCode      = 6
	bCode      = 7
	yCode      = 8
	l1Code     = 9
	l2Code     = 10
	r1Code     = 11
	r2Code     = 12
	selectCode = 13
	menuCode   = 14
	startCode  = 15
)

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(1280, 720, "TrimUI Chess")
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	var (
		gamePadId  int32 = 0
		shouldExit       = false
	)

	for !rl.WindowShouldClose() && !shouldExit {

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		//exit
		if rl.IsGamepadButtonDown(gamePadId, menuCode) && rl.IsGamepadButtonDown(gamePadId, startCode) {
			shouldExit = true //see WindowShouldClose, it checks if KeyEscape pressed or Close icon pressed
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
