package splendor

import (
	"os"

	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	path string

	fonts map[int]*ttf.Font
}

func OpenFont(path string) (*Font, error) {
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	font := &Font{
		path:  path,
		fonts: map[int]*ttf.Font{},
	}
	// var f *ttf.Font
	// for i := 5; i < 60; i++ {
	// 	f, err = ttf.OpenFont(path, i)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	font.fonts[i] = f
	// }
	return font, nil
}

func (f *Font) Size(size int) *ttf.Font {
	font := f.fonts[size]
	if font == nil {
		var err error
		font, err = ttf.OpenFont(f.path, size)
		if err != nil {
			panic(err)
		}
		f.fonts[size] = font
	}
	return font
}

func (f *Font) Close() {
	for _, font := range f.fonts {
		font.Close()
	}
}
