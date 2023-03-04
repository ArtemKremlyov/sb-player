package main

import (
	"fmt"
	"github.com/ArtemKremlyov/player"
)

func main() {
	mPlayer := player.NewPlayer()
	mPlayer.AddSong(player.Song{
		Duration: 10,
		Name:     "test",
	})

	err := mPlayer.Play()

	if err != nil {
		fmt.Errorf("play: %v", err)
	}
}
