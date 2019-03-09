# gosi

GOlang Space Invaders emulator


#### GO lang reference

https://github.com/a8m/go-lang-cheat-sheet


#### Hardware reference

http://computerarcheology.com/Arcade/SpaceInvaders/Hardware.html

Midway 8080 black/white arcade machine

MAIN CPU memory address space

| Address (15-bits)   | Dir | Data       | Description                               |
| ------------------- | --- | ---------- | ----------------------------------------- |
| `x0xxxxxxxxxxxxx`   | R   | `xxxxxxxx` | Program ROM (various amounts populated)   |
| `-1xxxxxxxxxxxxx`   | R/W | `xxxxxxxx` | Video RAM (256x256x1 bit display)         |
|                     |     |            | Portion in VBLANK region used as work RAM |


#### Intel 8085 CPU

https://www.wikiwand.com/en/Intel_8085  
https://stackoverflow.com/questions/2165914/how-do-interrupts-work-on-the-intel-8080  


#### SDL library

https://www.libsdl.org  
https://github.com/veandco/go-sdl2  


#### MAME

https://github.com/mamedev/mame/blob/master/src/mame/includes/mw8080bw.h  
https://github.com/mamedev/mame/blob/master/src/mame/machine/mw8080bw.cpp  
https://github.com/mamedev/mame/blob/master/src/mame/drivers/mw8080bw.cpp  
https://github.com/mamedev/mame/blob/master/src/mame/video/mw8080bw.cpp  
https://github.com/mamedev/mame/blob/master/src/mame/audio/mw8080bw.cpp  
https://github.com/mamedev/mame/blob/master/src/mame/layout/invaders.lay  
https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.h  
https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.cpp  

https://www.dorkbotpdx.org/blog/skinny/use_mames_debugger_to_reverse_engineer_and_extend_old_games

`$ ./mame64 -window -nomaximize -debug roms/invaders.zip`


Game driver (src/mame/drivers/mw8080bw.cpp:3245)

```c
/* PCB #        year  rom         parent    machine   inp       init              monitor,company,fullname,flags */
/* 739 */ GAMEL( 1978, invaders,   0,        invaders, invaders, mw8080bw_state, empty_init, ROT270, "Taito / Midway", "Space Invaders / Space Invaders M", MACHINE_SUPPORTS_SAVE, layout_invaders )
```


Memory address space (src/mame/drivers/mw8080bw.cpp:3168)

```c
ROM_START( invaders )
    ROM_REGION( 0x10000, "maincpu", 0 )
    ROM_LOAD( "invaders.h", 0x0000, 0x0800, CRC(734f5ad8) SHA1(ff6200af4c9110d8181249cbcef1a8a40fa40b7f) )
    ROM_LOAD( "invaders.g", 0x0800, 0x0800, CRC(6bfaca4a) SHA1(16f48649b531bdef8c2d1446c429b5f414524350) )
    ROM_LOAD( "invaders.f", 0x1000, 0x0800, CRC(0ccead96) SHA1(537aef03468f63c5b9e11dd61e253f7ae17d9743) )
    ROM_LOAD( "invaders.e", 0x1800, 0x0800, CRC(14e538b0) SHA1(1d6ca0c99f9df71e2990b610deb9d7da0125e2d8) )
ROM_END
```


Interrupt timing (src/mame/includes/mw8080bw.h:19)

```c
#define MW8080BW_MASTER_CLOCK             (19968000.0)
#define MW8080BW_CPU_CLOCK                (MW8080BW_MASTER_CLOCK / 10)
#define MW8080BW_PIXEL_CLOCK              (MW8080BW_MASTER_CLOCK / 4)
#define MW8080BW_HTOTAL                   (0x140)
#define MW8080BW_HBEND                    (0x000)
#define MW8080BW_HBSTART                  (0x100)
#define MW8080BW_VTOTAL                   (0x106)
#define MW8080BW_VBEND                    (0x000)
#define MW8080BW_VBSTART                  (0x0e0)
#define MW8080BW_VCOUNTER_START_NO_VBLANK (0x020)
#define MW8080BW_VCOUNTER_START_VBLANK    (0x0da)
#define MW8080BW_INT_TRIGGER_COUNT_1      (0x080)
#define MW8080BW_INT_TRIGGER_VBLANK_1     (0)
#define MW8080BW_INT_TRIGGER_COUNT_2      MW8080BW_VCOUNTER_START_VBLANK
#define MW8080BW_INT_TRIGGER_VBLANK_2     (1)
#define MW8080BW_60HZ                     (MW8080BW_PIXEL_CLOCK / MW8080BW_HTOTAL / MW8080BW_VTOTAL)
```

10 cpu cycles / 4 pixel cycles = 2.5  
128 \* 320 = 40960 -> 40960 / 2.5 = 16384  
218 \* 320 = 69760 -> 69760 / 2.5 = 27904  
262 \* 320 = 83840 -> 83840 / 2.5 = 33536  


