package splendor

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type CardTile struct {
	Tile

	Type SkullType
	Cost Skulls

	WP int
}

func (c *CardTile) Draw(r *sdl.Renderer) {
	c.Tile.Draw(r)

	set_color(DARK_GREEN)
	r.FillRect(&sdl.Rect{
		c.X + 40,
		c.Y + 75,
		20,
		c.H - 76,
	})

	for i, st := range AllVSkulls {
		x := c.X + 42
		y := c.Y + 75 + int32(i)*20
		draw_text(fmt.Sprintf("%d", c.Cost[st]), SkullColor[st], &sdl.Point{x, y}, 25)
	}

	set_color(DARK_GREEN)
	r.FillRect(&sdl.Rect{
		c.X + 95,
		c.Y + 35,
		20,
		20,
	})

	draw_text("*", SkullColor[c.Type], &sdl.Point{c.X + 100, c.Y + 30}, 25)

	set_color(DARK_GREEN)
	r.FillRect(&sdl.Rect{
		c.X + 27,
		c.Y + 15,
		20,
		20,
	})

	draw_text(fmt.Sprintf("%d", c.WP), WHITE, &sdl.Point{c.X + 30, c.Y + 10}, 25)
}
