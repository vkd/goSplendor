package splendor

import "github.com/veandco/go-sdl2/sdl"

type Layouter interface {
	Click(int32, int32)
	Move(int32, int32)
	Draw(*sdl.Renderer)
	Inside(int32, int32) bool
}

type Layout struct {
	sdl.Rect

	Objs []Layouter
}

func (l *Layout) Add(layout Layouter) {
	l.Objs = append(l.Objs, layout)
}

func (l *Layout) Draw(r *sdl.Renderer) {
	for i := range l.Objs {
		l.Objs[i].Draw(r)
	}
}

func (l *Layout) Click(x, y int32) {
	for i := range l.Objs {
		if l.Objs[i].Inside(x-l.Rect.X, y-l.Rect.Y) {
			l.Objs[i].Click(x-l.Rect.X, y-l.Rect.Y)
		}
	}
}

func (l *Layout) Move(x, y int32) {
	for i := range l.Objs {
		if l.Objs[i].Inside(x-l.Rect.X, y-l.Rect.Y) {
			l.Objs[i].Move(x-l.Rect.X, y-l.Rect.Y)
		}
	}
}

func (l *Layout) Inside(x, y int32) bool {
	return l.X <= x && l.Y <= y && x < l.X+l.W && y < l.Y+l.H
}
