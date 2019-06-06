package splendor

import "github.com/veandco/go-sdl2/sdl"

type Button struct {
	Tile
	Text string
}

func (b *Button) Draw(r *sdl.Renderer) {
	r.SetDrawColor(255, 255, 255, 255)
	r.DrawRect(&b.Rect)
	draw_text(b.Text, WHITE, &sdl.Point{b.X + b.W/2, b.Y + b.H/2}, 25)
}
