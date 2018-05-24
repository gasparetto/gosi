/*
	https://www.wikiwand.com/en/Intel_8085
	https://stackoverflow.com/questions/2165914/how-do-interrupts-work-on-the-intel-8080
	https://github.com/mamedev/mame/blob/master/src/devices/cpu/i8085/i8085.cpp
*/

package processors

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

func (I8085) Step() int {

	op := rom[regs.PC]
	cycles := 0

	switch op {

	case 0x00:
		cycles += opNOP()
	case 0x01:
		cycles += opLXI_BC()
	case 0x02:
		cycles += opSTAX_BC()
	case 0x03:
		cycles += opINX_BC()
	case 0x04:
		cycles += opINR_B()
	case 0x05:
		cycles += opDCR_B()
	case 0x06:
		cycles += opMVI_B()
	case 0x09:
		cycles += opDAD_BC()
	case 0x0a:
		cycles += opLDAX_BC()
	case 0x0b:
		cycles += opDCX_BC()
	case 0x0c:
		cycles += opINR_C()
	case 0x0d:
		cycles += opDCR_C()
	case 0x0e:
		cycles += opMVI_C()
	case 0x0f:
		cycles += opRRC()
	case 0x11:
		cycles += opLXI_DE()
	case 0x12:
		cycles += opSTAX_DE()
	case 0x13:
		cycles += opINX_DE()
	case 0x14:
		cycles += opINR_D()
	case 0x15:
		cycles += opDCR_D()
	case 0x16:
		cycles += opMVI_D()
	case 0x19:
		cycles += opDAD_DE()
	case 0x1a:
		cycles += opLDAX_DE()
	case 0x1b:
		cycles += opDCX_DE()
	case 0x1c:
		cycles += opINR_E()
	case 0x1d:
		cycles += opDCR_E()
	case 0x1e:
		cycles += opMVI_E()
	case 0x21:
		cycles += opLXI_HL()
	case 0x23:
		cycles += opINX_HL()
	case 0x24:
		cycles += opINR_H()
	case 0x25:
		cycles += opDCR_H()
	case 0x26:
		cycles += opMVI_H()
	case 0x29:
		cycles += opDAD_HL()
	case 0x2b:
		cycles += opDCX_HL()
	case 0x2c:
		cycles += opINR_L()
	case 0x2d:
		cycles += opDCR_L()
	case 0x2e:
		cycles += opMVI_L()
	case 0x31:
		cycles += opLXI_SP()
	case 0x32:
		cycles += opSTA()
	case 0x33:
		cycles += opINX_SP()
	case 0x36:
		cycles += opMVI_M()
	case 0x39:
		cycles += opDAD_SP()
	case 0x3a:
		cycles += opLDA()
	case 0x3b:
		cycles += opDCX_SP()
	case 0x3c:
		cycles += opINR_A()
	case 0x3d:
		cycles += opDCR_A()
	case 0x3e:
		cycles += opMVI_A()
	case 0x40:
		cycles += opMOV_B_B()
	case 0x41:
		cycles += opMOV_C_B()
	case 0x42:
		cycles += opMOV_D_B()
	case 0x43:
		cycles += opMOV_E_B()
	case 0x44:
		cycles += opMOV_H_B()
	case 0x45:
		cycles += opMOV_L_B()
	case 0x46:
		cycles += opMOV_B_M()
	case 0x47:
		cycles += opMOV_B_A()
	case 0x48:
		cycles += opMOV_B_C()
	case 0x49:
		cycles += opMOV_C_C()
	case 0x4a:
		cycles += opMOV_D_C()
	case 0x4b:
		cycles += opMOV_E_C()
	case 0x4c:
		cycles += opMOV_H_C()
	case 0x4d:
		cycles += opMOV_L_C()
	case 0x4e:
		cycles += opMOV_C_M()
	case 0x4f:
		cycles += opMOV_C_A()
	case 0x50:
		cycles += opMOV_B_D()
	case 0x51:
		cycles += opMOV_C_D()
	case 0x52:
		cycles += opMOV_D_D()
	case 0x53:
		cycles += opMOV_E_D()
	case 0x54:
		cycles += opMOV_H_D()
	case 0x55:
		cycles += opMOV_L_D()
	case 0x56:
		cycles += opMOV_D_M()
	case 0x57:
		cycles += opMOV_D_A()
	case 0x58:
		cycles += opMOV_B_E()
	case 0x59:
		cycles += opMOV_C_E()
	case 0x5a:
		cycles += opMOV_D_E()
	case 0x5b:
		cycles += opMOV_E_E()
	case 0x5c:
		cycles += opMOV_H_E()
	case 0x5d:
		cycles += opMOV_L_E()
	case 0x5e:
		cycles += opMOV_E_M()
	case 0x5f:
		cycles += opMOV_E_A()
	case 0x60:
		cycles += opMOV_B_H()
	case 0x61:
		cycles += opMOV_C_H()
	case 0x62:
		cycles += opMOV_D_H()
	case 0x63:
		cycles += opMOV_E_H()
	case 0x64:
		cycles += opMOV_H_H()
	case 0x65:
		cycles += opMOV_L_H()
	case 0x66:
		cycles += opMOV_H_M()
	case 0x67:
		cycles += opMOV_H_A()
	case 0x68:
		cycles += opMOV_B_L()
	case 0x69:
		cycles += opMOV_C_L()
	case 0x6a:
		cycles += opMOV_D_L()
	case 0x6b:
		cycles += opMOV_E_L()
	case 0x6c:
		cycles += opMOV_H_L()
	case 0x6d:
		cycles += opMOV_L_L()
	case 0x6e:
		cycles += opMOV_L_M()
	case 0x6f:
		cycles += opMOV_L_A()
	case 0x70:
		cycles += opMOV_M_B()
	case 0x71:
		cycles += opMOV_M_C()
	case 0x72:
		cycles += opMOV_M_D()
	case 0x73:
		cycles += opMOV_M_E()
	case 0x74:
		cycles += opMOV_M_H()
	case 0x75:
		cycles += opMOV_M_L()
	case 0x77:
		cycles += opMOV_M_A()
	case 0x78:
		cycles += opMOV_A_B()
	case 0x79:
		cycles += opMOV_A_C()
	case 0x7a:
		cycles += opMOV_A_D()
	case 0x7b:
		cycles += opMOV_A_E()
	case 0x7c:
		cycles += opMOV_A_H()
	case 0x7d:
		cycles += opMOV_A_L()
	case 0x7e:
		cycles += opMOV_A_M()
	case 0x7f:
		cycles += opMOV_A_A()
	case 0xa0:
		cycles += opANA_B()
	case 0xa1:
		cycles += opANA_C()
	case 0xa2:
		cycles += opANA_D()
	case 0xa3:
		cycles += opANA_E()
	case 0xa4:
		cycles += opANA_H()
	case 0xa5:
		cycles += opANA_L()
	case 0xa6:
		cycles += opANA_M()
	case 0xa7:
		cycles += opANA_A()
	case 0xa8:
		cycles += opXRA_B()
	case 0xa9:
		cycles += opXRA_C()
	case 0xaa:
		cycles += opXRA_D()
	case 0xab:
		cycles += opXRA_E()
	case 0xac:
		cycles += opXRA_H()
	case 0xad:
		cycles += opXRA_L()
	case 0xae:
		cycles += opXRA_M()
	case 0xaf:
		cycles += opXRA_A()
	case 0xc1:
		cycles += opPOP_BC()
	case 0xc2:
		cycles += opJNZ()
	case 0xc3:
		cycles += opJMP()
	case 0xc5:
		cycles += opPUSH_BC()
	case 0xc6:
		cycles += opADI()
	case 0xc9:
		cycles += opRET()
	case 0xca:
		cycles += opJZ()
	case 0xcd:
		cycles += opCALL()
	case 0xd1:
		cycles += opPOP_DE()
	case 0xd2:
		cycles += opJNC()
	case 0xd3:
		cycles += opOUT()
	case 0xd5:
		cycles += opPUSH_DE()
	case 0xda:
		cycles += opJC()
	case 0xe1:
		cycles += opPOP_HL()
	case 0xe2:
		cycles += opJPO()
	case 0xe5:
		cycles += opPUSH_HL()
	case 0xe6:
		cycles += opANI()
	case 0xea:
		cycles += opJPE()
	case 0xeb:
		cycles += opXCHG()
	case 0xf1:
		cycles += opPOP_PSW()
	case 0xf2:
		cycles += opJP()
	case 0xf5:
		cycles += opPUSH_PSW()
	case 0xfa:
		cycles += opJM()
	case 0xfb:
		cycles += opEI()
	case 0xfe:
		cycles += opCPI()
	default:
		trace()
		//breakpoint = true
		log.Fatalf("ERROR: Unknown opcode 0x%02x\n", op)
	}

	regs.PC++

	return cycles
}

