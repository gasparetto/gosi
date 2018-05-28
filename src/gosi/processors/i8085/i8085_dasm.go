package i8085

import (
	"fmt"
	"log"
)

func printOp() {

	op := memRead(regs.PC)

	switch op {
	case 0x00:
		_op("NOP")
	case 0x01:
		_opRegValVal("LXI", "BC", _byte3(), _byte2())
	case 0x02:
		_opReg("STAX", "BC")
	case 0x03:
		_opReg("INX", "BC")
	case 0x04:
		_opReg("INR", "B")
	case 0x05:
		_opReg("DCR", "B")
	case 0x06:
		_opRegVal("MVI", "B", _byte2())
	case 0x07:
		_op("RLC")
	case 0x08:
		_logFatal("undocumented NOP")
	case 0x09:
		_opReg("DAD", "BC")
	case 0x0a:
		_opReg("LDAX", "BC")
	case 0x0b:
		_opReg("DCX", "BC")
	case 0x0c:
		_opReg("INR", "C")
	case 0x0d:
		_opReg("DCR", "C")
	case 0x0e:
		_opRegVal("MVI", "C", _byte2())
	case 0x0f:
		_op("RRC")
	case 0x10:
		_logFatal("undocumented NOP")
	case 0x11:
		_opRegValVal("LXI", "DE", _byte3(), _byte2())
	case 0x12:
		_opReg("STAX", "DE")
	case 0x13:
		_opReg("INX", "DE")
	case 0x14:
		_opReg("INR", "D")
	case 0x15:
		_opReg("DCR", "D")
	case 0x16:
		_opRegVal("MVI", "D", _byte2())
	case 0x17:
		_op("RAL")
	case 0x18:
		_logFatal("undocumented NOP")
	case 0x19:
		_opReg("DAD", "DE")
	case 0x1a:
		_opReg("LDAX", "DE")
	case 0x1b:
		_opReg("DCX", "DE")
	case 0x1c:
		_opReg("INR", "E")
	case 0x1d:
		_opReg("DCR", "E")
	case 0x1e:
		_opRegVal("MVI", "E", _byte2())
	case 0x1f:
		_op("RAR")
	case 0x20:
		_logFatal("undocumented NOP")
	case 0x21:
		_opRegValVal("LXI", "HL", _byte3(), _byte2())
	case 0x22:
		_opAddr("SHLD", _addr())
	case 0x23:
		_opReg("INX", "HL")
	case 0x24:
		_opReg("INR", "H")
	case 0x25:
		_opReg("DCR", "H")
	case 0x26:
		_opRegVal("MVI", "H", _byte2())
	case 0x27:
		_op("DAA")
	case 0x28:
		_logFatal("undocumented NOP")
	case 0x29:
		_opReg("DAD", "HL")
	case 0x2a:
		_opAddr("LHLD", _addr())
	case 0x2b:
		_opReg("DCX", "HL")
	case 0x2c:
		_opReg("INR", "L")
	case 0x2d:
		_opReg("DCR", "L")
	case 0x2e:
		_opRegVal("MVI", "L", _byte2())
	case 0x2f:
		_op("CMA")
	case 0x30:
		_logFatal("undocumented NOP")
	case 0x31:
		_opRegAddr("LXI", "SP", _addr())
	case 0x32:
		_opAddr("STA", _addr())
	case 0x33:
		_opReg("INX", "SP")
	case 0x34:
		_opReg("INR", "M")
	case 0x35:
		_opReg("DCR", "M")
	case 0x36:
		_opRegAddr("MVI", "M", _addr())
	case 0x37:
		_op("STC")
	case 0x38:
		_logFatal("undocumented NOP")
	case 0x39:
		_opReg("DAD", "SP")
	case 0x3a:
		_opAddr("LDA", _addr())
	case 0x3b:
		_opReg("DCX", "SP")
	case 0x3c:
		_opReg("INR", "A")
	case 0x3d:
		_opReg("DCR", "A")
	case 0x3e:
		_opRegVal("MVI", "A", _byte2())
	case 0x3f:
		_op("CMC")
	case 0x40:
		_opRegReg("MOV", "B", "B")
	case 0x41:
		_opRegReg("MOV", "C", "B")
	case 0x42:
		_opRegReg("MOV", "D", "B")
	case 0x43:
		_opRegReg("MOV", "E", "B")
	case 0x44:
		_opRegReg("MOV", "H", "B")
	case 0x45:
		_opRegReg("MOV", "L", "B")
	case 0x46:
		_opRegAddr("MOV", "B", _addr())
	case 0x47:
		_opRegReg("MOV", "B", "A")
	case 0x48:
		_opRegReg("MOV", "B", "C")
	case 0x49:
		_opRegReg("MOV", "C", "C")
	case 0x4a:
		_opRegReg("MOV", "D", "C")
	case 0x4b:
		_opRegReg("MOV", "E", "C")
	case 0x4c:
		_opRegReg("MOV", "H", "C")
	case 0x4d:
		_opRegReg("MOV", "L", "C")
	case 0x4e:
		_opRegAddr("MOV", "C", _addr())
	case 0x4f:
		_opRegReg("MOV", "C", "A")
	case 0x50:
		_opRegReg("MOV", "B", "D")
	case 0x51:
		_opRegReg("MOV", "C", "D")
	case 0x52:
		_opRegReg("MOV", "D", "D")
	case 0x53:
		_opRegReg("MOV", "E", "D")
	case 0x54:
		_opRegReg("MOV", "H", "D")
	case 0x55:
		_opRegReg("MOV", "L", "D")
	case 0x56:
		_opRegAddr("MOV", "D", _addr())
	case 0x57:
		_opRegReg("MOV", "D", "A")
	case 0x58:
		_opRegReg("MOV", "B", "E")
	case 0x59:
		_opRegReg("MOV", "C", "E")
	case 0x5a:
		_opRegReg("MOV", "D", "E")
	case 0x5b:
		_opRegReg("MOV", "E", "E")
	case 0x5c:
		_opRegReg("MOV", "H", "E")
	case 0x5d:
		_opRegReg("MOV", "L", "E")
	case 0x5e:
		_opRegAddr("MOV", "E", _addr())
	case 0x5f:
		_opRegReg("MOV", "E", "A")
	case 0x60:
		_opRegReg("MOV", "B", "H")
	case 0x61:
		_opRegReg("MOV", "C", "H")
	case 0x62:
		_opRegReg("MOV", "D", "H")
	case 0x63:
		_opRegReg("MOV", "E", "H")
	case 0x64:
		_opRegReg("MOV", "H", "H")
	case 0x65:
		_opRegReg("MOV", "L", "H")
	case 0x66:
		_opRegAddr("MOV", "H", _addr())
	case 0x67:
		_opRegReg("MOV", "H", "A")
	case 0x68:
		_opRegReg("MOV", "B", "L")
	case 0x69:
		_opRegReg("MOV", "C", "L")
	case 0x6a:
		_opRegReg("MOV", "D", "L")
	case 0x6b:
		_opRegReg("MOV", "E", "L")
	case 0x6c:
		_opRegReg("MOV", "H", "L")
	case 0x6d:
		_opRegReg("MOV", "L", "L")
	case 0x6e:
		_opRegAddr("MOV", "L", _addr())
	case 0x6f:
		_opRegReg("MOV", "L", "A")
	case 0x70:
		_opAddrReg("MOV", _addr(), "B")
	case 0x71:
		_opAddrReg("MOV", _addr(), "C")
	case 0x72:
		_opAddrReg("MOV", _addr(), "D")
	case 0x73:
		_opAddrReg("MOV", _addr(), "E")
	case 0x74:
		_opAddrReg("MOV", _addr(), "H")
	case 0x75:
		_opAddrReg("MOV", _addr(), "L")
	case 0x76:
		_op("HLT")
	case 0x77:
		_opAddrReg("MOV", _addr(), "A")
	case 0x78:
		_opRegReg("MOV", "A", "B")
	case 0x79:
		_opRegReg("MOV", "A", "C")
	case 0x7a:
		_opRegReg("MOV", "A", "D")
	case 0x7b:
		_opRegReg("MOV", "A", "E")
	case 0x7c:
		_opRegReg("MOV", "A", "H")
	case 0x7d:
		_opRegReg("MOV", "A", "L")
	case 0x7e:
		_opRegAddr("MOV", "A", _addr())
	case 0x7f:
		_opRegReg("MOV", "A", "A")
	case 0x80:
		_opReg("ADD", "B")
	case 0x81:
		_opReg("ADD", "C")
	case 0x82:
		_opReg("ADD", "D")
	case 0x83:
		_opReg("ADD", "E")
	case 0x84:
		_opReg("ADD", "H")
	case 0x85:
		_opReg("ADD", "L")
	case 0x86:
		_opAddr("ADD", _addr())
	case 0x87:
		_opReg("ADD", "A")
	case 0x88:
		_opReg("ADC", "B")
	case 0x89:
		_opReg("ADC", "C")
	case 0x8a:
		_opReg("ADC", "D")
	case 0x8b:
		_opReg("ADC", "E")
	case 0x8c:
		_opReg("ADC", "H")
	case 0x8d:
		_opReg("ADC", "L")
	case 0x8e:
		_opAddr("ADC", _addr())
	case 0x8f:
		_opReg("ADC", "A")
	case 0x90:
		_opReg("SUB", "B")
	case 0x91:
		_opReg("SUB", "C")
	case 0x92:
		_opReg("SUB", "D")
	case 0x93:
		_opReg("SUB", "E")
	case 0x94:
		_opReg("SUB", "H")
	case 0x95:
		_opReg("SUB", "L")
	case 0x96:
		_opAddr("SUB", _addr())
	case 0x97:
		_opReg("SUB", "A")
	case 0x98:
		_opReg("SBB", "B")
	case 0x99:
		_opReg("SBB", "C")
	case 0x9a:
		_opReg("SBB", "D")
	case 0x9b:
		_opReg("SBB", "E")
	case 0x9c:
		_opReg("SBB", "H")
	case 0x9d:
		_opReg("SBB", "L")
	case 0x9e:
		_opAddr("SBB", _addr())
	case 0x9f:
		_opReg("SBB", "A")
	case 0xa0:
		_opReg("ANA", "B")
	case 0xa1:
		_opReg("ANA", "C")
	case 0xa2:
		_opReg("ANA", "D")
	case 0xa3:
		_opReg("ANA", "E")
	case 0xa4:
		_opReg("ANA", "H")
	case 0xa5:
		_opReg("ANA", "L")
	case 0xa6:
		_opAddr("ANA", _addr())
	case 0xa7:
		_opReg("ANA", "A")
	case 0xa8:
		_opReg("XRA", "B")
	case 0xa9:
		_opReg("XRA", "C")
	case 0xaa:
		_opReg("XRA", "D")
	case 0xab:
		_opReg("XRA", "E")
	case 0xac:
		_opReg("XRA", "H")
	case 0xad:
		_opReg("XRA", "L")
	case 0xae:
		_opAddr("XRA", _addr())
	case 0xaf:
		_opReg("XRA", "A")
	case 0xb0:
		_opReg("ORA", "B")
	case 0xb1:
		_opReg("ORA", "C")
	case 0xb2:
		_opReg("ORA", "D")
	case 0xb3:
		_opReg("ORA", "E")
	case 0xb4:
		_opReg("ORA", "H")
	case 0xb5:
		_opReg("ORA", "L")
	case 0xb6:
		_opAddr("ORA", _addr())
	case 0xb7:
		_opReg("ORA", "A")
	case 0xb8:
		_opReg("CMP", "B")
	case 0xb9:
		_opReg("CMP", "C")
	case 0xba:
		_opReg("CMP", "D")
	case 0xbb:
		_opReg("CMP", "E")
	case 0xbc:
		_opReg("CMP", "H")
	case 0xbd:
		_opReg("CMP", "L")
	case 0xbe:
		_opAddr("CMP", _addr())
	case 0xbf:
		_opReg("CMP", "A")
	case 0xc0:
		_op("RNZ")
	case 0xc1:
		_opReg("POP", "BC")
	case 0xc2:
		_opAddr("JNZ", _addr())
	case 0xc3:
		_opAddr("JMP", _addr())
	case 0xc4:
		_opAddr("CNZ", _addr())
	case 0xc5:
		_opReg("PUSH", "BC")
	case 0xc6:
		_opVal("ADI", _byte2())
	case 0xc7:
		_opReg("RST", "0")
	case 0xc8:
		_op("RZ")
	case 0xc9:
		_op("RET")
	case 0xca:
		_opAddr("JZ", _addr())
	case 0xcb:
		_logFatal("undocumented JMP nnnn")
	case 0xcc:
		_opAddr("CZ", _addr())
	case 0xcd:
		_opAddr("CALL", _addr())
	case 0xce:
		_opVal("ACI", _byte2())
	case 0xcf:
		_opReg("RST", "1")
	case 0xd0:
		_op("RNC")
	case 0xd1:
		_opReg("POP", "DE")
	case 0xd2:
		_opAddr("JNC", _addr())
	case 0xd3:
		_opVal("OUT", _byte2())
	case 0xd4:
		_opAddr("CNC", _addr())
	case 0xd5:
		_opReg("PUSH", "DE")
	case 0xd6:
		_opVal("SUI", _byte2())
	case 0xd7:
		_opReg("RST", "2")
	case 0xd8:
		_op("RC")
	case 0xd9:
		_logFatal("undocumented RET")
	case 0xda:
		_opAddr("JC", _addr())
	case 0xdb:
		_opVal("IN", _byte2())
	case 0xdc:
		_opAddr("CC", _addr())
	case 0xde:
		_opVal("SBI", _byte2())
	case 0xdf:
		_opReg("RST", "3")
	case 0xe0:
		_op("RPO")
	case 0xe1:
		_opReg("POP", "HL")
	case 0xe2:
		_opAddr("JPO", _addr())
	case 0xe3:
		_op("XTHL")
	case 0xe4:
		_opAddr("CPO", _addr())
	case 0xe5:
		_opReg("PUSH", "HL")
	case 0xe6:
		_opVal("ANI", _byte2())
	case 0xe7:
		_opReg("RST", "4")
	case 0xe8:
		_op("RPE")
	case 0xe9:
		_op("PCHL")
	case 0xea:
		_opAddr("JPE", _addr())
	case 0xeb:
		_op("XCHG")
	case 0xec:
		_opAddr("CPE", _addr())
	case 0xed:
		_logFatal("undocumented CALL nnnn")
	case 0xee:
		_opVal("XRI", _byte2())
	case 0xef:
		_opReg("RST", "5")
	case 0xf0:
		_op("RP")
	case 0xf1:
		_opReg("POP", "PSW")
	case 0xf2:
		_opAddr("JP", _addr())
	case 0xf3:
		_op("DI")
	case 0xf4:
		_opAddr("CP", _addr())
	case 0xf5:
		_opReg("PUSH", "PSW")
	case 0xf6:
		_opVal("ORI", _byte2())
	case 0xf7:
		_opReg("RST", "6")
	case 0xf8:
		_op("RM")
	case 0xf9:
		_op("SPHL")
	case 0xfa:
		_opAddr("JM", _addr())
	case 0xfb:
		_op("EI")
	case 0xfc:
		_opAddr("CM", _addr())
	case 0xfd:
		_logFatal("undocumented CALL nnnn")
	case 0xfe:
		_opVal("CPI", _byte2())
	case 0xff:
		_opReg("RST", "7")
	}
}

