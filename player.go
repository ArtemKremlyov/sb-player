package player

import (
	"container/list"
	"context"
	"errors"
	"log"
	"time"
)

type (
	Song struct {
		Name     string
		Duration time.Duration
	}

	MusicPlayer struct {
		ctx      context.Context
		playlist *list.List
		current  *list.Element
		curTime  time.Duration
		playing  bool
		cancel   context.CancelFunc
	}

	Player interface {
		Play() error
		Pause()
		Next()
		Prev()
		AddSong(s Song)
	}

	playState int
)

const (
	statePlaying playState = iota
	statePaused
)

func New(ctx context.Context, cancelFunc context.CancelFunc) Player {
	return &MusicPlayer{
		ctx:      ctx,
		playlist: list.New(),
		playing:  false,
		cancel:   cancelFunc,
	}
}

func (p *MusicPlayer) Play() error {
	if p.playlist.Len() == 0 {
		return errors.New("error: playlist is empty, please try add track, and try again")
	}

	if p.playing {
		return errors.New("error: player is already playing")
	}

	p.playing = true

	go func() {
		if p.current == nil {
			p.current = p.playlist.Front()
		}

		for {
			select {
			case <-p.ctx.Done():
				p.Pause()
				return
			default:
				song := p.current.Value.(Song)
				time.Sleep(1 * time.Second)

				if p.playing {
					// если песня прослушана
					if p.curTime > song.Duration {
						p.Next()
						p.curTime = 0
					}

					// чтобы слушатель не умирал
					if p.current == nil {
						p.Pause()
						p.current = p.playlist.Front()
						return
					}
					p.curTime += time.Second
				}

				log.Printf("Playing song: %s, duration: %v\n, status %v, %d", song.Name, song.Duration, p.playing, p.curTime/time.Second)
			}
		}
	}()

	return nil
}

func (p *MusicPlayer) Pause() {
	log.Println("paused")
	if p.playing {
		p.playing = false
	}
}

func (p *MusicPlayer) Next() {
	p.curTime = 0
	if p.current != nil {
		p.current = p.current.Next()
	}
}

func (p *MusicPlayer) Prev() {
	p.curTime = 0
	if p.current != nil {
		p.current = p.current.Prev()
	}
}

func (p *MusicPlayer) AddSong(s Song) {
	log.Printf("added song: %s, duration: %v\n", s.Name, s.Duration)
	p.playlist.PushBack(s)
}
