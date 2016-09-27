package splendor

import "github.com/veandco/go-sdl2/sdl"

func draw_text(text string, color *sdl.Color, p *sdl.Point, size int) {
	s, err := font.Size(size).RenderUTF8_Blended(text, *color)
	if err != nil {
		panic(err)
	}
	defer s.Free()
	text_texture, err := r.CreateTextureFromSurface(s)
	if err != nil {
		panic(err)
	}
	defer text_texture.Destroy()
	err = r.Copy(text_texture, nil, &sdl.Rect{p.X, p.Y, s.W, s.H})
	if err != nil {
		panic(err)
	}
}

func set_color(c *sdl.Color) {
	r.SetDrawColor(c.R, c.G, c.B, c.A)
}
