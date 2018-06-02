package i8085

import (
	"fmt"
	"log"
)

type I8085 struct{}

type Regs struct{
	PC               uint16 // 16-bit program counter register
	SP               uint16 // 16-bit stack pointer register
	B, C, D, E, H, L uint8  // 8-bit general purpose registers
	A                uint8  // ALU 8-bit accumulator
}

var regs = Regs{0x0000, 0x0000, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

var flags = struct {
	Z, S, P, CY, AC bool // ALU condition flags: Zero, Sign, Parity, Carry, Auxiliary Carry
}{false, false, false, false, false}

var interrupt bool // enable/disable the interrupt system

var ram []byte
var ramStart, ramEnd uint16

var rom []byte
var romStart, romEnd uint16

var ports_in = make(map[uint8]func() uint8)
var ports_out = make(map[uint8]func(v uint8))

func (I8085) GetRegs() Regs {
	return regs
}

func (I8085) GetRegDE() uint16 {
	return uint16(regs.D)<<8 | uint16(regs.E)
}

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

func (I8085) AttachPortIn(fp func() uint8, port uint8) {

	fmt.Printf("--- CPU i8085 :: ATTACH func %p to PORT IN 0x%02x ---\n", fp, port)

	ports_in[port] = fp
}

func (I8085) AttachPortOut(fp func(v uint8), port uint8) {

	fmt.Printf("--- CPU i8085 :: ATTACH func %p to PORT OUT 0x%02x ---\n", fp, port)

	ports_out[port] = fp
}

func (I8085) SetProgramCounter(addr uint16) {

	regs.PC = addr
}

func (I8085) GetProgramCounter() uint16 {

	return regs.PC
}

func (I8085) Step() int {

	//printOp()
	//printInternalState()

	op := fetch()
	fp := decode(op)
	return execute(fp)
}

func (I8085) Interrupt(op uint8) int {

	if interrupt {

		//debugPC = 0x9999

		interrupt = false

		fp := decode(op)
		cycles := execute(fp)
		//traceState()
		return cycles
	}
	return 0
}

func (I8085) DebugPrintNextOperation() {
	printOp()
}

func (I8085) DebugPrintInternalState() {
	printInternalState()
}

func setBC(v uint16) {
	regs.B = uint8(v >> 8)
	regs.C = uint8(v)
}

func getBC() uint16 {
	return uint16(regs.B)<<8 | uint16(regs.C)
}

func setDE(v uint16) {
	regs.D = uint8(v >> 8)
	regs.E = uint8(v)
}

func getDE() uint16 {
	return uint16(regs.D)<<8 | uint16(regs.E)
}

func setHL(v uint16) {
	regs.H = uint8(v >> 8)
	regs.L = uint8(v)
}

func getHL() uint16 {
	return uint16(regs.H)<<8 | uint16(regs.L)
}

func setFlags(v uint8) {
	flags.S = v&0x80 > 0
	flags.Z = v&0x40 > 0
	flags.AC = v&0x10 > 0
	flags.P = v&0x04 > 0
	flags.CY = v&0x01 > 0
}

func getFlags() uint8 {
	return uint8(btoi(flags.S)<<7 | btoi(flags.Z)<<6 | btoi(flags.AC)<<4 | btoi(flags.P)<<2 | 0x02 | btoi(flags.CY))
}

func memWrite(addr uint16, v uint8) {
	switch {
	case inBetween(addr, ramStart, ramEnd):
		ram[addr-ramStart] = v
	case inBetween(addr-0x2000, ramStart, ramEnd): // RAM mirror
		ram[addr-0x2000-ramStart] = v
	case inBetween(addr, romStart, romEnd):
		log.Fatalf("Attempted write to readonly memory 0x%04x\n", addr)
	default:
		log.Fatalf("Attempted write to unknown memory 0x%04x\n", addr)
	}
}

func memWrite16(addr uint16, v uint16) {
	memWrite(addr, uint8(v))
	memWrite(addr+1, uint8(v>>8))
}

func memRead(addr uint16) uint8 {
	switch {
	case inBetween(addr, ramStart, ramEnd):
		return ram[addr-ramStart]
	case inBetween(addr-0x2000, ramStart, ramEnd): // RAM mirror
		return ram[addr-0x2000-ramStart]
	case inBetween(addr, romStart, romEnd):
		return rom[addr-romStart]
	default:
		log.Fatalf("Attempted read to unknown memory 0x%04x\n", addr)
		return 0
	}
}

func memRead16(addr uint16) uint16 {
	return uint16(memRead(addr)) | uint16(memRead(addr+1))<<8
}
