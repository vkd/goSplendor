package splendor

const (
	MAX_PLAYERS = 4
)

type Board struct {
	Skulls Skulls

	Players            []*Player
	CurrentPlayer      *Player
	currentPlayerIndex int
}

func (b *Board) Initialize() {
	b.Skulls = Skulls{
		SBlack: 7,
		SBlue:  7,
		SGreen: 7,
		SRed:   7,
		SWhite: 7,
		SGold:  5,
	}

	b.Players = make([]*Player, 4)
	for i := 0; i < 4; i++ {
		b.Players[i] = NewPlayer()
	}
	b.CurrentPlayer = b.Players[b.currentPlayerIndex]
}

func (b *Board) CheckSkulls(skulls ...SkullType) bool {
	for _, st := range skulls {
		if st == SGold {
			return false
		}
		if b.Skulls[st] <= 0 {
			return false
		}
	}
	if b.CurrentPlayer.skulls.Count()+len(skulls) > 10 {
		return false
	}
	switch len(skulls) {
	case 1:
		// if b.Skulls[skulls[0]] <= 0 {
		// 	return false
		// }
		// return skulls[0] != SkullYellow
		return true
	case 2:
		if b.Skulls[skulls[0]] < 4 {
			return false
		}
		return skulls[0] == skulls[1]
	case 3:
		if skulls[0] == skulls[1] {
			return false
		}
		if skulls[0] == skulls[2] {
			return false
		}
		if skulls[1] == skulls[2] {
			return false
		}
		return true
	}
	return false
}

func (b *Board) GetSkulls(skulls ...SkullType) {
	if b.CheckSkulls(skulls...) {
		for _, s := range skulls {
			b.Skulls[s]--
		}
		b.CurrentPlayer.AddSkulls(skulls)

		// b.currentPlayerIndex = (b.currentPlayerIndex + 1) % MAX_PLAYERS
		b.CurrentPlayer = b.Players[b.currentPlayerIndex]
	}
}
