package entity

type Paddle struct {
	X     float64
	Y     float64
	width int
}

func (p *Paddle) MoveDown(delta float64) {
	p.Y += delta
}

func (p *Paddle) MoveUp(delta float64) {
	p.Y -= delta
}
