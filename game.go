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
	dino    *man
	cactus  []*cactus
}

const startSpeed = 50
const scoreStep = 30
const speedStep = 5
const speedUpLimit = 120

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
		if g.dino.direction != stay {
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
	g.drawBigTest(GOStr, 10)

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

	g.drawBigTest(score, 4)

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
	g.drawBigTest(NGStr, 10)

	start := "Press SPACE to start\n"
	scorePos := g.maxX/2 - len(start)/2
	moveCursor(position{scorePos, g.maxY/2 - 4})
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
	g.drawBigTest(pause, 10)
}

func (g game) drawBigTest(text [5]string, height int) {
	NGPosX := g.maxX/2 - len(text[0])/2
	NGPosY := g.maxY / 2
	for i := 0; i < len(text); i++ {
		moveCursor(position{NGPosX, NGPosY - height + i})
		draw(text[i])
	}
}
