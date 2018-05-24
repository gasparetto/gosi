package processors

import (
	"fmt"
)

type I8085 struct{}

var regs = struct {
	PC               uint16 // 16-bit program counter register
	SP               uint16 // 16-bit stack pointer register
	B, C, D, E, H, L uint8  // 8-bit general purpose registers
	A                uint8  // ALU 8-bit accumulator
}{0x0000, 0x0000, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

var flags = struct {
	Z, S, P, CY, AC bool // ALU condition flags: Zero, Sign, Parity, Carry, Auxiliary Carry
}{false, false, false, false, false}

var interrupt bool // enable/disable the interrupt system

var ram []byte
var ramStart, ramEnd uint16

var rom []byte
var romStart, romEnd uint16

func (I8085) AttachRam(buf []byte, offset uint16) {

	fmt.Printf("--- CPU i8085 :: ATTACH RAM (%d bytes @ 0x%04x) ---\n", len(buf), offset)

	ram = buf
	ramStart = offset
	ramEnd = offset + uint16(len(buf)) - 1
}

func (I8085) AttachRom(buf []byte, offset uint16) {

	fmt.Printf("--- CPU i8085 :: ATTACH ROM (%d bytes @ 0x%04x) ---\n", len(buf), offset)

	rom = buf
	romStart = offset
	romEnd = offset + uint16(len(buf)) - 1
}

func (I8085) Step() int {

	debugPC = regs.PC

	op := fetch()
	fp := decode(op)
	return execute(fp)
}

func (I8085) Interrupt(op uint8) int {

	if interrupt {

		debugPC = 0x9999

		interrupt = false

		fp := decode(op)
		return execute(fp)
	}
	return 0
}

func (I8085) DebugPrint() {
	trace()
}

func (I8085) DebugBreakpoint() bool {
	//return debugPC == 0x9999
	return false
}
