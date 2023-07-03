package main

type man struct {
	position      position
	direction     direction
	jumpIteration int
	foots         string
}

const jumpHeight = 7

var manHead = fgRgb(200, 200, 0, "⢠⡄")
var manBody = fgRgb(200, 200, 0, "⢺⡗")
var manFoot = fgRgb(200, 200, 0, "⢸⡇")
var manJump = fgRgb(200, 200, 0, "⡎⢱")
var manLeft = fgRgb(200, 200, 0, "⡎⡇")
var manRight = fgRgb(200, 200, 0, "⢸⢱")
var steps = [3]string{manLeft, manRight}

func (d man) draw(isDead bool) {
	if isDead {
		d.drawDeadman()
	} else {
		moveCursor(d.position)
		draw(fgRgb(200, 200, 0, d.foots))
		moveCursor([2]int{d.position[0], d.position[1] - 1})
		draw(manBody)
		moveCursor([2]int{d.position[0], d.position[1] - 2})
		draw(manHead)
	}
}

func (d man) drawDeadman() {
	moveCursor(d.position)
	draw(fgRgb(200, 0, 0, d.foots))
	moveCursor([2]int{d.position[0], d.position[1] - 1})
	draw(fgRgb(200, 0, 0, "⢺⡗"))
	moveCursor([2]int{d.position[0], d.position[1] - 2})
	draw(fgRgb(200, 0, 0, "⢠⡄"))
}

func (d *man) move(maxX int, maxY int, step bool) {
	if step {
		d.foots = steps[1]
	} else {
		d.foots = steps[0]
	}
	if d.direction == up && d.jumpIteration < jumpHeight {
		d.foots = manJump
		d.moveUp()
		d.jumpIteration++
		if d.jumpIteration == jumpHeight {
			d.direction = down
		}
	}
	if d.direction == down && d.position[1] < maxY {
		d.foots = manJump
		d.moveDown(maxY)
		d.jumpIteration--
		if d.position[1] == maxY {
			d.direction = stay
		}
	}
	if d.direction == left {
		d.moveLeft()
		d.direction = stay
	}
	if d.direction == right {
		d.moveRight(maxX)
		d.direction = stay
	}
}
func (d *man) moveUp() {
	d.position[1]--
}
func (d *man) moveDown(maxY int) {
	if d.position[1] < maxY {
		d.position[1]++
	}
}
func (d *man) moveRight(maxX int) {
	if d.position[0] < maxX {
		d.position[0]++
	}
}
func (d *man) moveLeft() {
	if d.position[0] > 0 {
		d.position[0]--
	}
}

func (d man) checkAssOnCactus(c *cactus) bool {
	return d.position[0] == c.position[0] &&
		d.position[1] >= c.position[1]-c.height
}
