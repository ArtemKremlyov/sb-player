package player

import (
	"container/list"
	"fmt"
	"log"
	"sync"
	"time"
)

type (
	Song struct {
		Name     string
		Duration int
	}

	MusicPlayer struct {
		playlist *list.List
		current  *list.Element
		playing  bool
		mutex    sync.Mutex
	}

	Player interface {
		Play() error
		Pause()
		Next()
		Prev()
		AddSong(s Song)
	}
)

func NewPlayer() Player {
	return &MusicPlayer{
		playlist: list.New(),
		playing:  false,
	}
}

func (p *MusicPlayer) Play() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.playlist.Len() == 0 {
		return fmt.Errorf("error: playlist is empty, please try add track, and try again")
	}

	if !p.playing {
		p.playing = true
		log.Println(p.current)
		go func() {
			for {
				if p.current == nil {
					p.current = p.playlist.Front()
				}

				select {
				case <-time.After(time.Duration(p.current.Value.(*Song).Duration)):
					p.Next()
				default:
					time.Sleep(50 * time.Millisecond)
				}

			}
		}()
	}

	return nil
}

func (p *MusicPlayer) Pause() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.playing {
		p.playing = false
	}
}

func (p *MusicPlayer) Next() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.current != nil {
		p.current = p.current.Next()
		p.Play()
	}
}

func (p *MusicPlayer) Prev() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.current != nil {
		p.current = p.current.Prev()
		p.Play()
	}
}

func (p *MusicPlayer) AddSong(s Song) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.playlist.PushBack(s)
}
