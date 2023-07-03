package main

import (
	"math/rand"
)

type position [2]int

type direction int

const (
	stay direction = iota
	up
	down
	right
	left
)

func main() {
	game := newGame()
	go game.listenForKeyPress()
	game.beforeGame()
	game.NewGameScreen()

	N, footCounter, step := 0, 0, false
	for {
		if game.started && game.paused {
			game.pause()
			continue
		}
		if !game.started {
			continue
		}
		if footCounter == 8 {
			footCounter = 0
			step = !step
		}
		game.man.move(game.maxX, game.maxY, step)

		for _, c := range game.cactus {
			if game.man.checkAssOnCactus(c) {
				game.over()
			}
			c.move()
			if c.checkHide() {
				game.cactus = game.cactus[1:]
				game.score++
				game.checkScore()
			}
		}

		if N > 10 && rand.Intn(10) == 9 {
			height := rand.Intn(4)
			switch height {
			case 3:
				game.newCactus(cactus3Stages, height, 3)
			case 2:
				game.newCactus(cactus2Stages, height, 3)
			case 1:
				game.newCactus(cactus1Stages, height, 2)
			}
			N = 0
		}

		game.draw()
		N++
		footCounter++
	}
}

func newGame() *game {
	maxX, maxY := getSize()

	game := &game{
		started: false,
		paused:  true,
		score:   0,
		maxX:    maxX,
		maxY:    maxY,
		speed:   startSpeed,
		man:     newDino(position{10, maxY}),
		cactus:  []*cactus{},
	}

	return game
}

func newDino(pos position) *man {
	return &man{
		position:  pos,
		direction: stay,
		foots:     steps[1],
	}
}
