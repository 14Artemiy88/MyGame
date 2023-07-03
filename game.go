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
	paused  bool
	maxX    int
	maxY    int
	score   int
	speed   time.Duration
	man     *man
	cactus  []*cactus
}

const startSpeed = 50
const scoreStep = 30
const speedStep = 5
const speedUpLimit = 120

func (g *game) draw(isDead bool) {
	clear()

	status := "score: " + strconv.Itoa(g.score)
	statusXPos := g.maxX/2 - len(status)/2
	moveCursor(position{statusXPos, 0})
	draw(status)

	for _, c := range g.cactus {
		c.draw()
	}
	g.man.draw(isDead)

	render()
	time.Sleep(time.Millisecond * g.speed)
}

func (g *game) listenForKeyPress() {
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func(t *tty.TTY) {
		err := t.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(t)

	for {
		if g.man.direction != stay {
			continue
		}
		char, err := t.ReadRune()
		if err != nil {
			panic(err)
		}

		// UP, DOWN, RIGHT, LEFT == [A, [B, [C, [D
		// we ignore the escape character [
		switch char {
		case ' ':
			g.paused = !g.paused
			g.started = true
		case 'A':
			g.man.direction = up
		//case 'B':
		//	g.man.direction = down
		case 'C':
			g.man.direction = right
		case 'D':
			g.man.direction = left
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

func (g *game) newCactus(cactusStages [][]string, height int, width int) {
	for i := 0; i < width; i++ {
		cactusPart := &cactus{
			stages:   cactusStages[i],
			height:   height,
			position: position{g.maxX - i, g.maxY},
		}
		g.cactus = append(g.cactus, cactusPart)
	}
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
		fmt.Sprintf("%v  ## ### ###### ## # ## ####     ##  ## ##  ## ####  #####   %v",
			manFoot,
			fgRgb(0, 200, 0, cactus3Stages[2][2]+cactus3Stages[1][2]+cactus3Stages[0][2]),
		),
		fmt.Sprintf("%v  ##  ## ##  ## ##   ## ##       ##  ##  ####  ##    ##  ##  %v  %v",
			manBody,
			fgRgb(0, 200, 0, cactus3Stages[2][1]+cactus3Stages[1][1]+cactus3Stages[0][1]),
			fgRgb(0, 200, 0, cactus2Stages[2][1]+cactus2Stages[1][1]+cactus2Stages[0][1]),
		),
		fmt.Sprintf("%v   ####  ##  ## ##   ## #####     ####    ##   ##### ##  ##  %v  %v  %v",
			manFoot,
			fgRgb(0, 200, 0, cactus3Stages[2][0]+cactus3Stages[1][0]+cactus3Stages[0][0]),
			fgRgb(0, 200, 0, cactus2Stages[2][0]+cactus2Stages[1][0]+cactus2Stages[0][0]),
			fgRgb(0, 200, 0, cactus1Stages[1][0]+cactus1Stages[0][0]),
		),
	}
	drawBigTest(GOStr, 10, g.maxX, g.maxY)

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
	score := [5]string{
		fmt.Sprintf(" ####   ####   ####  #####  #####        %v", bigScoreStrArr[0]),
		fmt.Sprintf("##     ##  ## ##  ## ##  ## ##      ##   %v", bigScoreStrArr[1]),
		fmt.Sprintf(" ####  ##     ##  ## #####  ####         %v", bigScoreStrArr[2]),
		fmt.Sprintf("    ## ##  ## ##  ## ##  ## ##      ##   %v", bigScoreStrArr[3]),
		fmt.Sprintf(" ####   ####   ####  ##  ## #####        %v\n", bigScoreStrArr[4]),
	}
	drawBigTest(score, 4, g.maxX, g.maxY)

	render()

	os.Exit(0)
}

func (g game) NewGameScreen() {
	NGStr := [5]string{
		"     ##   ## ##  ##     ####   ####  ##   ## #####      ",
		"     ### ###  ####     ##     ##  ## ### ### ##        ",
		fmt.Sprintf(
			"%v   ## # ##   ##      ## ### ###### ## # ## ####   %v",
			manFoot,
			fgRgb(0, 200, 0, cactus3Stages[2][2]+cactus3Stages[1][2]+cactus3Stages[0][2]),
		),
		fmt.Sprintf(
			"%v   ##   ##   ##      ##  ## ##  ## ##   ## ##     %v  %v",
			manBody,
			fgRgb(0, 200, 0, cactus3Stages[2][1]+cactus3Stages[1][1]+cactus3Stages[0][1]),
			fgRgb(0, 200, 0, cactus2Stages[2][1]+cactus2Stages[1][1]+cactus2Stages[0][1]),
		),
		fmt.Sprintf(
			"%v   ##   ##   ##       ####  ##  ## ##   ## #####  %v  %v  %v",
			manFoot,
			fgRgb(0, 200, 0, cactus3Stages[2][0]+cactus3Stages[1][0]+cactus3Stages[0][0]),
			fgRgb(0, 200, 0, cactus2Stages[2][0]+cactus2Stages[1][0]+cactus2Stages[0][0]),
			fgRgb(0, 200, 0, cactus1Stages[1][0]+cactus1Stages[0][0]),
		),
	}
	drawBigTest(NGStr, 10, g.maxX, g.maxY)

	drawByCenter("↑ - Jump", g.maxX, g.maxY/2-4)
	drawByCenter("→ - Right", g.maxX, g.maxY/2-3)
	drawByCenter("← - Left", g.maxX, g.maxY/2-2)
	drawByCenter("Press SPACE to start\n", g.maxX, g.maxY/2)

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
	drawBigTest(pause, 10, g.maxX, g.maxY)
}
