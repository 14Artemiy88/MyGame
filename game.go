package main

import "C"
import (
	"fmt"
	"github.com/mattn/go-tty"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type game struct {
	started bool
	maxX    int
	maxY    int
	score   int
	speed   time.Duration
	dino    *dino
	cactus  []*cactus
}

const startSpeed = 50
const scoreStep = 10
const speedStep = 5
const speedUpLimit = 40

func (g *game) draw() {
	clear()

	status := "score: " + strconv.Itoa(g.score)
	statusXPos := g.maxX/2 - len(status)/2

	moveCursor(position{statusXPos, 0})
	draw(status)

	g.dino.draw()
	for _, c := range g.cactus {
		c.draw()
	}

	render()
	time.Sleep(time.Millisecond * g.speed)
}

func (g *game) listenForKeyPress() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		if g.dino.direction != stay {
			continue
		}
		char, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		// UP, DOWN, RIGHT, LEFT == [A, [B, [C, [D
		// we ignore the escape character [
		switch char {
		case ' ':
			if !g.started {
				g.started = true
			}
		case 'A':
			g.dino.direction = up
		case 'B':
			g.dino.direction = down
		case 'C':
			g.dino.direction = right
		case 'D':
			g.dino.direction = left
		}
	}
}

func (g *game) beforeGame() {
	clear()
	hideCursor()

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			g.over()
		}
	}()
}
func (g *game) checkScore() {
	if g.score <= speedUpLimit && g.score%scoreStep == 0 {
		g.speed -= speedStep
	}
}
func (g *game) over() {
	showCursor()
	clear()

	var GOStr string
	GOStr = "     ####   ####  ##   ## #####     ####  ##  ## ##### #####      "
	GOPosX := g.maxX/2 - len(GOStr)/2
	GOPosY := g.maxY / 2
	moveCursor(position{GOPosX, GOPosY - 10})
	draw(GOStr)
	GOStr = "    ##     ##  ## ### ### ##       ##  ## ##  ## ##    ##  ##     "
	moveCursor(position{GOPosX, GOPosY - 9})
	draw(GOStr)
	GOStr = fmt.Sprintf("%v  ## ### ###### ## # ## ####     ##  ## ##  ## ####  #####      %v", manFoot, cactusSeg)
	moveCursor(position{GOPosX, GOPosY - 8})
	draw(GOStr)
	GOStr = fmt.Sprintf("%v  ##  ## ##  ## ##   ## ##       ##  ##  ####  ##    ##  ##  %v  %v", manBody, cactusSeg, cactusSeg)
	moveCursor(position{GOPosX, GOPosY - 7})
	draw(GOStr)
	GOStr = fmt.Sprintf("%v   ####  ##  ## ##   ## #####     ####    ##   ##### ##  ##  %v  %v  %v", manFoot, cactusSeg, cactusSeg, cactusSeg)
	moveCursor(position{GOPosX, GOPosY - 6})
	draw(GOStr)

	score := "score: " + strconv.Itoa(g.score) + "\n"
	scorePos := g.maxX/2 - len(score)/2
	moveCursor(position{scorePos, GOPosY - 4})
	draw(score)

	render()

	os.Exit(0)
}

func (g game) NewGameScreen() {
	var NGStr string
	NGStr = "     ##   ## ##  ##     ####   ####  ##   ## #####      "
	GOPosX := g.maxX/2 - len(NGStr)/2
	GOPosY := g.maxY / 2
	moveCursor(position{GOPosX, GOPosY - 10})
	draw(NGStr)
	NGStr = "     ### ###  ####     ##     ##  ## ### ### ##        "
	moveCursor(position{GOPosX, GOPosY - 9})
	draw(NGStr)
	NGStr = fmt.Sprintf("%v   ## # ##   ##      ## ### ###### ## # ## ####      %v", manFoot, cactusSeg)
	moveCursor(position{GOPosX, GOPosY - 8})
	draw(NGStr)
	NGStr = fmt.Sprintf("%v   ##   ##   ##      ##  ## ##  ## ##   ## ##     %v  %v", manBody, cactusSeg, cactusSeg)
	moveCursor(position{GOPosX, GOPosY - 7})
	draw(NGStr)
	NGStr = fmt.Sprintf("%v   ##   ##   ##       ####  ##  ## ##   ## #####  %v  %v  %v", manFoot, cactusSeg, cactusSeg, cactusSeg)
	moveCursor(position{GOPosX, GOPosY - 6})
	draw(NGStr)

	score := "Press SPACE to start\n"
	scorePos := g.maxX/2 - len(score)/2
	moveCursor(position{scorePos, GOPosY - 4})
	draw(score)

	render()
}
