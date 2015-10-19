package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	// unicode for ant
	antRune = '\U0001F41C'
)

type ant struct {
	id  int
	loc location
}

type location struct {
	x, y int
}

func draw(ants []ant) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	defer termbox.Flush()

	for _, ant := range ants {
		termbox.SetCell(ant.loc.x, ant.loc.y, antRune, termbox.ColorRed, termbox.ColorDefault)
	}
}

func generateAnts(n, width, height int) []ant {
	rand.Seed(time.Now().UnixNano())
	var ants []ant
	for i := 0; i < n; i++ {
		ants = append(ants, ant{id: i, loc: location{x: rand.Intn(width), y: rand.Intn(height)}})
	}
	return ants
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	w, h := termbox.Size()
	ants := generateAnts(10, w, h)
	draw(ants)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventError:
			panic(ev.Err)
		default:
			fmt.Println(ev.Type)
		}
	}
}
