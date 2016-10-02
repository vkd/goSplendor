package splendor

var (
	cards_1 = []*CardTile{
		// 0
		{Type: SBlue, Cost: Skulls{SGreen: 2, SBlack: 2}},
		{Type: SWhite, Cost: Skulls{SGreen: 4}, WP: 1},
		{Type: SWhite, Cost: Skulls{SWhite: 3, SBlue: 1, SBlack: 1}},
		{Type: SBlue, Cost: Skulls{SRed: 4}, WP: 1},
		{Type: SWhite, Cost: Skulls{SBlue: 2, SGreen: 2, SBlack: 1}},
		// 5
		{Type: SBlue, Cost: Skulls{SWhite: 1, SBlack: 2}},
		{Type: SBlue, Cost: Skulls{SWhite: 1, SGreen: 1, SRed: 1, SBlack: 1}},
		{Type: SWhite, Cost: Skulls{SBlue: 1, SGreen: 2, SRed: 1, SBlack: 1}},
		{Type: SBlue, Cost: Skulls{SBlack: 3}},
		{Type: SBlue, Cost: Skulls{SBlue: 1, SGreen: 3, SRed: 1}},
		// 10
		{Type: SGreen, Cost: Skulls{SWhite: 1, SBlue: 1, SRed: 1, SBlack: 2}},
		{Type: SGreen, Cost: Skulls{SWhite: 1, SBlue: 1, SRed: 1, SBlack: 1}},
		{Type: SBlue, Cost: Skulls{SWhite: 1, SGreen: 1, SRed: 2, SBlack: 1}},
		{Type: SGreen, Cost: Skulls{SBlue: 2, SRed: 2}},
		{Type: SGreen, Cost: Skulls{SWhite: 2, SBlue: 1}},
		// 15
		{Type: SBlue, Cost: Skulls{SWhite: 1, SGreen: 2, SRed: 2}},
		{Type: SGreen, Cost: Skulls{SBlack: 4}, WP: 1},
		{Type: SGreen, Cost: Skulls{SRed: 3}},
		{Type: SRed, Cost: Skulls{SBlue: 2, SGreen: 1}},
		{Type: SRed, Cost: Skulls{SWhite: 2, SRed: 2}},
		// 20
		{Type: SRed, Cost: Skulls{SWhite: 2, SGreen: 1, SBlack: 2}},
		{Type: SGreen, Cost: Skulls{SBlue: 1, SRed: 2, SBlack: 2}},
		{Type: SRed, Cost: Skulls{SWhite: 3}},
		{Type: SRed, Cost: Skulls{SWhite: 4}, WP: 1},
		{Type: SGreen, Cost: Skulls{SWhite: 1, SBlue: 3, SGreen: 1}},
		// 25
		{Type: SRed, Cost: Skulls{SWhite: 1, SBlue: 1, SGreen: 1, SBlack: 1}},
		{Type: SRed, Cost: Skulls{SWhite: 2, SBlue: 1, SGreen: 1, SBlack: 1}},
		{Type: SBlack, Cost: Skulls{SGreen: 3}},
		{Type: SBlack, Cost: Skulls{SBlue: 4}, WP: 1},
		{Type: SBlack, Cost: Skulls{SGreen: 1, SRed: 3, SBlack: 1}},
		// 30
		{Type: SRed, Cost: Skulls{SWhite: 1, SRed: 1, SBlack: 3}},
		{Type: SBlack, Cost: Skulls{SWhite: 1, SBlue: 1, SGreen: 1, SRed: 1}},
		{Type: SBlack, Cost: Skulls{SWhite: 1, SBlue: 2, SGreen: 1, SRed: 1}},
		{Type: SBlack, Cost: Skulls{SGreen: 2, SRed: 1}},
		{Type: SBlack, Cost: Skulls{SWhite: 2, SGreen: 2}},
		// 35
		{Type: SBlack, Cost: Skulls{SWhite: 2, SBlue: 2, SRed: 1}},
		{Type: SWhite, Cost: Skulls{SRed: 2, SBlack: 1}},
		{Type: SWhite, Cost: Skulls{SBlue: 2, SBlack: 2}},
		{Type: SWhite, Cost: Skulls{SBlue: 3}},
		{Type: SWhite, Cost: Skulls{SBlue: 1, SGreen: 1, SRed: 1, SBlack: 1}},
		// 40
	}
	cards_2 = []*CardTile{
		// 0
		{Type: SWhite, Cost: Skulls{SWhite: 6}, WP: 3},
		{Type: SWhite, Cost: Skulls{SWhite: 2, SBlue: 3, SRed: 3}, WP: 1},
		{Type: SBlue, Cost: Skulls{SGreen: 6}, WP: 3},
		{Type: SBlack, Cost: Skulls{SGreen: 5, SRed: 3}, WP: 2},
		{Type: SWhite, Cost: Skulls{SGreen: 3, SRed: 2, SBlack: 2}, WP: 1},
		// 5
		{Type: SWhite, Cost: Skulls{SRed: 5, SBlack: 3}, WP: 2},
		{Type: SWhite, Cost: Skulls{SRed: 5}, WP: 2},
		{Type: SWhite, Cost: Skulls{SGreen: 1, SRed: 4, SBlack: 2}, WP: 2},
		{Type: SBlue, Cost: Skulls{SGreen: 5}, WP: 2},
		{Type: SBlue, Cost: Skulls{SWhite: 4, SBlue: 2, SBlack: 1}, WP: 2},
		// 10
		{Type: SGreen, Cost: Skulls{SGreen: 6}, WP: 3},
		{Type: SGreen, Cost: Skulls{SWhite: 4, SBlue: 2, SBlack: 1}, WP: 2},
		{Type: SBlue, Cost: Skulls{SWhite: 2, SBlue: 3, SBlack: 2}, WP: 1},
		{Type: SBlue, Cost: Skulls{SBlue: 5, SGreen: 3}, WP: 2},
		{Type: SGreen, Cost: Skulls{SWhite: 2, SBlue: 3, SBlack: 2}, WP: 1},
		// 15
		{Type: SBlue, Cost: Skulls{SWhite: 3, SGreen: 2, SRed: 3}, WP: 1},
		{Type: SGreen, Cost: Skulls{SGreen: 5}, WP: 2},
		{Type: SGreen, Cost: Skulls{SWhite: 3, SGreen: 2, SRed: 3}, WP: 1},
		{Type: SRed, Cost: Skulls{SRed: 6}, WP: 3},
		{Type: SBlack, Cost: Skulls{SBlack: 6}, WP: 3},
		// 20
		{Type: SRed, Cost: Skulls{SBlue: 3, SRed: 2, SBlack: 3}, WP: 1},
		{Type: SGreen, Cost: Skulls{SBlue: 5, SGreen: 3}, WP: 2},
		{Type: SRed, Cost: Skulls{SWhite: 3, SBlack: 5}, WP: 2},
		{Type: SRed, Cost: Skulls{SWhite: 2, SRed: 3, SBlack: 3}, WP: 1},
		{Type: SRed, Cost: Skulls{SBlack: 5}, WP: 2},
		// 25
		{Type: SBlack, Cost: Skulls{SBlack: 5}, WP: 2},
		{Type: SRed, Cost: Skulls{SWhite: 1, SBlue: 4, SGreen: 2}, WP: 2},
		{Type: SBlack, Cost: Skulls{SWhite: 3, SBlue: 2, SGreen: 2}, WP: 1},
		{Type: SBlack, Cost: Skulls{SWhite: 3, SGreen: 3, SBlack: 2}, WP: 1},
		{Type: SBlack, Cost: Skulls{SBlue: 1, SGreen: 4, SRed: 2}, WP: 2},
		// 30
	}
	cards_3 = []*CardTile{
		// 0
		{Type: SGreen, Cost: Skulls{SWhite: 3, SBlue: 6, SGreen: 3}, WP: 4},
		{Type: SWhite, Cost: Skulls{SWhite: 3, SRed: 3, SBlack: 6}, WP: 4},
		{Type: SBlue, Cost: Skulls{SWhite: 6, SBlue: 3, SBlack: 3}, WP: 4},
		{Type: SWhite, Cost: Skulls{SBlue: 3, SGreen: 3, SRed: 5, SBlack: 3}, WP: 3},
		{Type: SWhite, Cost: Skulls{SBlack: 7}, WP: 4},
		// 5
		{Type: SRed, Cost: Skulls{SWhite: 3, SBlue: 5, SGreen: 3, SBlack: 3}, WP: 3},
		{Type: SBlue, Cost: Skulls{SWhite: 7, SBlue: 3}, WP: 5},
		{Type: SWhite, Cost: Skulls{SWhite: 3, SBlack: 7}, WP: 5},
		{Type: SBlue, Cost: Skulls{SWhite: 7}, WP: 4},
		{Type: SGreen, Cost: Skulls{SBlue: 7}, WP: 4},
		// 10
		{Type: SRed, Cost: Skulls{SGreen: 7, SRed: 3}, WP: 5},
		{Type: SBlack, Cost: Skulls{SRed: 7, SBlack: 3}, WP: 5},
		{Type: SBlue, Cost: Skulls{SWhite: 3, SGreen: 3, SRed: 3, SBlack: 5}, WP: 3},
		{Type: SRed, Cost: Skulls{SGreen: 7}, WP: 4},
		{Type: SRed, Cost: Skulls{SBlue: 3, SGreen: 6, SRed: 3}, WP: 4},
		// 15
		{Type: SGreen, Cost: Skulls{SWhite: 5, SBlue: 3, SRed: 3, SBlack: 3}, WP: 3},
		{Type: SGreen, Cost: Skulls{SBlue: 7, SGreen: 3}, WP: 5},
		{Type: SBlack, Cost: Skulls{SRed: 7}, WP: 4},
		{Type: SBlack, Cost: Skulls{SGreen: 3, SRed: 6, SBlack: 3}, WP: 4},
		{Type: SBlack, Cost: Skulls{SWhite: 3, SBlue: 3, SGreen: 5, SRed: 3}, WP: 3},
	}
	goals = []*Goal{
		// 0
		{Cost: Skulls{SBlue: 3, SGreen: 3, SRed: 3}, WP: 3},
		{Cost: Skulls{SWhite: 4, SBlue: 4}, WP: 3},
		{Cost: Skulls{SWhite: 3, SBlue: 3, SBlack: 3}, WP: 3},
		{Cost: Skulls{SRed: 4, SBlack: 4}, WP: 3},
		{Cost: Skulls{SGreen: 4, SRed: 4}, WP: 3},
		// 5
		{Cost: Skulls{SBlue: 4, SGreen: 4}, WP: 3},
		{Cost: Skulls{SGreen: 3, SRed: 3, SBlack: 3}, WP: 3},
		{Cost: Skulls{SWhite: 3, SRed: 3, SBlack: 3}, WP: 3},
		{Cost: Skulls{SWhite: 3, SBlue: 3, SGreen: 3}, WP: 3},
		{Cost: Skulls{SWhite: 4, SBlack: 4}, WP: 3},
	}
)
