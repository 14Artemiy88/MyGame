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
			g.started = !g.started
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

	GOStr := [5]string{
		"     ####   ####  ##   ## #####     ####  ##  ## ##### #####      ",
		"    ##     ##  ## ### ### ##       ##  ## ##  ## ##    ##  ##     ",
		fmt.Sprintf("%v  ## ### ###### ## # ## ####     ##  ## ##  ## ####  #####      %v", manFoot, cactusSeg),
		fmt.Sprintf("%v  ##  ## ##  ## ##   ## ##       ##  ##  ####  ##    ##  ##  %v  %v", manBody, cactusSeg, cactusSeg),
		fmt.Sprintf("%v   ####  ##  ## ##   ## #####     ####    ##   ##### ##  ##  %v  %v  %v", manFoot, cactusSeg, cactusSeg, cactusSeg),
	}
	GOPosX := g.maxX/2 - len(GOStr[0])/2
	GOPosY := g.maxY / 2
	for i := 0; i < len(GOStr); i++ {
		moveCursor(position{GOPosX, GOPosY - 10 + i})
		draw(GOStr[i])
	}

	bigScore := getBigNum(g.score)
	var bigScoreStrArr []string
	var bigScoreStr string
	for j := 0; j < 5; j++ {
		bigScoreStr = ""
		for i := range bigScore {
			bigScoreStr += " " + bigScore[i][j]
		}
		bigScoreStrArr = append(bigScoreStrArr, bigScoreStr)
	}
	moveCursor(position{10, 10})
	score := [5]string{
		fmt.Sprintf(" ####   ####   ####  #####  #####        %v", bigScoreStrArr[0]),
		fmt.Sprintf("##     ##  ## ##  ## ##  ## ##      ##   %v", bigScoreStrArr[1]),
		fmt.Sprintf(" ####  ##     ##  ## #####  ####         %v", bigScoreStrArr[2]),
		fmt.Sprintf("    ## ##  ## ##  ## ##  ## ##      ##   %v", bigScoreStrArr[3]),
		fmt.Sprintf(" ####   ####   ####  ##  ## #####        %v\n", bigScoreStrArr[4]),
	}
	scorePosX := g.maxX/2 - len(score[0])/2
	for i := 0; i < len(score); i++ {
		moveCursor(position{scorePosX, GOPosY - 4 + i})
		draw(score[i])
	}

	render()

	os.Exit(0)
}

func (g game) NewGameScreen() {
	NGStr := [5]string{
		"     ##   ## ##  ##     ####   ####  ##   ## #####      ",
		"     ### ###  ####     ##     ##  ## ### ### ##        ",
		fmt.Sprintf("%v   ## # ##   ##      ## ### ###### ## # ## ####      %v", manFoot, cactusSeg),
		fmt.Sprintf("%v   ##   ##   ##      ##  ## ##  ## ##   ## ##     %v  %v", manBody, cactusSeg, cactusSeg),
		fmt.Sprintf("%v   ##   ##   ##       ####  ##  ## ##   ## #####  %v  %v  %v", manFoot, cactusSeg, cactusSeg, cactusSeg),
	}
	NGPosX := g.maxX/2 - len(NGStr[0])/2
	NGPosY := g.maxY / 2
	for i := 0; i < len(NGStr); i++ {
		moveCursor(position{NGPosX, NGPosY - 10 + i})
		draw(NGStr[i])
	}

	start := "Press SPACE to start\n"
	scorePos := g.maxX/2 - len(start)/2
	moveCursor(position{scorePos, NGPosY - 4})
	draw(start)
	render()
}

func (g game) pause() {
	pause := [5]string{
		"#####   ####  ##  ##  ####  #####",
		"##  ## ##  ## ##  ## ##     ##",
		"#####  ###### ##  ##  ####  ####",
		"##     ##  ## ##  ##     ## ##",
		"##     ##  ##  ####   ####  #####",
	}
	NGPosX := g.maxX/2 - len(pause[0])/2
	NGPosY := g.maxY / 2
	for i := 0; i < len(pause); i++ {
		moveCursor(position{NGPosX, NGPosY - 10 + i})
		draw(pause[i])
	}
}
