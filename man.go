package main

type man struct {
	position      position
	direction     direction
	jumpIteration int
	foots         string
	piu           []*piu
}

const jumpHeight = 7

var manHead = fgRgb(200, 200, 0, "⢠⡄")
var manBody = fgRgb(200, 200, 0, "⢺⡗")
var manFoot = fgRgb(200, 200, 0, "⢸⡇")
var manJump = fgRgb(200, 200, 0, "⡎⢱")
var manLeft = fgRgb(200, 200, 0, "⡎⡇")
var manRight = fgRgb(200, 200, 0, "⢸⢱")
var steps = [3]string{manLeft, manRight}

func (m man) draw(isDead bool) {
	if isDead {
		m.drawDeadman()
	} else {
		moveCursor(m.position)
		draw(fgRgb(200, 200, 0, m.foots))
		moveCursor([2]int{m.position[0], m.position[1] - 1})
		draw(manBody)
		moveCursor([2]int{m.position[0], m.position[1] - 2})
		draw(manHead)
	}
}

func (m man) drawDeadman() {
	moveCursor(m.position)
	draw(fgRgb(200, 0, 0, m.foots))
	moveCursor([2]int{m.position[0], m.position[1] - 1})
	draw(fgRgb(200, 0, 0, "⢺⡗"))
	moveCursor([2]int{m.position[0], m.position[1] - 2})
	draw(fgRgb(200, 0, 0, "⢠⡄"))
}

func (m *man) move(maxX int, maxY int, step bool) {
	if step {
		m.foots = steps[1]
	} else {
		m.foots = steps[0]
	}
	if m.direction == up && m.jumpIteration < jumpHeight {
		m.foots = manJump
		m.moveUp()
		m.jumpIteration++
		if m.jumpIteration == jumpHeight {
			m.direction = down
		}
	}
	if m.direction == down && m.position[1] < maxY {
		m.foots = manJump
		m.moveDown(maxY)
		m.jumpIteration--
		if m.position[1] == maxY {
			m.direction = stay
		}
	}
	if m.direction == left {
		m.moveLeft()
		m.direction = stay
	}
	if m.direction == right {
		m.moveRight(maxX)
		m.direction = stay
	}
}
func (m *man) moveUp() {
	m.position[1]--
}
func (m *man) moveDown(maxY int) {
	if m.position[1] < maxY {
		m.position[1]++
	}
}
func (m *man) moveRight(maxX int) {
	if m.position[0] < maxX {
		m.position[0]++
	}
}
func (m *man) moveLeft() {
	if m.position[0] > 0 {
		m.position[0]--
	}
}

func (m man) checkAssOnCactus(c *cactus) bool {
	return m.position[0] == c.position[0] &&
		m.position[1] >= c.position[1]-c.height
}

func (m *man) piuPiu() {
	m.piu = append(m.piu, &piu{position: [2]int{m.position[0] + 1, m.position[1] - 1}})
}

func (m *man) delPiu(key int) {
	// Remove the element at index i from a.
	m.piu[key] = m.piu[len(m.piu)-1] // Copy last element to index i.
	m.piu = m.piu[:len(m.piu)-1]     // Truncate slice.
}