func _byte2() uint8 {
	return memRead(regs.PC+1)
}

func _byte3() uint8 {
	return memRead(regs.PC+2)
}

func _addr() uint16 {
	return uint16(_byte2()) | uint16(_byte3())<<8
}

func _op(op string) {
	_log(op)
}

func _opReg(op string, r string) {
	_log(fmt.Sprintf("%s\t%s", op, r))
}

func _opVal(op string, v uint8) {
	_log(fmt.Sprintf("%s\t0x%02x", op, v))
}

func _opRegVal(op string, r string, v uint8) {
	_log(fmt.Sprintf("%s\t%s, 0x%02x", op, r, v))
}

func _opRegReg(op string, r1 string, r2 string) {
	_log(fmt.Sprintf("%s\t%s, %s", op, r1, r2))
}

func _opRegValVal(op string, r string, v1 uint8, v2 uint8) {
	_log(fmt.Sprintf("%s\t%s, 0x%02x%02x", op, r, v1, v2))
}

func _opAddr(op string, addr uint16) {
	_log(fmt.Sprintf("%s\t0x%04x", op, addr))
}

func _opRegAddr(op string, r string, addr uint16) {
	_log(fmt.Sprintf("%s\t%s, 0x%04x", op, r, addr))
}

func _opAddrReg(op string, addr uint16, r string) {
	_log(fmt.Sprintf("%s\t0x%04x, %s", op, addr, r))
}

