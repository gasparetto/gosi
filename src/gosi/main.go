/*
	https://github.com/a8m/go-lang-cheat-sheet
*/

/*
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
	"fmt"
	"os"
	"log"
	"bytes"
	"./cpu"
)

func main() {

	fmt.Println("> start")

	fmt.Println("> load buffer files")
	buffer := new(bytes.Buffer)
	read("res/invaders/invaders.h", buffer)
	read("res/invaders/invaders.g", buffer)
	read("res/invaders/invaders.f", buffer)
	read("res/invaders/invaders.e", buffer)
	fmt.Printf("> read %d bytes\n", buffer.Len())

	var c = new(cpu.I8085)

	c.AttachRom(buffer.Bytes(), 0x0000)

	c.AttachRam(make([]byte, 0x2000), 0x2000)

	for i := 0; i < 20; i++ {

		c.Step()
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
