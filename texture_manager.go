package splendor

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/veandco/go-sdl2/img"

	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	tm     map[string]*sdl.Texture
	skulls map[SkullType]*sdl.Texture

	r *sdl.Renderer
}

func (t *TextureManager) Initialize(r *sdl.Renderer, dirpath string) error {
	t.r = r
	t.tm = map[string]*sdl.Texture{}
	t.skulls = map[SkullType]*sdl.Texture{}

	fis, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		if !fi.IsDir() {
			err = t.loadTexture(path.Join(dirpath, fi.Name()), strings.Split(fi.Name(), ".")[0])
			if err != nil {
				return err
			}
		}
	}

	t.skulls[SBlack] = t.tm["t-black"]
	t.skulls[SBlue] = t.tm["t-blue"]
	t.skulls[SGreen] = t.tm["t-green"]
	t.skulls[SRed] = t.tm["t-red"]
	t.skulls[SWhite] = t.tm["t-white"]
	t.skulls[SGold] = t.tm["t-yellow"]
	return nil
}

func (t *TextureManager) Get(key string) *sdl.Texture {
	return t.tm[key]
}

func (t *TextureManager) loadTexture(filepath string, name string) error {
	s, err := img.Load(filepath)
	if err != nil {
		return err
	}

	t.tm[name], err = t.r.CreateTextureFromSurface(s)
	if err != nil {
		return err
	}
	s.Free()
	return nil
}

func (t *TextureManager) Destroy() {
	for _, tx := range t.tm {
		tx.Destroy()
	}
}

func (t *TextureManager) Skull(st SkullType) *sdl.Texture {
	return t.skulls[st]
}
