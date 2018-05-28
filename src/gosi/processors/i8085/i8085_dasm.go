package i8085

import (
	"fmt"
	"log"
)

func printOp() {

	op := memRead(regs.PC)

	switch op {
	case 0x00:
		dasmOp("NOP")
	case 0x01:
		dasmOpRegVal1Val2("LXI", "BC", readByte3(), readByte2())
	case 0x02:
		dasmOpReg("STAX", "BC")
	case 0x03:
		dasmOpReg("INX", "BC")
	case 0x04:
		dasmOpReg("INR", "B")
	case 0x05:
		dasmOpReg("DCR", "B")
	case 0x06:
		dasmOpRegVal("MVI", "B", readByte2())
	case 0x07:
		dasmOp("RLC")
	case 0x08:
		log.Fatal("undocumented NOP")
	case 0x09:
		dasmOpReg("DAD", "BC")
	case 0x0a:
		dasmOpReg("LDAX", "BC")
	case 0x0b:
		dasmOpReg("DCX", "BC")
	case 0x0c:
		dasmOpReg("INR", "C")
	case 0x0d:
		dasmOpReg("DCR", "C")
	case 0x0e:
		dasmOpRegVal("MVI", "C", readByte2())
	case 0x0f:
		dasmOp("RRC")
	case 0x10:
		log.Fatal("undocumented NOP")
	case 0x11:
		dasmOpRegVal1Val2("LXI", "DE", readByte3(), readByte2())
	case 0x12:
		dasmOpReg("STAX", "DE")
	case 0x13:
		dasmOpReg("INX", "DE")
	case 0x14:
		dasmOpReg("INR", "D")
	case 0x15:
		dasmOpReg("DCR", "D")
	case 0x16:
		dasmOpRegVal("MVI", "D", readByte2())
	case 0x17:
		dasmOp("RAL")
	case 0x18:
		log.Fatal("undocumented NOP")
	case 0x19:
		dasmOpReg("DAD", "DE")
	case 0x1a:
		dasmOpReg("LDAX", "DE")
	case 0x1b:
		dasmOpReg("DCX", "DE")
	case 0x1c:
		dasmOpReg("INR", "E")
	case 0x1d:
		dasmOpReg("DCR", "E")
	case 0x1e:
		dasmOpRegVal("MVI", "E", readByte2())
	case 0x1f:
		dasmOp("RAR")
	case 0x20:
		log.Fatal("undocumented NOP")
	case 0x21:
		dasmOpRegVal1Val2("LXI", "HL", readByte3(), readByte2())
	case 0x22:
		dasmOpAddr("SHLD", readAddr())
	case 0x23:
		dasmOpReg("INX", "HL")
	case 0x24:
		dasmOpReg("INR", "H")
	case 0x25:
		dasmOpReg("DCR", "H")
	case 0x26:
		dasmOpRegVal("MVI", "H", readByte2())
	//case 0x27: fixme DAA
	case 0x28:
		log.Fatal("undocumented NOP")
	case 0x29:
		dasmOpReg("DAD", "HL")
	case 0x2a:
		dasmOpAddr("LHLD", readAddr())
	case 0x2b:
		dasmOpReg("DCX", "HL")
	case 0x2c:
		dasmOpReg("INR", "L")
	case 0x2d:
		dasmOpReg("DCR", "L")
	case 0x2e:
		dasmOpRegVal("MVI", "L", readByte2())
	case 0x2f:
		dasmOp("CMA")
	case 0x30:
		log.Fatal("undocumented NOP")
	case 0x31:
		dasmOpRegAddr("LXI", "SP", readAddr())
	case 0x32:
		dasmOpAddr("STA", readAddr())
	case 0x33:
		dasmOpReg("INX", "SP")
	case 0x34:
		dasmOpReg("INR", "M")
	case 0x35:
		dasmOpReg("DCR", "M")
	case 0x36:
		dasmOpRegAddr("MVI", "M", readAddr())
	case 0x37:
		dasmOp("STC")
	case 0x38:
		log.Fatal("undocumented NOP")
	case 0x39:
		dasmOpReg("DAD", "SP")
	case 0x3a:
		dasmOpAddr("LDA", readAddr())
	case 0x3b:
		dasmOpReg("DCX", "SP")
	case 0x3c:
		dasmOpReg("INR", "A")
	case 0x3d:
		dasmOpReg("DCR", "A")
	case 0x3e:
		dasmOpRegVal("MVI", "A", readByte2())
	//case 0x3f: fixme CMC
	case 0x40:
		dasmOpReg1Reg2("MOV", "B", "B")
	case 0x41:
		dasmOpReg1Reg2("MOV", "C", "B")
	case 0x42:
		dasmOpReg1Reg2("MOV", "D", "B")
	case 0x43:
		dasmOpReg1Reg2("MOV", "E", "B")
	case 0x44:
		dasmOpReg1Reg2("MOV", "H", "B")
	case 0x45:
		dasmOpReg1Reg2("MOV", "L", "B")
	case 0x46:
		dasmOpRegAddr("MOV", "B", readAddr())
	case 0x47:
		dasmOpReg1Reg2("MOV", "B", "A")
	case 0x48:
		dasmOpReg1Reg2("MOV", "B", "C")
	case 0x49:
		dasmOpReg1Reg2("MOV", "C", "C")
	case 0x4a:
		dasmOpReg1Reg2("MOV", "D", "C")
	case 0x4b:
		dasmOpReg1Reg2("MOV", "E", "C")
	case 0x4c:
		dasmOpReg1Reg2("MOV", "H", "C")
	case 0x4d:
		dasmOpReg1Reg2("MOV", "L", "C")
	case 0x4e:
		dasmOpRegAddr("MOV", "C", readAddr())
	case 0x4f:
		dasmOpReg1Reg2("MOV", "C", "A")
	case 0x50:
		dasmOpReg1Reg2("MOV", "B", "D")
	case 0x51:
		dasmOpReg1Reg2("MOV", "C", "D")
	case 0x52:
		dasmOpReg1Reg2("MOV", "D", "D")
	case 0x53:
		dasmOpReg1Reg2("MOV", "E", "D")
	case 0x54:
		dasmOpReg1Reg2("MOV", "H", "D")
	case 0x55:
		dasmOpReg1Reg2("MOV", "L", "D")
	case 0x56:
		dasmOpRegAddr("MOV", "D", readAddr())
	case 0x57:
		dasmOpReg1Reg2("MOV", "D", "A")
	case 0x58:
		dasmOpReg1Reg2("MOV", "B", "E")
	case 0x59:
		dasmOpReg1Reg2("MOV", "C", "E")
	case 0x5a:
		dasmOpReg1Reg2("MOV", "D", "E")
	case 0x5b:
		dasmOpReg1Reg2("MOV", "E", "E")
	case 0x5c:
		dasmOpReg1Reg2("MOV", "H", "E")
	case 0x5d:
		dasmOpReg1Reg2("MOV", "L", "E")
	case 0x5e:
		dasmOpRegAddr("MOV", "E", readAddr())
	case 0x5f:
		dasmOpReg1Reg2("MOV", "E", "A")
	case 0x60:
		dasmOpReg1Reg2("MOV", "B", "H")
	case 0x61:
		dasmOpReg1Reg2("MOV", "C", "H")
	case 0x62:
		dasmOpReg1Reg2("MOV", "D", "H")
	case 0x63:
		dasmOpReg1Reg2("MOV", "E", "H")
	case 0x64:
		dasmOpReg1Reg2("MOV", "H", "H")
	case 0x65:
		dasmOpReg1Reg2("MOV", "L", "H")
	case 0x66:
		dasmOpRegAddr("MOV", "H", readAddr())
	case 0x67:
		dasmOpReg1Reg2("MOV", "H", "A")
	case 0x68:
		dasmOpReg1Reg2("MOV", "B", "L")
	case 0x69:
		dasmOpReg1Reg2("MOV", "C", "L")
	case 0x6a:
		dasmOpReg1Reg2("MOV", "D", "L")
	case 0x6b:
		dasmOpReg1Reg2("MOV", "E", "L")
	case 0x6c:
		dasmOpReg1Reg2("MOV", "H", "L")
	case 0x6d:
		dasmOpReg1Reg2("MOV", "L", "L")
	case 0x6e:
		dasmOpRegAddr("MOV", "L", readAddr())
	case 0x6f:
		dasmOpReg1Reg2("MOV", "L", "A")
	case 0x70:
		dasmOpAddrReg("MOV", readAddr(), "B")
	case 0x71:
		dasmOpAddrReg("MOV", readAddr(), "C")
	case 0x72:
		dasmOpAddrReg("MOV", readAddr(), "D")
	case 0x73:
		dasmOpAddrReg("MOV", readAddr(), "E")
	case 0x74:
		dasmOpAddrReg("MOV", readAddr(), "H")
	case 0x75:
		dasmOpAddrReg("MOV", readAddr(), "L")
	case 0x76:
		dasmOp("HLT")
	case 0x77:
		dasmOpAddrReg("MOV", readAddr(), "A")
	case 0x78:
		dasmOpReg1Reg2("MOV", "A", "B")
	case 0x79:
		dasmOpReg1Reg2("MOV", "A", "C")
	case 0x7a:
		dasmOpReg1Reg2("MOV", "A", "D")
	case 0x7b:
		dasmOpReg1Reg2("MOV", "A", "E")
	case 0x7c:
		dasmOpReg1Reg2("MOV", "A", "H")
	case 0x7d:
		dasmOpReg1Reg2("MOV", "A", "L")
	case 0x7e:
		dasmOpRegAddr("MOV", "A", readAddr())
	case 0x7f:
		dasmOpReg1Reg2("MOV", "A", "A")
	case 0x80:
		dasmOpReg("ADD", "B")
	case 0x81:
		dasmOpReg("ADD", "C")
	case 0x82:
		dasmOpReg("ADD", "D")
	case 0x83:
		dasmOpReg("ADD", "E")
	case 0x84:
		dasmOpReg("ADD", "H")
	case 0x85:
		dasmOpReg("ADD", "L")
	case 0x86:
		dasmOpAddr("ADD", readAddr())
	case 0x87:
		dasmOpReg("ADD", "A")
	case 0x88:
		dasmOpReg("ADC", "B")
	case 0x89:
		dasmOpReg("ADC", "C")
	case 0x8a:
		dasmOpReg("ADC", "D")
	case 0x8b:
		dasmOpReg("ADC", "E")
	case 0x8c:
		dasmOpReg("ADC", "H")
	case 0x8d:
		dasmOpReg("ADC", "L")
	case 0x8e:
		dasmOpAddr("ADC", readAddr())
	case 0x8f:
		dasmOpReg("ADC", "A")
	case 0x90:
		dasmOpReg("SUB", "B")
	case 0x91:
		dasmOpReg("SUB", "C")
	case 0x92:
		dasmOpReg("SUB", "D")
	case 0x93:
		dasmOpReg("SUB", "E")
	case 0x94:
		dasmOpReg("SUB", "H")
	case 0x95:
		dasmOpReg("SUB", "L")
	case 0x96:
		dasmOpAddr("SUB", readAddr())
	case 0x97:
		dasmOpReg("SUB", "A")
	case 0x98:
		dasmOpReg("SBB", "B")
	case 0x99:
		dasmOpReg("SBB", "C")
	case 0x9a:
		dasmOpReg("SBB", "D")
	case 0x9b:
		dasmOpReg("SBB", "E")
	case 0x9c:
		dasmOpReg("SBB", "H")
	case 0x9d:
		dasmOpReg("SBB", "L")
	case 0x9e:
		dasmOpAddr("SBB", readAddr())
	case 0x9f:
		dasmOpReg("SBB", "A")
	case 0xa0:
		dasmOpReg("ANA", "B")
	case 0xa1:
		dasmOpReg("ANA", "C")
	case 0xa2:
		dasmOpReg("ANA", "D")
	case 0xa3:
		dasmOpReg("ANA", "E")
	case 0xa4:
		dasmOpReg("ANA", "H")
	case 0xa5:
		dasmOpReg("ANA", "L")
	case 0xa6:
		dasmOpAddr("ANA", readAddr())
	case 0xa7:
		dasmOpReg("ANA", "A")
	case 0xa8:
		dasmOpReg("XRA", "B")
	case 0xa9:
		dasmOpReg("XRA", "C")
	case 0xaa:
		dasmOpReg("XRA", "D")
	case 0xab:
		dasmOpReg("XRA", "E")
	case 0xac:
		dasmOpReg("XRA", "H")
	case 0xad:
		dasmOpReg("XRA", "L")
	case 0xae:
		dasmOpAddr("XRA", readAddr())
	case 0xaf:
		dasmOpReg("XRA", "A")
	case 0xb0:
		dasmOpReg("ORA", "B")
	case 0xb1:
		dasmOpReg("ORA", "C")
	case 0xb2:
		dasmOpReg("ORA", "D")
	case 0xb3:
		dasmOpReg("ORA", "E")
	case 0xb4:
		dasmOpReg("ORA", "H")
	case 0xb5:
		dasmOpReg("ORA", "L")
	case 0xb6:
		dasmOpAddr("ORA", readAddr())
	case 0xb7:
		dasmOpReg("ORA", "A")
	case 0xb8:
		dasmOpReg("CMP", "B")
	case 0xb9:
		dasmOpReg("CMP", "C")
	case 0xba:
		dasmOpReg("CMP", "D")
	case 0xbb:
		dasmOpReg("CMP", "E")
	case 0xbc:
		dasmOpReg("CMP", "H")
	case 0xbd:
		dasmOpReg("CMP", "L")
	case 0xbe:
		dasmOpAddr("CMP", readAddr())
	case 0xbf:
		dasmOpReg("CMP", "A")
	case 0xc0:
		dasmOp("RNZ")
	case 0xc1:
		dasmOpReg("POP", "BC")
	case 0xc2:
		dasmOpAddr("JNZ", readAddr())
	case 0xc3:
		dasmOpAddr("JMP", readAddr())
	case 0xc4:
		dasmOpAddr("CNZ", readAddr())
	case 0xc5:
		dasmOpReg("PUSH", "BC")
	case 0xc6:
		dasmOpVal("ADI", readByte2())
	case 0xc7:
		dasmOpReg("RST", "0")
	case 0xc8:
		dasmOp("RZ")
	case 0xc9:
		dasmOp("RET")
	case 0xca:
		dasmOpAddr("JZ", readAddr())
	case 0xcb:
		log.Fatal("undocumented JMP nnnn")
	case 0xcc:
		dasmOpAddr("CZ", readAddr())
	case 0xcd:
		dasmOpAddr("CALL", readAddr())
	case 0xce:
		dasmOpVal("ACI", readByte2())
	case 0xcf:
		dasmOpReg("RST", "1")
	case 0xd0:
		dasmOp("RNC")
	case 0xd1:
		dasmOpReg("POP", "DE")
	case 0xd2:
		dasmOpAddr("JNC", readAddr())
	case 0xd3:
		dasmOpVal("OUT", readByte2())
	case 0xd4:
		dasmOpAddr("CNC", readAddr())
	case 0xd5:
		dasmOpReg("PUSH", "DE")
	case 0xd6:
		dasmOpVal("SUI", readByte2())
	case 0xd7:
		dasmOpReg("RST", "2")
	case 0xd8:
		dasmOp("RC")
	case 0xd9:
		log.Fatal("undocumented RET")
	case 0xda:
		dasmOpAddr("JC", readAddr())
	case 0xdb:
		dasmOpVal("IN", readByte2())
	case 0xdc:
		dasmOpAddr("CC", readAddr())
	case 0xde:
		dasmOpVal("SBI", readByte2())
	case 0xdf:
		dasmOpReg("RST", "3")
	case 0xe0:
		dasmOp("RPO")
	case 0xe1:
		dasmOpReg("POP", "HL")
	case 0xe2:
		dasmOpAddr("JPO", readAddr())
	case 0xe3:
		dasmOp("XTHL")
	case 0xe4:
		dasmOpAddr("CPO", readAddr())
	case 0xe5:
		dasmOpReg("PUSH", "HL")
	case 0xe6:
		dasmOpVal("ANI", readByte2())
	case 0xe7:
		dasmOpReg("RST", "4")
	case 0xe8:
		dasmOp("RPE")
	case 0xe9:
		dasmOp("PCHL")
	case 0xea:
		dasmOpAddr("JPE", readAddr())
	case 0xeb:
		dasmOp("XCHG")
	case 0xec:
		dasmOpAddr("CPE", readAddr())
	case 0xed:
		log.Fatal("undocumented CALL nnnn")
	//case 0xee: fixme XRI nn
	case 0xef:
		dasmOpReg("RST", "5")
	case 0xf0:
		dasmOp("RP")
	case 0xf1:
		dasmOpReg("POP", "PSW")
	case 0xf2:
		dasmOpAddr("JP", readAddr())
	case 0xf3:
		dasmOp("DI")
	case 0xf4:
		dasmOpAddr("CP", readAddr())
	case 0xf5:
		dasmOpReg("PUSH", "PSW")
	case 0xf6:
		dasmOpVal("ORI", readByte2())
	case 0xf7:
		dasmOpReg("RST", "6")
	case 0xf8:
		dasmOp("RM")
	//case 0xf9: fixme SPHL
	case 0xfa:
		dasmOpAddr("JM", readAddr())
	case 0xfb:
		dasmOp("EI")
	case 0xfc:
		dasmOpAddr("CM", readAddr())
	case 0xfd:
		log.Fatal("undocumented CALL nnnn")
	case 0xfe:
		dasmOpVal("CPI", readByte2())
	case 0xff:
		dasmOpReg("RST", "7")
	}
}

