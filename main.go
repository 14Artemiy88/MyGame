package main

import (
	"math/rand"
	"time"
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
		if !game.started {
			//game.pause()
			continue
		}
		if footCounter == 8 {
			footCounter = 0
			step = !step
		}
		game.dino.move(game.maxX, game.maxY, step)

		for _, c := range game.cactus {
			if game.dino.checkAssOnCactus(c) {
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
			game.cactus = append(game.cactus, newCactus(position{game.maxX, game.maxY}))
			N = 0
		}

		game.draw()
		N++
		footCounter++
	}
}

func newGame() *game {
	rand.Seed(time.Now().UnixNano())
	maxX, maxY := getSize()

	game := &game{
		started: false,
		score:   0,
		maxX:    maxX,
		maxY:    maxY,
		speed:   startSpeed,
		dino:    newDino(position{10, maxY}),
		cactus:  append([]*cactus{}, newCactus(position{maxX, maxY})),
	}

	return game
}

func newDino(pos position) *dino {
	return &dino{
		position:  pos,
		direction: stay,
		foots:     steps[1],
	}
}

func newCactus(pos position) *cactus {
	return &cactus{
		position: pos,
		height:   rand.Intn(3),
		width:    rand.Intn(3),
	}
}
