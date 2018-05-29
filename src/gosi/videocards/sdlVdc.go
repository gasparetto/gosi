package videocards

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type SdlDisplayController struct{}

var ram []byte

var renderer *sdl.Renderer

func (SdlDisplayController) Renderer(r *sdl.Renderer) {
	renderer = r
}

func (SdlDisplayController) Render() {

	//fmt.Println("--- VIDEO DISPLAY CONTROLLER :: RENDER ---")

	renderer.SetDrawColor(0, 0, 0, sdl.ALPHA_OPAQUE)
	renderer.Clear()

	renderer.SetDrawColor(255, 255, 255, sdl.ALPHA_OPAQUE)

	var x, y int32

	for y = 0; y < 224; y++ {
		for x = 0; x < 32; x++ { // 256/8=32
			b := ram[x+y*32]
			if b != 0 {
				renderByte(b, x, y)
			}
		}
	}

	renderer.Present()
}

func renderByte(b uint8, x int32, y int32) {

	x8 := x * 8
	if (b & 0x01) != 0 {
		renderPixel(x8, y)
	}
	if (b & 0x02) != 0 {
		renderPixel(x8+1, y)
	}
	if (b & 0x04) != 0 {
		renderPixel(x8+2, y)
	}
	if (b & 0x08) != 0 {
		renderPixel(x8+3, y)
	}
	if (b & 0x10) != 0 {
		renderPixel(x8+4, y)
	}
	if (b & 0x20) != 0 {
		renderPixel(x8+5, y)
	}
	if (b & 0x40) != 0 {
		renderPixel(x8+6, y)
	}
	if (b & 0x80) != 0 {
		renderPixel(x8+7, y)
	}
}

func renderPixel(x int32, y int32) {

	//renderer.DrawPoint(x, y)
	renderer.DrawPoint(y, 256 - x) // rotate 90° left
}

func (SdlDisplayController) AttachRam(buf []byte) {

	fmt.Printf("--- VIDEO DISPLAY CONTROLLER :: ATTACH RAM (%d bytes) ---\n", len(buf))

	ram = buf
}

func (SdlDisplayController) DasmTrace() {

	for y := 0; y < 224; y++ {
		x := y * 32
		addr := 0x2400 + x
		fmt.Printf(
			"%04x   %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x %02x\n",
			addr,
			ram[x], ram[x+1], ram[x+2], ram[x+3], ram[x+4], ram[x+5], ram[x+6], ram[x+7], ram[x+8], ram[x+9],
			ram[x+10], ram[x+11], ram[x+12], ram[x+13], ram[x+14], ram[x+15], ram[x+16], ram[x+17], ram[x+18], ram[x+19],
			ram[x+20], ram[x+21], ram[x+22], ram[x+23], ram[x+24], ram[x+25], ram[x+26], ram[x+27], ram[x+28], ram[x+29],
			ram[x+30], ram[x+31])
	}
}