func readByte2() uint8 {
	return memRead(regs.PC+1)
}

func readByte3() uint8 {
	return memRead(regs.PC+2)
}

func readAddr() uint16 {
	return uint16(readByte2()) | uint16(readByte3())<<8
}

func dasmOp(op string) {
	printLine(op)
}

func dasmOpReg(op string, r string) {
	printLine(fmt.Sprintf("%s\t%s", op, r))
}

func dasmOpVal(op string, v uint8) {
	printLine(fmt.Sprintf("%s\t0x%02x", op, v))
}

func dasmOpRegVal(op string, r string, v uint8) {
	printLine(fmt.Sprintf("%s\t%s, 0x%02x", op, r, v))
}

func dasmOpReg1Reg2(op string, r1 string, r2 string) {
	printLine(fmt.Sprintf("%s\t%s, %s", op, r1, r2))
}

func dasmOpRegVal1Val2(op string, r string, v1 uint8, v2 uint8) {
	printLine(fmt.Sprintf("%s\t%s, 0x%02x%02x", op, r, v1, v2))
}

func dasmOpAddr(op string, addr uint16) {
	printLine(fmt.Sprintf("%s\t0x%04x", op, addr))
}

func dasmOpRegAddr(op string, r string, addr uint16) {
	printLine(fmt.Sprintf("%s\t%s, 0x%04x", op, r, addr))
}

