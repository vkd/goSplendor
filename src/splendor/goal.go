package splendor

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Goal struct {
	Tile

	Cost Skulls
	WP   int
}

func (g *Goal) Draw(r *sdl.Renderer) {
	g.Tile.Draw(r)

	for i, st := range AllVSkulls {
		x := g.X + 50 + int32(i)*15
		y := g.Y
		draw_text(fmt.Sprintf("%d", g.Cost[st]), SkullColor[st], &sdl.Point{x, y}, 25)
		i++
	}
}
