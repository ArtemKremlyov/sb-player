package main

import (
	"context"
	"log"
	"time"

	"github.com/ArtemKremlyov/player"
)

func main() {
	// создание плейлиста
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := player.New(ctx, cancel)

	// добавить песни в плейлист
	p.AddSong(player.Song{Name: "Song 1", Duration: 5 * time.Second})
	p.AddSong(player.Song{Name: "Song 2", Duration: 4 * time.Second})
	p.AddSong(player.Song{Name: "Song 3", Duration: 7 * time.Second})

	if err := p.Play(); err != nil {
		log.Fatalf("error while playing playlist: %v", err)
	}

	time.Sleep(3 * time.Second)

	p.Pause()
	time.Sleep(10 * time.Second)
	if err := p.Play(); err != nil {
		log.Println(err)
	}
	time.Sleep(10 * time.Second)
	if err := p.Play(); err != nil {
		log.Println(err)
	}

	time.Sleep(40 * time.Second)
}