func dasmOpAddrReg(op string, addr uint16, r string) {
	printLine(fmt.Sprintf("%s\t0x%04x, %s", op, addr, r))
}

func printLine(message string) {
	fmt.Printf("0x%04x   %s\n", regs.PC, message)
}

//func traceState() {
//	if trace {
//		fmt.Printf("> PC = %04x   mem = %s\n", debugPC, traceMem(debugPC))
//		fmt.Printf("> SP = %04x   mem = %s\n", regs.SP, traceMem(regs.SP))
//		fmt.Printf("> regs A:%02x B:%02x C:%02x D:%02x E:%02x H:%02x L:%02x\n",
//			regs.A, regs.B, regs.C, regs.D, regs.E, regs.H, regs.L)
//		fmt.Printf("> flags Z:%t S:%t P:%t AC:%t CY:%t\n", flags.Z, flags.S, flags.P, flags.AC, flags.CY)
//	}
//}
//
//func traceMem(addr uint16) string {
//	return fmt.Sprintf("%02x %02x [%02x] %02x %02x", memReadUnsafe(addr-2), memReadUnsafe(addr-1),
//		memReadUnsafe(addr), memReadUnsafe(addr+1), memReadUnsafe(addr+2))
//}
//
//func memReadUnsafe(addr uint16) uint8 {
//	switch {
//	case inBetween(addr, ramStart, ramEnd):
//		return ram[addr-ramStart]
//	case inBetween(addr, romStart, romEnd):
//		return rom[addr-romStart]
//	default:
//		return 0
//	}
//}
