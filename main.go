package main

import (
	"log"

	"github.com/BismarckDD/Caocao/c2game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	game, succeed, err := c2game.CreateGame()

	if !succeed || err != nil {
		log.Fatalf("Failed to create the CCZ-Game %s\n", "ERROR")
		return
	}

	log.Println("Start to run game.")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal("RunGame Error.", err)
	}

}
