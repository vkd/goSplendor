package splendor

import (
	"strconv"

	"github.com/veandco/go-sdl2/sdl"
)

type SkullToken struct {
	Tile

	Type SkullType

	skulls Skulls

	value_point *sdl.Point
}

func (s *SkullToken) Draw(r *sdl.Renderer) {
	s.Tile.Draw(r)
	if s.value_point == nil {
		s.value_point = &sdl.Point{s.X + 74, s.Y + 7}
	}
	draw_text(strconv.Itoa(s.skulls[s.Type]), SkullColor[s.Type], s.value_point, 40)
}
