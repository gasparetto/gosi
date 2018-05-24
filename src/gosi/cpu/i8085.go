/*
	https://www.wikiwand.com/en/Intel_8085
	https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.cpp
*/

package cpu

import (
	"fmt"
	"log"
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

func (I8085) Step() {

	op := rom[regs.PC]

	switch op {

	case 0x00:
		opNOP()
	case 0x01:
		opLXI_BC_data16()
	case 0x03:
		opINX_BC()
	case 0x04:
		opINR_B()
	case 0x05:
		opDCR_B()
	case 0x06:
		opMVI_B_data()
	case 0x0a:
		opLDAX_BC()
	case 0x0b:
		opDCX_BC()
	case 0x0c:
		opINR_C()
	case 0x0d:
		opDCR_C()
	case 0x0e:
		opMVI_C_data()
	case 0x11:
		opLXI_DE_data16()
	case 0x13:
		opINX_DE()
	case 0x14:
		opINR_D()
	case 0x15:
		opDCR_D()
	case 0x16:
		opMVI_D_data()
	case 0x1a:
		opLDAX_DE()
	case 0x1b:
		opDCX_DE()
	case 0x1c:
		opINR_D()
	case 0x1d:
		opDCR_D()
	case 0x1e:
		opMVI_E_data()
	case 0x21:
		opLXI_HL_data16()
	case 0x23:
		opINX_HL()
	case 0x24:
		opINR_H()
	case 0x25:
		opDCR_H()
	case 0x26:
		opMVI_H_data()
	case 0x2b:
		opDCX_HL()
	case 0x2c:
		opINR_L()
	case 0x2d:
		opDCR_L()
	case 0x2e:
		opMVI_L_data()
	case 0x31:
		opLXI_SP_data16()
	case 0x33:
		opINX_SP()
	case 0x36:
		opMVI_M_data()
	case 0x3b:
		opDCX_SP()
	case 0x40:
		opMOV_B_B()
	case 0x41:
		opMOV_C_B()
	case 0x42:
		opMOV_D_B()
	case 0x43:
		opMOV_E_B()
	case 0x44:
		opMOV_H_B()
	case 0x45:
		opMOV_L_B()
	case 0x46:
		opMOV_B_M()
	case 0x48:
		opMOV_B_C()
	case 0x49:
		opMOV_C_C()
	case 0x4a:
		opMOV_D_C()
	case 0x4b:
		opMOV_E_C()
	case 0x4c:
		opMOV_H_C()
	case 0x4d:
		opMOV_L_C()
	case 0x4e:
		opMOV_C_M()
	case 0x50:
		opMOV_B_D()
	case 0x51:
		opMOV_C_D()
	case 0x52:
		opMOV_D_D()
	case 0x53:
		opMOV_E_D()
	case 0x54:
		opMOV_H_D()
	case 0x55:
		opMOV_L_D()
	case 0x56:
		opMOV_D_M()
	case 0x58:
		opMOV_B_E()
	case 0x59:
		opMOV_C_E()
	case 0x5a:
		opMOV_D_E()
	case 0x5b:
		opMOV_E_E()
	case 0x5c:
		opMOV_H_E()
	case 0x5d:
		opMOV_L_E()
	case 0x5e:
		opMOV_E_M()
	case 0x60:
		opMOV_B_H()
	case 0x61:
		opMOV_C_H()
	case 0x62:
		opMOV_D_H()
	case 0x63:
		opMOV_E_H()
	case 0x64:
		opMOV_H_H()
	case 0x65:
		opMOV_L_H()
	case 0x66:
		opMOV_H_M()
	case 0x68:
		opMOV_B_L()
	case 0x69:
		opMOV_C_L()
	case 0x6a:
		opMOV_D_L()
	case 0x6b:
		opMOV_E_L()
	case 0x6c:
		opMOV_H_L()
	case 0x6d:
		opMOV_L_L()
	case 0x6e:
		opMOV_L_M()
	case 0x70:
		opMOV_M_B()
	case 0x71:
		opMOV_M_C()
	case 0x72:
		opMOV_M_D()
	case 0x73:
		opMOV_M_E()
	case 0x74:
		opMOV_M_H()
	case 0x75:
		opMOV_M_L()
	case 0x77:
		opMOV_M_A()
	case 0x7e:
		opMOV_A_M()
	case 0xc3:
		opJMP()
	case 0xcd:
		opCALL()
	default:
		log.Fatalf("ERROR: Unknown opcode 0x%02x\n", op)
	}

	regs.PC++
}

func bcSet(addr uint16) {
	regs.C = uint8(addr)
	regs.B = uint8(addr >> 8)
}

func bcGet() uint16 {
	return uint16(regs.C) | uint16(regs.B)<<8
}

func deSet(addr uint16) {
	regs.E = uint8(addr)
	regs.D = uint8(addr >> 8)
}

func deGet() uint16 {
	return uint16(regs.E) | uint16(regs.D)<<8
}

func hlSet(addr uint16) {
	regs.L = uint8(addr)
	regs.H = uint8(addr >> 8)
}

func hlGet() uint16 {
	return uint16(regs.L) | uint16(regs.H)<<8
}

func ramSet(addr uint16, v uint8) {
	switch {
	case inBetween(addr, ramStart, ramEnd):
		ram[addr-ramStart] = v
	case inBetween(addr, romStart, romEnd):
		log.Fatal("Attempted write to readonly memory")
	default:
		log.Fatal("Attempted write to unknown memory")
	}
}

func ramGet(addr uint16) uint8 {
	switch {
	case inBetween(addr, ramStart, ramEnd):
		return ram[addr-ramStart]
	case inBetween(addr, romStart, romEnd):
		return rom[addr-romStart]
	default:
		log.Fatal("Attempted read to unknown memory")
		return 0x0000
	}
}

func inBetween(i, min, max uint16) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func add(a uint8, b uint8) uint8 {
	res16 := uint16(a) + uint16(b)
	res8 := uint8(res16)
	flags.Z = res8 == 0
	flags.S = (res8 & 0x80) != 0
	flags.P = (res8 & 0x01) == 0
	flags.CY = (res16 & 0xff00) != 0
	//todo flags.AC
	return res8
}

func sub(a uint8, b uint8) uint8 {
	res16 := uint16(a) - uint16(b)
	res8 := uint8(res16)
	flags.Z = res8 == 0
	flags.S = (res8 & 0x80) != 0
	flags.P = (res8 & 0x01) == 0
	flags.CY = (res16 & 0xff00) != 0
	//todo flags.AC
	return res8
}
