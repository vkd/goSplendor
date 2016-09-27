package splendor

type Player struct {
	skulls Skulls
}

func NewPlayer() *Player {
	p := &Player{}
	p.skulls = Skulls{}
	return p
}

func (p *Player) Skull(st SkullType) int {
	return p.skulls[st]
}

func (p *Player) AddSkulls(sts []SkullType) {
	for _, st := range sts {
		p.skulls[st]++
	}
}
