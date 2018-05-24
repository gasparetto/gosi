package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"github.com/veandco/go-sdl2/sdl"
	"gosi/videocards"
	"gosi/processors"
	"strconv"
)

func main() {

	fmt.Println("> start")

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("GOSI", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		256, 224, sdl.WINDOW_SHOWN)
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
	readRom("res/invaders/invaders.h", buffer)
	readRom("res/invaders/invaders.g", buffer)
	readRom("res/invaders/invaders.f", buffer)
	readRom("res/invaders/invaders.e", buffer)
	fmt.Printf("> read %d bytes\n", buffer.Len())

	var cpu = new(processors.I8085)

	var vdc = new(videocards.SdlDisplayController)
	vdc.Renderer(renderer)

	cpu.AttachRom(buffer.Bytes(), 0x0000)

	ram := make([]byte, 0x2000)
	cpu.AttachRam(ram, 0x2000)
	vdc.AttachRam(ram[0x0400:])

	mainLoop(cpu, vdc)
	//debugLoop(cpu, vdc)

	fmt.Println("> end")
}

func mainLoop(cpu processors.Cpu, vdc videocards.DisplayController) {

	cycles := 0

	running := true
	int64V := false
	int128V := false
	for running {

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
					break
				}
			}
		}

		cycles += cpu.Step()

		// 320*128=40960 pixels -> /2,5=16384 cpu clock cycles
		if cycles > 16384 && !int64V {

			int64V = true
			cycles += cpu.Interrupt(0xcf) // RST 1
		}

		// 320*218=69760 pixels -> /2,5=27904 cpu clock cycles
		if cycles > 27904 && !int128V {

			int128V = true
			cycles += cpu.Interrupt(0xd7) // RST 2
		}

		// 320*262=83840 pixels -> /2,5=33536 cpu clock cycles
		if cycles > 33536 {

			cycles = 0
			int64V = false
			int128V = false

			vdc.Render()
		}

		if cpu.DebugBreakpoint() {
			fmt.Println("--- Stopped at breakpoint ---")
			cpu.DebugPrint()
			vdc.Render()
			//vdc.DasmTrace()
			debugLoop(cpu, vdc)
		}
	}
}

func debugLoop(cpu processors.Cpu, vdc videocards.DisplayController) {

	running := true
	for running {

		event := sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.QuitEvent:
			running = false
			break
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
					break
				}
				count := getIntFromKeycode(t.Keysym.Sym)
				for i := 0; i < count; i++ {
					cpu.Step()
				}
				if count > 0 {
					cpu.DebugPrint()
					vdc.Render()
				}
			}
			break
		}
	}
}

func readRom(name string, buffer *bytes.Buffer) {

	fmt.Printf("> load rom file: %s\n", name)

	filerc, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buffer.ReadFrom(filerc)
}

func getIntFromKeycode(kc sdl.Keycode) int {
	if kc == sdl.K_z {
		return 100
	}
	if kc == sdl.K_x {
		return 500
	}
	if kc == sdl.K_c {
		return 1000
	}
	if kc == sdl.K_v {
		return 5000
	}
	if kc == sdl.K_b {
		return 10000
	}
	keyName := sdl.GetKeyName(kc)
	digit, err := strconv.ParseInt(keyName, 10, 8)
	if err != nil {
		return 0
	}
	return int(digit)
}
