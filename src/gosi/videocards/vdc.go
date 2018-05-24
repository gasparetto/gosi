package videocards

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

// https://github.com/faiface/pixel

type DisplayController struct{}

var ram []byte

var renderer *sdl.Renderer

func (DisplayController) Render() {

	//fmt.Println("--- VIDEO DISPLAY CONTROLLER :: RENDER ---")

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)

	//renderer.DrawLine(112, 40, 184, 216)
	//renderer.DrawLine(184, 216, 40, 216)
	//renderer.DrawLine(40, 216, 112, 40)

	var x, y int32
	for y = 0; y < 256; y++ {
		for x = 0; x < 28; x++ { // 224/8=24
			b := ram[x+y*28]
			switch {
			case (b & 0x01) != 0:
				renderer.DrawPoint(x*8, y)
			case (b & 0x02) != 0:
				renderer.DrawPoint(x*8+1, y)
			case (b & 0x04) != 0:
				renderer.DrawPoint(x*8+2, y)
			case (b & 0x08) != 0:
				renderer.DrawPoint(x*8+3, y)
			case (b & 0x10) != 0:
				renderer.DrawPoint(x*8+4, y)
			case (b & 0x20) != 0:
				renderer.DrawPoint(x*8+5, y)
			case (b & 0x40) != 0:
				renderer.DrawPoint(x*8+6, y)
			case (b & 0x80) != 0:
				renderer.DrawPoint(x*8+7, y)
			}
		}
	}

	renderer.Present()
}

func (DisplayController) Renderer(r *sdl.Renderer) {
	renderer = r
}

func (DisplayController) AttachRam(buf []byte) {

	fmt.Printf("--- VIDEO DISPLAY CONTROLLER :: ATTACH RAM (%d bytes) ---\n", len(buf))

	ram = buf
}
