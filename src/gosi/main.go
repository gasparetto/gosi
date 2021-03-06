package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/veandco/go-sdl2/sdl"
	"gosi/videocards"
	"gosi/processors"
	"gosi/processors/i8085"
	"gosi/processors/mb14241"
	"strconv"
)

var cpu processors.Cpu
var vdc videocards.DisplayController
var breakpoints []uint16
//var breakpoints = []uint16 {0x0000}

const PORT0_DEF uint8 = 0x0e

var port0 = PORT0_DEF

const PORT1_DEF uint8 = 0x09

var port1 = PORT1_DEF

func main() {

	//defer profile.Start().Stop()

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
	readRom("res/invaders/invaders.h", buffer)
	readRom("res/invaders/invaders.g", buffer)
	readRom("res/invaders/invaders.f", buffer)
	readRom("res/invaders/invaders.e", buffer)
	fmt.Printf("> read %d bytes\n", buffer.Len())

	cpu = new(i8085.I8085)

	sdlVdc := new(videocards.SdlDisplayController)
	sdlVdc.Renderer(renderer)
	vdc = sdlVdc

	cpu.AttachRom(buffer.Bytes(), 0x0000)

	ram := make([]byte, 0x2000)
	cpu.AttachRam(ram, 0x2000)
	vdc.AttachRam(ram[0x0400:])

	var shifter = new(mb14241.MB14241)
	cpu.AttachPortIn(getPort0, 0x00)
	cpu.AttachPortIn(getPort1, 0x01)
	cpu.AttachPortIn(shifter.ReadResult, 0x03)
	cpu.AttachPortOut(shifter.WriteCount, 0x02)
	cpu.AttachPortOut(shifter.WriteData, 0x04)

	mainLoop()
}

func getPort0() uint8 {
	//fmt.Printf("read port 0: 0x%02x\n", port0)
	v := port0
	port0 = PORT0_DEF
	return v
}

func getPort1() uint8 {
	//fmt.Printf("read port 1: 0x%02x\n", port1)
	v := port1
	port1 = PORT1_DEF
	return v
}

func mainLoop() {

	cycles := 0
	frames := 0
	nextFps := time.Now().Add(time.Second)

	// timing and interrupts
	int64V := false
	int128V := false

	running := true
	for running {

		pc := cpu.GetProgramCounter()

		// check for breakpoints
		for i, breakpoint := range breakpoints {
			if breakpoint == pc {
				fmt.Printf("--- Stopped at breakpoint %d ---\n", i+1)
				debugLoop()
			}
		}

		//cpu.DebugPrintNextOperation()
		//cpu.DebugPrintInternalState()

		cycles += cpu.Step()

		// interrupt 1
		// 320*128=40960 pixels -> /2,5=16384 cpu clock cycles
		if cycles > 16384 && !int64V {
			int64V = true
			cycles += cpu.Interrupt(0xcf) // RST 1
		}

		// interrupt 2 & vertical blank start
		// 320*224=71680 pixels -> /2,5=28672 cpu clock cycles
		if cycles > 28672 && !int128V {
			int128V = true
			cycles += cpu.Interrupt(0xd7) // RST 2

			vdc.Render()
		}

		// vertical total
		// 320*262=83840 pixels -> /2,5=33536 cpu clock cycles
		if cycles > 33536 {
			cycles = 0
			frames++
			//fmt.Printf("%s --- Frame %d ---\n", time.Now(), frames)
			int64V = false
			int128V = false

			switch checkKeyPressed() {
			case 1: // quit
				running = false
			case 2: // coin
				port1 = port1 & 0xfe
			case 3: // start 1
				port1 = port1 | 0x04
			case 4: // shot
				port0 = port0 | 0x10
			case 5: // left
				port0 = port0 | 0x20
			case 6: // right
				port0 = port0 | 0x40
			}

			time.Sleep(12 * time.Millisecond)
			//time.Sleep(500 * time.Millisecond)
		}

		// fps
		now := time.Now()
		if now.After(nextFps) {
			//fmt.Printf("%s --- FPS %d ---\n", time.Now().Format(time.StampMicro), frames)
			frames = 0
			nextFps = now.Add(time.Second)
		}
	}
}

func checkKeyPressed() int {

	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return 1
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_ESCAPE {
				return 1
			}
			if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_5 {
				fmt.Println("coin inserted")
				return 2 //fixme
			}
			if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_1 {
				fmt.Println("player 1 start")
				return 3 //fixme
			}
			if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_SPACE {
				fmt.Println("player 1 fire")
				return 4 //fixme
			}
			if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_LEFT {
				fmt.Println("player 1 left")
				return 5 //fixme
			}
			if t.Type == sdl.KEYDOWN && t.Keysym.Sym == sdl.K_RIGHT {
				fmt.Println("player 1 right")
				return 6 //fixme
			}
		}
	}
	return 0
}

func debugLoop() {

	cpu.DebugPrintNextOperation()
	cpu.DebugPrintInternalState()
	vdc.Render()

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
				if count > 0 {
					for i := 0; i < count; i++ {
						cpu.Step()
						cpu.DebugPrintNextOperation()
					}
					cpu.DebugPrintInternalState()
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
