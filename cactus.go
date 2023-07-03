package main

/*
┃┃
┗┫┃ ┗┫

	┣┛  ┣┛ ┗┫
*/
type cactus struct {
	position position
	stages   []string
	height   int
}

var cactus3Stages = [][]string{{"┛", "┃", ""}, {"┣", "┫", "┃"}, {" ", "┗", "┃"}}
var cactus2Stages = [][]string{{"┛", ""}, {"┣", "┫"}, {" ", "┗"}}
var cactus1Stages = [][]string{{"┫"}, {"┗"}}

func (c *cactus) move() {
	c.position[0]--
}

func (c cactus) draw() {
	moveCursor(c.position)
	for i := 0; i < c.height; i++ {
		moveCursor([2]int{c.position[0], c.position[1] - i})
		draw(fgRgb(0, 200, 0, c.stages[i]))
	}
}

func (c cactus) checkHide() bool {
	return c.position[0] <= 0
}
