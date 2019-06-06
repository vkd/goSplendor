package splendor

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	FPS = 60

	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 800
)

var (
	DELTA_FPS = time.Duration(int64(time.Second) / FPS)

	// font     *font_helper.Font
	renderer *sdl.Renderer

	// Elasund = core.NewElasund()

	BLACK      = &sdl.Color{0, 0, 0, 255}
	BLUE       = &sdl.Color{10, 10, 255, 255}
	DARK_GREEN = &sdl.Color{10, 100, 10, 255}
	GREEN      = &sdl.Color{100, 200, 100, 255}
	PURPLE     = &sdl.Color{100, 10, 100, 255}
	RED        = &sdl.Color{255, 100, 100, 255}
	SKY        = &sdl.Color{130, 130, 255, 255}
	WHITE      = &sdl.Color{255, 255, 255, 255}
	YELLOW     = &sdl.Color{255, 255, 100, 255}

	SkullColor = map[SkullType]*sdl.Color{
		SBlack: BLACK,
		SBlue:  SKY,
		SGreen: GREEN,
		SRed:   RED,
		SWhite: WHITE,
		SGold:  YELLOW,
	}

	TILE_SIZE        = 50
	TILE_BORDER_SIZE = 1
	TILE_STEP        = TILE_SIZE + TILE_BORDER_SIZE

	TILE_START_X = 153
	TILE_START_Y = 184

	LINE_WIDTH int32 = 3

	window *sdl.Window
	r      *sdl.Renderer
	font   *Font
)

