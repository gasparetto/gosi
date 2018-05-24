/*
	https://github.com/a8m/go-lang-cheat-sheet
	https://github.com/veandco/go-sdl2
*/

/*
	http://computerarcheology.com/Arcade/SpaceInvaders/Hardware.html
	https://www.dorkbotpdx.org/blog/skinny/use_mames_debugger_to_reverse_engineer_and_extend_old_games

	$ ./mame64 -window -nomaximize -debug roms/invaders.zip

	https://github.com/mamedev/mame/blob/master/src/mame/machine/mw8080bw.cpp

	Midway 8080-based black and white hardware

	https://github.com/mamedev/mame/blob/master/src/mame/drivers/mw8080bw.cpp

    MAIN CPU memory address space
    Address (15-bits) Dir Data     Description
    ----------------- --- -------- -----------------------
    x0xxxxxxxxxxxxx   R   xxxxxxxx Program ROM (various amounts populated)
    -1xxxxxxxxxxxxx   R/W xxxxxxxx Video RAM (256x256x1 bit display)
                                   Portion in VBLANK region used as work RAM

	2552: Space Invaders (PCB #739)

	PCB #        year  rom         parent    machine   inp       init              monitor,            company,          fullname,                            flags
	739   GAMEL( 1978, invaders,   0,        invaders, invaders, mw8080bw_state,   empty_init, ROT270, "Taito / Midway", "Space Invaders / Space Invaders M", MACHINE_SUPPORTS_SAVE, layout_invaders )

	ROM_REGION( 0x10000, "maincpu", 0 )
	ROM_LOAD( "invaders.h", 0x0000, 0x0800, CRC(734f5ad8) SHA1(ff6200af4c9110d8181249cbcef1a8a40fa40b7f) )
	ROM_LOAD( "invaders.g", 0x0800, 0x0800, CRC(6bfaca4a) SHA1(16f48649b531bdef8c2d1446c429b5f414524350) )
	ROM_LOAD( "invaders.f", 0x1000, 0x0800, CRC(0ccead96) SHA1(537aef03468f63c5b9e11dd61e253f7ae17d9743) )
	ROM_LOAD( "invaders.e", 0x1800, 0x0800, CRC(14e538b0) SHA1(1d6ca0c99f9df71e2990b610deb9d7da0125e2d8) )

	https://github.com/mamedev/mame/blob/master/src/mame/video/mw8080bw.cpp
	https://github.com/mamedev/mame/blob/master/src/mame/audio/mw8080bw.cpp
	https://github.com/mamedev/mame/blob/master/src/mame/layout/invaders.lay

	https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.h
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"github.com/veandco/go-sdl2/sdl"
	"gosi/videocards"
	"gosi/processors"
)

func main() {

	fmt.Println("> start")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("GOSI", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		224, 256, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	fmt.Println("> load buffer files")
	buffer := new(bytes.Buffer)
	read("res/invaders/invaders.h", buffer)
	read("res/invaders/invaders.g", buffer)
	read("res/invaders/invaders.f", buffer)
	read("res/invaders/invaders.e", buffer)
	fmt.Printf("> read %d bytes\n", buffer.Len())

	var cpu = new(processors.I8085)

	var vdc = new(videocards.DisplayController)
	vdc.Renderer(renderer)

	cpu.AttachRom(buffer.Bytes(), 0x0000)

	ram := make([]byte, 0x2000)
	cpu.AttachRam(ram, 0x2000)
	vdc.AttachRam(ram[0x0400:])

	cycles := 0

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		cycles += cpu.Step()
		fmt.Printf("cycles: %d\n", cycles)

		if cycles > 12000 { //fixme 12000

			vdc.Render()

			//window.UpdateSurface()

			cycles = 0

			//log.Fatal("quit")
		}
	}

	fmt.Println("> end")
}

func read(name string, buffer *bytes.Buffer) {

	fmt.Printf("> load rom file: %s\n", name)

	filerc, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buffer.ReadFrom(filerc)
}
