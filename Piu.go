package main

var symbol = fgRgb(200, 0, 0, "━")

type piu struct {
	position position
	deleted  bool
}

func (p *piu) move() {
	p.position[0]++
}
func (p *piu) draw() {
	if !p.deleted {
		moveCursor(p.position)
		draw(fgRgb(255, 0, 0, symbol))
	}
}

func (p *piu) checkCactus(cactuses []*cactus) bool {
	for _, c := range cactuses {
		if !p.deleted &&
			c.height > 1 &&
			p.position[0] == c.position[0] &&
			p.position[1] >= c.position[1]-c.height {
			c.height--
			p.deleted = true

			return true
		}
	}

	return false
}
