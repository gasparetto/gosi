package processors

import (
	"log"
)

func fetch() uint8 {

	v := memRead(regs.PC)
	regs.PC++
	return v
}

func fetch16() uint16 {

	v := memRead16(regs.PC)
	regs.PC += 2
	return v
}

func setBC(v uint16) {
	regs.C = uint8(v)
	regs.B = uint8(v >> 8)
}

func getBC() uint16 {
	return uint16(regs.C) | uint16(regs.B)<<8
}

func setDE(v uint16) {
	regs.E = uint8(v)
	regs.D = uint8(v >> 8)
}

func getDE() uint16 {
	return uint16(regs.E) | uint16(regs.D)<<8
}

func setHL(v uint16) {
	regs.L = uint8(v)
	regs.H = uint8(v >> 8)
}

func getHL() uint16 {
	return uint16(regs.L) | uint16(regs.H)<<8
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
	case addr == 0x0002 || addr == 0x0004:
		log.Fatal("Attempted write to MB14241 shifter") //fixme
	case inBetween(addr, ramStart, ramEnd):
		ram[addr-ramStart] = v
	case inBetween(addr, romStart, romEnd):
		log.Fatal("Attempted write to readonly memory")
	default:
		log.Fatal("Attempted write to unknown memory")
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
	case inBetween(addr, romStart, romEnd):
		return rom[addr-romStart]
	default:
		log.Fatal("Attempted read to unknown memory")
		return 0
	}
}

func memRead16(addr uint16) uint16 {
	return uint16(memRead(addr)) | uint16(memRead(addr+1))<<8
}

func btoi(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

func inBetween(i, min, max uint16) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