func _log(message string) {
	fmt.Printf("0x%04x   %s\n", regs.PC, message)
}

func _logFatal(message string) {
	log.Fatalf("0x%04x   %s\n", regs.PC, message)
}

func printInternalState() {
	fmt.Printf("> PC = %04x   mem = %s\n", regs.PC, _logMem(regs.PC))
	fmt.Printf("> SP = %04x   mem = %s\n", regs.SP, _logMem(regs.SP))
	fmt.Printf("> regs A:%02x B:%02x C:%02x D:%02x E:%02x H:%02x L:%02x\n",
		regs.A, regs.B, regs.C, regs.D, regs.E, regs.H, regs.L)
	fmt.Printf("> flags Z:%t S:%t P:%t AC:%t CY:%t\n", flags.Z, flags.S, flags.P, flags.AC, flags.CY)
}

func _logMem(addr uint16) string {
	return fmt.Sprintf("%02x %02x [%02x] %02x %02x", _memReadUnsafe(addr-2), _memReadUnsafe(addr-1),
		_memReadUnsafe(addr), _memReadUnsafe(addr+1), _memReadUnsafe(addr+2))
}

func _memReadUnsafe(addr uint16) uint8 {
	switch {
	case inBetween(addr, ramStart, ramEnd):
		return ram[addr-ramStart]
	case inBetween(addr, romStart, romEnd):
		return rom[addr-romStart]
	default:
		return 0
	}
}
