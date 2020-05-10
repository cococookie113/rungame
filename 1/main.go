package main

import (
	"log"
	"math/rand"
	"time"

	"global"
	"scenemanager"
	"scenes"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 3.0, "rungame")

	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
