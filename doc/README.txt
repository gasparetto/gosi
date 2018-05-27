--------------------------------------------------------------------------------
GOSI :: GOlang Space Invaders emulator
--------------------------------------------------------------------------------

# GO lang

https://github.com/a8m/go-lang-cheat-sheet


# Midway 8080 black/white arcade machine

http://computerarcheology.com/Arcade/SpaceInvaders/Hardware.html

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


## Interrupt timing

MW8080BW_MASTER_CLOCK 19968000
MW8080BW_CPU_CLOCK    1996800   (MW8080BW_MASTER_CLOCK / 10)
MW8080BW_PIXEL_CLOCK  4992000   (MW8080BW_MASTER_CLOCK / 4)
MW8080BW_HTOTAL       320       (0x140)
MW8080BW_HBEND        0         (0x000)
MW8080BW_HBSTART      256       (0x100)
MW8080BW_VTOTAL       262       (0x106)
MW8080BW_VBEND        0         (0x000)
MW8080BW_VBSTART      224       (0x0e0)
MW8080BW_VCOUNTER_START_NO_VBLANK 32  (0x020)
MW8080BW_VCOUNTER_START_VBLANK    218 (0x0da)
MW8080BW_60HZ         59,541984 (MW8080BW_PIXEL_CLOCK / MW8080BW_HTOTAL / MW8080BW_VTOTAL)
MW8080BW_INT_TRIGGER_COUNT_1 128 (0x080)
MW8080BW_INT_TRIGGER_COUNT_2 218 (0x0da)

10 cpu cycles / 4 pixel cycles = 2,5

128*320=40960 -> /2,5=16384
218*320=69760 -> /2,5=27904
262*320=83840 -> /2,5=33536


# MAME

https://github.com/mamedev/mame/blob/master/src/mame/includes/mw8080bw.h
https://github.com/mamedev/mame/blob/master/src/mame/machine/mw8080bw.cpp
https://github.com/mamedev/mame/blob/master/src/mame/drivers/mw8080bw.cpp
https://github.com/mamedev/mame/blob/master/src/mame/video/mw8080bw.cpp
https://github.com/mamedev/mame/blob/master/src/mame/audio/mw8080bw.cpp
https://github.com/mamedev/mame/blob/master/src/mame/layout/invaders.lay

https://www.dorkbotpdx.org/blog/skinny/use_mames_debugger_to_reverse_engineer_and_extend_old_games

$ ./mame64 -window -nomaximize -debug roms/invaders.zip


# Intel 8085 CPU

https://www.wikiwand.com/en/Intel_8085
https://stackoverflow.com/questions/2165914/how-do-interrupts-work-on-the-intel-8080

https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.h
https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.cpp


# SDL library

https://github.com/veandco/go-sdl2