func (I8085) DasmTrace() {
	trace()
}

func (I8085) Breakpoint() bool {

	//return regs.PC == 0x0a98
	return false
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

func btoi(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func memWrite(addr uint16, v uint8) {
	switch {
	case addr == 0x0002 || addr == 0x0004:
		log.Fatal("Attempted write to MB14241 shifter")
		//case inBetween(addr, 0x2400, 0x4000):
		//	log.Fatal("Attempted write to video memory")
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
	case addr == 0x0003:
		log.Fatal("Attempted read from MB14241 shifter")
		return 0
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
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	return res8
}

func add_CY(a uint8, b uint8) uint8 {
	res16 := uint16(a) + uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	//todo flags.CY
	return res8
}

func sub(a uint8, b uint8) uint8 {
	res16 := uint16(a) - uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	return res8
}

func sub_CY(a uint8, b uint8) uint8 {
	res16 := uint16(a) - uint16(b)
	res8 := uint8(res16)
	flags_Z_S_P(res8)
	flags.AC = (res16 & 0xff00) != 0
	//todo flags.CY
	return res8
}

func flags_Z_S_P(v uint8) {
	flags.Z = v == 0
	flags.S = (v & 0x80) != 0
	flags.P = evenParity(v)
}

func evenParity(v uint8) bool {
	b0 := v & 0x01
	b1 := v & 0x02
	b2 := v & 0x04
	b3 := v & 0x08
	b4 := v & 0x10
	b5 := v & 0x20
	b6 := v & 0x40
	b7 := v & 0x80
	r := b7 + b6 + b5 + b4 + b3 + b2 + b1 + b0
	return (r & 0x01) == 0
}
