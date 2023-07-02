package main

type cactus struct {
	position position
	height   int
	width    int
}

var cactusSeg = fgRgb(0, 200, 0, "#")
var cactusStages = [3]string{cactusSeg, cactusSeg, cactusSeg}

func (c *cactus) move() {
	c.position[0]--
}

func (c cactus) draw() {
	moveCursor(c.position)
	for i := 0; i <= c.height; i++ {
		moveCursor([2]int{c.position[0], c.position[1] - i})
		draw(fgRgb(0, 200, 0, cactusSeg))
	}
}

func (c cactus) checkHide() bool {
	return c.position[0] <= 0
}
