package splendor

import "github.com/veandco/go-sdl2/sdl"

type Tile struct {
	sdl.Rect

	Texture *sdl.Texture

	OnClick F
	OnMove  F

	is_inside bool
}

type F func()

func (t *Tile) Inside(x, y int32) bool {
	t.is_inside = t.X <= x && t.Y <= y && x < t.X+t.W && y < t.Y+t.H
	return t.is_inside
}

func (t *Tile) Draw(r *sdl.Renderer) {
	if t.Texture != nil {
		r.Copy(t.Texture, nil, &t.Rect)
	}

	if t.is_inside {
		set_color(RED)
		r.DrawRect(&t.Rect)
	}
}

func (t *Tile) Move(x, y int32) {
	if t.OnMove != nil {
		t.OnMove()
	}
}

func (t *Tile) Click(x, y int32) {
	if t.OnClick != nil {
		t.OnClick()
	}
}