func Run() {
	var err error

	runtime.LockOSThread()

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	window, err = sdl.CreateWindow("Splendor", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	font, err = OpenFont("../../font/cmuntb.ttf")

	// font, err = ttf.OpenFont("../font/cmuntb.ttf", 20)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	r, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer r.Destroy()

	tm := TextureManager{}
	err = tm.Initialize(r, "../../textures")
	if err != nil {
		panic(err)
	}

	desk1 := CardDeck{}
	for i := 1; i <= 40; i++ {
		c := cards_1[i-1]
		c.Texture = tm.Get("1-" + strconv.Itoa(i))
		desk1.Put(c)
	}

	line_1 := CardDeck{}
	for i := 0; i < 4; i++ {
		line_1.Put(desk1.PopRandom())
	}
	// log.Printf("desk1: %d", desk1.Len())

	desk2 := CardDeck{}
	for i := 1; i <= 30; i++ {
		c := cards_2[i-1]
		c.Texture = tm.Get("2-" + strconv.Itoa(i))
		desk2.Put(c)
	}

	line_2 := CardDeck{}
	for i := 0; i < 4; i++ {
		line_2.Put(desk2.PopRandom())
	}
	// log.Printf("desk2: %d", desk2.Len())

	desk3 := CardDeck{}
	for i := 1; i <= 20; i++ {
		c := cards_3[i-1]
		c.Texture = tm.Get("3-" + strconv.Itoa(i))
		desk3.Put(c)
	}

	line_3 := CardDeck{}
	for i := 0; i < 4; i++ {
		line_3.Put(desk3.PopRandom())
	}
	// log.Printf("desk3: %d", desk3.Len())

	hand := CardDeck{}
	hand.Put(desk1.PopRandom())
	hand.Put(desk2.PopRandom())
	hand.Put(desk3.PopRandom())

	var e sdl.Event
	var stop bool

	// t := tm.Get("1-1")
	back1 := tm.Get("1-0")
	back2 := tm.Get("2-0")
	back3 := tm.Get("3-0")
	// g1 := tm.Get("g1")

	goal_width := int32(128)
	goal_height := int32(128)

	padding := int32(10)
	start_x := int32(400)
	start_y := int32(goal_height + padding)

	width := int32(128)
	height := int32(182)

	start_y_line_1 := start_y + 2*(height+padding)
	start_y_line_2 := start_y + 1*(height+padding)
	start_y_line_3 := start_y

	mouse_x := int32(0)
	mouse_y := int32(0)

	deck_goals := CardDeck{}
	for i := 1; i <= 10; i++ {
		c := goals[i-1]
		c.Texture = tm.Get("g" + strconv.Itoa(i))
		deck_goals.Put(c)
	}

	goals := CardDeck{}
	for i := 0; i < 5; i++ {
		goals.Put(deck_goals.PopRandom())
	}

	board := Board{}
	board.Initialize()

	player_skull_x := int32(WINDOW_WIDTH / 2)
	player_skull_y := int32(WINDOW_HEIGHT - 64)
	player_skull_padding := int32(80)

	to_take_skulls := make([]SkullType, 0, 3)

	ls := []Layouter{}
	for i, st := range AllSkulls {
		skull_type := st
		ls = append(ls, &SkullToken{
			Tile: Tile{
				Rect:    sdl.Rect{50, 100 + int32(i)*90, 64, 64},
				Texture: tm.Skull(skull_type),
				OnClick: func() {
					if skull_type == SGold {
						return
					}
					if len(to_take_skulls) == 1 {
						if to_take_skulls[0] == skull_type {
							to_take_skulls = append(to_take_skulls, skull_type)
							board.GetSkulls(to_take_skulls...)
							to_take_skulls = to_take_skulls[:0]
							return
						}
					}
					for _, st := range to_take_skulls {
						if st == skull_type {
							return
						}
					}
					to_take_skulls = append(to_take_skulls, skull_type)
					if len(to_take_skulls) == 3 {
						board.GetSkulls(to_take_skulls...)
						to_take_skulls = to_take_skulls[:0]
						return
					}
					// if to_take_skulls[skull_type] > 0 {
					// 	board.GetSkulls(to_take_skulls)
					// }
					// to_take_skulls = append(to_take_skulls, skull_type)
				},
			},
			Type:   skull_type,
			skulls: board.Skulls,
		})
	}
	ls = append(ls, &Button{
		Text: "OK",
		Tile: Tile{
			Rect: sdl.Rect{50, 100 + 6*90, 64, 64},
			OnClick: func() {
				board.GetSkulls(to_take_skulls...)
				to_take_skulls = to_take_skulls[:0]
			},
		},
	})
	ls = append(ls, &Button{
		Text: "Clear",
		Tile: Tile{
			Rect: sdl.Rect{50, 100 + 7*90, 64, 64},
			OnClick: func() {
				// board.GetSkulls(to_take_skulls.ToArray()...)
				to_take_skulls = to_take_skulls[:0]
			},
		},
	})
	root := Layout{
		sdl.Rect{0, 0, 260, WINDOW_HEIGHT},
		ls,
	}

	now := time.Now()

	for range time.Tick(DELTA_FPS) {
		// r.SetDrawColor(100, 10, 100, 255)
		r.SetDrawColor(10, 100, 10, 255)
		r.Clear()

		for e = sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e := e.(type) {
			case *sdl.QuitEvent:
				stop = true
			case *sdl.KeyboardEvent: // KeyDownEvent:
				if e.Keysym.Sym == sdl.K_ESCAPE {
					stop = true
				}
				if e.Keysym.Sym == sdl.K_q {
					stop = true
				}
			case *sdl.MouseMotionEvent:
				mouse_x = e.X
				mouse_y = e.Y
				root.Move(e.X, e.Y)
			case *sdl.MouseWheelEvent:
				// player_skull_padding -= e.Y
				// player_skull_x -= 4 * e.X
			case *sdl.MouseButtonEvent:
				if e.Button == sdl.BUTTON_RIGHT {
					log.Printf("right: %v", e.Type)
				}
				if e.Type == sdl.MOUSEBUTTONDOWN {
					if e.Button == sdl.BUTTON_LEFT {
						root.Click(e.X, e.Y)
						// 		if e.X >= 50 && e.X <= 50+64 {
						// 			if e.Y >= 100 && e.Y <= 100+64 {
						// 				board.CurrentPlayer.AddSkulls([]SkullType{SkullBlack})
						// 			}
						// 		}
					}
				}
			}
		}

		for i := 0; i < line_1.Len(); i++ {
			c := line_1.Get(i).(*CardTile)
			c.Rect = sdl.Rect{
				start_x + int32(i)*(width+padding),
				start_y_line_1,
				width,
				height,
			}
			c.Draw(r)
		}
		for i := 0; i < line_2.Len(); i++ {
			c := line_2.Get(i).(*CardTile)
			c.Rect = sdl.Rect{
				start_x + int32(i)*(width+padding),
				start_y_line_2,
				width,
				height,
			}
			c.Draw(r)
		}
		for i := 0; i < line_3.Len(); i++ {
			c := line_3.Get(i).(*CardTile)
			c.Rect = sdl.Rect{
				start_x + int32(i)*(width+padding),
				start_y_line_3,
				width,
				height,
			}
			c.Draw(r)
		}

		back_padding := int32(4)
		per_card := float64(4)
		for i := 0; float64(i) < float64(desk1.Len())/per_card; i++ {
			r.Copy(back1, nil, &sdl.Rect{
				start_x - width - padding - int32(i)*back_padding,
				start_y + 2*(height+padding) + int32(i)*back_padding,
				width,
				height,
			})
		}
		for i := 0; float64(i) < float64(desk2.Len())/per_card; i++ {
			r.Copy(back2, nil, &sdl.Rect{
				start_x - width - padding - int32(i)*back_padding,
				start_y + 1*(height+padding) + int32(i)*back_padding,
				width,
				height,
			})
		}
		for i := 0; float64(i) < float64(desk3.Len())/per_card; i++ {
			r.Copy(back3, nil, &sdl.Rect{
				start_x - width - padding - int32(i)*back_padding,
				start_y + int32(i)*back_padding,
				width,
				height,
			})
		}

		for i := 0; i < goals.Len(); i++ {
			g := goals.Get(i).(*Goal)
			g.Rect = sdl.Rect{
				start_x - (goal_width + padding) + int32(i)*(goal_width+padding),
				start_y - (goal_height + padding),
				goal_width,
				goal_height,
			}
			g.Draw(r)
		}

		for i := 0; i < hand.Len(); i++ {
			r.Copy(hand.Get(i).(*CardTile).Texture, nil, &sdl.Rect{
				1050 + int32(i)*41,
				330 + int32(i)*35,
				width,
				height,
			})
		}

		// r.Copy(tm.Get("t-blue"), nil, &sdl.Rect{
		// 	760,
		// 	720,
		// 	64,
		// 	64,
		// })

		// draw_text("0", SKY, &sdl.Point{834, 732}, 30)

		// draw_text("Mouse:", sdl.Color{100, 100, 255, 255}, &sdl.Point{50, 50}, 30)

		fps := 1 / time.Since(now).Seconds()
		now = time.Now()
		draw_text(fmt.Sprintf("fps: %f", fps), WHITE, &sdl.Point{0, 0}, 15)

		draw_text(fmt.Sprintf("mouse: (%d;%d)", mouse_x, mouse_y), WHITE, &sdl.Point{0, 15}, 15)
		draw_text(fmt.Sprintf("to_take_skulls: %v", to_take_skulls), WHITE, &sdl.Point{0, 30}, 15)

		draw_text(fmt.Sprintf("player_index: %d", board.currentPlayerIndex), WHITE, &sdl.Point{1000, 0}, 15)
		for i := 0; i < MAX_PLAYERS; i++ {
			draw_text(fmt.Sprintf("%s", board.Players[i].skulls.String()), WHITE, &sdl.Point{1000, 15 + int32(i)*15}, 15)
		}
		// draw_text(fmt.Sprintf("player_skull_x: %d", player_skull_x), WHITE, &sdl.Point{0, 45}, 15)
		// draw_text(fmt.Sprintf("player_skull_y: %d", player_skull_y), WHITE, &sdl.Point{0, 60}, 15)

		// for i, st := range AllSkulls {
		// 	r.Copy(tm.Skull(st), nil, &sdl.Rect{
		// 		50,
		// 		100 + int32(i)*90,
		// 		64,
		// 		64,
		// 	})
		// 	draw_text(strconv.Itoa(board.Skulls[st]), SkullColor[st], &sdl.Point{50 + 64 + 10, 100 + int32(i)*90 + 7}, 40)
		// }

		for j, st := range AllSkulls {
			i := j - 3
			r.Copy(tm.Skull(st), nil, &sdl.Rect{
				player_skull_x + int32(i)*(64+player_skull_padding),
				player_skull_y,
				64,
				64,
			})
			draw_text(
				fmt.Sprintf("%d|%d", board.CurrentPlayer.Skull(st), 0),
				SkullColor[st],
				&sdl.Point{
					player_skull_x + int32(i)*(64+player_skull_padding) + 64 + 10,
					player_skull_y + 17,
				},
				30,
			)
		}

		root.Draw(r)

		r.Present()

		if stop {
			break
		}
	}

	sdl.Quit()
}
