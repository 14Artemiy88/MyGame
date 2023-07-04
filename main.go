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
		for key, piu := range game.man.piu {
			piu.move()
			needDelete := piu.checkCactus(game.cactus)
			if needDelete || piu.position[0] > game.maxX {
				if len(game.man.piu) > key {
					game.man.delPiu(key)
				}
			}
		}

		for _, cactus := range game.cactus {
			if game.man.checkAssOnCactus(cactus) {
				game.draw(true)
				time.Sleep(time.Millisecond * 1000)
				game.over()
			}
			cactus.move()
			if cactus.checkHide() {
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

		game.draw(false)
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
		man:     newMan(position{10, maxY}),
		cactus:  []*cactus{},
		star:    newStars(maxX, maxY/2),
	}

	return game
}

func newMan(pos position) *man {
	return &man{
		position:  pos,
		direction: stay,
		foots:     steps[1],
	}
}

func newStars(maxX int, maxY int) [starsCount]star {
	var stars []star
	for i := 0; i <= starsCount; i++ {
		x := rand.Intn(maxX)
		y := rand.Intn(maxY)
		stars = append(stars, star{position: position{x, y}})
	}

	return [starsCount]star(stars)
}
