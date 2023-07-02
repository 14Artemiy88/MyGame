package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/term"
)

var screen = bufio.NewWriter(os.Stdout)

func hideCursor() {
	fmt.Fprint(screen, "\033[?25l")
}

func showCursor() {
	fmt.Fprint(screen, "\033[?25h")
}

func moveCursor(pos [2]int) {
	fmt.Fprintf(screen, "\033[%d;%dH", pos[1], pos[0])
}

func clear() {
	fmt.Fprint(screen, "\033[2J")
}

func bgRgb(r int, g int, b int) {
	fmt.Fprintf(screen, "\033[48;2;%d;%d;%dm", r, g, b)
}

func fgRgb(r int, g int, b int, symbol string) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + symbol + "\033[0m"
}

func draw(str string) {
	fmt.Fprint(screen, str)
}

func drawF(str string, params ...string) {
	fmt.Fprintf(screen, str, params)
}

func render() {
	screen.Flush()
}

func getSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	return width, height
}
